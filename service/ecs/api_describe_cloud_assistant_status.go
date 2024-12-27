// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeCloudAssistantStatusCommon = "DescribeCloudAssistantStatus"

// DescribeCloudAssistantStatusCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeCloudAssistantStatusCommon operation. The "output" return
// value will be populated with the DescribeCloudAssistantStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCloudAssistantStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCloudAssistantStatusCommon Send returns without error.
//
// See DescribeCloudAssistantStatusCommon for more information on using the DescribeCloudAssistantStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCloudAssistantStatusCommonRequest method.
//    req, resp := client.DescribeCloudAssistantStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeCloudAssistantStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCloudAssistantStatusCommon,
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

// DescribeCloudAssistantStatusCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeCloudAssistantStatusCommon for usage and error information.
func (c *ECS) DescribeCloudAssistantStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCloudAssistantStatusCommonRequest(input)
	return out, req.Send()
}

// DescribeCloudAssistantStatusCommonWithContext is the same as DescribeCloudAssistantStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCloudAssistantStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeCloudAssistantStatusCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCloudAssistantStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCloudAssistantStatus = "DescribeCloudAssistantStatus"

// DescribeCloudAssistantStatusRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeCloudAssistantStatus operation. The "output" return
// value will be populated with the DescribeCloudAssistantStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCloudAssistantStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCloudAssistantStatusCommon Send returns without error.
//
// See DescribeCloudAssistantStatus for more information on using the DescribeCloudAssistantStatus
// API call, and error handling.
//
//    // Example sending a request using the DescribeCloudAssistantStatusRequest method.
//    req, resp := client.DescribeCloudAssistantStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeCloudAssistantStatusRequest(input *DescribeCloudAssistantStatusInput) (req *request.Request, output *DescribeCloudAssistantStatusOutput) {
	op := &request.Operation{
		Name:       opDescribeCloudAssistantStatus,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCloudAssistantStatusInput{}
	}

	output = &DescribeCloudAssistantStatusOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCloudAssistantStatus API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeCloudAssistantStatus for usage and error information.
func (c *ECS) DescribeCloudAssistantStatus(input *DescribeCloudAssistantStatusInput) (*DescribeCloudAssistantStatusOutput, error) {
	req, out := c.DescribeCloudAssistantStatusRequest(input)
	return out, req.Send()
}

// DescribeCloudAssistantStatusWithContext is the same as DescribeCloudAssistantStatus with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCloudAssistantStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeCloudAssistantStatusWithContext(ctx byteplus.Context, input *DescribeCloudAssistantStatusInput, opts ...request.Option) (*DescribeCloudAssistantStatusOutput, error) {
	req, out := c.DescribeCloudAssistantStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeCloudAssistantStatusInput struct {
	_ struct{} `type:"structure"`

	ClientVersion *string `type:"string"`

	InstanceIds []*string `type:"list"`

	OSType *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeCloudAssistantStatusInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCloudAssistantStatusInput) GoString() string {
	return s.String()
}

// SetClientVersion sets the ClientVersion field's value.
func (s *DescribeCloudAssistantStatusInput) SetClientVersion(v string) *DescribeCloudAssistantStatusInput {
	s.ClientVersion = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DescribeCloudAssistantStatusInput) SetInstanceIds(v []*string) *DescribeCloudAssistantStatusInput {
	s.InstanceIds = v
	return s
}

// SetOSType sets the OSType field's value.
func (s *DescribeCloudAssistantStatusInput) SetOSType(v string) *DescribeCloudAssistantStatusInput {
	s.OSType = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCloudAssistantStatusInput) SetPageNumber(v int32) *DescribeCloudAssistantStatusInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCloudAssistantStatusInput) SetPageSize(v int32) *DescribeCloudAssistantStatusInput {
	s.PageSize = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeCloudAssistantStatusInput) SetStatus(v string) *DescribeCloudAssistantStatusInput {
	s.Status = &v
	return s
}

type DescribeCloudAssistantStatusOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Instances []*InstanceForDescribeCloudAssistantStatusOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeCloudAssistantStatusOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCloudAssistantStatusOutput) GoString() string {
	return s.String()
}

// SetInstances sets the Instances field's value.
func (s *DescribeCloudAssistantStatusOutput) SetInstances(v []*InstanceForDescribeCloudAssistantStatusOutput) *DescribeCloudAssistantStatusOutput {
	s.Instances = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCloudAssistantStatusOutput) SetPageNumber(v int32) *DescribeCloudAssistantStatusOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCloudAssistantStatusOutput) SetPageSize(v int32) *DescribeCloudAssistantStatusOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCloudAssistantStatusOutput) SetTotalCount(v int32) *DescribeCloudAssistantStatusOutput {
	s.TotalCount = &v
	return s
}

type InstanceForDescribeCloudAssistantStatusOutput struct {
	_ struct{} `type:"structure"`

	ClientVersion *string `type:"string"`

	HostName *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	LastHeartbeatTime *string `type:"string"`

	OSType *string `type:"string"`

	OSVersion *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s InstanceForDescribeCloudAssistantStatusOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForDescribeCloudAssistantStatusOutput) GoString() string {
	return s.String()
}

// SetClientVersion sets the ClientVersion field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetClientVersion(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.ClientVersion = &v
	return s
}

// SetHostName sets the HostName field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetHostName(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.HostName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetInstanceId(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetInstanceName(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.InstanceName = &v
	return s
}

// SetLastHeartbeatTime sets the LastHeartbeatTime field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetLastHeartbeatTime(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.LastHeartbeatTime = &v
	return s
}

// SetOSType sets the OSType field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetOSType(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.OSType = &v
	return s
}

// SetOSVersion sets the OSVersion field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetOSVersion(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.OSVersion = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *InstanceForDescribeCloudAssistantStatusOutput) SetStatus(v string) *InstanceForDescribeCloudAssistantStatusOutput {
	s.Status = &v
	return s
}
