// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opAssociateHaVipCommon = "AssociateHaVip"

// AssociateHaVipCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the AssociateHaVipCommon operation. The "output" return
// value will be populated with the AssociateHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateHaVipCommon Send returns without error.
//
// See AssociateHaVipCommon for more information on using the AssociateHaVipCommon
// API call, and error handling.
//
//    // Example sending a request using the AssociateHaVipCommonRequest method.
//    req, resp := client.AssociateHaVipCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AssociateHaVipCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateHaVipCommon,
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

// AssociateHaVipCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation AssociateHaVipCommon for usage and error information.
func (c *VPC) AssociateHaVipCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateHaVipCommonRequest(input)
	return out, req.Send()
}

// AssociateHaVipCommonWithContext is the same as AssociateHaVipCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateHaVipCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateHaVipCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateHaVipCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateHaVip = "AssociateHaVip"

// AssociateHaVipRequest generates a "byteplus/request.Request" representing the
// client's request for the AssociateHaVip operation. The "output" return
// value will be populated with the AssociateHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateHaVipCommon Send returns without error.
//
// See AssociateHaVip for more information on using the AssociateHaVip
// API call, and error handling.
//
//    // Example sending a request using the AssociateHaVipRequest method.
//    req, resp := client.AssociateHaVipRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AssociateHaVipRequest(input *AssociateHaVipInput) (req *request.Request, output *AssociateHaVipOutput) {
	op := &request.Operation{
		Name:       opAssociateHaVip,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateHaVipInput{}
	}

	output = &AssociateHaVipOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateHaVip API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation AssociateHaVip for usage and error information.
func (c *VPC) AssociateHaVip(input *AssociateHaVipInput) (*AssociateHaVipOutput, error) {
	req, out := c.AssociateHaVipRequest(input)
	return out, req.Send()
}

// AssociateHaVipWithContext is the same as AssociateHaVip with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateHaVip for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateHaVipWithContext(ctx byteplus.Context, input *AssociateHaVipInput, opts ...request.Option) (*AssociateHaVipOutput, error) {
	req, out := c.AssociateHaVipRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateHaVipInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// HaVipId is a required field
	HaVipId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	InstanceType *string `type:"string" enum:"InstanceTypeForAssociateHaVipInput"`
}

// String returns the string representation
func (s AssociateHaVipInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateHaVipInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateHaVipInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssociateHaVipInput"}
	if s.HaVipId == nil {
		invalidParams.Add(request.NewErrParamRequired("HaVipId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *AssociateHaVipInput) SetClientToken(v string) *AssociateHaVipInput {
	s.ClientToken = &v
	return s
}

// SetHaVipId sets the HaVipId field's value.
func (s *AssociateHaVipInput) SetHaVipId(v string) *AssociateHaVipInput {
	s.HaVipId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *AssociateHaVipInput) SetInstanceId(v string) *AssociateHaVipInput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *AssociateHaVipInput) SetInstanceType(v string) *AssociateHaVipInput {
	s.InstanceType = &v
	return s
}

type AssociateHaVipOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AssociateHaVipOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateHaVipOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AssociateHaVipOutput) SetRequestId(v string) *AssociateHaVipOutput {
	s.RequestId = &v
	return s
}

const (
	// InstanceTypeForAssociateHaVipInputNetworkInterface is a InstanceTypeForAssociateHaVipInput enum value
	InstanceTypeForAssociateHaVipInputNetworkInterface = "NetworkInterface"

	// InstanceTypeForAssociateHaVipInputEcsInstance is a InstanceTypeForAssociateHaVipInput enum value
	InstanceTypeForAssociateHaVipInputEcsInstance = "EcsInstance"
)
