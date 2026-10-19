package universal_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/universal"
)

var invalidLabels = []string{
	"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
	"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\rbeijing",
	"cn\x00beijing", "evil.com", "cn%2fbeijing", "cn\\beijing", "北京",
	"cn-北京", "cn-beijing\n",
	"1cn-beijing", "1", "-cn-beijing", "_ecs", "cn-beijing-", "rds_mysql_", "-", "_",
}

type recordingTransport struct {
	mu   sync.Mutex
	reqs []*http.Request
}

func (r *recordingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r.mu.Lock()
	r.reqs = append(r.reqs, req)
	r.mu.Unlock()
	return &http.Response{
		StatusCode: 200,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       ioutil.NopCloser(strings.NewReader(`{"ResponseMetadata":{"RequestId":"test"},"Result":{}}`)),
		Request:    req,
	}, nil
}

func (r *recordingTransport) count() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.reqs)
}

func newSession(t *testing.T, region string, tr http.RoundTripper, endpoint string) *session.Session {
	t.Helper()
	cfg := byteplus.NewConfig().
		WithRegion(region).
		WithUseDualStack(false).
		WithCredentials(credentials.NewStaticCredentials("ak", "sk", "")).
		WithHTTPClient(&http.Client{Transport: tr}).
		WithMaxRetries(0)
	if endpoint != "" {
		cfg.WithEndpoint(endpoint)
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		t.Fatalf("new session: %v", err)
	}
	return sess
}

func info(service string) universal.RequestUniversal {
	return universal.RequestUniversal{
		ServiceName: service,
		Action:      "DescribeRegions",
		Version:     "2020-04-01",
		HttpMethod:  universal.GET,
		ContentType: universal.Default,
	}
}

// call runs both universal entry points and returns their errors.
func call(t *testing.T, u *universal.Universal, service string) []error {
	t.Helper()
	var errs []error
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("service %q: panic: %v", service, r)
			}
		}()
		out, err := u.DoCall(info(service), &map[string]interface{}{})
		if out == nil {
			t.Fatalf("service %q: DoCall output must not be nil", service)
		}
		errs = append(errs, err)
		out2 := map[string]interface{}{}
		errs = append(errs, u.DoCallWithType(info(service), &map[string]interface{}{}, &out2))
	}()
	return errs
}

func assertCode(t *testing.T, name string, err error, code string, want error) {
	t.Helper()
	var bpErr bytepluserr.Error
	if !errors.As(err, &bpErr) || bpErr.Code() != code {
		t.Fatalf("%s: expect code %s, got %v", name, code, err)
	}
	if bpErr.Error() != want.Error() {
		t.Fatalf("%s: expect %q, got %q", name, want, bpErr)
	}
}

func TestUniversalRejectsInvalidRegionBeforeSignAndSend(t *testing.T) {
	for _, region := range invalidLabels {
		tr := &recordingTransport{}
		u := universal.New(newSession(t, region, tr, ""))
		for _, err := range call(t, u, "ecs") {
			assertCode(t, fmt.Sprintf("region %q", region), err, endpoints.ErrCodeInvalidRegion, endpoints.ValidateRegion(region))
			if full := fmt.Sprintf("byteplus session: resolve client config for service %q: %s",
				"ecs", endpoints.ValidateRegion(region)); err.Error() != full {
				t.Fatalf("region %q: expect %q, got %q", region, full, err)
			}
		}
		if n := tr.count(); n != 0 {
			t.Fatalf("region %q: %d network request(s) were sent", region, n)
		}
	}
}

// ServiceName is supplied by the caller for every call and builds the
// endpoint host: it is rejected by the session before anything is signed or
// sent.
func TestUniversalRejectsInvalidServiceBeforeSignAndSend(t *testing.T) {
	for _, service := range invalidLabels {
		tr := &recordingTransport{}
		u := universal.New(newSession(t, "cn-beijing", tr, ""))
		for _, err := range call(t, u, service) {
			want := endpoints.ValidateService(service)
			assertCode(t, fmt.Sprintf("service %q", service), err, endpoints.ErrCodeInvalidService, want)
			if full := fmt.Sprintf("byteplus session: resolve client config for service %q: %s",
				service, want); err.Error() != full {
				t.Fatalf("service %q: expect %q, got %q", service, full, err)
			}
		}
		if n := tr.count(); n != 0 {
			t.Fatalf("service %q: %d network request(s) were sent", service, n)
		}
	}
}

func TestUniversalValidInputIsSignedAndSent(t *testing.T) {
	tr := &recordingTransport{}
	u := universal.New(newSession(t, "cn-beijing", tr, ""))
	for _, err := range call(t, u, "ecs") {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if tr.count() != 2 {
		t.Fatalf("expect 2 requests, got %d", tr.count())
	}
	for _, req := range tr.reqs {
		if req.URL.Host != "ecs.cn-beijing.byteplusapi.com.cn" {
			t.Fatalf("unexpected host %q", req.URL.Host)
		}
		if req.Header.Get("Authorization") == "" || req.Header.Get("X-Date") == "" {
			t.Fatalf("valid request must be signed")
		}
	}
}

// The universal client resolves the config on every call, so a region changed
// on the session after the client is built is checked too.
func TestUniversalChecksRegionChangedAfterConstruction(t *testing.T) {
	tr := &recordingTransport{}
	sess := newSession(t, "cn-beijing", tr, "")
	u := universal.New(sess)

	sess.Config.WithRegion("cn-beijing@127.0.0.1")
	for _, err := range call(t, u, "ecs") {
		assertCode(t, "changed region", err, endpoints.ErrCodeInvalidRegion, endpoints.ValidateRegion("cn-beijing@127.0.0.1"))
	}
	if tr.count() != 0 {
		t.Fatalf("%d network request(s) were sent", tr.count())
	}

	sess.Config.WithRegion("ap-southeast-1")
	for _, err := range call(t, u, "ecs") {
		if err != nil {
			t.Fatalf("restored region: unexpected error %v", err)
		}
	}
	if tr.count() != 2 || tr.reqs[0].URL.Host != "ecs.ap-southeast-1.byteplusapi.com" {
		t.Fatalf("restored region: requests must be sent to the original host")
	}
}

// With an explicit endpoint the host is built neither from the region nor
// from ServiceName: requests are not affected and go to the configured host.
func TestUniversalExplicitEndpointUnchanged(t *testing.T) {
	for _, c := range []struct{ region, service string }{
		{"cn/beijing", "ecs"},
		{"cn-beijing", "ecs@127.0.0.1"},
	} {
		tr := &recordingTransport{}
		u := universal.New(newSession(t, c.region, tr, "custom.example.com"))
		for _, err := range call(t, u, c.service) {
			if err != nil {
				t.Fatalf("%q/%q: unexpected error: %v", c.region, c.service, err)
			}
		}
		if tr.count() != 2 || tr.reqs[0].URL.Host != "custom.example.com" {
			t.Fatalf("%q/%q: requests must be sent to the explicit host", c.region, c.service)
		}
	}
}
