package clicreds

import (
	"bytes"
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
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/internal/shareddefaults"
)

const (
	defaultConsoleLoginRegion        = "ap-southeast-1"
	consoleLoginCacheDirEnv          = "BYTEPLUS_LOGIN_CACHE_DIRECTORY"
	consoleLoginCacheDirRelativePath = "login/cache"
)

// LoginTokenCache is the on-disk cache written by byteplus-cli console-login
// and consumed by this provider.
type LoginTokenCache struct {
	LoginSession string          `json:"login_session"`
	AccessToken  json.RawMessage `json:"access_token"`
	RefreshToken string          `json:"refresh_token,omitempty"`
	IDToken      string          `json:"id_token,omitempty"`
	Scope        string          `json:"scope"`
	ClientID     string          `json:"client_id"`
	EndpointURL  string          `json:"endpoint_url,omitempty"`
	IssuedAt     string          `json:"issued_at"`
	ExpiresIn    int             `json:"expires_in"`
	TokenType    string          `json:"token_type"`

	// 兼容早期 SDK 分支误写出的字段；读取时作为兜底，刷新后继续保留额外字段不影响 CLI。
	ExpiresAt        string `json:"expires_at,omitempty"`
	RefreshExpiresAt string `json:"refresh_expires_at,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
	Region           string `json:"region,omitempty"`
}

func (p *CliProvider) retrieveConsoleLogin(profile *cliProfile, profileName, configPath string) (credentials.Value, error) {
	loginSession := strings.TrimSpace(profile.LoginSession)
	if loginSession == "" {
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigLoginSessionMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain login-session", profileName, configPath),
			nil,
		)
	}

	cacheDir, err := getConsoleLoginCacheDir(configPath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}
	if p.delegate == nil {
		p.delegate = newConsoleLoginRefreshableProvider(loginSession, cacheDir)
	}
	value, err := p.delegate.Retrieve()
	if err == nil {
		p.retrieved = true
	}
	return value, err
}

type consoleLoginOAuthClientAPI interface {
	ExchangeToken(context.Context, *ConsoleTokenRequest) (*ConsoleTokenResponse, error)
}

type ConsoleLoginRefreshableProvider struct {
	loginSession string
	cacheDir     string

	oauthClientFactory func(endpointURL string) consoleLoginOAuthClientAPI

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

func defaultConsoleLoginOAuthFactory(endpointURL string) consoleLoginOAuthClientAPI {
	return NewConsoleOAuthClient(&ConsoleOAuthClientConfig{Endpoint: endpointURL})
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
		return p.applyCredentialsLocked(exp)
	}

	err := p.refreshTokenWithCachedRT(context.Background())
	if err == nil {
		return nil
	}
	if !isConsoleRefreshTokenInvalidErr(err) {
		return err
	}

	diskCache, diskErr := p.loadCacheFromDisk()
	if diskErr != nil {
		return wrapConsoleLoginRefreshErr(err, diskErr)
	}
	if strings.TrimSpace(diskCache.RefreshToken) == "" || diskCache.RefreshToken == p.cached.RefreshToken {
		return bytepluserr.New(
			"CliConsoleLoginRefreshTokenExpired",
			"console-login refresh token has been rejected by signin service; please run 'bp login' to re-authenticate",
			err,
		)
	}

	p.cached = diskCache
	if exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached); ok {
		return p.applyCredentialsLocked(exp)
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

	client := p.oauthClientFactory(consoleLoginCacheEndpointURL(p.cached))
	resp, err := client.ExchangeToken(ctx, &ConsoleTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     p.cached.ClientID,
		Scope:        p.cached.Scope,
		RefreshToken: p.cached.RefreshToken,
	})
	if err != nil {
		if isConsoleRefreshTokenInvalidErr(err) {
			return err
		}
		return bytepluserr.New(
			"CliConsoleLoginRefreshFailed",
			"failed to refresh console-login access token; please run 'bp login' to re-authenticate",
			err,
		)
	}
	if resp == nil || strings.TrimSpace(resp.AccessToken) == "" {
		return bytepluserr.New(
			"CliConsoleLoginRefreshEmpty",
			"console-login refresh response did not include access_token",
			nil,
		)
	}
	if resp.ExpiresIn <= 0 {
		return bytepluserr.New(
			"CliConsoleLoginRefreshExpiresIn",
			"console-login refresh response did not include valid expires_in",
			nil,
		)
	}

	p.cached.AccessToken = json.RawMessage(resp.AccessToken)
	if strings.TrimSpace(resp.RefreshToken) != "" {
		p.cached.RefreshToken = resp.RefreshToken
	}
	if strings.TrimSpace(resp.TokenType) != "" {
		p.cached.TokenType = resp.TokenType
	}
	p.cached.IssuedAt = time.Now().UTC().Format(time.RFC3339)
	p.cached.ExpiresIn = int(resp.ExpiresIn)

	exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached)
	if !ok {
		return bytepluserr.New(
			"CliConsoleLoginAccessTokenInvalid",
			"console-login refresh succeeded but the new access_token could not be parsed into STS credentials",
			nil,
		)
	}
	return p.applyCredentialsLocked(exp)
}

func (p *ConsoleLoginRefreshableProvider) applyCredentialsLocked(exp time.Time) error {
	cachePath, _ := p.cachePath()
	creds, err := parseConsoleSTSCredentials(p.cached.AccessToken, cachePath)
	if err != nil {
		return err
	}
	p.creds = credentials.Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		ProviderName:    CliProviderName,
	}
	p.expiration = exp
	p.initialized = true
	return nil
}

func (p *ConsoleLoginRefreshableProvider) loadCacheFromDisk() (*LoginTokenCache, error) {
	cachePath, err := p.cachePath()
	if err != nil {
		return nil, err
	}
	return loadConsoleLoginTokenCache(cachePath)
}

func (p *ConsoleLoginRefreshableProvider) cachePath() (string, error) {
	if strings.TrimSpace(p.cacheDir) == "" {
		return "", bytepluserr.New(
			"CliConsoleLoginCacheDirMissing",
			fmt.Sprintf("console-login provider for login_session %q has no cache directory", p.loginSession),
			nil,
		)
	}
	hash := sha1.Sum([]byte(strings.TrimSpace(p.loginSession)))
	return filepath.Join(p.cacheDir, fmt.Sprintf("%x.json", hash)), nil
}

func consoleLoginCacheFilePath(loginSession, configPath string) (string, error) {
	dir, err := getConsoleLoginCacheDir(configPath)
	if err != nil {
		return "", err
	}
	hash := sha1.Sum([]byte(strings.TrimSpace(loginSession)))
	return filepath.Join(dir, fmt.Sprintf("%x.json", hash)), nil
}

func getConsoleLoginCacheDir(configPath string) (string, error) {
	if env := strings.TrimSpace(credentials.GetEnvWithFallback(consoleLoginCacheDirEnv)); env != "" {
		return env, nil
	}
	if strings.TrimSpace(configPath) != "" {
		return filepath.Join(filepath.Dir(configPath), consoleLoginCacheDirRelativePath), nil
	}
	home := shareddefaults.UserHomeDir()
	if home == "" {
		return "", bytepluserr.New(
			"CliConsoleLoginCacheDirMissing",
			"failed to resolve console-login cache directory: home directory not found",
			nil,
		)
	}
	return filepath.Join(home, ".byteplus", consoleLoginCacheDirRelativePath), nil
}

func loadConsoleLoginTokenCache(path string) (*LoginTokenCache, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, bytepluserr.New(
				"CliConsoleLoginCacheMissing",
				fmt.Sprintf("console-login cache file %s does not exist; please run 'bp login' first", path),
				err,
			)
		}
		return nil, bytepluserr.New(
			"CliConsoleLoginCacheLoad",
			fmt.Sprintf("failed to load console-login cache file %s", path),
			err,
		)
	}
	if len(strings.TrimSpace(string(data))) == 0 {
		return nil, bytepluserr.New(
			"CliConsoleLoginCacheEmpty",
			fmt.Sprintf("console-login cache file %s was empty", path),
			nil,
		)
	}
	cache := &LoginTokenCache{}
	if err := json.Unmarshal(data, cache); err != nil {
		return nil, bytepluserr.New(
			"CliConsoleLoginCacheUnmarshal",
			fmt.Sprintf("failed to unmarshal console-login cache file %s", path),
			err,
		)
	}
	return cache, nil
}

func parseConsoleSTSCredentials(accessToken json.RawMessage, cachePath string) (*STSCredentials, error) {
	payload, err := consoleAccessTokenPayload(accessToken)
	if err != nil {
		return nil, bytepluserr.New(
			"CliConsoleLoginSTSParse",
			fmt.Sprintf("failed to normalize console-login access_token in %s", cachePath),
			err,
		)
	}
	creds, err := ParseSTSCredentials(payload)
	if err != nil {
		return nil, bytepluserr.New(
			"CliConsoleLoginSTSParse",
			fmt.Sprintf("failed to parse console-login sts credentials in %s", cachePath),
			err,
		)
	}
	return creds, nil
}

func consoleLoginCacheExpiration(cache *LoginTokenCache) (time.Time, bool) {
	if cache == nil {
		return time.Time{}, false
	}
	if exp, ok, err := consoleLoginIssuedExpiration(cache); err == nil && ok {
		return exp, true
	}
	if s := strings.TrimSpace(cache.ExpiresAt); s != "" {
		if exp, err := parseTokenExpiration(s); err == nil {
			return exp, true
		}
	}
	if payload, err := consoleAccessTokenPayload(cache.AccessToken); err == nil {
		creds, err := ParseSTSCredentials(payload)
		if err != nil {
			return time.Time{}, false
		}
		if s := strings.TrimSpace(creds.ExpiredTime); s != "" {
			if exp, err := parseTokenExpiration(s); err == nil {
				return exp, true
			}
		}
	}
	return time.Time{}, false
}

func consoleLoginIssuedExpiration(cache *LoginTokenCache) (time.Time, bool, error) {
	if cache == nil || strings.TrimSpace(cache.IssuedAt) == "" || cache.ExpiresIn <= 0 {
		return time.Time{}, false, nil
	}
	issuedAt, err := parseTokenExpiration(cache.IssuedAt)
	if err != nil {
		return time.Time{}, true, err
	}
	return issuedAt.Add(time.Duration(cache.ExpiresIn) * time.Second), true, nil
}

func tryParseConsoleLoginAccessTokenAndExpiration(cache *LoginTokenCache) (time.Time, bool) {
	if cache == nil || len(bytes.TrimSpace(cache.AccessToken)) == 0 {
		return time.Time{}, false
	}
	exp, ok := consoleLoginCacheExpiration(cache)
	if !ok || !time.Now().Before(exp.Add(-time.Minute)) {
		return time.Time{}, false
	}
	if _, err := parseConsoleSTSCredentials(cache.AccessToken, "console-login cache"); err != nil {
		return time.Time{}, false
	}
	return exp, true
}

func consoleLoginCacheEndpointURL(cache *LoginTokenCache) string {
	if cache == nil {
		return ""
	}
	if endpoint := strings.TrimSpace(cache.EndpointURL); endpoint != "" {
		return endpoint
	}
	return strings.TrimSpace(cache.Endpoint)
}

func isConsoleRefreshTokenInvalidErr(err error) bool {
	if err == nil {
		return false
	}
	type unwrapper interface {
		Unwrap() error
	}
	for cur := err; cur != nil; {
		if apiErr, ok := cur.(*ConsoleOAuthAPIError); ok {
			return apiErr.IsRefreshTokenInvalid()
		}
		uw, ok := cur.(unwrapper)
		if !ok {
			return false
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

func consoleAccessTokenPayload(accessToken json.RawMessage) (string, error) {
	raw := bytes.TrimSpace(accessToken)
	if len(raw) == 0 {
		return "", fmt.Errorf("console access token is empty")
	}
	var encoded string
	if err := json.Unmarshal(raw, &encoded); err == nil {
		encoded = strings.TrimSpace(encoded)
		if encoded == "" {
			return "", fmt.Errorf("console access token string is empty")
		}
		return encoded, nil
	}
	return string(raw), nil
}

// consoleAccessTokenString returns the on-disk JSON representation of the
// access token. Exposed for testing purposes.
func consoleAccessTokenString(creds *STSCredentials) (string, error) {
	if creds == nil {
		return "", fmt.Errorf("sts credentials are nil")
	}
	data, err := json.Marshal(creds)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
