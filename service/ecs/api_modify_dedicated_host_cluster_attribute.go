// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyDedicatedHostClusterAttributeCommon = "ModifyDedicatedHostClusterAttribute"

// ModifyDedicatedHostClusterAttributeCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyDedicatedHostClusterAttributeCommon operation. The "output" return
// value will be populated with the ModifyDedicatedHostClusterAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDedicatedHostClusterAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDedicatedHostClusterAttributeCommon Send returns without error.
//
// See ModifyDedicatedHostClusterAttributeCommon for more information on using the ModifyDedicatedHostClusterAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDedicatedHostClusterAttributeCommonRequest method.
//    req, resp := client.ModifyDedicatedHostClusterAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyDedicatedHostClusterAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDedicatedHostClusterAttributeCommon,
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

// ModifyDedicatedHostClusterAttributeCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyDedicatedHostClusterAttributeCommon for usage and error information.
func (c *ECS) ModifyDedicatedHostClusterAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDedicatedHostClusterAttributeCommonRequest(input)
	return out, req.Send()
}

// ModifyDedicatedHostClusterAttributeCommonWithContext is the same as ModifyDedicatedHostClusterAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDedicatedHostClusterAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyDedicatedHostClusterAttributeCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDedicatedHostClusterAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDedicatedHostClusterAttribute = "ModifyDedicatedHostClusterAttribute"

// ModifyDedicatedHostClusterAttributeRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyDedicatedHostClusterAttribute operation. The "output" return
// value will be populated with the ModifyDedicatedHostClusterAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDedicatedHostClusterAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDedicatedHostClusterAttributeCommon Send returns without error.
//
// See ModifyDedicatedHostClusterAttribute for more information on using the ModifyDedicatedHostClusterAttribute
// API call, and error handling.
//
//    // Example sending a request using the ModifyDedicatedHostClusterAttributeRequest method.
//    req, resp := client.ModifyDedicatedHostClusterAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyDedicatedHostClusterAttributeRequest(input *ModifyDedicatedHostClusterAttributeInput) (req *request.Request, output *ModifyDedicatedHostClusterAttributeOutput) {
	op := &request.Operation{
		Name:       opModifyDedicatedHostClusterAttribute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDedicatedHostClusterAttributeInput{}
	}

	output = &ModifyDedicatedHostClusterAttributeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyDedicatedHostClusterAttribute API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyDedicatedHostClusterAttribute for usage and error information.
func (c *ECS) ModifyDedicatedHostClusterAttribute(input *ModifyDedicatedHostClusterAttributeInput) (*ModifyDedicatedHostClusterAttributeOutput, error) {
	req, out := c.ModifyDedicatedHostClusterAttributeRequest(input)
	return out, req.Send()
}

// ModifyDedicatedHostClusterAttributeWithContext is the same as ModifyDedicatedHostClusterAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDedicatedHostClusterAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyDedicatedHostClusterAttributeWithContext(ctx byteplus.Context, input *ModifyDedicatedHostClusterAttributeInput, opts ...request.Option) (*ModifyDedicatedHostClusterAttributeOutput, error) {
	req, out := c.ModifyDedicatedHostClusterAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDedicatedHostClusterAttributeInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	DedicatedHostClusterId *string `type:"string"`

	DedicatedHostClusterName *string `type:"string"`

	Description *string `type:"string"`
}

// String returns the string representation
func (s ModifyDedicatedHostClusterAttributeInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDedicatedHostClusterAttributeInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyDedicatedHostClusterAttributeInput) SetClientToken(v string) *ModifyDedicatedHostClusterAttributeInput {
	s.ClientToken = &v
	return s
}

// SetDedicatedHostClusterId sets the DedicatedHostClusterId field's value.
func (s *ModifyDedicatedHostClusterAttributeInput) SetDedicatedHostClusterId(v string) *ModifyDedicatedHostClusterAttributeInput {
	s.DedicatedHostClusterId = &v
	return s
}

// SetDedicatedHostClusterName sets the DedicatedHostClusterName field's value.
func (s *ModifyDedicatedHostClusterAttributeInput) SetDedicatedHostClusterName(v string) *ModifyDedicatedHostClusterAttributeInput {
	s.DedicatedHostClusterName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyDedicatedHostClusterAttributeInput) SetDescription(v string) *ModifyDedicatedHostClusterAttributeInput {
	s.Description = &v
	return s
}

type ModifyDedicatedHostClusterAttributeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDedicatedHostClusterAttributeOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDedicatedHostClusterAttributeOutput) GoString() string {
	return s.String()
}
