package clicreds

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	}
	accessToken, err := consoleAccessTokenString(sts)
	if err != nil {
		t.Fatalf("build access token: %v", err)
	}

	writeConsoleLoginCache(t, configPath, "sess-123", &LoginTokenCache{
		LoginSession: "sess-123",
		AccessToken:  json.RawMessage(accessToken),
		Scope:        consoleScopeAllAll,
		ClientID:     consoleClientIDSameDevice,
		IssuedAt:     time.Now().UTC().Format(time.RFC3339),
		ExpiresIn:    3600,
		TokenType:    "urn:ietf:params:oauth:token-type:access_token_sts",
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
	}
	accessToken, err := consoleAccessTokenString(sts)
	if err != nil {
		t.Fatalf("build access token: %v", err)
	}
	writeConsoleLoginCache(t, configPath, "sess-456", &LoginTokenCache{
		LoginSession: "sess-456",
		AccessToken:  json.RawMessage(accessToken),
		Scope:        consoleScopeAllAll,
		ClientID:     consoleClientIDSameDevice,
		IssuedAt:     time.Now().Add(-time.Hour).UTC().Format(time.RFC3339),
		ExpiresIn:    1,
		TokenType:    "urn:ietf:params:oauth:token-type:access_token_sts",
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

func TestConsoleLoginRefreshesExpiredCLICache(t *testing.T) {
	dir := t.TempDir()
	configPath := writeConsoleLoginConfig(t, dir, "default", map[string]interface{}{
		"mode":          modeConsoleLogin,
		"login-session": "sess-refresh",
	})

	newSTS := &STSCredentials{AccessKeyID: "NEWAK", SecretAccessKey: "NEWSK", SessionToken: "NEWTOKEN"}
	newAccessToken, err := consoleAccessTokenString(newSTS)
	if err != nil {
		t.Fatalf("build refreshed access token: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != consoleTokenPath {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Content-Type"); !strings.Contains(got, "application/x-www-form-urlencoded") {
			t.Fatalf("content-type = %q, want form-urlencoded", got)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		if r.Form.Get("grant_type") != "refresh_token" {
			t.Fatalf("grant_type = %q", r.Form.Get("grant_type"))
		}
		if r.Form.Get("client_id") != consoleClientIDSameDevice {
			t.Fatalf("client_id = %q", r.Form.Get("client_id"))
		}
		if r.Form.Get("scope") != consoleScopeAllAll {
			t.Fatalf("scope = %q", r.Form.Get("scope"))
		}
		if r.Form.Get("refresh_token") != "old-refresh" {
			t.Fatalf("refresh_token = %q", r.Form.Get("refresh_token"))
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token":  newAccessToken,
			"refresh_token": "new-refresh",
			"id_token":      "new-id-token",
			"token_type":    "urn:ietf:params:oauth:token-type:access_token_sts",
			"expires_in":    900,
			"scope":         consoleScopeAllAll,
		}); err != nil {
			t.Fatalf("encode response: %v", err)
		}
	}))
	defer server.Close()

	oldSTS := &STSCredentials{AccessKeyID: "OLD", SecretAccessKey: "OLD", SessionToken: "OLD"}
	oldAccessToken, err := consoleAccessTokenString(oldSTS)
	if err != nil {
		t.Fatalf("build old access token: %v", err)
	}

	cachePath := writeConsoleLoginCache(t, configPath, "sess-refresh", &LoginTokenCache{
		LoginSession: "sess-refresh",
		AccessToken:  json.RawMessage(oldAccessToken),
		RefreshToken: "old-refresh",
		Scope:        consoleScopeAllAll,
		ClientID:     consoleClientIDSameDevice,
		EndpointURL:  server.URL,
		IssuedAt:     time.Now().Add(-time.Hour).UTC().Format(time.RFC3339),
		ExpiresIn:    1,
		TokenType:    "urn:ietf:params:oauth:token-type:access_token_sts",
	})

	provider := NewCliProvider(configPath, "default")
	value, err := provider.Retrieve()
	if err != nil {
		t.Fatalf("retrieve: %v", err)
	}
	if value.AccessKeyID != newSTS.AccessKeyID || value.SecretAccessKey != newSTS.SecretAccessKey || value.SessionToken != newSTS.SessionToken {
		t.Fatalf("unexpected refreshed credential value: %+v", value)
	}

	data, err := os.ReadFile(cachePath)
	if err != nil {
		t.Fatalf("read refreshed cache: %v", err)
	}
	var refreshed LoginTokenCache
	if err := json.Unmarshal(data, &refreshed); err != nil {
		t.Fatalf("unmarshal refreshed cache: %v", err)
	}
	if refreshed.RefreshToken != "new-refresh" || refreshed.IDToken != "new-id-token" {
		t.Fatalf("refresh metadata not persisted: %+v", refreshed)
	}
	if strings.TrimSpace(refreshed.IssuedAt) == "" || refreshed.ExpiresIn != 900 {
		t.Fatalf("CLI expiration fields not persisted: %+v", refreshed)
	}
	if strings.Contains(string(data), `"expires_at"`) {
		t.Fatalf("refreshed cache should use CLI issued_at/expires_in fields, got %s", string(data))
	}
}
