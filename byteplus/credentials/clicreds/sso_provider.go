package clicreds

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
)

// SsoRefreshableProvider keeps the SSO token cache in memory while refreshing.
//
// It reads the CLI cache on bootstrap and once more only when OAuth rejects the
// in-memory refresh token with invalid_grant. The SDK never writes the SSO
// cache file; byteplus-cli remains the sole writer.
type SsoRefreshableProvider struct {
	tokenPath   string
	profile     *cliProfile
	profileName string
	configPath  string
	region      string

	mu          sync.Mutex
	cached      *SsoTokenCache
	creds       credentials.Value
	expiration  time.Time
	initialized bool

	oauthClientFactory  func(region string) OAuthClientAPI
	portalClientFactory func(region string) PortalClientAPI
}

func defaultSsoOAuthFactory(region string) OAuthClientAPI {
	return NewOAuthClient(&OAuthClientConfig{Region: region})
}

func defaultSsoPortalFactory(region string) PortalClientAPI {
	return NewPortalClient(&PortalClientConfig{Region: region})
}

func newSsoRefreshableProvider(tokenPath string, profile *cliProfile, profileName, configPath, region string) *SsoRefreshableProvider {
	return &SsoRefreshableProvider{
		tokenPath:           tokenPath,
		profile:             profile,
		profileName:         profileName,
		configPath:          configPath,
		region:              region,
		oauthClientFactory:  defaultSsoOAuthFactory,
		portalClientFactory: defaultSsoPortalFactory,
	}
}

func (p *SsoRefreshableProvider) Retrieve() (credentials.Value, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.initialized || p.isExpiredLocked() {
		if err := p.refreshLocked(); err != nil {
			return credentials.Value{ProviderName: CliProviderName}, err
		}
	}
	return p.creds, nil
}

func (p *SsoRefreshableProvider) IsExpired() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.initialized {
		return true
	}
	return p.isExpiredLocked()
}

func (p *SsoRefreshableProvider) isExpiredLocked() bool {
	if p.expiration.IsZero() {
		return true
	}
	return !time.Now().Before(p.expiration.Add(-time.Minute))
}

func (p *SsoRefreshableProvider) refreshLocked() error {
	if p.cached == nil {
		cache, err := loadSsoTokenCache(p.tokenPath)
		if err != nil {
			return err
		}
		p.cached = cache
	}

	if !p.isAccessTokenExpiredLocked() {
		return p.fetchRoleCredentialsLocked()
	}

	err := p.refreshAccessTokenLocked(context.Background())
	if err == nil {
		return p.fetchRoleCredentialsLocked()
	}
	if !isSsoRefreshTokenInvalidErr(err) {
		return err
	}

	diskCache, diskErr := loadSsoTokenCache(p.tokenPath)
	if diskErr != nil {
		return wrapSsoRefreshErr(err, diskErr)
	}
	if strings.TrimSpace(diskCache.RefreshToken) == "" ||
		diskCache.RefreshToken == p.cached.RefreshToken {
		return bytepluserr.New(
			"CliSsoTokenRefreshExpired",
			"SSO refresh token has been rejected by the authorization server; please run 'bp sso login' to re-authenticate",
			err,
		)
	}

	p.cached = diskCache
	if !p.isAccessTokenExpiredLocked() {
		return p.fetchRoleCredentialsLocked()
	}
	if err2 := p.refreshAccessTokenLocked(context.Background()); err2 != nil {
		return bytepluserr.New(
			"CliSsoTokenRefreshExpired",
			"SSO refresh token rejected; reloaded cache from disk but new refresh token also failed; please run 'bp sso login'",
			err2,
		)
	}
	return p.fetchRoleCredentialsLocked()
}

func (p *SsoRefreshableProvider) isAccessTokenExpiredLocked() bool {
	if p.cached == nil || strings.TrimSpace(p.cached.AccessToken) == "" {
		return true
	}
	exp, err := parseTokenExpiration(p.cached.ExpiresAt)
	if err != nil {
		return true
	}
	return !time.Now().Before(exp.Add(-time.Minute))
}

func (p *SsoRefreshableProvider) refreshAccessTokenLocked(ctx context.Context) error {
	if p.cached == nil || strings.TrimSpace(p.cached.RefreshToken) == "" {
		return bytepluserr.New(
			"CliSsoTokenRefreshMissing",
			fmt.Sprintf("SSO token cache %s did not contain refresh token; please run 'bp sso login' to re-authenticate", p.tokenPath),
			nil,
		)
	}
	refreshExpired, err := isRefreshTokenExpired(p.cached)
	if err != nil {
		return bytepluserr.New(
			"CliSsoTokenRefreshExpiresParse",
			fmt.Sprintf("failed to parse refresh token expiration in %s; please run 'bp sso login' to re-authenticate", p.tokenPath),
			err,
		)
	}
	if refreshExpired {
		return bytepluserr.New(
			"CliSsoTokenRefreshExpired",
			fmt.Sprintf("SSO refresh token in %s has expired; please run 'bp sso login' to re-authenticate", p.tokenPath),
			nil,
		)
	}
	if strings.TrimSpace(p.cached.ClientId) == "" || strings.TrimSpace(p.cached.ClientSecret) == "" {
		return bytepluserr.New(
			"CliSsoTokenClientMissing",
			fmt.Sprintf("SSO token cache %s did not contain client id/secret; please run 'bp sso login' to re-authenticate", p.tokenPath),
			nil,
		)
	}

	oauthClient := p.oauthClientFactory(p.region)
	tokenResp, err := oauthClient.CreateToken(ctx, &CreateTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     p.cached.ClientId,
		ClientSecret: p.cached.ClientSecret,
		RefreshToken: p.cached.RefreshToken,
	})
	if err != nil {
		if isSsoRefreshTokenInvalidErr(err) {
			return err
		}
		return bytepluserr.New(
			"CliSsoTokenRefreshFailed",
			"failed to refresh SSO access token; please run 'bp sso login' to re-authenticate",
			err,
		)
	}
	if tokenResp == nil || strings.TrimSpace(tokenResp.AccessToken) == "" {
		return bytepluserr.New(
			"CliSsoTokenRefreshEmpty",
			"SSO refresh token response did not include access token; please run 'bp sso login' to re-authenticate",
			nil,
		)
	}
	if tokenResp.ExpiresIn <= 0 {
		return bytepluserr.New(
			"CliSsoTokenRefreshExpiresIn",
			"SSO refresh token response did not include expires_in; please run 'bp sso login' to re-authenticate",
			nil,
		)
	}

	p.cached.AccessToken = tokenResp.AccessToken
	if strings.TrimSpace(tokenResp.RefreshToken) != "" {
		p.cached.RefreshToken = tokenResp.RefreshToken
	}
	p.cached.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second).UTC().Format(time.RFC3339)
	return nil
}

func (p *SsoRefreshableProvider) fetchRoleCredentialsLocked() error {
	accountID := strings.TrimSpace(p.profile.AccountId)
	if accountID == "" {
		return bytepluserr.New(
			"CliConfigAccountIDMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain account-id", p.profileName, p.configPath),
			nil,
		)
	}
	roleName := strings.TrimSpace(p.profile.RoleName)
	if roleName == "" {
		return bytepluserr.New(
			"CliConfigRoleNameMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain role-name", p.profileName, p.configPath),
			nil,
		)
	}

	portalClient := p.portalClientFactory(p.region)
	resp, err := portalClient.GetRoleCredentials(context.Background(), &GetRoleCredentialsRequest{
		AccessToken: p.cached.AccessToken,
		AccountID:   accountID,
		RoleName:    roleName,
	})
	if err != nil {
		return bytepluserr.New(
			"CliSsoPortalCredentials",
			"failed to get SSO role credentials; please run 'bp sso login' to re-authenticate",
			err,
		)
	}

	creds := resp.RoleCredentials
	if strings.TrimSpace(creds.AccessKeyID) == "" || strings.TrimSpace(creds.SecretAccessKey) == "" {
		return bytepluserr.New(
			"CliSsoPortalCredentialsEmpty",
			"SSO portal credentials response did not include access key or secret key",
			nil,
		)
	}

	exp := time.Time{}
	if creds.Expiration > 0 {
		exp = UnixTimestampToTime(creds.Expiration)
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

func isSsoRefreshTokenInvalidErr(err error) bool {
	if err == nil {
		return false
	}
	type unwrapper interface {
		Unwrap() error
	}
	cur := err
	for cur != nil {
		if apiErr, ok := cur.(*OAuthAPIError); ok {
			return apiErr.StatusCode == 400 &&
				strings.EqualFold(apiErr.Response.Error, "invalid_grant")
		}
		uw, ok := cur.(unwrapper)
		if !ok {
			break
		}
		cur = uw.Unwrap()
	}
	return false
}

func wrapSsoRefreshErr(refreshErr, diskErr error) error {
	return bytepluserr.New(
		"CliSsoTokenRefreshFailed",
		fmt.Sprintf("SSO refresh token rejected (%v); failed to reload cache from disk: %v; please run 'bp sso login'", refreshErr, diskErr),
		refreshErr,
	)
}
