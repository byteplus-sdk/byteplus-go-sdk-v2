package client

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/client/metadata"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

// A Config provides configuration to a service client instance.
type Config struct {
	Config        *byteplus.Config
	Handlers      request.Handlers
	Endpoint      string
	SigningRegion string
	SigningName   string

	// States that the signing name did not come from a modeled source but
	// was derived based on other data. Used by service client constructors
	// to determine if the signin name can be overridden based on metadata the
	// service has.
	SigningNameDerived bool
}

// ConfigProvider provides a generic way for a service client to receive
// the ClientConfig without circular dependencies.
type ConfigProvider interface {
	ClientConfig(serviceName string, cfgs ...*byteplus.Config) Config
}

// ConfigNoResolveEndpointProvider same as ConfigProvider except it will not
// resolve the endpoint automatically. The service client's endpoint must be
// provided via the byteplus.Config.Endpoint field.
type ConfigNoResolveEndpointProvider interface {
	ClientConfigNoResolveEndpoint(cfgs ...*byteplus.Config) Config
}

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	request.Retryer
	metadata.ClientInfo

	Config   byteplus.Config
	Handlers request.Handlers
}

// New will return a pointer to a new initialized service client.
func New(cfg byteplus.Config, info metadata.ClientInfo, handlers request.Handlers, options ...func(*Client)) *Client {
	svc := &Client{
		Config:     cfg,
		ClientInfo: info,
		Handlers:   handlers.Copy(),
	}

	switch retryer, ok := cfg.Retryer.(request.Retryer); {
	case ok:
		svc.Retryer = retryer
	case cfg.Retryer != nil && cfg.Logger != nil:
		s := fmt.Sprintf("WARNING: %T does not implement request.Retryer; using DefaultRetryer instead", cfg.Retryer)
		cfg.Logger.Log(s)
		fallthrough
	default:
		maxRetries := byteplus.IntValue(cfg.MaxRetries)
		if cfg.MaxRetries == nil || maxRetries == byteplus.UseServiceDefaultRetries {
			maxRetries = DefaultRetryerMaxNumRetries
		}
		svc.Retryer = DefaultRetryer{NumMaxRetries: maxRetries}
	}

	svc.AddDebugHandlers()

	for _, option := range options {
		option(svc)
	}

	return svc
}

// NewRequest returns a new Request pointer for the service API
// operation and parameters.
func (c *Client) NewRequest(operation *request.Operation, params interface{}, data interface{}) *request.Request {
	return request.New(c.Config, c.ClientInfo, c.Handlers, c.Retryer, operation, params, data)
}

// AddDebugHandlers injects debug logging handlers into the service to log request
// debug information.
func (c *Client) AddDebugHandlers() {
	if !c.Config.LogLevel.AtLeast(byteplus.LogDebug) {
		return
	}
	if c.Config.LogLevel.Matches(byteplus.LogInfoWithInputAndOutput) ||
		c.Config.LogLevel.Matches(byteplus.LogDebugWithInputAndOutput) {
		c.Handlers.Send.PushFrontNamed(LogInputHandler)
		c.Handlers.Complete.PushBackNamed(LogOutHandler)
		return
	}

	c.Handlers.Send.PushFrontNamed(LogHTTPRequestHandler)
	c.Handlers.Send.PushBackNamed(LogHTTPResponseHandler)

}
