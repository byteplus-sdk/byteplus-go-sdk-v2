package session

import (
	"errors"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/client"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

var endpointLabelInvalidInputs = []string{
	"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
	"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\rbeijing",
	"cn\x00beijing", "evil.com", "cn%2fbeijing", "cn\\beijing", "北京",
	"cn-北京", "cn-beijing\n",
	"1cn-beijing", "1", "-cn-beijing", "_ecs", "cn-beijing-", "rds_mysql_", "-", "_",
}

func newEndpointTestSession(t *testing.T, region string) *Session {
	t.Helper()
	sess, err := NewSession(byteplus.NewConfig().
		WithRegion(region).
		WithUseDualStack(false).
		WithCredentials(credentials.NewStaticCredentials("ak", "sk", "")))
	if err != nil {
		t.Fatalf("new session: %v", err)
	}
	return sess
}

func assertSameError(t *testing.T, name string, got, want error) {
	t.Helper()
	if want == nil {
		t.Fatalf("%s: test case is expected to be invalid", name)
	}
	if got == nil || got.Error() != want.Error() {
		t.Fatalf("%s: expect error %q, got %v", name, want, got)
	}
}

// assertDeferredError checks that the config returned on failure can be
// dereferenced and that its Validate handlers fail requests with want.
func assertDeferredError(t *testing.T, name string, cfg client.Config, want error) {
	t.Helper()
	if cfg.Config == nil {
		t.Fatalf("%s: returned client.Config.Config must not be nil", name)
	}
	if cfg.Endpoint != "" {
		t.Fatalf("%s: no endpoint must be resolved, got %q", name, cfg.Endpoint)
	}
	r := &request.Request{}
	cfg.Handlers.Validate.Run(r)
	assertSameError(t, name+": validate handler", r.Error, want)
}

func TestClientConfigRejectsInvalidRegionAndService(t *testing.T) {
	for _, v := range endpointLabelInvalidInputs {
		for _, c := range []struct {
			kind, service, region, code string
			want                        error
		}{
			{"region", "ecs", v, endpoints.ErrCodeInvalidRegion, endpoints.ValidateRegion(v)},
			{"service", v, "cn-beijing", endpoints.ErrCodeInvalidService, endpoints.ValidateService(v)},
		} {
			name := c.kind + " " + v
			cfg, err := newEndpointTestSession(t, c.region).clientConfigWithErr(c.service)
			assertSameError(t, name, err, c.want)
			if cfg.Config == nil {
				t.Fatalf("%s: returned client.Config.Config must not be nil", name)
			}

			// The public entry points keep the error code reachable with
			// errors.As and keep their historic prefix.
			wrapped := fmt.Errorf("byteplus session: resolve client config for service %q: %s",
				c.service, c.want.Error())

			cfg = newEndpointTestSession(t, c.region).ClientConfig(c.service)
			assertDeferredError(t, name+" ClientConfig", cfg, wrapped)

			cfg, err = newEndpointTestSession(t, c.region).ClientConfigWithError(c.service)
			assertSameError(t, name+" ClientConfigWithError", err, wrapped)
			var bpErr bytepluserr.Error
			if !errors.As(err, &bpErr) || bpErr.Code() != c.code {
				t.Fatalf("%s: expect code %s, got %v", name, c.code, err)
			}
		}
	}
}

// A user supplied EndpointResolver must not be called with an invalid region
// or service: the session level check runs before any resolver.
func TestClientConfigChecksBeforeCustomResolver(t *testing.T) {
	resolver := endpoints.ResolverFunc(
		func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
			t.Fatalf("resolver must not be called for %q/%q", service, region)
			return endpoints.ResolvedEndpoint{}, nil
		})
	for _, v := range endpointLabelInvalidInputs {
		s := newEndpointTestSession(t, v)
		s.Config.WithEndpointResolver(resolver)
		_, err := s.clientConfigWithErr("ecs")
		assertSameError(t, "region "+v, err, endpoints.ValidateRegion(v))

		s = newEndpointTestSession(t, "cn-beijing")
		s.Config.WithEndpointResolver(resolver)
		_, err = s.clientConfigWithErr(v)
		assertSameError(t, "service "+v, err, endpoints.ValidateService(v))
	}

	// A valid input still reaches the resolver.
	called := 0
	s := newEndpointTestSession(t, "cn-beijing")
	s.Config.WithEndpointResolver(endpoints.ResolverFunc(
		func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
			called++
			return endpoints.ResolvedEndpoint{URL: "resolved.example.com"}, nil
		}))
	cfg, err := s.clientConfigWithErr("ecs")
	if err != nil || called != 1 || cfg.Endpoint != "https://resolved.example.com" {
		t.Fatalf("unexpected %q, %d, %v", cfg.Endpoint, called, err)
	}
}

// An explicitly configured endpoint is not built from region / service, the
// client is built as before and the host is kept.
func TestClientConfigExplicitEndpointUnchanged(t *testing.T) {
	for _, v := range endpointLabelInvalidInputs {
		s := newEndpointTestSession(t, v)
		s.Config.WithEndpoint("custom.example.com")
		cfg, err := s.clientConfigWithErr("ecs")
		if err != nil {
			t.Fatalf("region %q: custom endpoint should pass: %v", v, err)
		}
		if cfg.Endpoint != "https://custom.example.com" || cfg.SigningRegion != v {
			t.Fatalf("region %q: unexpected endpoint %q / signing region %q", v, cfg.Endpoint, cfg.SigningRegion)
		}
	}
}

func TestClientConfigValidInputsUnchanged(t *testing.T) {
	cases := []struct{ service, region, want string }{
		{"ecs", "cn-beijing", "https://ecs.cn-beijing.byteplusapi.com.cn"},
		{"ecs", "ap-southeast-1", "https://ecs.ap-southeast-1.byteplusapi.com"},
		{"rds_mysql", "ap-southeast-1", "https://rds-mysql.ap-southeast-1.byteplusapi.com"},
		{"iam", "ap-southeast-1", "https://iam.byteplusapi.com"},
	}
	for _, c := range cases {
		cfg, err := newEndpointTestSession(t, c.region).clientConfigWithErr(c.service)
		if err != nil || cfg.Endpoint != c.want || cfg.SigningRegion != c.region {
			t.Fatalf("%s/%s: want %q, got %q, %v", c.service, c.region, c.want, cfg.Endpoint, err)
		}
	}

	s := newEndpointTestSession(t, "cn-beijing")
	s.Config.WithEndpointResolver(endpoints.NewStandardEndpointResolver())
	cfg, err := s.clientConfigWithErr("ecs")
	if err != nil || cfg.Endpoint != "https://ecs.cn-beijing.byteplusapi.com.cn" {
		t.Fatalf("standard resolver: unexpected %q, %v", cfg.Endpoint, err)
	}
}
