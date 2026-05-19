package clicreds

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeConsoleLoginConfig(t *testing.T, dir, profile string, profileBody map[string]interface{}) string {
	t.Helper()
	cfg := map[string]interface{}{
		"current": profile,
		"profiles": map[string]interface{}{
			profile: profileBody,
		},
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal cli config: %v", err)
	}
	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatalf("write cli config: %v", err)
	}
	return path
}

func writeConsoleLoginCache(t *testing.T, configPath, loginSession string, cache *LoginTokenCache) string {
	t.Helper()
	cachePath, err := consoleLoginCacheFilePath(loginSession, configPath)
	if err != nil {
		t.Fatalf("compute cache path: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(cachePath), 0700); err != nil {
		t.Fatalf("mkdir cache: %v", err)
	}
	data, err := json.Marshal(cache)
	if err != nil {
		t.Fatalf("marshal cache: %v", err)
	}
	if err := os.WriteFile(cachePath, data, 0600); err != nil {
		t.Fatalf("write cache: %v", err)
	}
	return cachePath
}

func TestConsoleLoginUsesCachedCredentials(t *testing.T) {
	dir := t.TempDir()
	configPath := writeConsoleLoginConfig(t, dir, "default", map[string]interface{}{
		"mode":          modeConsoleLogin,
		"login-session": "sess-123",
	})

	sts := &STSCredentials{
		AccessKeyID:     "AKIA",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		CurrentTime:     time.Now().UTC().Format(time.RFC3339),
		ExpiredTime:     time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
	}
	accessToken, err := consoleAccessTokenString(sts)
	if err != nil {
		t.Fatalf("build access token: %v", err)
	}

	writeConsoleLoginCache(t, configPath, "sess-123", &LoginTokenCache{
		LoginSession: "sess-123",
		AccessToken:  accessToken,
		ExpiresAt:    time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
	})

	provider := NewCliProvider(configPath, "default")
	value, err := provider.Retrieve()
	if err != nil {
		t.Fatalf("retrieve: %v", err)
	}
	if value.AccessKeyID != sts.AccessKeyID || value.SecretAccessKey != sts.SecretAccessKey || value.SessionToken != sts.SessionToken {
		t.Fatalf("unexpected credential value: %+v", value)
	}
}

func TestConsoleLoginMissingLoginSession(t *testing.T) {
	dir := t.TempDir()
	configPath := writeConsoleLoginConfig(t, dir, "default", map[string]interface{}{
		"mode": modeConsoleLogin,
	})

	provider := NewCliProvider(configPath, "default")
	_, err := provider.Retrieve()
	if err == nil {
		t.Fatalf("expected error for missing login-session")
	}
	if !strings.Contains(err.Error(), "login-session") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestConsoleLoginExpiredCacheWithoutRefreshToken(t *testing.T) {
	dir := t.TempDir()
	configPath := writeConsoleLoginConfig(t, dir, "default", map[string]interface{}{
		"mode":          modeConsoleLogin,
		"login-session": "sess-456",
	})

	sts := &STSCredentials{
		AccessKeyID:     "AKIA",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		ExpiredTime:     time.Now().Add(-time.Hour).UTC().Format(time.RFC3339),
	}
	accessToken, err := consoleAccessTokenString(sts)
	if err != nil {
		t.Fatalf("build access token: %v", err)
	}
	writeConsoleLoginCache(t, configPath, "sess-456", &LoginTokenCache{
		LoginSession: "sess-456",
		AccessToken:  accessToken,
		ExpiresAt:    time.Now().Add(-time.Hour).UTC().Format(time.RFC3339),
	})

	provider := NewCliProvider(configPath, "default")
	_, err = provider.Retrieve()
	if err == nil {
		t.Fatalf("expected error for expired cache without refresh token")
	}
	if !strings.Contains(err.Error(), "refresh_token") && !strings.Contains(err.Error(), "refresh token") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestConsoleAccessTokenJSONString(t *testing.T) {
	creds := &STSCredentials{AccessKeyID: "AK", SecretAccessKey: "SK", SessionToken: "TK"}
	s, err := consoleAccessTokenString(creds)
	if err != nil {
		t.Fatalf("consoleAccessTokenString: %v", err)
	}
	parsed, err := ParseSTSCredentials(s)
	if err != nil {
		t.Fatalf("ParseSTSCredentials: %v", err)
	}
	if parsed.AccessKeyID != creds.AccessKeyID || parsed.SecretAccessKey != creds.SecretAccessKey || parsed.SessionToken != creds.SessionToken {
		t.Fatalf("round-trip mismatch: %+v vs %+v", parsed, creds)
	}
}
