// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opGetScheduledInstanceLatestReleaseAtCommon = "GetScheduledInstanceLatestReleaseAt"

// GetScheduledInstanceLatestReleaseAtCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the GetScheduledInstanceLatestReleaseAtCommon operation. The "output" return
// value will be populated with the GetScheduledInstanceLatestReleaseAtCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetScheduledInstanceLatestReleaseAtCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetScheduledInstanceLatestReleaseAtCommon Send returns without error.
//
// See GetScheduledInstanceLatestReleaseAtCommon for more information on using the GetScheduledInstanceLatestReleaseAtCommon
// API call, and error handling.
//
//    // Example sending a request using the GetScheduledInstanceLatestReleaseAtCommonRequest method.
//    req, resp := client.GetScheduledInstanceLatestReleaseAtCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) GetScheduledInstanceLatestReleaseAtCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetScheduledInstanceLatestReleaseAtCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetScheduledInstanceLatestReleaseAtCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation GetScheduledInstanceLatestReleaseAtCommon for usage and error information.
func (c *ECS) GetScheduledInstanceLatestReleaseAtCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetScheduledInstanceLatestReleaseAtCommonRequest(input)
	return out, req.Send()
}

// GetScheduledInstanceLatestReleaseAtCommonWithContext is the same as GetScheduledInstanceLatestReleaseAtCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetScheduledInstanceLatestReleaseAtCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) GetScheduledInstanceLatestReleaseAtCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetScheduledInstanceLatestReleaseAtCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetScheduledInstanceLatestReleaseAt = "GetScheduledInstanceLatestReleaseAt"

// GetScheduledInstanceLatestReleaseAtRequest generates a "byteplus/request.Request" representing the
// client's request for the GetScheduledInstanceLatestReleaseAt operation. The "output" return
// value will be populated with the GetScheduledInstanceLatestReleaseAtCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetScheduledInstanceLatestReleaseAtCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetScheduledInstanceLatestReleaseAtCommon Send returns without error.
//
// See GetScheduledInstanceLatestReleaseAt for more information on using the GetScheduledInstanceLatestReleaseAt
// API call, and error handling.
//
//    // Example sending a request using the GetScheduledInstanceLatestReleaseAtRequest method.
//    req, resp := client.GetScheduledInstanceLatestReleaseAtRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) GetScheduledInstanceLatestReleaseAtRequest(input *GetScheduledInstanceLatestReleaseAtInput) (req *request.Request, output *GetScheduledInstanceLatestReleaseAtOutput) {
	op := &request.Operation{
		Name:       opGetScheduledInstanceLatestReleaseAt,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetScheduledInstanceLatestReleaseAtInput{}
	}

	output = &GetScheduledInstanceLatestReleaseAtOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetScheduledInstanceLatestReleaseAt API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation GetScheduledInstanceLatestReleaseAt for usage and error information.
func (c *ECS) GetScheduledInstanceLatestReleaseAt(input *GetScheduledInstanceLatestReleaseAtInput) (*GetScheduledInstanceLatestReleaseAtOutput, error) {
	req, out := c.GetScheduledInstanceLatestReleaseAtRequest(input)
	return out, req.Send()
}

// GetScheduledInstanceLatestReleaseAtWithContext is the same as GetScheduledInstanceLatestReleaseAt with the addition of
// the ability to pass a context and additional request options.
//
// See GetScheduledInstanceLatestReleaseAt for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) GetScheduledInstanceLatestReleaseAtWithContext(ctx byteplus.Context, input *GetScheduledInstanceLatestReleaseAtInput, opts ...request.Option) (*GetScheduledInstanceLatestReleaseAtOutput, error) {
	req, out := c.GetScheduledInstanceLatestReleaseAtRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetScheduledInstanceLatestReleaseAtInput struct {
	_ struct{} `type:"structure"`

	DeliveryType *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	StartDeliveryAt *string `type:"string"`

	VolumeType *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s GetScheduledInstanceLatestReleaseAtInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s GetScheduledInstanceLatestReleaseAtInput) GoString() string {
	return s.String()
}

// SetDeliveryType sets the DeliveryType field's value.
func (s *GetScheduledInstanceLatestReleaseAtInput) SetDeliveryType(v string) *GetScheduledInstanceLatestReleaseAtInput {
	s.DeliveryType = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *GetScheduledInstanceLatestReleaseAtInput) SetInstanceTypeId(v string) *GetScheduledInstanceLatestReleaseAtInput {
	s.InstanceTypeId = &v
	return s
}

// SetStartDeliveryAt sets the StartDeliveryAt field's value.
func (s *GetScheduledInstanceLatestReleaseAtInput) SetStartDeliveryAt(v string) *GetScheduledInstanceLatestReleaseAtInput {
	s.StartDeliveryAt = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *GetScheduledInstanceLatestReleaseAtInput) SetVolumeType(v string) *GetScheduledInstanceLatestReleaseAtInput {
	s.VolumeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *GetScheduledInstanceLatestReleaseAtInput) SetZoneId(v string) *GetScheduledInstanceLatestReleaseAtInput {
	s.ZoneId = &v
	return s
}

type GetScheduledInstanceLatestReleaseAtOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LatestReleaseAt *string `type:"string"`
}

// String returns the string representation
func (s GetScheduledInstanceLatestReleaseAtOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s GetScheduledInstanceLatestReleaseAtOutput) GoString() string {
	return s.String()
}

// SetLatestReleaseAt sets the LatestReleaseAt field's value.
func (s *GetScheduledInstanceLatestReleaseAtOutput) SetLatestReleaseAt(v string) *GetScheduledInstanceLatestReleaseAtOutput {
	s.LatestReleaseAt = &v
	return s
}
