// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeTrafficMirrorTargetsCommon = "DescribeTrafficMirrorTargets"

// DescribeTrafficMirrorTargetsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeTrafficMirrorTargetsCommon operation. The "output" return
// value will be populated with the DescribeTrafficMirrorTargetsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTrafficMirrorTargetsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTrafficMirrorTargetsCommon Send returns without error.
//
// See DescribeTrafficMirrorTargetsCommon for more information on using the DescribeTrafficMirrorTargetsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTrafficMirrorTargetsCommonRequest method.
//    req, resp := client.DescribeTrafficMirrorTargetsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeTrafficMirrorTargetsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTrafficMirrorTargetsCommon,
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

// DescribeTrafficMirrorTargetsCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeTrafficMirrorTargetsCommon for usage and error information.
func (c *VPC) DescribeTrafficMirrorTargetsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTrafficMirrorTargetsCommonRequest(input)
	return out, req.Send()
}

// DescribeTrafficMirrorTargetsCommonWithContext is the same as DescribeTrafficMirrorTargetsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTrafficMirrorTargetsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeTrafficMirrorTargetsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTrafficMirrorTargetsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTrafficMirrorTargets = "DescribeTrafficMirrorTargets"

// DescribeTrafficMirrorTargetsRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeTrafficMirrorTargets operation. The "output" return
// value will be populated with the DescribeTrafficMirrorTargetsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTrafficMirrorTargetsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTrafficMirrorTargetsCommon Send returns without error.
//
// See DescribeTrafficMirrorTargets for more information on using the DescribeTrafficMirrorTargets
// API call, and error handling.
//
//    // Example sending a request using the DescribeTrafficMirrorTargetsRequest method.
//    req, resp := client.DescribeTrafficMirrorTargetsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeTrafficMirrorTargetsRequest(input *DescribeTrafficMirrorTargetsInput) (req *request.Request, output *DescribeTrafficMirrorTargetsOutput) {
	op := &request.Operation{
		Name:       opDescribeTrafficMirrorTargets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTrafficMirrorTargetsInput{}
	}

	output = &DescribeTrafficMirrorTargetsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTrafficMirrorTargets API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeTrafficMirrorTargets for usage and error information.
func (c *VPC) DescribeTrafficMirrorTargets(input *DescribeTrafficMirrorTargetsInput) (*DescribeTrafficMirrorTargetsOutput, error) {
	req, out := c.DescribeTrafficMirrorTargetsRequest(input)
	return out, req.Send()
}

// DescribeTrafficMirrorTargetsWithContext is the same as DescribeTrafficMirrorTargets with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTrafficMirrorTargets for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeTrafficMirrorTargetsWithContext(ctx byteplus.Context, input *DescribeTrafficMirrorTargetsInput, opts ...request.Option) (*DescribeTrafficMirrorTargetsOutput, error) {
	req, out := c.DescribeTrafficMirrorTargetsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTrafficMirrorTargetsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `type:"integer"`

	NextToken *string `type:"string"`

	ProjectName *string `type:"string"`

	TagFilters []*TagFilterForDescribeTrafficMirrorTargetsInput `type:"list"`

	TrafficMirrorTargetIds *string `type:"string"`

	TrafficMirrorTargetName *string `type:"string"`
}

// String returns the string representation
func (s DescribeTrafficMirrorTargetsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTrafficMirrorTargetsInput) GoString() string {
	return s.String()
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeTrafficMirrorTargetsInput) SetMaxResults(v int64) *DescribeTrafficMirrorTargetsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeTrafficMirrorTargetsInput) SetNextToken(v string) *DescribeTrafficMirrorTargetsInput {
	s.NextToken = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeTrafficMirrorTargetsInput) SetProjectName(v string) *DescribeTrafficMirrorTargetsInput {
	s.ProjectName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeTrafficMirrorTargetsInput) SetTagFilters(v []*TagFilterForDescribeTrafficMirrorTargetsInput) *DescribeTrafficMirrorTargetsInput {
	s.TagFilters = v
	return s
}

// SetTrafficMirrorTargetIds sets the TrafficMirrorTargetIds field's value.
func (s *DescribeTrafficMirrorTargetsInput) SetTrafficMirrorTargetIds(v string) *DescribeTrafficMirrorTargetsInput {
	s.TrafficMirrorTargetIds = &v
	return s
}

// SetTrafficMirrorTargetName sets the TrafficMirrorTargetName field's value.
func (s *DescribeTrafficMirrorTargetsInput) SetTrafficMirrorTargetName(v string) *DescribeTrafficMirrorTargetsInput {
	s.TrafficMirrorTargetName = &v
	return s
}

type DescribeTrafficMirrorTargetsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	RequestId *string `type:"string"`

	TrafficMirrorTargets []*TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput `type:"list"`
}

// String returns the string representation
func (s DescribeTrafficMirrorTargetsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTrafficMirrorTargetsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeTrafficMirrorTargetsOutput) SetNextToken(v string) *DescribeTrafficMirrorTargetsOutput {
	s.NextToken = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeTrafficMirrorTargetsOutput) SetRequestId(v string) *DescribeTrafficMirrorTargetsOutput {
	s.RequestId = &v
	return s
}

// SetTrafficMirrorTargets sets the TrafficMirrorTargets field's value.
func (s *DescribeTrafficMirrorTargetsOutput) SetTrafficMirrorTargets(v []*TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) *DescribeTrafficMirrorTargetsOutput {
	s.TrafficMirrorTargets = v
	return s
}

type TagFilterForDescribeTrafficMirrorTargetsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeTrafficMirrorTargetsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeTrafficMirrorTargetsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeTrafficMirrorTargetsInput) SetKey(v string) *TagFilterForDescribeTrafficMirrorTargetsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeTrafficMirrorTargetsInput) SetValues(v []*string) *TagFilterForDescribeTrafficMirrorTargetsInput {
	s.Values = v
	return s
}

type TagForDescribeTrafficMirrorTargetsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeTrafficMirrorTargetsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeTrafficMirrorTargetsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeTrafficMirrorTargetsOutput) SetKey(v string) *TagForDescribeTrafficMirrorTargetsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeTrafficMirrorTargetsOutput) SetValue(v string) *TagForDescribeTrafficMirrorTargetsOutput {
	s.Value = &v
	return s
}

type TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceType *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeTrafficMirrorTargetsOutput `type:"list"`

	TrafficMirrorTargetId *string `type:"string"`

	TrafficMirrorTargetName *string `type:"string"`
}

// String returns the string representation
func (s TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetCreatedAt(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetDescription(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.Description = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetInstanceId(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetInstanceType(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.InstanceType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetProjectName(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetStatus(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetTags(v []*TagForDescribeTrafficMirrorTargetsOutput) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.Tags = v
	return s
}

// SetTrafficMirrorTargetId sets the TrafficMirrorTargetId field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetTrafficMirrorTargetId(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.TrafficMirrorTargetId = &v
	return s
}

// SetTrafficMirrorTargetName sets the TrafficMirrorTargetName field's value.
func (s *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput) SetTrafficMirrorTargetName(v string) *TrafficMirrorTargetForDescribeTrafficMirrorTargetsOutput {
	s.TrafficMirrorTargetName = &v
	return s
}
