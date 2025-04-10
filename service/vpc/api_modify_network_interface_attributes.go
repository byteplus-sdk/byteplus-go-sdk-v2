// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyNetworkInterfaceAttributesCommon = "ModifyNetworkInterfaceAttributes"

// ModifyNetworkInterfaceAttributesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyNetworkInterfaceAttributesCommon operation. The "output" return
// value will be populated with the ModifyNetworkInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyNetworkInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyNetworkInterfaceAttributesCommon Send returns without error.
//
// See ModifyNetworkInterfaceAttributesCommon for more information on using the ModifyNetworkInterfaceAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyNetworkInterfaceAttributesCommonRequest method.
//    req, resp := client.ModifyNetworkInterfaceAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyNetworkInterfaceAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyNetworkInterfaceAttributesCommon,
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

// ModifyNetworkInterfaceAttributesCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyNetworkInterfaceAttributesCommon for usage and error information.
func (c *VPC) ModifyNetworkInterfaceAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyNetworkInterfaceAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyNetworkInterfaceAttributesCommonWithContext is the same as ModifyNetworkInterfaceAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyNetworkInterfaceAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyNetworkInterfaceAttributesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyNetworkInterfaceAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyNetworkInterfaceAttributes = "ModifyNetworkInterfaceAttributes"

// ModifyNetworkInterfaceAttributesRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyNetworkInterfaceAttributes operation. The "output" return
// value will be populated with the ModifyNetworkInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyNetworkInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyNetworkInterfaceAttributesCommon Send returns without error.
//
// See ModifyNetworkInterfaceAttributes for more information on using the ModifyNetworkInterfaceAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyNetworkInterfaceAttributesRequest method.
//    req, resp := client.ModifyNetworkInterfaceAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyNetworkInterfaceAttributesRequest(input *ModifyNetworkInterfaceAttributesInput) (req *request.Request, output *ModifyNetworkInterfaceAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyNetworkInterfaceAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyNetworkInterfaceAttributesInput{}
	}

	output = &ModifyNetworkInterfaceAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyNetworkInterfaceAttributes API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyNetworkInterfaceAttributes for usage and error information.
func (c *VPC) ModifyNetworkInterfaceAttributes(input *ModifyNetworkInterfaceAttributesInput) (*ModifyNetworkInterfaceAttributesOutput, error) {
	req, out := c.ModifyNetworkInterfaceAttributesRequest(input)
	return out, req.Send()
}

// ModifyNetworkInterfaceAttributesWithContext is the same as ModifyNetworkInterfaceAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyNetworkInterfaceAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyNetworkInterfaceAttributesWithContext(ctx byteplus.Context, input *ModifyNetworkInterfaceAttributesInput, opts ...request.Option) (*ModifyNetworkInterfaceAttributesOutput, error) {
	req, out := c.ModifyNetworkInterfaceAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyNetworkInterfaceAttributesInput struct {
	_ struct{} `type:"structure"`

	DeleteOnTermination *bool `type:"boolean"`

	Description *string `min:"1" max:"255" type:"string"`

	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`

	NetworkInterfaceName *string `min:"1" max:"128" type:"string"`

	PortSecurityEnabled *bool `type:"boolean"`

	SecurityGroupIds []*string `type:"list"`
}

// String returns the string representation
func (s ModifyNetworkInterfaceAttributesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyNetworkInterfaceAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyNetworkInterfaceAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyNetworkInterfaceAttributesInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.NetworkInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkInterfaceId"))
	}
	if s.NetworkInterfaceName != nil && len(*s.NetworkInterfaceName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("NetworkInterfaceName", 1))
	}
	if s.NetworkInterfaceName != nil && len(*s.NetworkInterfaceName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("NetworkInterfaceName", 128, *s.NetworkInterfaceName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDeleteOnTermination sets the DeleteOnTermination field's value.
func (s *ModifyNetworkInterfaceAttributesInput) SetDeleteOnTermination(v bool) *ModifyNetworkInterfaceAttributesInput {
	s.DeleteOnTermination = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyNetworkInterfaceAttributesInput) SetDescription(v string) *ModifyNetworkInterfaceAttributesInput {
	s.Description = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *ModifyNetworkInterfaceAttributesInput) SetNetworkInterfaceId(v string) *ModifyNetworkInterfaceAttributesInput {
	s.NetworkInterfaceId = &v
	return s
}

// SetNetworkInterfaceName sets the NetworkInterfaceName field's value.
func (s *ModifyNetworkInterfaceAttributesInput) SetNetworkInterfaceName(v string) *ModifyNetworkInterfaceAttributesInput {
	s.NetworkInterfaceName = &v
	return s
}

// SetPortSecurityEnabled sets the PortSecurityEnabled field's value.
func (s *ModifyNetworkInterfaceAttributesInput) SetPortSecurityEnabled(v bool) *ModifyNetworkInterfaceAttributesInput {
	s.PortSecurityEnabled = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *ModifyNetworkInterfaceAttributesInput) SetSecurityGroupIds(v []*string) *ModifyNetworkInterfaceAttributesInput {
	s.SecurityGroupIds = v
	return s
}

type ModifyNetworkInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyNetworkInterfaceAttributesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyNetworkInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyNetworkInterfaceAttributesOutput) SetRequestId(v string) *ModifyNetworkInterfaceAttributesOutput {
	s.RequestId = &v
	return s
}
