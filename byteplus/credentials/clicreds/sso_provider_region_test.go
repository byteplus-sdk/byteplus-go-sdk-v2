package clicreds

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
)

type failIfCalledOAuthClient struct{ t *testing.T }

func (c failIfCalledOAuthClient) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	c.t.Fatalf("OAuth client must not be used with an invalid region")
	return nil, nil
}

type failIfCalledPortalClient struct{ t *testing.T }

func (c failIfCalledPortalClient) GetRoleCredentials(context.Context, *GetRoleCredentialsRequest) (*GetRoleCredentialsResponse, error) {
	c.t.Fatalf("Portal client must not be used with an invalid region")
	return nil, nil
}

func TestSsoInvalidRegionIsRejectedWithoutNetwork(t *testing.T) {
	regions := []string{
		"@127.0.0.1:9/x#",
		"127.0.0.1:9/#",
		"evil.com/#",
		"a.b",
		"cn beijing",
		"1-region",
		"region-",
	}
	for _, region := range regions {
		region := region
		t.Run(region, func(t *testing.T) {
			dir := t.TempDir()
			tokenPath := filepath.Join(dir, "token.json")
			writeSsoTokenCacheForTest(t, tokenPath, &SsoTokenCache{
				AccessToken:           "access",
				ExpiresAt:             time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
				ClientId:              "client-id",
				ClientSecret:          "client-secret",
				ClientSecretExpiresAt: time.Now().Add(time.Hour).Unix(),
				RefreshToken:          "refresh",
			})
			profile := &cliProfile{AccountId: "12345678", RoleName: "RoleA"}
			p := newSsoRefreshableProvider(tokenPath, profile, "default", filepath.Join(dir, "config.json"), region)
			p.oauthClientFactory = func(string) OAuthClientAPI { return failIfCalledOAuthClient{t} }
			p.portalClientFactory = func(string) PortalClientAPI { return failIfCalledPortalClient{t} }

			_, err := p.Retrieve()
			if err == nil {
				t.Fatalf("Retrieve with region %q: expected error", region)
			}
			be, ok := err.(bytepluserr.Error)
			if !ok || be.Code() != endpoints.ErrCodeInvalidRegion {
				t.Fatalf("Retrieve with region %q: error = %v, want code %s", region, err, endpoints.ErrCodeInvalidRegion)
			}
		})
	}
}
