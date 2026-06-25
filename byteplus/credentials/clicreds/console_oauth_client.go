package clicreds

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultConsoleLoginEndpoint = "https://signin.byteplus.com"
	consoleLoginTokenPath       = "/authorize/oauth/token"
	consoleLoginRequestTimeout  = 30 * time.Second
	consoleLoginRetryAttempts   = 3
)

// ConsoleLoginOAuthClientConfig configures ConsoleLoginOAuthClient.
type ConsoleLoginOAuthClientConfig struct {
	EndpointURL string
	HTTPClient  *http.Client
}

// ConsoleLoginOAuthClient calls the signin OAuth token endpoint for the SDK
// console-login refresh path. Interactive authorization is handled only by CLI.
type ConsoleLoginOAuthClient struct {
	tokenURL   string
	httpClient *http.Client
}

// ConsoleLoginOAuthClientAPI is the refresh-only interface used by tests.
type ConsoleLoginOAuthClientAPI interface {
	RefreshToken(ctx context.Context, req *ConsoleLoginRefreshTokenRequest) (*ConsoleLoginTokenResponse, error)
}

var _ ConsoleLoginOAuthClientAPI = (*ConsoleLoginOAuthClient)(nil)

// ConsoleLoginRefreshTokenRequest is the refresh_token grant request.
type ConsoleLoginRefreshTokenRequest struct {
	ClientID     string
	Scope        string
	RefreshToken string
}

// ConsoleLoginTokenResponse mirrors fields returned by the signin token API.
type ConsoleLoginTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

type consoleLoginOAuthErrorBody struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorURI         string `json:"error_uri,omitempty"`
}

// ConsoleLoginOAuthAPIError wraps non-2xx signin OAuth responses so callers can
// detect invalid_grant and run the disk-reload fallback.
type ConsoleLoginOAuthAPIError struct {
	StatusCode int
	ErrorCode  string
	ErrorDesc  string
	RawBody    string
	RequestID  string
}

func (e *ConsoleLoginOAuthAPIError) Error() string {
	if e == nil {
		return ""
	}
	var parts []string
	if e.ErrorCode != "" {
		parts = append(parts, e.ErrorCode)
	}
	if e.ErrorDesc != "" {
		parts = append(parts, e.ErrorDesc)
	}
	msg := strings.Join(parts, ": ")
	if msg == "" {
		if e.RawBody != "" {
			msg = e.RawBody
		} else {
			msg = "unknown error"
		}
	}
	suffix := fmt.Sprintf("[status %d", e.StatusCode)
	if e.RequestID != "" {
		suffix += ", requestId: " + e.RequestID
	}
	suffix += "]"
	return fmt.Sprintf("console login oauth request failed: %s %s", msg, suffix)
}

// IsRefreshTokenInvalid reports whether signin rejected the refresh_token.
func (e *ConsoleLoginOAuthAPIError) IsRefreshTokenInvalid() bool {
	if e == nil {
		return false
	}
	return e.StatusCode == http.StatusBadRequest && e.ErrorCode == "invalid_grant"
}

// NewConsoleLoginOAuthClient creates a refresh-only OAuth client.
func NewConsoleLoginOAuthClient(cfg *ConsoleLoginOAuthClientConfig) *ConsoleLoginOAuthClient {
	endpoint := defaultConsoleLoginEndpoint
	if cfg != nil && strings.TrimSpace(cfg.EndpointURL) != "" {
		endpoint = strings.TrimSpace(cfg.EndpointURL)
	}
	endpoint = strings.TrimRight(endpoint, "/")

	client := &http.Client{Timeout: consoleLoginRequestTimeout}
	if cfg != nil && cfg.HTTPClient != nil {
		client = cfg.HTTPClient
	}

	return &ConsoleLoginOAuthClient{
		tokenURL:   endpoint + consoleLoginTokenPath,
		httpClient: client,
	}
}

// RefreshToken exchanges a refresh_token for a new access_token in memory.
func (c *ConsoleLoginOAuthClient) RefreshToken(ctx context.Context, req *ConsoleLoginRefreshTokenRequest) (*ConsoleLoginTokenResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if strings.TrimSpace(req.ClientID) == "" {
		return nil, fmt.Errorf("client_id is required")
	}
	if strings.TrimSpace(req.RefreshToken) == "" {
		return nil, fmt.Errorf("refresh_token is required")
	}

	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("client_id", req.ClientID)
	form.Set("refresh_token", req.RefreshToken)
	if strings.TrimSpace(req.Scope) != "" {
		form.Set("scope", req.Scope)
	}
	body := form.Encode()

	var resp ConsoleLoginTokenResponse
	err := doWithRetry(ctx, retryOptions{maxAttempts: consoleLoginRetryAttempts}, func() error {
		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.tokenURL, strings.NewReader(body))
		if err != nil {
			return fmt.Errorf("failed to build request: %w", err)
		}
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		httpResp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}
		defer httpResp.Body.Close()

		respBytes, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}
		requestID := httpResp.Header.Get("X-Tt-Logid")

		if httpResp.StatusCode/100 != 2 {
			apiErr := &ConsoleLoginOAuthAPIError{
				StatusCode: httpResp.StatusCode,
				RawBody:    string(respBytes),
				RequestID:  requestID,
			}
			if len(respBytes) > 0 {
				var errBody consoleLoginOAuthErrorBody
				if json.Unmarshal(respBytes, &errBody) == nil {
					apiErr.ErrorCode = errBody.Error
					apiErr.ErrorDesc = errBody.ErrorDescription
				}
			}
			return apiErr
		}

		if len(respBytes) > 0 {
			if err := json.Unmarshal(respBytes, &resp); err != nil {
				return fmt.Errorf("failed to decode token response (status %d, requestId: %s): %w", httpResp.StatusCode, requestID, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(resp.AccessToken) == "" {
		return nil, fmt.Errorf("RefreshToken succeeded but access_token was empty")
	}
	if resp.ExpiresIn <= 0 {
		return nil, fmt.Errorf("RefreshToken succeeded but expires_in was missing or invalid")
	}
	return &resp, nil
}
