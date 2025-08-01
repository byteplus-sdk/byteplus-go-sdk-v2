// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyInstanceMetadataOptionsCommon = "ModifyInstanceMetadataOptions"

// ModifyInstanceMetadataOptionsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyInstanceMetadataOptionsCommon operation. The "output" return
// value will be populated with the ModifyInstanceMetadataOptionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceMetadataOptionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceMetadataOptionsCommon Send returns without error.
//
// See ModifyInstanceMetadataOptionsCommon for more information on using the ModifyInstanceMetadataOptionsCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceMetadataOptionsCommonRequest method.
//    req, resp := client.ModifyInstanceMetadataOptionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceMetadataOptionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyInstanceMetadataOptionsCommon,
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

// ModifyInstanceMetadataOptionsCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyInstanceMetadataOptionsCommon for usage and error information.
func (c *ECS) ModifyInstanceMetadataOptionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceMetadataOptionsCommonRequest(input)
	return out, req.Send()
}

// ModifyInstanceMetadataOptionsCommonWithContext is the same as ModifyInstanceMetadataOptionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceMetadataOptionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceMetadataOptionsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceMetadataOptionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyInstanceMetadataOptions = "ModifyInstanceMetadataOptions"

// ModifyInstanceMetadataOptionsRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyInstanceMetadataOptions operation. The "output" return
// value will be populated with the ModifyInstanceMetadataOptionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceMetadataOptionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceMetadataOptionsCommon Send returns without error.
//
// See ModifyInstanceMetadataOptions for more information on using the ModifyInstanceMetadataOptions
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceMetadataOptionsRequest method.
//    req, resp := client.ModifyInstanceMetadataOptionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceMetadataOptionsRequest(input *ModifyInstanceMetadataOptionsInput) (req *request.Request, output *ModifyInstanceMetadataOptionsOutput) {
	op := &request.Operation{
		Name:       opModifyInstanceMetadataOptions,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceMetadataOptionsInput{}
	}

	output = &ModifyInstanceMetadataOptionsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyInstanceMetadataOptions API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyInstanceMetadataOptions for usage and error information.
func (c *ECS) ModifyInstanceMetadataOptions(input *ModifyInstanceMetadataOptionsInput) (*ModifyInstanceMetadataOptionsOutput, error) {
	req, out := c.ModifyInstanceMetadataOptionsRequest(input)
	return out, req.Send()
}

// ModifyInstanceMetadataOptionsWithContext is the same as ModifyInstanceMetadataOptions with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceMetadataOptions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceMetadataOptionsWithContext(ctx byteplus.Context, input *ModifyInstanceMetadataOptionsInput, opts ...request.Option) (*ModifyInstanceMetadataOptionsOutput, error) {
	req, out := c.ModifyInstanceMetadataOptionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyInstanceMetadataOptionsInput struct {
	_ struct{} `type:"structure"`

	HttpTokens *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyInstanceMetadataOptionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceMetadataOptionsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstanceMetadataOptionsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyInstanceMetadataOptionsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHttpTokens sets the HttpTokens field's value.
func (s *ModifyInstanceMetadataOptionsInput) SetHttpTokens(v string) *ModifyInstanceMetadataOptionsInput {
	s.HttpTokens = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyInstanceMetadataOptionsInput) SetInstanceId(v string) *ModifyInstanceMetadataOptionsInput {
	s.InstanceId = &v
	return s
}

type ModifyInstanceMetadataOptionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyInstanceMetadataOptionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceMetadataOptionsOutput) GoString() string {
	return s.String()
}
