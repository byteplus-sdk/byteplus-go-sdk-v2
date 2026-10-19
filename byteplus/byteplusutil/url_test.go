package byteplusutil

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
)

var injectionInputs = []string{
	"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
	"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\rbeijing",
	"cn\x00beijing", "evil.com", "cn%2fbeijing", "cn\\beijing", "北京",
	"cn-北京", "cn-beijing\n",
	"1cn-beijing", "1", "-cn-beijing", "_ecs", "cn-beijing-", "rds_mysql_", "-", "_",
}

// recordingLogger records the debug messages written by the resolver.
type recordingLogger struct {
	lines []string
}

func (l *recordingLogger) Log(args ...interface{})   { l.add(args...) }
func (l *recordingLogger) Debug(args ...interface{}) { l.add(args...) }
func (l *recordingLogger) DebugByLevel(_ byteplus.LogLevelType, args ...interface{}) {
	l.add(args...)
}
func (l *recordingLogger) SetDebug(*bool)                          {}
func (l *recordingLogger) SetDebugLogLevel(*byteplus.LogLevelType) {}
func (l *recordingLogger) Warn(args ...interface{})                { l.add(args...) }
func (l *recordingLogger) Error(args ...interface{})               { l.add(args...) }
func (l *recordingLogger) SetIoWriter(io.Writer)                   {}
func (l *recordingLogger) add(args ...interface{})                 { l.lines = append(l.lines, fmt.Sprint(args...)) }

func boolPtr(b bool) *bool { return &b }

func assertCode(t *testing.T, name string, err error, code string, want error) {
	t.Helper()
	var bpErr bytepluserr.Error
	if !errors.As(err, &bpErr) || bpErr.Code() != code {
		t.Fatalf("%s: expect code %s, got %v", name, code, err)
	}
	if err.Error() != want.Error() {
		t.Fatalf("%s: expect %q, got %q", name, want, err)
	}
}

func TestResolveDefaultEndpointRejectsInvalidRegion(t *testing.T) {
	// regional, global (region is not part of the host), cn suffix and dual
	// stack branches: the check runs before any of them.
	for _, service := range []string{"ecs", "iam", "billing", "ark"} {
		for _, dualStack := range []bool{false, true} {
			for _, region := range injectionInputs {
				got, err := ResolveDefaultEndpoint(service, region, boolPtr(dualStack), nil)
				if got != "" {
					t.Fatalf("%s/%q: unexpected endpoint %q", service, region, got)
				}
				assertCode(t, service+"/"+region, err, endpoints.ErrCodeInvalidRegion, endpoints.ValidateRegion(region))
			}
		}
	}
}

func TestResolveDefaultEndpointRejectsInvalidService(t *testing.T) {
	for _, service := range injectionInputs {
		for _, dualStack := range []bool{false, true} {
			got, err := ResolveDefaultEndpoint(service, "cn-beijing", boolPtr(dualStack), nil)
			if got != "" {
				t.Fatalf("%q: unexpected endpoint %q", service, got)
			}
			assertCode(t, service, err, endpoints.ErrCodeInvalidService, endpoints.ValidateService(service))
			// An invalid service must not be reported as merely unregistered.
			if errors.Is(err, ErrServiceEndpointInfoMissing) {
				t.Fatalf("%q: reported as missing service instead of invalid", service)
			}
		}
	}
}

func TestResolveDefaultEndpointValidInputsUnchanged(t *testing.T) {
	cases := []struct {
		service, region string
		dualStack       bool
		want            string
	}{
		{"ecs", "cn-beijing", false, "ecs.cn-beijing.byteplusapi.com.cn"},
		{"ecs", "CN-Beijing", false, "ecs.cn-beijing.byteplusapi.com.cn"},
		{"ecs", "ap-southeast-1", false, "ecs.ap-southeast-1.byteplusapi.com"},
		{"ecs", "cn-hongkong", false, "ecs.cn-hongkong.byteplusapi.com"},
		{"ecs", "ap-southeast-1", true, "ecs.ap-southeast-1.byteplus-api.com"},
		{"rds_mysql", "cn-beijing", false, "rds-mysql.cn-beijing.byteplusapi.com.cn"},
		{"ark", "cn-beijing", false, "ark.cn-beijing.byteplusapi.com"},
		{"iam", "cn-beijing", false, "iam.byteplusapi.com.cn"},
		{"iam", "", false, "iam.byteplusapi.com"},
		{"billing", "ap-southeast-1", true, "billing.byteplus-api.com"},
	}
	for _, c := range cases {
		got, err := ResolveDefaultEndpoint(c.service, c.region, boolPtr(c.dualStack), nil)
		if err != nil || got != c.want {
			t.Fatalf("%s/%s/%v: want %q, got %q, %v", c.service, c.region, c.dualStack, c.want, got, err)
		}
	}

	// Existing errors are unchanged.
	if _, err := ResolveDefaultEndpoint("not_registered", "cn-beijing", boolPtr(false), nil); !errors.Is(err, ErrServiceEndpointInfoMissing) {
		t.Fatalf("expect ErrServiceEndpointInfoMissing, got %v", err)
	}
	if _, err := ResolveDefaultEndpoint("ecs", "", boolPtr(false), nil); !errors.Is(err, ErrMissingRegion) {
		t.Fatalf("expect ErrMissingRegion, got %v", err)
	}
}

// GetDefaultEndpointByServiceInfo can not return an error: it falls back to
// the fixed default endpoint, and the debug log proves the fallback is caused
// by the label check.
func TestGetDefaultEndpointByServiceInfoFallsBackOnInvalidInput(t *testing.T) {
	for _, v := range injectionInputs {
		for _, c := range []struct{ service, region, code string }{
			{"ecs", v, endpoints.ErrCodeInvalidRegion},
			{v, "cn-beijing", endpoints.ErrCodeInvalidService},
		} {
			logger := &recordingLogger{}
			got := GetDefaultEndpointByServiceInfo(c.service, c.region, nil, boolPtr(false), logger)
			if got == nil || *got != "open.ap-southeast-1.byteplusapi.com" {
				t.Fatalf("%q/%q: expect fallback endpoint, got %v", c.service, c.region, got)
			}
			found := false
			for _, line := range logger.lines {
				if strings.Contains(line, "resolve default endpoint failed") && strings.Contains(line, c.code+": ") {
					found = true
				}
			}
			if !found {
				t.Fatalf("%q/%q: no fallback log with code %s, logs: %q", c.service, c.region, c.code, logger.lines)
			}
		}
	}

	got := GetDefaultEndpointByServiceInfo("ecs", "cn-beijing", nil, boolPtr(false), nil)
	if got == nil || *got != "ecs.cn-beijing.byteplusapi.com.cn" {
		t.Fatalf("valid input: unexpected endpoint %v", got)
	}
}

// TestBuiltinDefaultEndpointsAreValid makes sure every built-in service and
// region passes the new rule and still resolves.
func TestBuiltinDefaultEndpointsAreValid(t *testing.T) {
	regions := []string{regionCodeAPSouthEast2, regionCodeAPSouthEast3, regionCodeCNHongkong, "ap-southeast-1", "cn-beijing"}
	for region := range bootstrapRegion {
		regions = append(regions, region)
	}
	for region := range cnNonMainLandRegion {
		regions = append(regions, region)
	}
	if len(defaultEndpoint) == 0 {
		t.Fatalf("defaultEndpoint must not be empty")
	}
	for name, info := range defaultEndpoint {
		if err := endpoints.ValidateService(name); err != nil {
			t.Fatalf("built-in service %q rejected: %v", name, err)
		}
		if err := endpoints.ValidateService(info.Service); err != nil {
			t.Fatalf("built-in service %q rejected: %v", info.Service, err)
		}
		for _, region := range regions {
			if err := endpoints.ValidateRegion(region); err != nil {
				t.Fatalf("built-in region %q rejected: %v", region, err)
			}
			for _, dualStack := range []bool{false, true} {
				got, err := ResolveDefaultEndpoint(name, region, boolPtr(dualStack), nil)
				if err != nil {
					t.Fatalf("built-in %s/%s: unexpected error %v", name, region, err)
				}
				want := standardizeDomainServiceCode(name)
				if !info.IsGlobal {
					want += "." + region
				}
				if dualStack {
					want += ".byteplus-api.com"
				} else {
					want += ".byteplusapi.com"
				}
				if info.GoChinaEnabled && isCNRegion(region) {
					want += ".cn"
				}
				if got != want {
					t.Fatalf("built-in %s/%s: want %q, got %q", name, region, want, got)
				}
			}
		}
	}
}
