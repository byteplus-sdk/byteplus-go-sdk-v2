package session_test

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/ecs"
)

var invalidLabels = []string{
	"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
	"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\rbeijing",
	"cn\x00beijing", "evil.com", "cn%2fbeijing", "cn\\beijing", "北京",
	"cn-北京", "cn-beijing\n",
	"1cn-beijing", "1", "-cn-beijing", "_ecs", "cn-beijing-", "rds_mysql_", "-", "_",
}

// recordingTransport counts the requests sent over the network and records
// them. It answers with an empty successful response.
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

func newSession(t *testing.T, region string, tr http.RoundTripper, cfgs ...*byteplus.Config) *session.Session {
	t.Helper()
	cfg := byteplus.NewConfig().
		WithRegion(region).
		WithUseDualStack(false).
		WithCredentials(credentials.NewStaticCredentials("ak", "sk", "")).
		WithHTTPClient(&http.Client{Transport: tr}).
		WithMaxRetries(0)
	sess, err := session.NewSession(append([]*byteplus.Config{cfg}, cfgs...)...)
	if err != nil {
		t.Fatalf("new session: %v", err)
	}
	return sess
}

func assertWrappedLabelError(t *testing.T, name, service string, err error, code string, want error) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s: expect error with code %s, got nil", name, code)
	}
	var bpErr bytepluserr.Error
	if !errors.As(err, &bpErr) || bpErr.Code() != code {
		t.Fatalf("%s: expect error with code %s, got %v", name, code, err)
	}
	if bpErr.Error() != want.Error() {
		t.Fatalf("%s: expect %q, got %q", name, want, bpErr)
	}
	if full := fmt.Sprintf("byteplus session: resolve client config for service %q: %s", service, want); err.Error() != full {
		t.Fatalf("%s: expect %q, got %q", name, full, err)
	}
}

func assertNotSigned(t *testing.T, name string, h http.Header) {
	t.Helper()
	for _, k := range []string{"Authorization", "X-Date", "X-Content-Sha256"} {
		if v := h.Get(k); v != "" {
			t.Fatalf("%s: request was signed, header %s=%q", name, k, v)
		}
	}
}

// The generated client path (service.New -> request -> Send) must not panic
// for an invalid region and must fail before anything is signed or sent,
// whatever resolver is configured.
func TestServiceClientRejectsInvalidRegionBeforeSignAndSend(t *testing.T) {
	resolvers := map[string]func(t *testing.T) endpoints.Resolver{
		"default":  func(*testing.T) endpoints.Resolver { return nil },
		"standard": func(*testing.T) endpoints.Resolver { return endpoints.NewStandardEndpointResolver() },
		"custom": func(t *testing.T) endpoints.Resolver {
			return endpoints.ResolverFunc(func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
				t.Fatalf("custom resolver must not be called for %q/%q", service, region)
				return endpoints.ResolvedEndpoint{}, nil
			})
		},
	}
	for name, newResolver := range resolvers {
		for _, region := range invalidLabels {
			tr := &recordingTransport{}
			var extra []*byteplus.Config
			if r := newResolver(t); r != nil {
				extra = append(extra, byteplus.NewConfig().WithEndpointResolver(r))
			}
			sess := newSession(t, region, tr, extra...)
			label := fmt.Sprintf("[%s] region %q", name, region)

			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Fatalf("%s: panic: %v", label, r)
					}
				}()
				svc := ecs.New(sess)
				if svc.Endpoint != "" {
					t.Fatalf("%s: unexpected endpoint %q", label, svc.Endpoint)
				}

				_, err := svc.DescribeRegions(&ecs.DescribeRegionsInput{})
				assertWrappedLabelError(t, label, "ecs", err, endpoints.ErrCodeInvalidRegion, endpoints.ValidateRegion(region))

				// Send the request built by the client directly: signing
				// mutates the request in place, so it must stay unsigned.
				req, _ := svc.DescribeRegionsRequest(&ecs.DescribeRegionsInput{})
				err = req.Send()
				assertWrappedLabelError(t, label+" Send", "ecs", err, endpoints.ErrCodeInvalidRegion, endpoints.ValidateRegion(region))
				assertNotSigned(t, label, req.HTTPRequest.Header)

				// Presign must not produce a signed URL either.
				req, _ = svc.DescribeRegionsRequest(&ecs.DescribeRegionsInput{})
				u, err := req.Presign(15 * time.Minute)
				if u != "" {
					t.Fatalf("%s: unexpected presigned url %q", label, u)
				}
				assertWrappedLabelError(t, label+" Presign", "ecs", err, endpoints.ErrCodeInvalidRegion, endpoints.ValidateRegion(region))
			}()

			if n := tr.count(); n != 0 {
				t.Fatalf("%s: %d network request(s) were sent", label, n)
			}
		}
	}
}

func TestServiceClientValidRegionIsSignedAndSent(t *testing.T) {
	tr := &recordingTransport{}
	svc := ecs.New(newSession(t, "cn-beijing", tr))
	if want := "https://ecs.cn-beijing.byteplusapi.com.cn"; svc.Endpoint != want {
		t.Fatalf("want endpoint %s, got %s", want, svc.Endpoint)
	}
	req, _ := svc.DescribeRegionsRequest(&ecs.DescribeRegionsInput{})
	if err := req.Send(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tr.count() != 1 {
		t.Fatalf("expect 1 request, got %d", tr.count())
	}
	sent := tr.reqs[0]
	if sent.URL.Host != "ecs.cn-beijing.byteplusapi.com.cn" {
		t.Fatalf("unexpected host %q", sent.URL.Host)
	}
	for _, k := range []string{"Authorization", "X-Date"} {
		if sent.Header.Get(k) == "" || req.HTTPRequest.Header.Get(k) == "" {
			t.Fatalf("valid request must be signed, missing %s", k)
		}
	}
	if !strings.Contains(sent.Header.Get("Authorization"), "/cn-beijing/ecs/request") {
		t.Fatalf("unexpected credential scope: %q", sent.Header.Get("Authorization"))
	}
}

// An explicitly configured endpoint is used as is: the client is built, the
// host is kept and requests are not affected by the label check.
func TestServiceClientExplicitEndpointUnchanged(t *testing.T) {
	for _, region := range []string{"cn-beijing", "cn/beijing"} {
		tr := &recordingTransport{}
		svc := ecs.New(newSession(t, region, tr, byteplus.NewConfig().WithEndpoint("custom.example.com")))
		if svc.Endpoint != "https://custom.example.com" {
			t.Fatalf("region %q: unexpected endpoint %q", region, svc.Endpoint)
		}
		if _, err := svc.DescribeRegions(&ecs.DescribeRegionsInput{}); err != nil {
			t.Fatalf("region %q: unexpected error %v", region, err)
		}
		if tr.count() != 1 || tr.reqs[0].URL.Host != "custom.example.com" {
			t.Fatalf("region %q: request must be sent to the explicit host", region)
		}
	}
}

// builtinServiceNames parses the generated service clients and returns their
// ServiceName / EndpointsID / ServiceID / SigningName constants.
func builtinServiceNames(t *testing.T) map[string]string {
	t.Helper()
	files, err := filepath.Glob(filepath.Join("..", "..", "service", "*", "service_*.go"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no service client found: %v", err)
	}
	names := map[string]string{}
	fset := token.NewFileSet()
	for _, f := range files {
		file, err := parser.ParseFile(fset, f, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", f, err)
		}
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.CONST {
				continue
			}
			for _, spec := range gen.Specs {
				vs := spec.(*ast.ValueSpec)
				for i, id := range vs.Names {
					switch id.Name {
					case "ServiceName", "EndpointsID", "ServiceID", "SigningName":
					default:
						continue
					}
					if i >= len(vs.Values) {
						continue
					}
					lit, ok := vs.Values[i].(*ast.BasicLit)
					if !ok || lit.Kind != token.STRING {
						continue
					}
					v, _ := strconv.Unquote(lit.Value)
					names[f+":"+id.Name] = v
				}
			}
		}
	}
	return names
}

// TestBuiltinServiceClientsStillResolve checks every generated service
// client: its names pass the label check, and its default client config is
// built for the same host as the default resolver computes.
func TestBuiltinServiceClientsStillResolve(t *testing.T) {
	names := builtinServiceNames(t)
	if len(names) < 50 {
		t.Fatalf("too few service names found: %d", len(names))
	}
	regions := []string{"ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "cn-beijing", "cn-hongkong"}
	for where, name := range names {
		if err := endpoints.ValidateService(name); err != nil {
			t.Fatalf("%s: built-in service %q rejected: %v", where, name, err)
		}
		if !strings.HasSuffix(where, ":ServiceName") {
			continue
		}
		for _, region := range regions {
			want, rerr := byteplusutil.ResolveDefaultEndpoint(name, region, byteplus.Bool(false), nil)
			cfg, err := newSession(t, region, &recordingTransport{}).ClientConfigWithError(name)
			if rerr != nil {
				// services without a default endpoint keep failing the same way.
				if !errors.Is(rerr, byteplusutil.ErrServiceEndpointInfoMissing) ||
					!errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing) {
					t.Fatalf("%s/%s: expect error %v, got %v", name, region, rerr, err)
				}
				continue
			}
			if err != nil || cfg.Endpoint != "https://"+want {
				t.Fatalf("%s/%s: want https://%s, got %q, %v", name, region, want, cfg.Endpoint, err)
			}
		}
	}
}
