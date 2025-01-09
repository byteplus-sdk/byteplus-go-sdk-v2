// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeInvocationInstancesCommon = "DescribeInvocationInstances"

// DescribeInvocationInstancesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInvocationInstancesCommon operation. The "output" return
// value will be populated with the DescribeInvocationInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInvocationInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInvocationInstancesCommon Send returns without error.
//
// See DescribeInvocationInstancesCommon for more information on using the DescribeInvocationInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInvocationInstancesCommonRequest method.
//    req, resp := client.DescribeInvocationInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInvocationInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInvocationInstancesCommon,
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

// DescribeInvocationInstancesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInvocationInstancesCommon for usage and error information.
func (c *ECS) DescribeInvocationInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInvocationInstancesCommonRequest(input)
	return out, req.Send()
}

// DescribeInvocationInstancesCommonWithContext is the same as DescribeInvocationInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInvocationInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInvocationInstancesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInvocationInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInvocationInstances = "DescribeInvocationInstances"

// DescribeInvocationInstancesRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInvocationInstances operation. The "output" return
// value will be populated with the DescribeInvocationInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInvocationInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInvocationInstancesCommon Send returns without error.
//
// See DescribeInvocationInstances for more information on using the DescribeInvocationInstances
// API call, and error handling.
//
//    // Example sending a request using the DescribeInvocationInstancesRequest method.
//    req, resp := client.DescribeInvocationInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInvocationInstancesRequest(input *DescribeInvocationInstancesInput) (req *request.Request, output *DescribeInvocationInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeInvocationInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInvocationInstancesInput{}
	}

	output = &DescribeInvocationInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInvocationInstances API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInvocationInstances for usage and error information.
func (c *ECS) DescribeInvocationInstances(input *DescribeInvocationInstancesInput) (*DescribeInvocationInstancesOutput, error) {
	req, out := c.DescribeInvocationInstancesRequest(input)
	return out, req.Send()
}

// DescribeInvocationInstancesWithContext is the same as DescribeInvocationInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInvocationInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInvocationInstancesWithContext(ctx byteplus.Context, input *DescribeInvocationInstancesInput, opts ...request.Option) (*DescribeInvocationInstancesOutput, error) {
	req, out := c.DescribeInvocationInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInvocationInstancesInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	InvocationId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeInvocationInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInvocationInstancesInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeInvocationInstancesInput) SetInstanceId(v string) *DescribeInvocationInstancesInput {
	s.InstanceId = &v
	return s
}

// SetInvocationId sets the InvocationId field's value.
func (s *DescribeInvocationInstancesInput) SetInvocationId(v string) *DescribeInvocationInstancesInput {
	s.InvocationId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeInvocationInstancesInput) SetPageNumber(v int32) *DescribeInvocationInstancesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeInvocationInstancesInput) SetPageSize(v int32) *DescribeInvocationInstancesInput {
	s.PageSize = &v
	return s
}

type DescribeInvocationInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InvocationInstances []*InvocationInstanceForDescribeInvocationInstancesOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeInvocationInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInvocationInstancesOutput) GoString() string {
	return s.String()
}

// SetInvocationInstances sets the InvocationInstances field's value.
func (s *DescribeInvocationInstancesOutput) SetInvocationInstances(v []*InvocationInstanceForDescribeInvocationInstancesOutput) *DescribeInvocationInstancesOutput {
	s.InvocationInstances = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeInvocationInstancesOutput) SetPageNumber(v int32) *DescribeInvocationInstancesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeInvocationInstancesOutput) SetPageSize(v int32) *DescribeInvocationInstancesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeInvocationInstancesOutput) SetTotalCount(v int32) *DescribeInvocationInstancesOutput {
	s.TotalCount = &v
	return s
}

type InvocationInstanceForDescribeInvocationInstancesOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	InvocationId *string `type:"string"`
}

// String returns the string representation
func (s InvocationInstanceForDescribeInvocationInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s InvocationInstanceForDescribeInvocationInstancesOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *InvocationInstanceForDescribeInvocationInstancesOutput) SetInstanceId(v string) *InvocationInstanceForDescribeInvocationInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInvocationId sets the InvocationId field's value.
func (s *InvocationInstanceForDescribeInvocationInstancesOutput) SetInvocationId(v string) *InvocationInstanceForDescribeInvocationInstancesOutput {
	s.InvocationId = &v
	return s
}