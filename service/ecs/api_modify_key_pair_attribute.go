// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyKeyPairAttributeCommon = "ModifyKeyPairAttribute"

// ModifyKeyPairAttributeCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyKeyPairAttributeCommon operation. The "output" return
// value will be populated with the ModifyKeyPairAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyKeyPairAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyKeyPairAttributeCommon Send returns without error.
//
// See ModifyKeyPairAttributeCommon for more information on using the ModifyKeyPairAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyKeyPairAttributeCommonRequest method.
//    req, resp := client.ModifyKeyPairAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyKeyPairAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyKeyPairAttributeCommon,
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

// ModifyKeyPairAttributeCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyKeyPairAttributeCommon for usage and error information.
func (c *ECS) ModifyKeyPairAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyKeyPairAttributeCommonRequest(input)
	return out, req.Send()
}

// ModifyKeyPairAttributeCommonWithContext is the same as ModifyKeyPairAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyKeyPairAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyKeyPairAttributeCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyKeyPairAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyKeyPairAttribute = "ModifyKeyPairAttribute"

// ModifyKeyPairAttributeRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyKeyPairAttribute operation. The "output" return
// value will be populated with the ModifyKeyPairAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyKeyPairAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyKeyPairAttributeCommon Send returns without error.
//
// See ModifyKeyPairAttribute for more information on using the ModifyKeyPairAttribute
// API call, and error handling.
//
//    // Example sending a request using the ModifyKeyPairAttributeRequest method.
//    req, resp := client.ModifyKeyPairAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyKeyPairAttributeRequest(input *ModifyKeyPairAttributeInput) (req *request.Request, output *ModifyKeyPairAttributeOutput) {
	op := &request.Operation{
		Name:       opModifyKeyPairAttribute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyKeyPairAttributeInput{}
	}

	output = &ModifyKeyPairAttributeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyKeyPairAttribute API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyKeyPairAttribute for usage and error information.
func (c *ECS) ModifyKeyPairAttribute(input *ModifyKeyPairAttributeInput) (*ModifyKeyPairAttributeOutput, error) {
	req, out := c.ModifyKeyPairAttributeRequest(input)
	return out, req.Send()
}

// ModifyKeyPairAttributeWithContext is the same as ModifyKeyPairAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyKeyPairAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyKeyPairAttributeWithContext(ctx byteplus.Context, input *ModifyKeyPairAttributeInput, opts ...request.Option) (*ModifyKeyPairAttributeOutput, error) {
	req, out := c.ModifyKeyPairAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyKeyPairAttributeInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	KeyPairId *string `type:"string"`

	KeyPairName *string `type:"string"`
}

// String returns the string representation
func (s ModifyKeyPairAttributeInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyKeyPairAttributeInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyKeyPairAttributeInput) SetClientToken(v string) *ModifyKeyPairAttributeInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyKeyPairAttributeInput) SetDescription(v string) *ModifyKeyPairAttributeInput {
	s.Description = &v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *ModifyKeyPairAttributeInput) SetKeyPairId(v string) *ModifyKeyPairAttributeInput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *ModifyKeyPairAttributeInput) SetKeyPairName(v string) *ModifyKeyPairAttributeInput {
	s.KeyPairName = &v
	return s
}

type ModifyKeyPairAttributeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	KeyPairName *string `type:"string"`
}

// String returns the string representation
func (s ModifyKeyPairAttributeOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyKeyPairAttributeOutput) GoString() string {
	return s.String()
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *ModifyKeyPairAttributeOutput) SetKeyPairName(v string) *ModifyKeyPairAttributeOutput {
	s.KeyPairName = &v
	return s
}
