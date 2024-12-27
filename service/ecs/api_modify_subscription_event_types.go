// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifySubscriptionEventTypesCommon = "ModifySubscriptionEventTypes"

// ModifySubscriptionEventTypesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifySubscriptionEventTypesCommon operation. The "output" return
// value will be populated with the ModifySubscriptionEventTypesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySubscriptionEventTypesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySubscriptionEventTypesCommon Send returns without error.
//
// See ModifySubscriptionEventTypesCommon for more information on using the ModifySubscriptionEventTypesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifySubscriptionEventTypesCommonRequest method.
//    req, resp := client.ModifySubscriptionEventTypesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifySubscriptionEventTypesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifySubscriptionEventTypesCommon,
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

// ModifySubscriptionEventTypesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifySubscriptionEventTypesCommon for usage and error information.
func (c *ECS) ModifySubscriptionEventTypesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifySubscriptionEventTypesCommonRequest(input)
	return out, req.Send()
}

// ModifySubscriptionEventTypesCommonWithContext is the same as ModifySubscriptionEventTypesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySubscriptionEventTypesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifySubscriptionEventTypesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifySubscriptionEventTypesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifySubscriptionEventTypes = "ModifySubscriptionEventTypes"

// ModifySubscriptionEventTypesRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifySubscriptionEventTypes operation. The "output" return
// value will be populated with the ModifySubscriptionEventTypesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySubscriptionEventTypesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySubscriptionEventTypesCommon Send returns without error.
//
// See ModifySubscriptionEventTypes for more information on using the ModifySubscriptionEventTypes
// API call, and error handling.
//
//    // Example sending a request using the ModifySubscriptionEventTypesRequest method.
//    req, resp := client.ModifySubscriptionEventTypesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifySubscriptionEventTypesRequest(input *ModifySubscriptionEventTypesInput) (req *request.Request, output *ModifySubscriptionEventTypesOutput) {
	op := &request.Operation{
		Name:       opModifySubscriptionEventTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySubscriptionEventTypesInput{}
	}

	output = &ModifySubscriptionEventTypesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySubscriptionEventTypes API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifySubscriptionEventTypes for usage and error information.
func (c *ECS) ModifySubscriptionEventTypes(input *ModifySubscriptionEventTypesInput) (*ModifySubscriptionEventTypesOutput, error) {
	req, out := c.ModifySubscriptionEventTypesRequest(input)
	return out, req.Send()
}

// ModifySubscriptionEventTypesWithContext is the same as ModifySubscriptionEventTypes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySubscriptionEventTypes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifySubscriptionEventTypesWithContext(ctx byteplus.Context, input *ModifySubscriptionEventTypesInput, opts ...request.Option) (*ModifySubscriptionEventTypesOutput, error) {
	req, out := c.ModifySubscriptionEventTypesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifySubscriptionEventTypesInput struct {
	_ struct{} `type:"structure"`

	EventTypes []*string `type:"list"`

	SubscriptionId *string `type:"string"`
}

// String returns the string representation
func (s ModifySubscriptionEventTypesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySubscriptionEventTypesInput) GoString() string {
	return s.String()
}

// SetEventTypes sets the EventTypes field's value.
func (s *ModifySubscriptionEventTypesInput) SetEventTypes(v []*string) *ModifySubscriptionEventTypesInput {
	s.EventTypes = v
	return s
}

// SetSubscriptionId sets the SubscriptionId field's value.
func (s *ModifySubscriptionEventTypesInput) SetSubscriptionId(v string) *ModifySubscriptionEventTypesInput {
	s.SubscriptionId = &v
	return s
}

type ModifySubscriptionEventTypesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	SubscriptionId *string `type:"string"`
}

// String returns the string representation
func (s ModifySubscriptionEventTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySubscriptionEventTypesOutput) GoString() string {
	return s.String()
}

// SetSubscriptionId sets the SubscriptionId field's value.
func (s *ModifySubscriptionEventTypesOutput) SetSubscriptionId(v string) *ModifySubscriptionEventTypesOutput {
	s.SubscriptionId = &v
	return s
}
