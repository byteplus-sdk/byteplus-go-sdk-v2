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
	LoginSession     string `json:"login_session"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	ExpiresAt        string `json:"expires_at,omitempty"`
	RefreshExpiresAt string `json:"refresh_expires_at,omitempty"`
	Scope            string `json:"scope,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
	Region           string `json:"region,omitempty"`
	ClientID         string `json:"client_id,omitempty"`
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

	cachePath, err := consoleLoginCacheFilePath(loginSession, configPath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}

	cache, err := loadConsoleLoginTokenCache(cachePath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}

	if err := ensureValidConsoleLoginToken(cache, cachePath); err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}

	creds, err := parseConsoleSTSCredentials(cache.AccessToken, cachePath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}

	if exp, ok := consoleLoginCacheExpiration(cache); ok {
		p.hasExpiration = true
		p.SetExpiration(exp, time.Minute)
	}

	p.retrieved = true
	return credentials.Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		ProviderName:    CliProviderName,
	}, nil
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
				fmt.Sprintf("console-login cache file %s does not exist; please run `byteplus configure login` first", path),
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

func saveConsoleLoginTokenCache(path string, cache *LoginTokenCache) error {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return bytepluserr.New(
			"CliConsoleLoginCacheDir",
			fmt.Sprintf("failed to create console-login cache directory for %s", path),
			err,
		)
	}
	if err := writeJSONFileAtomic(path, 0600, cache); err != nil {
		return bytepluserr.New(
			"CliConsoleLoginCacheWrite",
			fmt.Sprintf("failed to write console-login cache file %s", path),
			err,
		)
	}
	return nil
}

func ensureValidConsoleLoginToken(cache *LoginTokenCache, cachePath string) error {
	if cache == nil {
		return bytepluserr.New(
			"CliConsoleLoginCacheNil",
			fmt.Sprintf("console-login cache file %s is nil", cachePath),
			nil,
		)
	}
	if strings.TrimSpace(cache.AccessToken) == "" {
		return bytepluserr.New(
			"CliConsoleLoginAccessTokenMissing",
			fmt.Sprintf("console-login cache file %s did not contain access_token", cachePath),
			nil,
		)
	}

	expired, err := isConsoleLoginAccessTokenExpired(cache)
	if err != nil {
		return bytepluserr.New(
			"CliConsoleLoginExpiresParse",
			fmt.Sprintf("failed to parse console-login expiration in %s", cachePath),
			err,
		)
	}
	if !expired {
		return nil
	}

	if strings.TrimSpace(cache.RefreshToken) == "" {
		return bytepluserr.New(
			"CliConsoleLoginRefreshMissing",
			fmt.Sprintf("console-login cache file %s expired and did not contain refresh_token; please re-run `byteplus configure login`", cachePath),
			nil,
		)
	}
	refreshExpired, err := isConsoleLoginRefreshTokenExpired(cache)
	if err == nil && refreshExpired {
		return bytepluserr.New(
			"CliConsoleLoginRefreshExpired",
			fmt.Sprintf("console-login refresh token in %s has expired; please re-run `byteplus configure login`", cachePath),
			nil,
		)
	}

	clientID := strings.TrimSpace(cache.ClientID)
	if clientID == "" {
		clientID = consoleClientIDSameDevice
	}
	oauthClient := NewConsoleOAuthClient(&ConsoleOAuthClientConfig{Endpoint: cache.Endpoint})
	tokenResp, err := oauthClient.ExchangeToken(context.Background(), &ConsoleTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     clientID,
		RefreshToken: cache.RefreshToken,
	})
	if err != nil {
		return bytepluserr.New(
			"CliConsoleLoginRefreshFailed",
			"failed to refresh console-login access token",
			err,
		)
	}
	if tokenResp == nil || strings.TrimSpace(tokenResp.AccessToken) == "" {
		return bytepluserr.New(
			"CliConsoleLoginRefreshEmpty",
			"console-login refresh response did not include access_token",
			nil,
		)
	}

	cache.AccessToken = tokenResp.AccessToken
	if strings.TrimSpace(tokenResp.RefreshToken) != "" {
		cache.RefreshToken = tokenResp.RefreshToken
	}
	if tokenResp.TokenType != "" {
		cache.TokenType = tokenResp.TokenType
	}
	if tokenResp.ExpiresIn > 0 {
		cache.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second).UTC().Format(time.RFC3339)
	}
	if tokenResp.RefreshExpiresIn > 0 {
		cache.RefreshExpiresAt = time.Now().Add(time.Duration(tokenResp.RefreshExpiresIn) * time.Second).UTC().Format(time.RFC3339)
	}
	if tokenResp.Scope != "" {
		cache.Scope = tokenResp.Scope
	}
	return saveConsoleLoginTokenCache(cachePath, cache)
}

func parseConsoleSTSCredentials(accessToken, cachePath string) (*STSCredentials, error) {
	creds, err := ParseSTSCredentials(accessToken)
	if err != nil {
		return nil, bytepluserr.New(
			"CliConsoleLoginSTSParse",
			fmt.Sprintf("failed to parse console-login sts credentials in %s", cachePath),
			err,
		)
	}
	return creds, nil
}

func isConsoleLoginAccessTokenExpired(cache *LoginTokenCache) (bool, error) {
	if cache == nil {
		return true, fmt.Errorf("console-login cache is nil")
	}
	expiresAt := strings.TrimSpace(cache.ExpiresAt)
	if expiresAt == "" {
		// Without an explicit expiration, fall back to the STS payload's
		// ExpiredTime field if present.
		if creds, err := ParseSTSCredentials(cache.AccessToken); err == nil {
			if strings.TrimSpace(creds.ExpiredTime) != "" {
				exp, err := parseTokenExpiration(creds.ExpiredTime)
				if err == nil {
					return time.Now().After(exp), nil
				}
			}
		}
		return false, nil
	}
	exp, err := parseTokenExpiration(expiresAt)
	if err != nil {
		return true, err
	}
	return time.Now().After(exp), nil
}

func isConsoleLoginRefreshTokenExpired(cache *LoginTokenCache) (bool, error) {
	if cache == nil {
		return true, fmt.Errorf("console-login cache is nil")
	}
	expiresAt := strings.TrimSpace(cache.RefreshExpiresAt)
	if expiresAt == "" {
		return false, nil
	}
	exp, err := parseTokenExpiration(expiresAt)
	if err != nil {
		return true, err
	}
	return time.Now().After(exp), nil
}

func consoleLoginCacheExpiration(cache *LoginTokenCache) (time.Time, bool) {
	if cache == nil {
		return time.Time{}, false
	}
	if s := strings.TrimSpace(cache.ExpiresAt); s != "" {
		if exp, err := parseTokenExpiration(s); err == nil {
			return exp, true
		}
	}
	if creds, err := ParseSTSCredentials(cache.AccessToken); err == nil {
		if s := strings.TrimSpace(creds.ExpiredTime); s != "" {
			if exp, err := parseTokenExpiration(s); err == nil {
				return exp, true
			}
		}
	}
	return time.Time{}, false
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
