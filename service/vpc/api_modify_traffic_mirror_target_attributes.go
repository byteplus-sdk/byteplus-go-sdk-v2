// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyTrafficMirrorTargetAttributesCommon = "ModifyTrafficMirrorTargetAttributes"

// ModifyTrafficMirrorTargetAttributesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyTrafficMirrorTargetAttributesCommon operation. The "output" return
// value will be populated with the ModifyTrafficMirrorTargetAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTrafficMirrorTargetAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTrafficMirrorTargetAttributesCommon Send returns without error.
//
// See ModifyTrafficMirrorTargetAttributesCommon for more information on using the ModifyTrafficMirrorTargetAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTrafficMirrorTargetAttributesCommonRequest method.
//    req, resp := client.ModifyTrafficMirrorTargetAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyTrafficMirrorTargetAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTrafficMirrorTargetAttributesCommon,
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

// ModifyTrafficMirrorTargetAttributesCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyTrafficMirrorTargetAttributesCommon for usage and error information.
func (c *VPC) ModifyTrafficMirrorTargetAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTrafficMirrorTargetAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyTrafficMirrorTargetAttributesCommonWithContext is the same as ModifyTrafficMirrorTargetAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTrafficMirrorTargetAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyTrafficMirrorTargetAttributesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTrafficMirrorTargetAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTrafficMirrorTargetAttributes = "ModifyTrafficMirrorTargetAttributes"

// ModifyTrafficMirrorTargetAttributesRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyTrafficMirrorTargetAttributes operation. The "output" return
// value will be populated with the ModifyTrafficMirrorTargetAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTrafficMirrorTargetAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTrafficMirrorTargetAttributesCommon Send returns without error.
//
// See ModifyTrafficMirrorTargetAttributes for more information on using the ModifyTrafficMirrorTargetAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyTrafficMirrorTargetAttributesRequest method.
//    req, resp := client.ModifyTrafficMirrorTargetAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyTrafficMirrorTargetAttributesRequest(input *ModifyTrafficMirrorTargetAttributesInput) (req *request.Request, output *ModifyTrafficMirrorTargetAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyTrafficMirrorTargetAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTrafficMirrorTargetAttributesInput{}
	}

	output = &ModifyTrafficMirrorTargetAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyTrafficMirrorTargetAttributes API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyTrafficMirrorTargetAttributes for usage and error information.
func (c *VPC) ModifyTrafficMirrorTargetAttributes(input *ModifyTrafficMirrorTargetAttributesInput) (*ModifyTrafficMirrorTargetAttributesOutput, error) {
	req, out := c.ModifyTrafficMirrorTargetAttributesRequest(input)
	return out, req.Send()
}

// ModifyTrafficMirrorTargetAttributesWithContext is the same as ModifyTrafficMirrorTargetAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTrafficMirrorTargetAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyTrafficMirrorTargetAttributesWithContext(ctx byteplus.Context, input *ModifyTrafficMirrorTargetAttributesInput, opts ...request.Option) (*ModifyTrafficMirrorTargetAttributesOutput, error) {
	req, out := c.ModifyTrafficMirrorTargetAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTrafficMirrorTargetAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// TrafficMirrorTargetId is a required field
	TrafficMirrorTargetId *string `type:"string" required:"true"`

	TrafficMirrorTargetName *string `type:"string"`
}

// String returns the string representation
func (s ModifyTrafficMirrorTargetAttributesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTrafficMirrorTargetAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTrafficMirrorTargetAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTrafficMirrorTargetAttributesInput"}
	if s.TrafficMirrorTargetId == nil {
		invalidParams.Add(request.NewErrParamRequired("TrafficMirrorTargetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyTrafficMirrorTargetAttributesInput) SetDescription(v string) *ModifyTrafficMirrorTargetAttributesInput {
	s.Description = &v
	return s
}

// SetTrafficMirrorTargetId sets the TrafficMirrorTargetId field's value.
func (s *ModifyTrafficMirrorTargetAttributesInput) SetTrafficMirrorTargetId(v string) *ModifyTrafficMirrorTargetAttributesInput {
	s.TrafficMirrorTargetId = &v
	return s
}

// SetTrafficMirrorTargetName sets the TrafficMirrorTargetName field's value.
func (s *ModifyTrafficMirrorTargetAttributesInput) SetTrafficMirrorTargetName(v string) *ModifyTrafficMirrorTargetAttributesInput {
	s.TrafficMirrorTargetName = &v
	return s
}

type ModifyTrafficMirrorTargetAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyTrafficMirrorTargetAttributesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTrafficMirrorTargetAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyTrafficMirrorTargetAttributesOutput) SetRequestId(v string) *ModifyTrafficMirrorTargetAttributesOutput {
	s.RequestId = &v
	return s
}
