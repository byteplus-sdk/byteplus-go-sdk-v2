// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyIpAddressPoolAttributesCommon = "ModifyIpAddressPoolAttributes"

// ModifyIpAddressPoolAttributesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyIpAddressPoolAttributesCommon operation. The "output" return
// value will be populated with the ModifyIpAddressPoolAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyIpAddressPoolAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyIpAddressPoolAttributesCommon Send returns without error.
//
// See ModifyIpAddressPoolAttributesCommon for more information on using the ModifyIpAddressPoolAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyIpAddressPoolAttributesCommonRequest method.
//    req, resp := client.ModifyIpAddressPoolAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyIpAddressPoolAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyIpAddressPoolAttributesCommon,
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

// ModifyIpAddressPoolAttributesCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyIpAddressPoolAttributesCommon for usage and error information.
func (c *VPC) ModifyIpAddressPoolAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyIpAddressPoolAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyIpAddressPoolAttributesCommonWithContext is the same as ModifyIpAddressPoolAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyIpAddressPoolAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyIpAddressPoolAttributesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyIpAddressPoolAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyIpAddressPoolAttributes = "ModifyIpAddressPoolAttributes"

// ModifyIpAddressPoolAttributesRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyIpAddressPoolAttributes operation. The "output" return
// value will be populated with the ModifyIpAddressPoolAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyIpAddressPoolAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyIpAddressPoolAttributesCommon Send returns without error.
//
// See ModifyIpAddressPoolAttributes for more information on using the ModifyIpAddressPoolAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyIpAddressPoolAttributesRequest method.
//    req, resp := client.ModifyIpAddressPoolAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyIpAddressPoolAttributesRequest(input *ModifyIpAddressPoolAttributesInput) (req *request.Request, output *ModifyIpAddressPoolAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyIpAddressPoolAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyIpAddressPoolAttributesInput{}
	}

	output = &ModifyIpAddressPoolAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyIpAddressPoolAttributes API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyIpAddressPoolAttributes for usage and error information.
func (c *VPC) ModifyIpAddressPoolAttributes(input *ModifyIpAddressPoolAttributesInput) (*ModifyIpAddressPoolAttributesOutput, error) {
	req, out := c.ModifyIpAddressPoolAttributesRequest(input)
	return out, req.Send()
}

// ModifyIpAddressPoolAttributesWithContext is the same as ModifyIpAddressPoolAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyIpAddressPoolAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyIpAddressPoolAttributesWithContext(ctx byteplus.Context, input *ModifyIpAddressPoolAttributesInput, opts ...request.Option) (*ModifyIpAddressPoolAttributesOutput, error) {
	req, out := c.ModifyIpAddressPoolAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyIpAddressPoolAttributesInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	// IpAddressPoolId is a required field
	IpAddressPoolId *string `type:"string" required:"true"`

	Name *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifyIpAddressPoolAttributesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyIpAddressPoolAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyIpAddressPoolAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyIpAddressPoolAttributesInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.IpAddressPoolId == nil {
		invalidParams.Add(request.NewErrParamRequired("IpAddressPoolId"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 1))
	}
	if s.Name != nil && len(*s.Name) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("Name", 128, *s.Name))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyIpAddressPoolAttributesInput) SetClientToken(v string) *ModifyIpAddressPoolAttributesInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyIpAddressPoolAttributesInput) SetDescription(v string) *ModifyIpAddressPoolAttributesInput {
	s.Description = &v
	return s
}

// SetIpAddressPoolId sets the IpAddressPoolId field's value.
func (s *ModifyIpAddressPoolAttributesInput) SetIpAddressPoolId(v string) *ModifyIpAddressPoolAttributesInput {
	s.IpAddressPoolId = &v
	return s
}

// SetName sets the Name field's value.
func (s *ModifyIpAddressPoolAttributesInput) SetName(v string) *ModifyIpAddressPoolAttributesInput {
	s.Name = &v
	return s
}

type ModifyIpAddressPoolAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyIpAddressPoolAttributesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyIpAddressPoolAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyIpAddressPoolAttributesOutput) SetRequestId(v string) *ModifyIpAddressPoolAttributesOutput {
	s.RequestId = &v
	return s
}
