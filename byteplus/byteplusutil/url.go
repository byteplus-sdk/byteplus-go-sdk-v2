package byteplusutil

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

var endpoint = "open.ap-southeast-1.byteplusapi.com"

func (c *Endpoint) GetEndpoint() string {
	if c.CustomerEndpoint != "" {
		return c.CustomerEndpoint
	} else {
		return endpoint
	}
}

type RegionEndpointMap map[string]string

type ServiceEndpointInfo struct {
	Service         string
	IsGlobal        bool
	GlobalEndpoint  string
	DefaultEndpoint string
	RegionEndpointMap
}

var defaultEndpoint = map[string]*ServiceEndpointInfo{
	"billing": {
		Service:           "billing",
		IsGlobal:          true,
		GlobalEndpoint:    "",
		DefaultEndpoint:   "open.byteplusapi.com",
		RegionEndpointMap: nil,
	},
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
//
// Returns:
// - *string: A pointer to the endpoint string. It could be a global endpoint, a region-specific
// endpoint, or a default endpoint if the specified service or region does not have a defined endpoint.
//
// Example:
//
//	endpoint := GetDefaultEndpointByServiceInfo("exampleService", "cn-beijing")
//
// Note: Ensure the `defaultEndpoint` map is properly populated with service and region endpoint
// information before calling this function.
func GetDefaultEndpointByServiceInfo(service string, regionCode string) *string {
	defaultEndpointInfo, sExist := defaultEndpoint[service]
	if !sExist {
		return &endpoint
	}

	isGlobal := defaultEndpointInfo.IsGlobal
	if isGlobal {
		if len(defaultEndpointInfo.GlobalEndpoint) > 0 {
			return &defaultEndpointInfo.GlobalEndpoint
		}
	} else {
		regionEndpointMp := defaultEndpointInfo.RegionEndpointMap
		regionEndpointStr, rExist := regionEndpointMp[regionCode]
		if rExist {
			return &regionEndpointStr
		}
	}

	if len(defaultEndpointInfo.DefaultEndpoint) > 0 {
		return &defaultEndpointInfo.DefaultEndpoint
	}
	return &endpoint
}
