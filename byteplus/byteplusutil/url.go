package byteplusutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
)

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

type Endpoint struct {
	CustomerEndpoint string
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (c *Endpoint) WithCustomerEndpoint(customerEndpoint string) *Endpoint {
	c.CustomerEndpoint = customerEndpoint
	return c
}

type ServiceInfo struct {
	Service string
	Region  string
}

func (c *Endpoint) GetEndpoint() string {
	if c.CustomerEndpoint != "" {
		return c.CustomerEndpoint
	} else {
		return endpoint
	}
}

const (
	separator               = "."
	openPrefix              = "open"
	byteplusEndpointSuffix  = separator + "byteplusapi.com"
	apSoutheastPrefix       = separator + "ap-southeast-1"
	endpointSuffix          = apSoutheastPrefix + ".byteplusapi.com"
	dualstackEndpointSuffix = separator + "byteplus-api.com"
	cnSuffix                = separator + "cn"
)

var endpoint = openPrefix + endpointSuffix

const (
	regionCodeAPSouthEast2 = "ap-southeast-2"
	regionCodeAPSouthEast3 = "ap-southeast-3"
	regionCodeCNHongkong   = "cn-hongkong"
)

var cnNonMainLandRegion = map[string]struct{}{
	regionCodeCNHongkong: {},
}

// ErrServiceEndpointInfoMissing indicates that the requested service is not
// registered in defaultEndpoint. Surfaced by ResolveDefaultEndpoint and
// wrapped by the Session default addressing pipeline.
var ErrServiceEndpointInfoMissing = errors.New("byteplusutil: service endpoint info missing")

// ErrMissingRegion is returned by ResolveDefaultEndpoint when a Regional
// service is asked to resolve an endpoint without a region. Without this
// guard the SDK used to produce a double-dot host such as
// "vpc..byteplusapi.com" and return err=nil, which callers could not tell
// from a legitimate result.
var ErrMissingRegion = errors.New("byteplusutil: missing region for regional service")

type ServiceEndpointInfo struct {
	Service        string
	IsGlobal       bool
	GoChinaEnabled bool
}

var defaultEndpoint = map[string]*ServiceEndpointInfo{
	"vpc":            {Service: "vpc", IsGlobal: false, GoChinaEnabled: true},
	"vke":            {Service: "vke", IsGlobal: false, GoChinaEnabled: true},
	"storage_ebs":    {Service: "storage_ebs", IsGlobal: false, GoChinaEnabled: true},
	"vedbm":          {Service: "vedbm", IsGlobal: false, GoChinaEnabled: true},
	"privatelink":    {Service: "privatelink", IsGlobal: false, GoChinaEnabled: true},
	"clb":            {Service: "clb", IsGlobal: false, GoChinaEnabled: true},
	"transitrouter":  {Service: "transitrouter", IsGlobal: false, GoChinaEnabled: true},
	"directconnect":  {Service: "directconnect", IsGlobal: false, GoChinaEnabled: true},
	"vpn":            {Service: "vpn", IsGlobal: false, GoChinaEnabled: true},
	"natgateway":     {Service: "natgateway", IsGlobal: false, GoChinaEnabled: true},
	"rds_mysql":      {Service: "rds_mysql", IsGlobal: false, GoChinaEnabled: true},
	"smc":            {Service: "smc", IsGlobal: true, GoChinaEnabled: false},
	"iam":            {Service: "iam", IsGlobal: true, GoChinaEnabled: true},
	"vepfs":          {Service: "vepfs", IsGlobal: false, GoChinaEnabled: true},
	"kms":            {Service: "kms", IsGlobal: false, GoChinaEnabled: true},
	"ecs":            {Service: "ecs", IsGlobal: false, GoChinaEnabled: true},
	"mongodb":        {Service: "mongodb", IsGlobal: false, GoChinaEnabled: true},
	"private_zone":   {Service: "private_zone", IsGlobal: true, GoChinaEnabled: true},
	"rds_postgresql": {Service: "rds_postgresql", IsGlobal: false, GoChinaEnabled: true},
	"resource_share": {Service: "resource_share", IsGlobal: true, GoChinaEnabled: false},
	"vmp":            {Service: "vmp", IsGlobal: false, GoChinaEnabled: true},
	"tag":            {Service: "tag", IsGlobal: true, GoChinaEnabled: false},
	"cr":             {Service: "cr", IsGlobal: false, GoChinaEnabled: true},
	"alb":            {Service: "alb", IsGlobal: false, GoChinaEnabled: true},
	"sts":            {Service: "sts", IsGlobal: false, GoChinaEnabled: true},
	"hbase":          {Service: "hbase", IsGlobal: false, GoChinaEnabled: true},
	"rds_mssql":      {Service: "rds_mssql", IsGlobal: false, GoChinaEnabled: true},
	"ml_platform":    {Service: "ml_platform", IsGlobal: false, GoChinaEnabled: false},
	"apig":           {Service: "apig", IsGlobal: false, GoChinaEnabled: false},
	"ark":            {Service: "ark", IsGlobal: false, GoChinaEnabled: false},
	"waf":            {Service: "waf", IsGlobal: true, GoChinaEnabled: false},
	"quota":          {Service: "quota", IsGlobal: true, GoChinaEnabled: false},
	"dms":            {Service: "dms", IsGlobal: false, GoChinaEnabled: true},
	"vefaas":         {Service: "vefaas", IsGlobal: false, GoChinaEnabled: false},
	"cen":            {Service: "cen", IsGlobal: true, GoChinaEnabled: false},
	"cp":             {Service: "cp", IsGlobal: false, GoChinaEnabled: false},
	"cloudmonitor":   {Service: "cloudmonitor", IsGlobal: false, GoChinaEnabled: true},
	"eco_partner":    {Service: "eco_partner", IsGlobal: true, GoChinaEnabled: false},
	"milvus":         {Service: "milvus", IsGlobal: false, GoChinaEnabled: false},
	"llmshield":      {Service: "llmshield", IsGlobal: false, GoChinaEnabled: false},
	"billing":        {Service: "billing", IsGlobal: true, GoChinaEnabled: true},
	"id":             {Service: "id", IsGlobal: false, GoChinaEnabled: false},
	"clawsentry":     {Service: "clawsentry", IsGlobal: false, GoChinaEnabled: false},
	"resourcecenter": {Service: "resourcecenter", IsGlobal: true, GoChinaEnabled: false},
	"escloud":        {Service: "escloud", IsGlobal: false, GoChinaEnabled: false},
	"cpaas":          {Service: "cpaas", IsGlobal: true, GoChinaEnabled: false},
	"filenas":        {Service: "filenas", IsGlobal: false, GoChinaEnabled: true},
	"kafka":          {Service: "kafka", IsGlobal: false, GoChinaEnabled: true},
	"kickart":        {Service: "kickart", IsGlobal: true, GoChinaEnabled: false},
	"rabbitmq":       {Service: "rabbitmq", IsGlobal: false, GoChinaEnabled: false},
	"redis":          {Service: "redis", IsGlobal: false, GoChinaEnabled: true},
	"vod":            {Service: "vod", IsGlobal: false, GoChinaEnabled: false},
	"vs":             {Service: "vs", IsGlobal: true, GoChinaEnabled: false},
	"autoscaling":    {Service: "autoscaling", IsGlobal: false, GoChinaEnabled: true},
	"aidap":          {Service: "aidap", IsGlobal: false, GoChinaEnabled: false},
}

func standardizeDomainServiceCode(serviceCode string) string {
	return strings.ReplaceAll(strings.ToLower(serviceCode), "_", "-")
}

// GetDefaultEndpointByServiceInfo is the backwards-compatible entry point that
// preserves the original signature. Internally it delegates to
// ResolveDefaultEndpoint and, on any resolution error, falls back to the shared
// open endpoint so external callers that dereference the returned pointer keep
// working.
//
// Deprecated: the customBootstrapRegion parameter is retained for API
// compatibility only and is ignored at runtime; bootstrap-region-based routing
// no longer participates in default addressing. New SDK-internal call sites
// should prefer ResolveDefaultEndpoint so that a missing service registration
// surfaces as an error instead of being silently swallowed.
func GetDefaultEndpointByServiceInfo(service string, regionCode string,
	customBootstrapRegion map[string]struct{}, useDualStack *bool, logger byteplus.Logger) *string {
	resolved, err := ResolveDefaultEndpoint(service, regionCode, useDualStack, logger)
	if err != nil {
		debug(logger, fmt.Sprintf("resolve default endpoint failed for service %s: %v, fall back to %s",
			service, err, endpoint))
		fallback := endpoint
		return &fallback
	}
	return &resolved
}

// ResolveDefaultEndpoint is the SDK-internal entry point used by the default
// Session request pipeline when neither Config.Endpoint nor
// Config.EndpointResolver is provided.
//
// It implements the addressing decision described in the Endpoint auto-address
// spec: consult the defaultEndpoint map, then combine RegionType (IsGlobal)
// with GoChinaEnabled to produce a deterministic
// {service}[.{region}].byteplusapi.com[.cn] endpoint.
//
// Returns ErrServiceEndpointInfoMissing when the service is not registered in
// defaultEndpoint.
func ResolveDefaultEndpoint(service, regionCode string, useDualStack *bool, logger byteplus.Logger) (string, error) {
	debug(logger, fmt.Sprintf("ResolveDefaultEndpoint start, service: %s, region: %s", service, regionCode))

	info, ok := defaultEndpoint[service]
	if !ok {
		return "", fmt.Errorf("%w: service %q not registered", ErrServiceEndpointInfoMissing, service)
	}

	normalizedRegion := strings.ToLower(strings.TrimSpace(regionCode))

	baseSuffix := byteplusEndpointSuffix
	if hasEnableDualStack(useDualStack) {
		debug(logger, fmt.Sprintf("dual stack enabled, use suffix: %s", dualstackEndpointSuffix))
		baseSuffix = dualstackEndpointSuffix
	}

	useCN := info.GoChinaEnabled && isCNRegion(normalizedRegion)
	suffix := baseSuffix
	if useCN {
		suffix = baseSuffix + cnSuffix
	}

	svc := standardizeDomainServiceCode(service)
	var resultEndpoint string
	if info.IsGlobal {
		resultEndpoint = svc + suffix
	} else {
		if normalizedRegion == "" {
			return "", fmt.Errorf("%w: service %q requires a non-empty region",
				ErrMissingRegion, service)
		}
		resultEndpoint = svc + separator + normalizedRegion + suffix
	}
	debug(logger, fmt.Sprintf("resolved endpoint for service %s region %s: %s",
		service, normalizedRegion, resultEndpoint))
	return resultEndpoint, nil
}

var bootstrapRegion = map[string]struct{}{
	regionCodeAPSouthEast2: {},
	regionCodeAPSouthEast3: {},
}

// inBootstrapRegionList is retained for external callers that may still rely on
// it. The default addressing pipeline no longer consults it.
func inBootstrapRegionList(regionCode string, customBootstrapRegion map[string]struct{}) bool {
	regionCode = strings.TrimSpace(regionCode)
	bsRegionListPath := os.Getenv("BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF")
	if len(bsRegionListPath) > 0 {
		f, err := ioutil.ReadFile(filepath.Clean(bsRegionListPath))
		if err == nil {
			for _, l := range strings.Split(string(f), "\n") {
				l = strings.TrimSpace(l)
				if len(l) == 0 {
					continue
				}
				if l == regionCode {
					return true
				}
			}
		}
	}

	if len(bootstrapRegion) > 0 {
		_, ok := bootstrapRegion[regionCode]
		if ok {
			return ok
		}
	}

	if len(customBootstrapRegion) > 0 {
		_, ok := customBootstrapRegion[regionCode]
		return ok
	}

	return false
}

func hasEnableDualStack(useDualStack *bool) bool {
	if useDualStack == nil {
		return os.Getenv("BYTEPLUS_ENABLE_DUALSTACK") == "true"
	}
	return *useDualStack
}

func isCNRegion(region string) bool {
	hasCNPrefix := strings.HasPrefix(region, "cn-")
	if !hasCNPrefix {
		return false
	}

	_, isCNNonMainLangRegion := cnNonMainLandRegion[region]
	return !isCNNonMainLangRegion
}

func debug(logger byteplus.Logger, args ...interface{}) {
	if logger != nil {
		logger.DebugByLevel(byteplus.LogDebugWithEndpoint, append([]interface{}{"[Endpoint]"}, args...)...)
	}
}
