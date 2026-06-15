package clicreds

import (
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
	// CliProviderName provides a name of CLI config provider.
	CliProviderName        = "CliProvider"
	modeSSO                = "sso"
	modeAK                 = "ak"
	modeRamRoleArn         = "ramrolearn"
	modeOIDC               = "oidc"
	modeEcsRole            = "ecsrole"
	modeConsoleLogin       = "console-login"
	defaultRegion          = "ap-southeast-1"
	loginCacheDirectoryEnv = "BYTEPLUS_LOGIN_CACHE_DIRECTORY"
)

type cliConfigure struct {
	Current     string                 `json:"current"`
	Profiles    map[string]*cliProfile `json:"profiles"`
	SsoSessions map[string]*SsoSession `json:"sso-session"`
}

type cliProfile struct {
	Mode             string `json:"mode"`
	AccessKey        string `json:"access-key"`
	SecretKey        string `json:"secret-key"`
	SessionToken     string `json:"session-token"`
	StsExpiration    int64  `json:"sts-expiration"`
	SsoSessionName   string `json:"sso-session-name"`
	AccountId        string `json:"account-id"`
	RoleName         string `json:"role-name"`
	OIDCTokenFile    string `json:"oidc-token-file"`
	RoleTrn          string `json:"role-trn"`
	Region           string `json:"region"`
	DisableSSL       bool   `json:"disable-ssl"`
	EndpointResolver string `json:"endpoint-resolver,omitempty"`
	UseDualStack     *bool  `json:"use-dual-stack,omitempty"`
	LoginSession     string `json:"login-session,omitempty"`
}

type SsoSession struct {
	Name               string   `json:"name"`
	StartURL           string   `json:"start-url"`
	Region             string   `json:"region"`
	RegistrationScopes []string `json:"registration-scopes,omitempty"`
}

type SsoTokenCache struct {
	StartURL              string `json:"start_url"`
	SessionName           string `json:"session_name"`
	AccessToken           string `json:"access_token"`
	ExpiresAt             string `json:"expires_at"`
	ClientId              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	ClientIdIssuedAt      int64  `json:"client_id_issued_at,omitempty"`
	ClientSecretExpiresAt int64  `json:"client_secret_expires_at,omitempty"`
	RefreshToken          string `json:"refresh_token,omitempty"`
	Region                string `json:"region"`
}

type LoginTokenCache struct {
	LoginSession string          `json:"login_session"`
	AccessToken  json.RawMessage `json:"access_token"`
	IssuedAt     string          `json:"issued_at"`
	ExpiresIn    int64           `json:"expires_in"`
	TokenType    string          `json:"token_type"`
	RefreshToken string          `json:"refresh_token,omitempty"`
	IDToken      string          `json:"id_token,omitempty"`
	ClientID     string          `json:"client_id,omitempty"`
	Scope        string          `json:"scope,omitempty"`
	EndpointURL  string          `json:"endpoint_url,omitempty"`
}

type consoleLoginStsCredentials struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken    string `json:"session_token"`
}

// CliProvider retrieves credentials from byteplus-cli's config file
// (~/.byteplus/config.json).
//
// It supports reading access-key/secret-key/session-token, and will treat
// sts-expiration (Unix timestamp in seconds or milliseconds) as the credential's expiration time if present.
type CliProvider struct {
	credentials.Expiry

	configPath string
	cacheDir   string
	profile    string
	delegate   credentials.Provider

	retrieved     bool
	hasExpiration bool
}

// NewCliProvider returns a raw *CliProvider for use in the default credential
// chain. Unlike NewCliCredentials, it does not wrap in a Credentials object,
// allowing the chain to manage expiration and caching uniformly.
func NewCliProvider(configPath, profile string) *CliProvider {
	configPath = resolveCliConfigPath(configPath)
	return &CliProvider{
		configPath: configPath,
		cacheDir:   resolveSsoCacheDir(configPath),
		profile:    profile,
		Expiry:     credentials.Expiry{},
	}
}

// NewCliCredentials returns a pointer to a new Credentials object wrapping the
// byteplus-cli config provider.
func NewCliCredentials(configPath, profile string) *credentials.Credentials {
	return credentials.NewExpireAbleCredentials(NewCliProvider(configPath, profile))
}

func (p *CliProvider) Retrieve() (credentials.Value, error) {
	if p.delegate != nil && !p.delegate.IsExpired() {
		return p.delegate.Retrieve()
	}

	p.retrieved = false
	p.hasExpiration = false
	p.delegate = nil
	p.SetExpiration(time.Time{}, 0)

	configPath := p.configPath
	if configPath == "" {
		configPath = resolveCliConfigPath("")
		if configPath == "" {
			return credentials.Value{ProviderName: CliProviderName}, credentials.ErrSharedCredentialsHomeNotFound
		}
		p.configPath = configPath
		p.cacheDir = resolveSsoCacheDir(configPath)
	}

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigLoad",
			fmt.Sprintf("failed to load cli config file %s", configPath),
			err,
		)
	}

	cfg := &cliConfigure{}
	if len(strings.TrimSpace(string(b))) != 0 {
		if err := json.Unmarshal(b, cfg); err != nil {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigUnmarshal",
				fmt.Sprintf("failed to unmarshal cli config file %s", configPath),
				err,
			)
		}
	}

	profileName := p.getProfile(cfg)
	if cfg.Profiles == nil || len(cfg.Profiles) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigNoProfiles",
			fmt.Sprintf("cli config file %s did not contain any profiles", configPath),
			nil,
		)
	}

	profile, ok := cfg.Profiles[profileName]
	if !ok || profile == nil {
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigProfileNotFound",
			fmt.Sprintf("cli config file %s did not contain profile %s", configPath, profileName),
			nil,
		)
	}

	mode := strings.ToLower(strings.TrimSpace(profile.Mode))
	switch mode {
	case "", modeAK:
		return p.retrieveAK(profile, profileName, configPath)
	case modeSSO:
		return p.retrieveSSO(profile, profileName, configPath, cfg)
	case modeRamRoleArn:
		return p.retrieveRamRoleArn(profile, profileName, configPath)
	case modeOIDC:
		return p.retrieveOIDC(profile, profileName, configPath)
	case modeEcsRole:
		return p.retrieveEcsRole(profile, profileName, configPath)
	case modeConsoleLogin:
		return p.retrieveConsoleLogin(profile, profileName, configPath)
	default:
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigModeInvalid",
			fmt.Sprintf("cli config profile %s in %s contained unsupported mode %q", profileName, configPath, profile.Mode),
			nil,
		)
	}
}

func resolveCliConfigPath(configPath string) string {
	if configPath != "" {
		return configPath
	}
	if env := credentials.GetEnvWithFallback("BYTEPLUS_CLI_CONFIG_FILE"); env != "" {
		return env
	}
	home := shareddefaults.UserHomeDir()
	if home == "" {
		return ""
	}
	return filepath.Join(home, ".byteplus", "config.json")
}

func resolveSsoCacheDir(configPath string) string {
	if configPath == "" {
		return ""
	}
	return filepath.Join(filepath.Dir(configPath), "sso", "cache")
}

func (p *CliProvider) retrieveAK(profile *cliProfile, profileName, configPath string) (credentials.Value, error) {
	if len(profile.AccessKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigAccessKey",
			fmt.Sprintf("cli config profile %s in %s did not contain access-key", profileName, configPath),
			nil,
		)
	}
	if len(profile.SecretKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigSecretKey",
			fmt.Sprintf("cli config profile %s in %s did not contain secret-key", profileName, configPath),
			nil,
		)
	}

	p.retrieved = true
	return credentials.Value{
		AccessKeyID:     profile.AccessKey,
		SecretAccessKey: profile.SecretKey,
		SessionToken:    profile.SessionToken,
		ProviderName:    CliProviderName,
	}, nil
}

func (p *CliProvider) retrieveSSO(profile *cliProfile, profileName, configPath string, cfg *cliConfigure) (credentials.Value, error) {
	if value, ok, err := p.useStsCredentialsIfValid(profile, profileName, configPath); err != nil {
		return value, err
	} else if ok {
		return value, nil
	}

	if p.delegate == nil {
		sessionName, startURL, region, err := p.resolveSsoSession(profile, profileName, configPath, cfg)
		if err != nil {
			return credentials.Value{ProviderName: CliProviderName}, err
		}
		tokenPath, err := p.resolveTokenCachePath(startURL, sessionName, profileName, configPath)
		if err != nil {
			return credentials.Value{ProviderName: CliProviderName}, err
		}
		p.delegate = newSsoRefreshableProvider(tokenPath, profile, profileName, configPath, region)
	}
	return p.delegate.Retrieve()
}

func (p *CliProvider) useStsCredentialsIfValid(profile *cliProfile, profileName, configPath string) (credentials.Value, bool, error) {
	expTimestamp := profile.StsExpiration
	if expTimestamp == 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, nil
	}
	if expTimestamp < 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, bytepluserr.New(
			"CliConfigExpiration",
			fmt.Sprintf("cli config profile %s in %s contained invalid sts-expiration", profileName, configPath),
			nil,
		)
	}

	exp := UnixTimestampToTime(expTimestamp)
	if time.Now().After(exp) {
		return credentials.Value{ProviderName: CliProviderName}, false, nil
	}
	if len(profile.AccessKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, bytepluserr.New(
			"CliConfigAccessKey",
			fmt.Sprintf("cli config profile %s in %s did not contain access-key", profileName, configPath),
			nil,
		)
	}
	if len(profile.SecretKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, bytepluserr.New(
			"CliConfigSecretKey",
			fmt.Sprintf("cli config profile %s in %s did not contain secret-key", profileName, configPath),
			nil,
		)
	}

	p.hasExpiration = true
	p.SetExpiration(exp, time.Minute)
	p.retrieved = true
	return credentials.Value{
		AccessKeyID:     profile.AccessKey,
		SecretAccessKey: profile.SecretKey,
		SessionToken:    profile.SessionToken,
		ProviderName:    CliProviderName,
	}, true, nil
}

func (p *CliProvider) resolveSsoSession(profile *cliProfile, profileName, configPath string, cfg *cliConfigure) (string, string, string, error) {
	sessionName := strings.TrimSpace(profile.SsoSessionName)
	if sessionName == "" {
		return "", "", "", bytepluserr.New(
			"CliConfigSsoSessionNameMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain sso-session-name", profileName, configPath),
			nil,
		)
	}
	if cfg == nil || len(cfg.SsoSessions) == 0 {
		return "", "", "", bytepluserr.New(
			"CliConfigSsoSessionsMissing",
			fmt.Sprintf("cli config file %s did not contain any sso-sessions", configPath),
			nil,
		)
	}
	session := cfg.SsoSessions[sessionName]
	if session == nil {
		return "", "", "", bytepluserr.New(
			"CliConfigSsoSessionNotFound",
			fmt.Sprintf("cli config file %s did not contain sso session %s", configPath, sessionName),
			nil,
		)
	}
	startURL := strings.TrimSpace(session.StartURL)
	if startURL == "" {
		return "", "", "", bytepluserr.New(
			"CliConfigSsoStartURLMissing",
			fmt.Sprintf("sso session %s in %s did not contain start-url", sessionName, configPath),
			nil,
		)
	}

	region := strings.TrimSpace(session.Region)
	if region == "" {
		region = strings.TrimSpace(profile.Region)
	}
	if region == "" {
		region = defaultRegion
	}
	return sessionName, startURL, region, nil
}

func (p *CliProvider) resolveTokenCachePath(startURL, sessionName, profileName, configPath string) (string, error) {
	if p.cacheDir == "" {
		return "", bytepluserr.New(
			"CliSsoCacheDirMissing",
			fmt.Sprintf("cli config profile %s in %s did not resolve cache directory; please run 'bp sso login' to re-authenticate", profileName, configPath),
			nil,
		)
	}

	tokenPath := filepath.Join(p.cacheDir, tokenCacheFileName(startURL, sessionName))
	if _, err := os.Stat(tokenPath); err != nil {
		if os.IsNotExist(err) {
			return "", bytepluserr.New(
				"CliSsoTokenCacheMissing",
				fmt.Sprintf("sso token cache file %s does not exist; please run 'bp sso login' to re-authenticate", tokenPath),
				err,
			)
		}
		return "", bytepluserr.New(
			"CliSsoTokenCacheStat",
			fmt.Sprintf("failed to stat sso token cache file %s; please run 'bp sso login' to re-authenticate", tokenPath),
			err,
		)
	}

	return tokenPath, nil
}

func (p *CliProvider) retrieveRamRoleArn(profile *cliProfile, profileName, configPath string) (credentials.Value, error) {
	if p.delegate == nil {
		if len(profile.AccessKey) == 0 {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigAccessKey",
				fmt.Sprintf("cli config profile %s in %s did not contain access-key (required for RamRoleArn mode)", profileName, configPath),
				nil,
			)
		}
		if len(profile.SecretKey) == 0 {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigSecretKey",
				fmt.Sprintf("cli config profile %s in %s did not contain secret-key (required for RamRoleArn mode)", profileName, configPath),
				nil,
			)
		}
		if len(strings.TrimSpace(profile.RoleName)) == 0 {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigRoleNameMissing",
				fmt.Sprintf("cli config profile %s in %s did not contain role-name (required for RamRoleArn mode)", profileName, configPath),
				nil,
			)
		}
		if len(strings.TrimSpace(profile.AccountId)) == 0 {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigAccountIDMissing",
				fmt.Sprintf("cli config profile %s in %s did not contain account-id (required for RamRoleArn mode)", profileName, configPath),
				nil,
			)
		}
		region := strings.TrimSpace(profile.Region)
		if region == "" {
			region = defaultRegion
		}
		schema := "https"
		if profile.DisableSSL {
			schema = "http"
		}
		host := credentials.DefaultEndpoint
		stsProvider := &credentials.StsProvider{
			StsValue: credentials.StsValue{
				AccessKey:       profile.AccessKey,
				SecurityKey:     profile.SecretKey,
				SessionToken:    profile.SessionToken,
				RoleName:        profile.RoleName,
				AccountId:       profile.AccountId,
				Region:          region,
				Schema:          schema,
				Host:            host,
				Timeout:         5 * time.Second,
				DurationSeconds: 3600,
			},
		}
		p.delegate = stsProvider
	}
	return p.delegate.Retrieve()
}

func (p *CliProvider) retrieveOIDC(profile *cliProfile, profileName, configPath string) (credentials.Value, error) {
	if p.delegate == nil {
		tokenFile := strings.TrimSpace(profile.OIDCTokenFile)
		if tokenFile == "" {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigOIDCTokenFileMissing",
				fmt.Sprintf("cli config profile %s in %s did not contain oidc-token-file", profileName, configPath),
				nil,
			)
		}
		roleTrn := strings.TrimSpace(profile.RoleTrn)
		if roleTrn == "" {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigOIDCRoleTrnMissing",
				fmt.Sprintf("cli config profile %s in %s did not contain role-trn", profileName, configPath),
				nil,
			)
		}
		p.delegate = credentials.NewOIDCCredentialsProviderWithOptions(
			tokenFile,
			roleTrn,
			func(o *credentials.OIDCProviderOptions) {
				o.DurationSeconds = 3600
				if profile.DisableSSL {
					o.Schema = "http"
				}
			},
		)
	}
	return p.delegate.Retrieve()
}

func (p *CliProvider) retrieveEcsRole(profile *cliProfile, profileName, configPath string) (credentials.Value, error) {
	if p.delegate == nil {
		roleName := strings.TrimSpace(profile.RoleName)
		if roleName == "" {
			return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
				"CliConfigRoleNameMissing",
				fmt.Sprintf("cli config profile %s in %s did not contain role-name (required for EcsRole mode)", profileName, configPath),
				nil,
			)
		}
		p.delegate = credentials.NewEcsRoleProvider(roleName)
	}
	return p.delegate.Retrieve()
}

func (p *CliProvider) retrieveConsoleLogin(profile *cliProfile, profileName, configPath string) (credentials.Value, error) {
	loginSession := strings.TrimSpace(profile.LoginSession)
	if loginSession == "" {
		return credentials.Value{ProviderName: CliProviderName}, bytepluserr.New(
			"CliConfigLoginSessionMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain login-session; run 'bp login' first", profileName, configPath),
			nil,
		)
	}

	if p.delegate == nil {
		p.delegate = newConsoleLoginRefreshableProvider(loginSession, resolveLoginCacheDir(configPath))
	}
	return p.delegate.Retrieve()
}

func resolveLoginCacheDir(configPath string) string {
	if dir := os.Getenv(loginCacheDirectoryEnv); dir != "" {
		return dir
	}
	return filepath.Join(filepath.Dir(configPath), "login", "cache")
}

func loadSsoTokenCache(path string) (*SsoTokenCache, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, bytepluserr.New(
			"CliSsoTokenCacheLoad",
			fmt.Sprintf("failed to load sso token cache file %s; please run 'bp sso login' to re-authenticate", path),
			err,
		)
	}
	if len(strings.TrimSpace(string(data))) == 0 {
		return nil, bytepluserr.New(
			"CliSsoTokenCacheEmpty",
			fmt.Sprintf("sso token cache file %s was empty; please run 'bp sso login' to re-authenticate", path),
			nil,
		)
	}

	cache := &SsoTokenCache{}
	if err := json.Unmarshal(data, cache); err != nil {
		return nil, bytepluserr.New(
			"CliSsoTokenCacheUnmarshal",
			fmt.Sprintf("failed to unmarshal sso token cache file %s; please run 'bp sso login' to re-authenticate", path),
			err,
		)
	}
	return cache, nil
}

func isTokenExpired(expiresAt string) (bool, error) {
	exp, err := parseTokenExpiration(expiresAt)
	if err != nil {
		return true, err
	}
	return time.Now().After(exp), nil
}

func consoleLoginCacheExpiration(cache *LoginTokenCache) (time.Time, error) {
	if cache == nil {
		return time.Time{}, fmt.Errorf("token cache is nil")
	}
	issuedAt, err := parseTokenExpiration(cache.IssuedAt)
	if err != nil {
		return time.Time{}, err
	}
	if cache.ExpiresIn <= 0 {
		return time.Time{}, fmt.Errorf("expires_in is missing or invalid")
	}
	return issuedAt.Add(time.Duration(cache.ExpiresIn) * time.Second), nil
}

func parseConsoleLoginAccessToken(raw json.RawMessage) (*consoleLoginStsCredentials, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("access_token is empty")
	}

	tokenBytes := []byte(raw)
	var tokenString string
	if err := json.Unmarshal(raw, &tokenString); err == nil {
		tokenBytes = []byte(tokenString)
	}

	var stsCreds consoleLoginStsCredentials
	if err := json.Unmarshal(tokenBytes, &stsCreds); err != nil {
		return nil, err
	}
	if strings.TrimSpace(stsCreds.AccessKeyID) == "" {
		return nil, fmt.Errorf("parsed STS credentials missing access_key_id")
	}
	if strings.TrimSpace(stsCreds.SecretAccessKey) == "" {
		return nil, fmt.Errorf("parsed STS credentials missing secret_access_key")
	}
	if strings.TrimSpace(stsCreds.SessionToken) == "" {
		return nil, fmt.Errorf("parsed STS credentials missing session_token")
	}
	return &stsCreds, nil
}

func parseTokenExpiration(expiresAt string) (time.Time, error) {
	value := strings.TrimSpace(expiresAt)
	if value == "" {
		return time.Time{}, fmt.Errorf("expires_at is empty")
	}
	exp, err := time.Parse(time.RFC3339, value)
	if err == nil {
		return exp, nil
	}
	exp, err = time.Parse(time.RFC3339Nano, value)
	if err == nil {
		return exp, nil
	}
	return time.Time{}, fmt.Errorf("expires_at is invalid: %s", value)
}

func isRefreshTokenExpired(cache *SsoTokenCache) (bool, error) {
	if cache == nil {
		return true, fmt.Errorf("token cache is nil")
	}
	if cache.ClientSecretExpiresAt <= 0 {
		return true, fmt.Errorf("refresh token expiration is missing")
	}
	exp := UnixTimestampToTime(cache.ClientSecretExpiresAt)
	if exp.IsZero() {
		return true, fmt.Errorf("refresh token expiration is invalid")
	}
	return time.Now().After(exp), nil
}

func (p *CliProvider) IsExpired() bool {
	if p.delegate != nil {
		return p.delegate.IsExpired()
	}
	if !p.retrieved {
		return true
	}
	if p.hasExpiration {
		return p.Expiry.IsExpired()
	}
	return false
}

func (p *CliProvider) getProfile(cfg *cliConfigure) string {
	if p.profile == "" {
		p.profile = credentials.GetEnvWithFallback("BYTEPLUS_PROFILE", "BYTEPLUS_CLI_PROFILE")
	}
	if p.profile == "" && cfg != nil && cfg.Current != "" {
		p.profile = cfg.Current
	}
	if p.profile == "" {
		p.profile = "default"
	}
	return p.profile
}

func UnixTimestampToTime(ts int64) time.Time {
	switch {
	case ts >= 1e18:
		return time.Unix(0, ts)
	case ts >= 1e15:
		return time.Unix(0, ts*int64(time.Microsecond))
	case ts >= 1e12:
		sec := ts / 1000
		nsec := (ts % 1000) * int64(time.Millisecond)
		return time.Unix(sec, nsec)
	default:
		return time.Unix(ts, 0)
	}
}

func tokenCacheFileName(startURL, sessionName string) string {
	payload := struct {
		StartURL    string `json:"start_url"`
		SessionName string `json:"session_name"`
	}{
		StartURL:    startURL,
		SessionName: sessionName,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		data = []byte(startURL + "\n" + sessionName)
	}
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x.json", hash)
}
