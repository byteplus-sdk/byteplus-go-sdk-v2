package byteplusutil

import (
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
	regionCodeAPSouthEast3 = "ap-southeast-3"
	regionCodeCNHongkong   = "cn-hongkong"
)

var cnNonMainLandRegion = map[string]struct{}{
	regionCodeCNHongkong: {},
}

type RegionEndpointMap map[string]string

type ServiceEndpointInfo struct {
	Service        string
	IsGlobal       bool
	Prefix         string
	GlobalEndpoint string
	RegionEndpointMap
}

var defaultEndpoint = map[string]*ServiceEndpointInfo{
	"billing": {
		Service:           "billing",
		IsGlobal:          true,
		GlobalEndpoint:    "",
		Prefix:            "",
		RegionEndpointMap: nil,
	},
	"iam": {
		Service:           "iam",
		IsGlobal:          true,
		GlobalEndpoint:    "",
		Prefix:            "",
		RegionEndpointMap: nil,
	},
	"vpc": {
		Service:        "vpc",
		IsGlobal:       false,
		GlobalEndpoint: "",
		Prefix:         apSoutheastPrefix,
		RegionEndpointMap: RegionEndpointMap{
			regionCodeAPSouthEast3: "vpc" + separator + regionCodeAPSouthEast3 + byteplusEndpointSuffix,
		},
	},
	"ecs": {
		Service:        "ecs",
		IsGlobal:       false,
		GlobalEndpoint: "",
		Prefix:         apSoutheastPrefix,
		RegionEndpointMap: RegionEndpointMap{
			regionCodeAPSouthEast3: "ecs" + separator + regionCodeAPSouthEast3 + byteplusEndpointSuffix,
		},
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
func GetDefaultEndpointByServiceInfo(service string, regionCode string, customBootstrapRegion map[string]struct{}) *string {
	resultEndpoint := endpoint

	if !inBootstrapRegionList(regionCode, customBootstrapRegion) {
		return &resultEndpoint
	}

	defaultEndpointInfo, sExist := defaultEndpoint[service]
	if !sExist {
		return &resultEndpoint
	}

	suffix := byteplusEndpointSuffix
	if hasEnableDualStack() {
		suffix = dualstackEndpointSuffix
	}
	suffix = defaultEndpointInfo.Prefix + suffix

	if defaultEndpointInfo.IsGlobal {
		if len(defaultEndpointInfo.GlobalEndpoint) > 0 {
			resultEndpoint = defaultEndpointInfo.GlobalEndpoint
			return &resultEndpoint
		}
		resultEndpoint = standardizeDomainServiceCode(service) + suffix
		if isCNRegion(regionCode) {
			resultEndpoint += cnSuffix
		}
		return &resultEndpoint
	}

	// regional endpoint
	regionEndpointMp := defaultEndpointInfo.RegionEndpointMap
	regionEndpointStr, rExist := regionEndpointMp[regionCode]
	if rExist {
		resultEndpoint = regionEndpointStr
		return &resultEndpoint
	}

	resultEndpoint = standardizeDomainServiceCode(service) + separator + regionCode + suffix
	if isCNRegion(regionCode) {
		resultEndpoint += cnSuffix
	}
	return &resultEndpoint

}

var bootstrapRegion = map[string]struct{}{
	regionCodeAPSouthEast3: {},
}

func inBootstrapRegionList(regionCode string, customBootstrapRegion map[string]struct{}) bool {
	regionCode = strings.TrimSpace(regionCode)
	bsRegionListPath := os.Getenv("BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF")
	if len(bsRegionListPath) > 0 {
		f, err := os.ReadFile(filepath.Clean(bsRegionListPath))
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

func hasEnableDualStack() bool {
	return os.Getenv("BYTEPLUS_ENABLE_DUALSTACK") == "true"
}

func isCNRegion(region string) bool {
	hasCNPrefix := strings.HasPrefix(region, "cn-")
	if !hasCNPrefix {
		return false
	}

	_, isCNNonMainLangRegion := cnNonMainLandRegion[region]
	return !isCNNonMainLangRegion
}
