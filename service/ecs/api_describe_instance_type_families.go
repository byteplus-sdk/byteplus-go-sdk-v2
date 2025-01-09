// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeInstanceTypeFamiliesCommon = "DescribeInstanceTypeFamilies"

// DescribeInstanceTypeFamiliesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInstanceTypeFamiliesCommon operation. The "output" return
// value will be populated with the DescribeInstanceTypeFamiliesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceTypeFamiliesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceTypeFamiliesCommon Send returns without error.
//
// See DescribeInstanceTypeFamiliesCommon for more information on using the DescribeInstanceTypeFamiliesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceTypeFamiliesCommonRequest method.
//    req, resp := client.DescribeInstanceTypeFamiliesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstanceTypeFamiliesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstanceTypeFamiliesCommon,
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

// DescribeInstanceTypeFamiliesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInstanceTypeFamiliesCommon for usage and error information.
func (c *ECS) DescribeInstanceTypeFamiliesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceTypeFamiliesCommonRequest(input)
	return out, req.Send()
}

// DescribeInstanceTypeFamiliesCommonWithContext is the same as DescribeInstanceTypeFamiliesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceTypeFamiliesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstanceTypeFamiliesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceTypeFamiliesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstanceTypeFamilies = "DescribeInstanceTypeFamilies"

// DescribeInstanceTypeFamiliesRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInstanceTypeFamilies operation. The "output" return
// value will be populated with the DescribeInstanceTypeFamiliesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceTypeFamiliesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceTypeFamiliesCommon Send returns without error.
//
// See DescribeInstanceTypeFamilies for more information on using the DescribeInstanceTypeFamilies
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceTypeFamiliesRequest method.
//    req, resp := client.DescribeInstanceTypeFamiliesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstanceTypeFamiliesRequest(input *DescribeInstanceTypeFamiliesInput) (req *request.Request, output *DescribeInstanceTypeFamiliesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstanceTypeFamilies,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceTypeFamiliesInput{}
	}

	output = &DescribeInstanceTypeFamiliesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInstanceTypeFamilies API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInstanceTypeFamilies for usage and error information.
func (c *ECS) DescribeInstanceTypeFamilies(input *DescribeInstanceTypeFamiliesInput) (*DescribeInstanceTypeFamiliesOutput, error) {
	req, out := c.DescribeInstanceTypeFamiliesRequest(input)
	return out, req.Send()
}

// DescribeInstanceTypeFamiliesWithContext is the same as DescribeInstanceTypeFamilies with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceTypeFamilies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstanceTypeFamiliesWithContext(ctx byteplus.Context, input *DescribeInstanceTypeFamiliesInput, opts ...request.Option) (*DescribeInstanceTypeFamiliesOutput, error) {
	req, out := c.DescribeInstanceTypeFamiliesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInstanceTypeFamiliesInput struct {
	_ struct{} `type:"structure"`

	Generation *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceTypeFamiliesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceTypeFamiliesInput) GoString() string {
	return s.String()
}

// SetGeneration sets the Generation field's value.
func (s *DescribeInstanceTypeFamiliesInput) SetGeneration(v string) *DescribeInstanceTypeFamiliesInput {
	s.Generation = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeInstanceTypeFamiliesInput) SetZoneId(v string) *DescribeInstanceTypeFamiliesInput {
	s.ZoneId = &v
	return s
}

type DescribeInstanceTypeFamiliesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceTypeFamilies []*InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeInstanceTypeFamiliesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceTypeFamiliesOutput) GoString() string {
	return s.String()
}

// SetInstanceTypeFamilies sets the InstanceTypeFamilies field's value.
func (s *DescribeInstanceTypeFamiliesOutput) SetInstanceTypeFamilies(v []*InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput) *DescribeInstanceTypeFamiliesOutput {
	s.InstanceTypeFamilies = v
	return s
}

type InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput struct {
	_ struct{} `type:"structure"`

	Generation *string `type:"string"`

	InstanceTypeFamily *string `type:"string"`

	ZoneIds []*string `type:"list"`
}

// String returns the string representation
func (s InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput) GoString() string {
	return s.String()
}

// SetGeneration sets the Generation field's value.
func (s *InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput) SetGeneration(v string) *InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput {
	s.Generation = &v
	return s
}

// SetInstanceTypeFamily sets the InstanceTypeFamily field's value.
func (s *InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput) SetInstanceTypeFamily(v string) *InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput {
	s.InstanceTypeFamily = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput) SetZoneIds(v []*string) *InstanceTypeFamilyForDescribeInstanceTypeFamiliesOutput {
	s.ZoneIds = v
	return s
}