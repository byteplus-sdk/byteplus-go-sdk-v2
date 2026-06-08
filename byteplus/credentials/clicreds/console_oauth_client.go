package clicreds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultConsoleEndpoint = "https://signin.byteplus.com"

	consoleAuthorizePath = "/authorize/oauth/authorize"
	consoleTokenPath     = "/authorize/oauth/token"

	consoleClientIDSameDevice  = "trn:signin:::devtools/same-device"
	consoleClientIDCrossDevice = "trn:signin:::devtools/cross-device"

	consoleScopeAllAll = "Console:All:All"
)

// ConsoleOAuthClientConfig configures a ConsoleOAuthClient.
type ConsoleOAuthClientConfig struct {
	Endpoint   string
	HTTPClient *http.Client
	Timeout    time.Duration
}

// ConsoleOAuthClient talks to the BytePlus signin OAuth endpoints used by
// console-login.
type ConsoleOAuthClient struct {
	endpoint   string
	httpClient *http.Client
}

// NewConsoleOAuthClient builds a console-login OAuth client.
func NewConsoleOAuthClient(cfg *ConsoleOAuthClientConfig) *ConsoleOAuthClient {
	if cfg == nil {
		cfg = &ConsoleOAuthClientConfig{}
	}
	endpoint := strings.TrimRight(strings.TrimSpace(cfg.Endpoint), "/")
	if endpoint == "" {
		endpoint = defaultConsoleEndpoint
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		timeout := cfg.Timeout
		if timeout <= 0 {
			timeout = 30 * time.Second
		}
		httpClient = &http.Client{Timeout: timeout}
	}
	return &ConsoleOAuthClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
}

// ConsoleTokenRequest is the input for /authorize/oauth/token.
type ConsoleTokenRequest struct {
	GrantType    string
	ClientID     string
	Scope        string
	Code         string
	CodeVerifier string
	RedirectURI  string
	RefreshToken string
}

// ConsoleTokenResponse is the response payload from /authorize/oauth/token.
type ConsoleTokenResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	IDToken          string `json:"id_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	Scope            string `json:"scope"`
}

// STSCredentials contains the BytePlus STS credentials embedded in the
// console access token JSON payload.
type STSCredentials struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken    string `json:"session_token"`
	CurrentTime     string `json:"current_time,omitempty"`
	ExpiredTime     string `json:"expired_time,omitempty"`
}

// ConsoleOAuthAPIError represents a non-2xx response from the console OAuth
// endpoint.
type ConsoleOAuthAPIError struct {
	StatusCode  int
	ErrorCode   string `json:"error"`
	Description string `json:"error_description"`
	RawBody     string
}

func (e *ConsoleOAuthAPIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Description != "" {
		return fmt.Sprintf("console oauth: %s (%s)", e.Description, e.ErrorCode)
	}
	if e.ErrorCode != "" {
		return fmt.Sprintf("console oauth error: %s", e.ErrorCode)
	}
	return fmt.Sprintf("console oauth: unexpected status %d: %s", e.StatusCode, e.RawBody)
}

// IsRetryable reports whether the error indicates a transient failure.
func (e *ConsoleOAuthAPIError) IsRetryable() bool {
	if e == nil {
		return false
	}
	return isRetryableHTTPStatus(e.StatusCode)
}

// BuildAuthorizeURL builds the URL used to start the OAuth authorize flow.
func (c *ConsoleOAuthClient) BuildAuthorizeURL(clientID, redirectURI, state, codeChallenge, scope string) (string, error) {
	if strings.TrimSpace(c.endpoint) == "" {
		return "", fmt.Errorf("console oauth endpoint is empty")
	}
	u, err := url.Parse(c.endpoint + consoleAuthorizePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse authorize url: %w", err)
	}
	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", clientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("state", state)
	q.Set("code_challenge", codeChallenge)
	q.Set("code_challenge_method", "S256")
	if strings.TrimSpace(scope) == "" {
		scope = consoleScopeAllAll
	}
	q.Set("scope", scope)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// ExchangeToken calls /authorize/oauth/token with the supplied grant.
func (c *ConsoleOAuthClient) ExchangeToken(ctx context.Context, req *ConsoleTokenRequest) (*ConsoleTokenResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("console token request is nil")
	}
	if strings.TrimSpace(req.GrantType) == "" {
		return nil, fmt.Errorf("console token request: grant_type is required")
	}
	if strings.TrimSpace(req.ClientID) == "" {
		return nil, fmt.Errorf("console token request: client_id is required")
	}

	form := url.Values{}
	form.Set("grant_type", req.GrantType)
	form.Set("client_id", req.ClientID)
	if req.Scope != "" {
		form.Set("scope", req.Scope)
	}
	if req.Code != "" {
		form.Set("code", req.Code)
	}
	if req.CodeVerifier != "" {
		form.Set("code_verifier", req.CodeVerifier)
	}
	if req.RedirectURI != "" {
		form.Set("redirect_uri", req.RedirectURI)
	}
	if req.RefreshToken != "" {
		form.Set("refresh_token", req.RefreshToken)
	}

	var tokenResp ConsoleTokenResponse
	err := doWithRetry(ctx, retryOptions{maxAttempts: 3}, func() error {
		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint+consoleTokenPath, strings.NewReader(form.Encode()))
		if err != nil {
			return fmt.Errorf("failed to build console token request: %w", err)
		}
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		httpReq.Header.Set("Accept", "application/json")

		resp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, readErr := ioutil.ReadAll(io.Reader(resp.Body))
		if readErr != nil {
			return fmt.Errorf("failed to read console token response: %w", readErr)
		}
		if resp.StatusCode/100 != 2 {
			apiErr := &ConsoleOAuthAPIError{StatusCode: resp.StatusCode, RawBody: string(body)}
			_ = json.Unmarshal(body, apiErr)
			return apiErr
		}
		var parsed ConsoleTokenResponse
		if err := json.Unmarshal(body, &parsed); err != nil {
			return fmt.Errorf("failed to unmarshal console token response: %w", err)
		}
		if strings.TrimSpace(parsed.AccessToken) == "" {
			return fmt.Errorf("console token response missing access_token")
		}
		tokenResp = parsed
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &tokenResp, nil
}

// ParseSTSCredentials parses the JSON payload that the console embeds in the
// access token (the access token is a JSON-encoded STS document).
func ParseSTSCredentials(accessToken string) (*STSCredentials, error) {
	s := strings.TrimSpace(accessToken)
	if s == "" {
		return nil, fmt.Errorf("console access token is empty")
	}
	var creds STSCredentials
	if err := json.Unmarshal([]byte(s), &creds); err != nil {
		return nil, fmt.Errorf("failed to parse console sts credentials: %w", err)
	}
	if strings.TrimSpace(creds.AccessKeyID) == "" || strings.TrimSpace(creds.SecretAccessKey) == "" || strings.TrimSpace(creds.SessionToken) == "" {
		var legacy struct {
			AccessKeyID     string `json:"AccessKeyId"`
			SecretAccessKey string `json:"SecretAccessKey"`
			SessionToken    string `json:"SessionToken"`
			CurrentTime     string `json:"CurrentTime"`
			ExpiredTime     string `json:"ExpiredTime"`
		}
		if err := json.Unmarshal([]byte(s), &legacy); err == nil {
			if creds.AccessKeyID == "" {
				creds.AccessKeyID = legacy.AccessKeyID
			}
			if creds.SecretAccessKey == "" {
				creds.SecretAccessKey = legacy.SecretAccessKey
			}
			if creds.SessionToken == "" {
				creds.SessionToken = legacy.SessionToken
			}
			if creds.CurrentTime == "" {
				creds.CurrentTime = legacy.CurrentTime
			}
			if creds.ExpiredTime == "" {
				creds.ExpiredTime = legacy.ExpiredTime
			}
		}
	}
	if strings.TrimSpace(creds.AccessKeyID) == "" || strings.TrimSpace(creds.SecretAccessKey) == "" {
		return nil, fmt.Errorf("console sts credentials missing access key or secret key")
	}
	if strings.TrimSpace(creds.SessionToken) == "" {
		return nil, fmt.Errorf("console sts credentials missing session token")
	}
	return &creds, nil
}
