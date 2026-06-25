package clicreds

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type fakeSsoOAuthClient struct {
	t *testing.T
}

func (c fakeSsoOAuthClient) CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	c.t.Helper()
	if req.GrantType != "refresh_token" {
		c.t.Fatalf("GrantType = %q, want refresh_token", req.GrantType)
	}
	if req.ClientID != "client-id" || req.ClientSecret != "client-secret" || req.RefreshToken != "old-refresh" {
		c.t.Fatalf("unexpected refresh request: %+v", req)
	}
	return &CreateTokenResponse{
		AccessToken:  "new-access",
		TokenType:    "Bearer",
		RefreshToken: "new-refresh",
		ExpiresIn:    900,
	}, nil
}

type fakeSsoPortalClient struct {
	t *testing.T
}

func (c fakeSsoPortalClient) GetRoleCredentials(ctx context.Context, req *GetRoleCredentialsRequest) (*GetRoleCredentialsResponse, error) {
	c.t.Helper()
	if req.AccessToken != "new-access" || req.AccountID != "12345678" || req.RoleName != "RoleA" {
		c.t.Fatalf("unexpected portal request: %+v", req)
	}
	return &GetRoleCredentialsResponse{
		RoleCredentials: RoleCredentials{
			AccessKeyID:     "AK_REFRESHED",
			SecretAccessKey: "SK_REFRESHED",
			SessionToken:    "SESSION_REFRESHED",
			Expiration:      time.Now().Add(time.Hour).Unix(),
		},
	}, nil
}

func writeSsoTokenCacheForTest(t *testing.T, tokenPath string, cache *SsoTokenCache) string {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(tokenPath), 0700); err != nil {
		t.Fatalf("mkdir sso cache: %v", err)
	}
	data, err := json.Marshal(cache)
	if err != nil {
		t.Fatalf("marshal sso cache: %v", err)
	}
	if err := os.WriteFile(tokenPath, data, 0600); err != nil {
		t.Fatalf("write sso cache: %v", err)
	}
	return string(data)
}

func TestSsoRefreshDoesNotRewriteTokenCache(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "sso", "cache", "token.json")
	before := writeSsoTokenCacheForTest(t, tokenPath, &SsoTokenCache{
		StartURL:              "https://signin.byteplus.com/sso/start",
		SessionName:           "session-a",
		AccessToken:           "old-access",
		ExpiresAt:             time.Now().Add(-time.Hour).UTC().Format(time.RFC3339),
		ClientId:              "client-id",
		ClientSecret:          "client-secret",
		ClientSecretExpiresAt: time.Now().Add(time.Hour).Unix(),
		RefreshToken:          "old-refresh",
		Region:                "ap-southeast-1",
	})

	profile := &cliProfile{
		AccountId: "12345678",
		RoleName:  "RoleA",
	}
	provider := newSsoRefreshableProvider(tokenPath, profile, "default", filepath.Join(dir, "config.json"), "ap-southeast-1")
	provider.oauthClientFactory = func(region string) OAuthClientAPI {
		return fakeSsoOAuthClient{t: t}
	}
	provider.portalClientFactory = func(region string) PortalClientAPI {
		return fakeSsoPortalClient{t: t}
	}

	value, err := provider.Retrieve()
	if err != nil {
		t.Fatalf("Retrieve: %v", err)
	}
	if value.AccessKeyID != "AK_REFRESHED" || value.SecretAccessKey != "SK_REFRESHED" || value.SessionToken != "SESSION_REFRESHED" {
		t.Fatalf("unexpected credential value: %+v", value)
	}

	afterData, err := os.ReadFile(tokenPath)
	if err != nil {
		t.Fatalf("read sso cache after refresh: %v", err)
	}
	if string(afterData) != before {
		t.Fatalf("SDK must not rewrite SSO token cache.\nbefore: %s\nafter:  %s", before, string(afterData))
	}
}
