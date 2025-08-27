package byteplusutil

import (
	"fmt"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

type Endpoint struct {
	//UseSSL bool
	//Locate bool
	//UseInternal                 bool
	CustomerEndpoint string
	//CustomerDomainIgnoreService bool

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

type RegionEndpointMap map[string]string

type ServiceEndpointInfo struct {
	Service         string
	IsGlobal        bool
	DefaultEndpoint string
	GlobalEndpoint  string
	RegionEndpointMap
}

var defaultEndpoint = map[string]*ServiceEndpointInfo{
	"ark": {
		Service:         "ark",
		IsGlobal:        false,
		DefaultEndpoint: openPrefix + byteplusEndpointSuffix,
	},
	"billing": {
		Service:         "billing",
		IsGlobal:        true,
		DefaultEndpoint: openPrefix + byteplusEndpointSuffix,
	},
	"iam": {
		Service:         "iam",
		IsGlobal:        true,
		DefaultEndpoint: openPrefix + byteplusEndpointSuffix,
	},
	"vpc": {
		Service:         "vpc",
		IsGlobal:        false,
		DefaultEndpoint: endpoint,
	},
	"ecs": {
		Service:         "ecs",
		IsGlobal:        false,
		DefaultEndpoint: endpoint,
	},
	"kms": {
		Service:         "kms",
		IsGlobal:        false,
		DefaultEndpoint: endpoint,
	},
}

func standardizeDomainServiceCode(serviceCode string) string {
	return strings.ReplaceAll(strings.ToLower(serviceCode), "_", "-")
}

// GetDefaultEndpointByServiceInfo retrieves the default endpoint for a given service and region.
//
// This function takes in the service name and region code, and returns a pointer to the default
// endpoint string. It checks if the service has a global endpoint or a region-specific endpoint.
// If neither is found, it returns a pointer to the default endpoint.
//
// Parameters:
// - service: The name of the service for which to retrieve the endpoint.
// - regionCode: The region code to look up the region-specific endpoint.
// - customBootstrapRegion: The map which keys are bootstrapping region code and values are empty struct.
// - useDualStack: Whether to use the dualstack endpoint.
// Returns:
// - *string: A pointer to the endpoint string. It could be a global endpoint, a region-specific
// endpoint, or a default endpoint if the specified service or region does not have a defined endpoint.
//
// Example:
//
//	endpoint := GetDefaultEndpointByServiceInfo("exampleService", "cn-beijing", nil)
//
// Note: Ensure the `defaultEndpoint` map is properly populated with service and region endpoint
// information before calling this function.
func GetDefaultEndpointByServiceInfo(service string, regionCode string,
	customBootstrapRegion map[string]struct{}, useDualStack *bool, logger byteplus.Logger) *string {

	debug(logger, fmt.Sprintf("GetDefaultEndpointByServiceInfo start, service: %s, region: %s", service, regionCode))

	resultEndpoint := endpoint

	if !inBootstrapRegionList(regionCode, customBootstrapRegion) {
		debug(logger, fmt.Sprintf("region %s not in bootstrap region list, use default endpoint: %s", regionCode, resultEndpoint))
		defaultEndpointInfo, sExist := defaultEndpoint[service]
		if !sExist {
			debug(logger, fmt.Sprintf("service %s not found in default endpoint map, use default endpoint: %s", service, resultEndpoint))
			return &resultEndpoint
		}
		resultEndpoint = defaultEndpointInfo.DefaultEndpoint
		debug(logger, fmt.Sprintf("service %s found in default endpoint map, use default map endpoint: %s", service, resultEndpoint))
		return &resultEndpoint
	}

	suffix := byteplusEndpointSuffix
	if hasEnableDualStack(useDualStack) {
		debug(logger, fmt.Sprintf("dual stack enabled, use suffix: %s", suffix))
		suffix = dualstackEndpointSuffix
	}

	defaultEndpointInfo, sExist := defaultEndpoint[service]
	if !sExist {
		debug(logger, fmt.Sprintf("service %s not found in default endpoint map, use default endpoint: %s", service, resultEndpoint))
		return &resultEndpoint
	}

	if defaultEndpointInfo.IsGlobal {
		debug(logger, fmt.Sprintf("service %s is global", service))
		if len(defaultEndpointInfo.GlobalEndpoint) > 0 {
			resultEndpoint = defaultEndpointInfo.GlobalEndpoint
			debug(logger, fmt.Sprintf("use predefined global endpoint: %s", resultEndpoint))
			return &resultEndpoint
		}
		resultEndpoint = standardizeDomainServiceCode(service) + suffix
		debug(logger, fmt.Sprintf("use constructed global endpoint: %s", resultEndpoint))
		return &resultEndpoint
	}

	debug(logger, fmt.Sprintf("service %s is regional", service))
	// regional endpoint
	regionEndpointMp := defaultEndpointInfo.RegionEndpointMap
	if regionEndpointMp != nil {
		regionEndpointStr, rExist := regionEndpointMp[regionCode]
		if rExist {
			resultEndpoint = regionEndpointStr
			debug(logger, fmt.Sprintf("found predefined regional endpoint for region %s: %s", regionCode, resultEndpoint))
			return &resultEndpoint
		}
	}

	resultEndpoint = standardizeDomainServiceCode(service) + separator + regionCode + suffix
	if isCNRegion(regionCode) {
		resultEndpoint += cnSuffix
		debug(logger, fmt.Sprintf("region %s is cn region, add suffix: %s", regionCode, cnSuffix))
	}
	debug(logger, fmt.Sprintf("use constructed regional endpoint: %s", resultEndpoint))
	return &resultEndpoint

}

var bootstrapRegion = map[string]struct{}{
	regionCodeAPSouthEast2: {},
	regionCodeAPSouthEast3: {},
}

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
	logger.DebugByLevel(byteplus.LogDebugWithEndpoint, append([]interface{}{"[Endpoint]"}, args...)...)
}
