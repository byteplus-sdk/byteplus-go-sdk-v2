package endpoints

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
)

// labelRegexp is the reference definition of a valid endpoint label.
var labelRegexp = regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_-]*[A-Za-z0-9])?$`)

// injectionInputs can alter the structure of a host name built from them.
var injectionInputs = []string{
	"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
	"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\rbeijing",
	"cn\x00beijing", "evil.com", "cn%2fbeijing", "cn\\beijing", "北京",
	"cn-北京", "cn-beijing\n",
}

// boundaryInputs only use allowed characters but break the first / last
// character rules.
var boundaryInputs = []string{
	"1cn-beijing", "1", "-cn-beijing", "_ecs", "cn-beijing-", "rds_mysql_", "-", "_",
}

var validInputs = []string{
	"cn-beijing", "ap-southeast-1", "a", "rds_mysql", "Redis", "Rds_mysql-2", "",
}

const ruleSuffix = "must start with a letter, end with a letter or digit, " +
	"and contain only letters, digits, '-' and '_'"

func expectedMessage(kind, value, reason string) string {
	return fmt.Sprintf("invalid %s %q: %s, %s %s", kind, value, reason, kind, ruleSuffix)
}

// assertLabelError checks the error type, code, message and Error() output.
func assertLabelError(t *testing.T, err error, code, kind, value, reason string) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s %q: expect error, got nil", kind, value)
	}
	var bpErr bytepluserr.Error
	if !errors.As(err, &bpErr) {
		t.Fatalf("%s %q: expect bytepluserr.Error, got %T: %v", kind, value, err, err)
	}
	if bpErr.Code() != code {
		t.Fatalf("%s %q: expect code %s, got %s", kind, value, code, bpErr.Code())
	}
	wantMsg := expectedMessage(kind, value, reason)
	if bpErr.Message() != wantMsg {
		t.Fatalf("%s %q: expect message\n%q\ngot\n%q", kind, value, wantMsg, bpErr.Message())
	}
	if want := code + ": " + wantMsg; err.Error() != want {
		t.Fatalf("%s %q: expect Error()\n%q\ngot\n%q", kind, value, want, err.Error())
	}
	if bpErr.OrigErr() != nil {
		t.Fatalf("%s %q: unexpected orig error %v", kind, value, bpErr.OrigErr())
	}
}

func TestValidateRegionAndServiceReasons(t *testing.T) {
	cases := []struct {
		value  string
		reason string
	}{
		// illegal character, the index is the first offending byte.
		{"cn-beijing@127.0.0.1", "illegal character at index 10"},
		{"cn-beijing:443", "illegal character at index 10"},
		{"cn/beijing", "illegal character at index 2"},
		{"cn?beijing", "illegal character at index 2"},
		{"cn#beijing", "illegal character at index 2"},
		{"cn beijing", "illegal character at index 2"},
		{"cn\tbeijing", "illegal character at index 2"},
		{"cn\nbeijing", "illegal character at index 2"},
		{"cn\rbeijing", "illegal character at index 2"},
		{"cn\x00beijing", "illegal character at index 2"},
		{"evil.com", "illegal character at index 4"},
		{"cn%2fbeijing", "illegal character at index 2"},
		{"cn\\beijing", "illegal character at index 2"},
		{"北京", "illegal character at index 0"},
		{"cn-北京", "illegal character at index 3"},
		{"cn-beijing\n", "illegal character at index 10"},
		// the character set is checked before the first / last rules.
		{"1cn.beijing", "illegal character at index 3"},
		{"-.", "illegal character at index 1"},
		// first character
		{"1cn-beijing", "must start with a letter"},
		{"1", "must start with a letter"},
		{"-cn-beijing", "must start with a letter"},
		{"_ecs", "must start with a letter"},
		{"-", "must start with a letter"},
		{"_", "must start with a letter"},
		{"1abc-", "must start with a letter"},
		// last character
		{"cn-beijing-", "must not end with '-' or '_'"},
		{"rds_mysql_", "must not end with '-' or '_'"},
		{"a-", "must not end with '-' or '_'"},
	}
	for _, c := range cases {
		assertLabelError(t, ValidateRegion(c.value), ErrCodeInvalidRegion, "region", c.value, c.reason)
		assertLabelError(t, ValidateService(c.value), ErrCodeInvalidService, "service", c.value, c.reason)
	}

	// Every listed injection and boundary input must be covered above.
	covered := map[string]bool{}
	for _, c := range cases {
		covered[c.value] = true
	}
	for _, v := range append(append([]string{}, injectionInputs...), boundaryInputs...) {
		if !covered[v] {
			t.Fatalf("input %q has no explicit reason assertion", v)
		}
	}
}

func TestValidateRegionAndServiceAcceptValidInputs(t *testing.T) {
	for _, v := range validInputs {
		if err := ValidateRegion(v); err != nil {
			t.Fatalf("region %q: unexpected error %v", v, err)
		}
		if err := ValidateService(v); err != nil {
			t.Fatalf("service %q: unexpected error %v", v, err)
		}
		if err := ValidateServiceAndRegion(v, v); err != nil {
			t.Fatalf("service/region %q: unexpected error %v", v, err)
		}
	}
}

func TestValidateServiceAndRegionOrder(t *testing.T) {
	// service is checked first.
	err := ValidateServiceAndRegion("ecs@evil", "cn/beijing")
	assertLabelError(t, err, ErrCodeInvalidService, "service", "ecs@evil", "illegal character at index 3")

	err = ValidateServiceAndRegion("ecs", "cn/beijing")
	assertLabelError(t, err, ErrCodeInvalidRegion, "region", "cn/beijing", "illegal character at index 2")

	err = ValidateServiceAndRegion("", "cn/beijing")
	assertLabelError(t, err, ErrCodeInvalidRegion, "region", "cn/beijing", "illegal character at index 2")

	err = ValidateServiceAndRegion("ecs_", "")
	assertLabelError(t, err, ErrCodeInvalidService, "service", "ecs_", "must not end with '-' or '_'")
}

// checkEquivalence compares the validator with the reference regular
// expression, and checks that the reason matches the documented order.
func checkEquivalence(t *testing.T, s string) {
	t.Helper()
	regionErr := ValidateRegion(s)
	serviceErr := ValidateService(s)
	want := s == "" || labelRegexp.MatchString(s)
	if (regionErr == nil) != want || (serviceErr == nil) != want {
		t.Fatalf("input %q: regexp match %v, region err %v, service err %v", s, want, regionErr, serviceErr)
	}
	if want {
		return
	}
	reason := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_') {
			reason = fmt.Sprintf("illegal character at index %d", i)
			break
		}
	}
	if reason == "" {
		if c := s[0]; !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			reason = "must start with a letter"
		} else {
			reason = "must not end with '-' or '_'"
		}
	}
	if got := regionErr.(bytepluserr.Error).Message(); got != expectedMessage("region", s, reason) {
		t.Fatalf("input %q: unexpected region message %q", s, got)
	}
	if got := serviceErr.(bytepluserr.Error).Message(); got != expectedMessage("service", s, reason) {
		t.Fatalf("input %q: unexpected service message %q", s, got)
	}
}

// TestValidateEndpointLabelMatchesRegexpExhaustively compares the validator
// with the reference regular expression for every 1 and 2 byte input.
func TestValidateEndpointLabelMatchesRegexpExhaustively(t *testing.T) {
	checkEquivalence(t, "")
	for a := 0; a < 256; a++ {
		checkEquivalence(t, string([]byte{byte(a)}))
		for b := 0; b < 256; b++ {
			checkEquivalence(t, string([]byte{byte(a), byte(b)}))
		}
	}
	// 3 byte inputs over a representative alphabet, including every
	// boundary class and a non ASCII byte.
	alphabet := []byte{'a', 'Z', '0', '9', '-', '_', '.', '@', ':', '/', ' ', 0x00, 0x7f, 0x80, 0xff}
	for _, a := range alphabet {
		for _, b := range alphabet {
			for _, c := range alphabet {
				checkEquivalence(t, string([]byte{a, b, c}))
			}
		}
	}
}

func TestStandardEndpointResolverRejectsInvalidServiceAndRegion(t *testing.T) {
	r := NewStandardEndpointResolverWithOptions(&StandardEndpointOptions{
		Format:    defaultFormat,
		SiteStack: SiteStackByteplusIPv4,
		// Custom service names are rendered into the host as well.
		CustomServices: map[string]ServiceInfo{},
	})
	for _, v := range append(append([]string{}, injectionInputs...), boundaryInputs...) {
		r.CustomServices[v] = ServiceInfo{Service: v}

		res, err := r.EndpointFor("ecs", v)
		if res.URL != "" {
			t.Fatalf("region %q: unexpected url %q", v, res.URL)
		}
		if want := ValidateRegion(v); err == nil || err.Error() != want.Error() {
			t.Fatalf("region %q: expect %v, got %v", v, want, err)
		}

		res, err = r.EndpointFor(v, "cn-beijing")
		if res.URL != "" {
			t.Fatalf("service %q: unexpected url %q", v, res.URL)
		}
		if want := ValidateService(v); err == nil || err.Error() != want.Error() {
			t.Fatalf("service %q: expect %v, got %v", v, want, err)
		}
	}
}

func TestStandardEndpointResolverValidInputsUnchanged(t *testing.T) {
	r := NewStandardEndpointResolver()
	cases := []struct {
		service, region string
		opts            []func(*Options)
		want            string
	}{
		{"ecs", "cn-beijing", nil, "ecs.cn-beijing.byteplusapi.com.cn"},
		{"ecs", "ap-southeast-1", nil, "ecs.ap-southeast-1.byteplusapi.com"},
		{"ecs", "cn-hongkong", nil, "ecs.cn-hongkong.byteplusapi.com"},
		{"rds_mysql", "ap-southeast-1", nil, "rds-mysql.ap-southeast-1.byteplusapi.com"},
		{"Redis", "ap-southeast-1", nil, "redis.ap-southeast-1.byteplusapi.com"},
		{"iam", "cn-beijing", nil, "iam.byteplusapi.com"},
		{"ecs", "ap-southeast-1", []func(*Options){WithIPVersion(IPVersionDualStack)}, "ecs.ap-southeast-1.byteplus-api.com"},
		{"billing", "ap-southeast-1", []func(*Options){WithIPVersion(IPVersionDualStack)}, "billing.byteplus-api.com"},
	}
	for _, c := range cases {
		res, err := r.EndpointFor(c.service, c.region, c.opts...)
		if err != nil || res.URL != c.want {
			t.Fatalf("%s/%s: want %q, got %q, %v", c.service, c.region, c.want, res.URL, err)
		}
	}
}

func newTestPartition() partition {
	return partition{
		ID:          "byteplus",
		DNSSuffix:   "byteplusapi.com",
		RegionRegex: regionRegex{regexp.MustCompile(`^.*$`)},
		Defaults: endpoint{
			Hostname:  "{service}.{region}.{dnsSuffix}",
			Protocols: []string{"https"},
		},
		Services: services{
			"ecs": service{},
		},
	}
}

func TestPartitionEndpointForRejectsInvalidServiceAndRegion(t *testing.T) {
	p := newTestPartition()
	for _, v := range append(append([]string{}, injectionInputs...), boundaryInputs...) {
		res, err := p.EndpointFor("ecs", v)
		if want := ValidateRegion(v); err == nil || err.Error() != want.Error() || res.URL != "" {
			t.Fatalf("region %q: expect %v, got %q, %v", v, want, res.URL, err)
		}

		// Unknown services are resolved from their name when
		// ResolveUnknownService is set, they must be rejected too.
		res, err = p.EndpointFor(v, "cn-beijing", ResolveUnknownServiceOption)
		if want := ValidateService(v); err == nil || err.Error() != want.Error() || res.URL != "" {
			t.Fatalf("service %q: expect %v, got %q, %v", v, want, res.URL, err)
		}

		// The public wrappers go through partition.EndpointFor.
		_, err = partitions{p}.EndpointFor(v, "cn-beijing", ResolveUnknownServiceOption)
		if want := ValidateService(v); err == nil || err.Error() != want.Error() {
			t.Fatalf("partitions service %q: expect %v, got %v", v, want, err)
		}
		_, err = p.Partition().EndpointFor("ecs", v)
		if want := ValidateRegion(v); err == nil || err.Error() != want.Error() {
			t.Fatalf("Partition region %q: expect %v, got %v", v, want, err)
		}
	}
}

func TestPartitionEndpointForValidInputsUnchanged(t *testing.T) {
	p := newTestPartition()
	res, err := p.EndpointFor("ecs", "cn-beijing")
	if err != nil || res.URL != "https://ecs.cn-beijing.byteplusapi.com" {
		t.Fatalf("unexpected %q, %v", res.URL, err)
	}
	res, err = p.EndpointFor("vpc", "ap-southeast-1", ResolveUnknownServiceOption)
	if err != nil || res.URL != "https://vpc.ap-southeast-1.byteplusapi.com" {
		t.Fatalf("unexpected %q, %v", res.URL, err)
	}
	if _, err = p.EndpointFor("vpc", "ap-southeast-1"); err == nil {
		t.Fatalf("unknown service without ResolveUnknownService must still fail")
	}
}

// TestBuiltinServicesAndRegionsAreValid makes sure the new rule does not
// reject any service or region shipped with the SDK, and that they still
// resolve.
func TestBuiltinServicesAndRegionsAreValid(t *testing.T) {
	if len(ServiceInfos) == 0 || len(regionMatcher.WhiteRegions) == 0 {
		t.Fatalf("built-in lists must not be empty")
	}
	regionList := make([]string, 0, len(regionMatcher.WhiteRegions))
	for region := range regionMatcher.WhiteRegions {
		if err := ValidateRegion(region); err != nil {
			t.Fatalf("built-in region %q rejected: %v", region, err)
		}
		regionList = append(regionList, region)
	}
	for region := range cnNonMainLandRegion {
		if err := ValidateRegion(region); err != nil {
			t.Fatalf("built-in region %q rejected: %v", region, err)
		}
	}
	r := NewStandardEndpointResolver()
	for name, info := range ServiceInfos {
		if err := ValidateService(name); err != nil {
			t.Fatalf("built-in service %q rejected: %v", name, err)
		}
		if err := ValidateService(info.Service); err != nil {
			t.Fatalf("built-in service %q rejected: %v", info.Service, err)
		}
		for _, region := range regionList {
			if _, err := r.EndpointFor(name, region); err != nil {
				t.Fatalf("built-in %s/%s: unexpected error %v", name, region, err)
			}
		}
	}
	if ErrCodeInvalidRegion != "InvalidRegion" || ErrCodeInvalidService != "InvalidService" {
		t.Fatalf("unexpected error codes %q %q", ErrCodeInvalidRegion, ErrCodeInvalidService)
	}
}
