package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
)

func newTestSession(t *testing.T, ak, sk, region string, disableSSL *bool, endpoint ...string) *session.Session {
	t.Helper()
	cfg := byteplus.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithRegion(region)
	if disableSSL != nil {
		cfg = cfg.WithDisableSSL(*disableSSL)
	}
	if len(endpoint) > 0 {
		cfg = cfg.WithEndpoint(endpoint[0])
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}
	return sess
}

func TestBuildAuthToken(t *testing.T) {
	tests := []struct {
		name         string
		sess         *session.Session
		dbUser       string
		instanceId   string
		expires      int
		wantErr      bool
		errContains  string
		disableSSL   bool
		region       string
		expectedHost string
	}{
		{
			name:       "valid config and parameters with custom expires",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    1800,
			wantErr:    false,
			region:     "cn-beijing",
		},
		{
			name:       "valid config with default expires",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    0,
			wantErr:    false,
			region:     "cn-beijing",
		},
		{
			name:       "valid config with negative expires",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    -1,
			wantErr:    false,
			region:     "cn-beijing",
		},
		{
			name:       "valid config with DisableSSL true",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", byteplus.Bool(true)),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    900,
			wantErr:    false,
			disableSSL: true,
			region:     "cn-beijing",
		},
		{
			name:         "valid config with custom endpoint",
			sess:         newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil, "custom-rds.example.com"),
			dbUser:       "testuser",
			instanceId:   "mysql-instance-123",
			expires:      900,
			wantErr:      false,
			region:       "cn-beijing",
			expectedHost: "custom-rds.example.com",
		},
		{
			name:        "nil session",
			sess:        nil,
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "session must not be nil",
		},
		{
			name:        "empty dbUser",
			sess:        newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:      "",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "dbUser or instanceId must not be empty",
		},
		{
			name:        "empty instanceId",
			sess:        newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:      "testuser",
			instanceId:  "",
			expires:     900,
			wantErr:     true,
			errContains: "dbUser or instanceId must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			token, err := BuildAuthToken(ctx, tt.sess, tt.dbUser, tt.instanceId, tt.expires)

			if tt.wantErr {
				if err == nil {
					t.Errorf("BuildAuthToken() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("BuildAuthToken() error = %v, should contain %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("BuildAuthToken() unexpected error = %v", err)
				return
			}

			if token == "" {
				t.Error("BuildAuthToken() returned empty token")
			}

			// Verify token contains expected parameters
			if !strings.Contains(token, "Action=ConnectDatabase") {
				t.Error("BuildAuthToken() token should contain Action=ConnectDatabase")
			}
			if !strings.Contains(token, "DBUser="+tt.dbUser) {
				t.Errorf("BuildAuthToken() token should contain DBUser=%s", tt.dbUser)
			}
			if tt.expires > 0 && !strings.Contains(token, "X-Expires=") {
				t.Error("BuildAuthToken() token should contain X-Expires parameter when expires > 0")
			}
			if tt.expires <= 0 && !strings.Contains(token, "X-Expires=900") {
				t.Error("BuildAuthToken() token should contain X-Expires=900 when expires <= 0")
			}

			// Verify token does not contain scheme (MR 1042: host/scheme stripped)
			if strings.HasPrefix(token, "http://") || strings.HasPrefix(token, "https://") {
				t.Error("BuildAuthToken() token should not contain scheme after MR 1042 fix")
			}

			// Verify endpoint host is present in X-Host param
			expectedHost := tt.expectedHost
			if expectedHost == "" && tt.region != "" {
				expectedHost = "rds-mysql." + tt.region + ".byteplusapi.com.cn"
			}
			if expectedHost != "" && !strings.Contains(token, expectedHost) {
				t.Errorf("BuildAuthToken() token should contain endpoint host %s, got %s", expectedHost, token)
			}
		})
	}
}

// BuildAuthToken produces a presigned token: when the endpoint is resolved
// from the region, an invalid region must be rejected with its error code
// before any URL is built or signed.
func TestBuildAuthTokenRejectsInvalidRegion(t *testing.T) {
	invalid := []string{
		"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
		"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\rbeijing",
		"cn\x00beijing", "evil.com", "cn%2fbeijing", "cn\\beijing", "北京",
		"cn-北京", "cn-beijing\n",
		"1cn-beijing", "1", "-cn-beijing", "_ecs", "cn-beijing-", "rds_mysql_", "-", "_",
	}
	for _, region := range invalid {
		sess := newTestSession(t, "ak", "sk", region, nil)
		token, err := BuildAuthToken(context.Background(), sess, "user", "mysql-xxx", 0)
		if token != "" {
			t.Fatalf("region %q: expect empty token, got %q", region, token)
		}
		want := fmt.Sprintf("unable to resolve client config: byteplus session: resolve client config for service %q: %s",
			defaultService, endpoints.ValidateRegion(region))
		if err == nil || err.Error() != want {
			t.Fatalf("region %q: expect %q, got %v", region, want, err)
		}
		var bpErr bytepluserr.Error
		if !errors.As(err, &bpErr) || bpErr.Code() != endpoints.ErrCodeInvalidRegion {
			t.Fatalf("region %q: expect code %s, got %v", region, endpoints.ErrCodeInvalidRegion, err)
		}
	}

	// valid control: the token is built and carries the resolved host.
	token, err := BuildAuthToken(context.Background(),
		newTestSession(t, "ak", "sk", "ap-southeast-1", nil), "user", "mysql-xxx", 0)
	if err != nil || !strings.Contains(token, "X-HOST=https%3A%2F%2Frds-mysql.ap-southeast-1.byteplusapi.com&") ||
		!strings.Contains(token, "X-Signature=") {
		t.Fatalf("valid region: unexpected token %q, %v", token, err)
	}
}

// An explicitly configured endpoint is not built from the region, the token
// is built as before and carries the configured host.
func TestBuildAuthTokenExplicitEndpointUnchanged(t *testing.T) {
	for _, region := range []string{"ap-southeast-1", "cn/beijing"} {
		sess := newTestSession(t, "ak", "sk", region, nil, "custom.example.com")
		token, err := BuildAuthToken(context.Background(), sess, "user", "mysql-xxx", 0)
		if err != nil || !strings.Contains(token, "X-HOST=https%3A%2F%2Fcustom.example.com&") {
			t.Fatalf("region %q: unexpected token %q, %v", region, token, err)
		}
	}
}

// BuildAuthToken must only check the config resolution error, it must not run
// the Validate handlers registered on the session.
func TestBuildAuthTokenDoesNotRunSessionValidateHandlers(t *testing.T) {
	sess := newTestSession(t, "ak", "sk", "ap-southeast-1", nil)
	calls := 0
	sess.Handlers.Validate.PushBack(func(r *request.Request) { calls++ })

	if _, err := BuildAuthToken(context.Background(), sess, "user", "mysql-xxx", 900); err != nil {
		t.Fatalf("BuildAuthToken() unexpected error = %v", err)
	}
	if calls != 0 {
		t.Errorf("session Validate handler was run %d times, want 0", calls)
	}

	// An invalid region still fails, and still without running the handlers.
	bad := newTestSession(t, "ak", "sk", "evil.com/", nil)
	bad.Handlers.Validate.PushBack(func(r *request.Request) { calls++ })
	if _, err := BuildAuthToken(context.Background(), bad, "user", "mysql-xxx", 900); err == nil {
		t.Fatal("BuildAuthToken() expected error for invalid region")
	}
	if calls != 0 {
		t.Errorf("session Validate handler was run %d times, want 0", calls)
	}
}
