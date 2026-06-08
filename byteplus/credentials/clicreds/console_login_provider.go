package clicreds

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
)

// ConsoleLoginRefreshableProvider owns the in-memory console-login cache used
// by the SDK. It reads disk only on bootstrap and on invalid_grant fallback.
type ConsoleLoginRefreshableProvider struct {
	loginSession string
	cacheDir     string

	oauthClientFactory func(endpointURL string) ConsoleLoginOAuthClientAPI

	mu          sync.Mutex
	cached      *LoginTokenCache
	creds       credentials.Value
	expiration  time.Time
	initialized bool
}

func newConsoleLoginRefreshableProvider(loginSession, cacheDir string) *ConsoleLoginRefreshableProvider {
	return &ConsoleLoginRefreshableProvider{
		loginSession:       loginSession,
		cacheDir:           cacheDir,
		oauthClientFactory: defaultConsoleLoginOAuthFactory,
	}
}

func defaultConsoleLoginOAuthFactory(endpointURL string) ConsoleLoginOAuthClientAPI {
	return NewConsoleLoginOAuthClient(&ConsoleLoginOAuthClientConfig{
		EndpointURL: endpointURL,
	})
}

func (p *ConsoleLoginRefreshableProvider) Retrieve() (credentials.Value, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.initialized || p.isExpiredLocked() {
		if err := p.refreshLocked(); err != nil {
			return credentials.Value{ProviderName: CliProviderName}, err
		}
	}
	return p.creds, nil
}

func (p *ConsoleLoginRefreshableProvider) IsExpired() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.initialized {
		return true
	}
	return p.isExpiredLocked()
}

func (p *ConsoleLoginRefreshableProvider) isExpiredLocked() bool {
	if p.expiration.IsZero() {
		return true
	}
	return !time.Now().Before(p.expiration.Add(-time.Minute))
}

func (p *ConsoleLoginRefreshableProvider) refreshLocked() error {
	if p.cached == nil {
		cache, err := p.loadCacheFromDisk()
		if err != nil {
			return err
		}
		p.cached = cache
	}

	if exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached); ok {
		p.applyCredentials(exp)
		return nil
	}

	err := p.refreshTokenWithCachedRT(context.Background())
	if err == nil {
		return nil
	}
	if !isRefreshTokenInvalidErr(err) {
		return err
	}

	diskCache, diskErr := p.loadCacheFromDisk()
	if diskErr != nil {
		return wrapConsoleLoginRefreshErr(err, diskErr)
	}
	if strings.TrimSpace(diskCache.RefreshToken) == "" ||
		diskCache.RefreshToken == p.cached.RefreshToken {
		return bytepluserr.New(
			"CliConsoleLoginRefreshTokenExpired",
			"console-login refresh token has been rejected by signin service; please run 'bp login' to re-authenticate",
			err,
		)
	}

	p.cached = diskCache
	if exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached); ok {
		p.applyCredentials(exp)
		return nil
	}
	if err2 := p.refreshTokenWithCachedRT(context.Background()); err2 != nil {
		return bytepluserr.New(
			"CliConsoleLoginRefreshTokenExpired",
			"console-login refresh token rejected; reloaded cache from disk but new refresh token also failed; please run 'bp login'",
			err2,
		)
	}
	return nil
}

func (p *ConsoleLoginRefreshableProvider) refreshTokenWithCachedRT(ctx context.Context) error {
	if p.cached == nil || strings.TrimSpace(p.cached.RefreshToken) == "" {
		return bytepluserr.New(
			"CliConsoleLoginRefreshTokenMissing",
			fmt.Sprintf("console-login cache for login_session %q does not contain refresh_token; please run 'bp login' first", p.loginSession),
			nil,
		)
	}
	if strings.TrimSpace(p.cached.ClientID) == "" {
		return bytepluserr.New(
			"CliConsoleLoginClientIDMissing",
			fmt.Sprintf("console-login cache for login_session %q does not contain client_id; please run 'bp login' to regenerate", p.loginSession),
			nil,
		)
	}

	client := p.oauthClientFactory(p.cached.EndpointURL)
	resp, err := client.RefreshToken(ctx, &ConsoleLoginRefreshTokenRequest{
		ClientID:     p.cached.ClientID,
		Scope:        p.cached.Scope,
		RefreshToken: p.cached.RefreshToken,
	})
	if err != nil {
		if isRefreshTokenInvalidErr(err) {
			return err
		}
		return bytepluserr.New(
			"CliConsoleLoginRefreshTokenFailed",
			"failed to refresh console-login access token; please run 'bp login' to re-authenticate",
			err,
		)
	}

	p.cached.AccessToken = json.RawMessage(resp.AccessToken)
	if strings.TrimSpace(resp.RefreshToken) != "" {
		p.cached.RefreshToken = resp.RefreshToken
	}
	p.cached.IssuedAt = time.Now().UTC().Format(time.RFC3339)
	p.cached.ExpiresIn = int64(resp.ExpiresIn)
	if strings.TrimSpace(resp.TokenType) != "" {
		p.cached.TokenType = resp.TokenType
	}
	if strings.TrimSpace(resp.IDToken) != "" {
		p.cached.IDToken = resp.IDToken
	}
	if strings.TrimSpace(resp.Scope) != "" {
		p.cached.Scope = resp.Scope
	}

	exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached)
	if !ok {
		return bytepluserr.New(
			"CliConsoleLoginAccessTokenInvalid",
			"console-login refresh succeeded but the new access_token could not be parsed into STS credentials",
			nil,
		)
	}
	p.applyCredentials(exp)
	return nil
}

func (p *ConsoleLoginRefreshableProvider) applyCredentials(exp time.Time) {
	creds, err := parseConsoleLoginAccessToken(p.cached.AccessToken)
	if err != nil {
		p.creds = credentials.Value{ProviderName: CliProviderName}
		p.expiration = time.Time{}
		p.initialized = true
		return
	}
	p.creds = credentials.Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		ProviderName:    CliProviderName,
	}
	p.expiration = exp
	p.initialized = true
}

func (p *ConsoleLoginRefreshableProvider) loadCacheFromDisk() (*LoginTokenCache, error) {
	cachePath, err := p.cachePath()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(cachePath); err != nil {
		if os.IsNotExist(err) {
			return nil, bytepluserr.New(
				"CliConsoleLoginCacheMissing",
				fmt.Sprintf("console-login token cache file %s does not exist; please run 'bp login' first", cachePath),
				err,
			)
		}
		return nil, bytepluserr.New(
			"CliConsoleLoginCacheStat",
			fmt.Sprintf("failed to stat console-login token cache file %s", cachePath),
			err,
		)
	}
	data, err := ioutil.ReadFile(cachePath)
	if err != nil {
		return nil, bytepluserr.New(
			"CliConsoleLoginCacheLoad",
			fmt.Sprintf("failed to load console-login token cache file %s", cachePath),
			err,
		)
	}
	if len(strings.TrimSpace(string(data))) == 0 {
		return nil, bytepluserr.New(
			"CliConsoleLoginCacheEmpty",
			fmt.Sprintf("console-login token cache file %s was empty", cachePath),
			nil,
		)
	}
	var cache LoginTokenCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, bytepluserr.New(
			"CliConsoleLoginCacheUnmarshal",
			fmt.Sprintf("failed to unmarshal console-login token cache file %s", cachePath),
			err,
		)
	}
	return &cache, nil
}

func (p *ConsoleLoginRefreshableProvider) cachePath() (string, error) {
	if strings.TrimSpace(p.cacheDir) == "" {
		return "", bytepluserr.New(
			"CliConsoleLoginCacheDirMissing",
			fmt.Sprintf("console-login provider for login_session %q has no cache directory", p.loginSession),
			nil,
		)
	}
	hash := sha1.Sum([]byte(p.loginSession))
	return filepath.Join(p.cacheDir, fmt.Sprintf("%x.json", hash)), nil
}

func tryParseConsoleLoginAccessTokenAndExpiration(cache *LoginTokenCache) (time.Time, bool) {
	if cache == nil {
		return time.Time{}, false
	}
	exp, err := consoleLoginCacheExpiration(cache)
	if err != nil {
		return time.Time{}, false
	}
	if !time.Now().Before(exp.Add(-time.Minute)) {
		return time.Time{}, false
	}
	if _, err := parseConsoleLoginAccessToken(cache.AccessToken); err != nil {
		return time.Time{}, false
	}
	return exp, true
}

func isRefreshTokenInvalidErr(err error) bool {
	if err == nil {
		return false
	}
	type unwrapper interface {
		Unwrap() error
	}
	cur := err
	for cur != nil {
		if apiErr, ok := cur.(*ConsoleLoginOAuthAPIError); ok {
			return apiErr.IsRefreshTokenInvalid()
		}
		uw, ok := cur.(unwrapper)
		if !ok {
			break
		}
		cur = uw.Unwrap()
	}
	return false
}

func wrapConsoleLoginRefreshErr(refreshErr, diskErr error) error {
	return bytepluserr.New(
		"CliConsoleLoginRefreshTokenExpired",
		fmt.Sprintf("console-login refresh token rejected (%v); failed to reload cache from disk: %v; please run 'bp login'", refreshErr, diskErr),
		refreshErr,
	)
}
