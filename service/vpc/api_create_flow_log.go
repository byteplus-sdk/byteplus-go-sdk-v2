// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreateFlowLogCommon = "CreateFlowLog"

// CreateFlowLogCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateFlowLogCommon operation. The "output" return
// value will be populated with the CreateFlowLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFlowLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFlowLogCommon Send returns without error.
//
// See CreateFlowLogCommon for more information on using the CreateFlowLogCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateFlowLogCommonRequest method.
//    req, resp := client.CreateFlowLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateFlowLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateFlowLogCommon,
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

// CreateFlowLogCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreateFlowLogCommon for usage and error information.
func (c *VPC) CreateFlowLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateFlowLogCommonRequest(input)
	return out, req.Send()
}

// CreateFlowLogCommonWithContext is the same as CreateFlowLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFlowLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateFlowLogCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateFlowLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateFlowLog = "CreateFlowLog"

// CreateFlowLogRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateFlowLog operation. The "output" return
// value will be populated with the CreateFlowLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFlowLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFlowLogCommon Send returns without error.
//
// See CreateFlowLog for more information on using the CreateFlowLog
// API call, and error handling.
//
//    // Example sending a request using the CreateFlowLogRequest method.
//    req, resp := client.CreateFlowLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateFlowLogRequest(input *CreateFlowLogInput) (req *request.Request, output *CreateFlowLogOutput) {
	op := &request.Operation{
		Name:       opCreateFlowLog,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFlowLogInput{}
	}

	output = &CreateFlowLogOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateFlowLog API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreateFlowLog for usage and error information.
func (c *VPC) CreateFlowLog(input *CreateFlowLogInput) (*CreateFlowLogOutput, error) {
	req, out := c.CreateFlowLogRequest(input)
	return out, req.Send()
}

// CreateFlowLogWithContext is the same as CreateFlowLog with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFlowLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateFlowLogWithContext(ctx byteplus.Context, input *CreateFlowLogInput, opts ...request.Option) (*CreateFlowLogOutput, error) {
	req, out := c.CreateFlowLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateFlowLogInput struct {
	_ struct{} `type:"structure"`

	// AggregationInterval is a required field
	AggregationInterval *int64 `type:"integer" required:"true"`

	ClientToken *string `type:"string"`

	Description *string `max:"255" type:"string"`

	// FlowLogName is a required field
	FlowLogName *string `min:"1" max:"128" type:"string" required:"true"`

	// LogProjectName is a required field
	LogProjectName *string `type:"string" required:"true"`

	// LogTopicName is a required field
	LogTopicName *string `type:"string" required:"true"`

	ProjectName *string `type:"string"`

	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// ResourceType is a required field
	ResourceType *string `type:"string" required:"true"`

	Tags []*TagForCreateFlowLogInput `type:"list"`

	// TrafficType is a required field
	TrafficType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateFlowLogInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFlowLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFlowLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateFlowLogInput"}
	if s.AggregationInterval == nil {
		invalidParams.Add(request.NewErrParamRequired("AggregationInterval"))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.FlowLogName == nil {
		invalidParams.Add(request.NewErrParamRequired("FlowLogName"))
	}
	if s.FlowLogName != nil && len(*s.FlowLogName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("FlowLogName", 1))
	}
	if s.FlowLogName != nil && len(*s.FlowLogName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("FlowLogName", 128, *s.FlowLogName))
	}
	if s.LogProjectName == nil {
		invalidParams.Add(request.NewErrParamRequired("LogProjectName"))
	}
	if s.LogTopicName == nil {
		invalidParams.Add(request.NewErrParamRequired("LogTopicName"))
	}
	if s.ResourceId == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceType"))
	}
	if s.TrafficType == nil {
		invalidParams.Add(request.NewErrParamRequired("TrafficType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAggregationInterval sets the AggregationInterval field's value.
func (s *CreateFlowLogInput) SetAggregationInterval(v int64) *CreateFlowLogInput {
	s.AggregationInterval = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateFlowLogInput) SetClientToken(v string) *CreateFlowLogInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateFlowLogInput) SetDescription(v string) *CreateFlowLogInput {
	s.Description = &v
	return s
}

// SetFlowLogName sets the FlowLogName field's value.
func (s *CreateFlowLogInput) SetFlowLogName(v string) *CreateFlowLogInput {
	s.FlowLogName = &v
	return s
}

// SetLogProjectName sets the LogProjectName field's value.
func (s *CreateFlowLogInput) SetLogProjectName(v string) *CreateFlowLogInput {
	s.LogProjectName = &v
	return s
}

// SetLogTopicName sets the LogTopicName field's value.
func (s *CreateFlowLogInput) SetLogTopicName(v string) *CreateFlowLogInput {
	s.LogTopicName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateFlowLogInput) SetProjectName(v string) *CreateFlowLogInput {
	s.ProjectName = &v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *CreateFlowLogInput) SetResourceId(v string) *CreateFlowLogInput {
	s.ResourceId = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *CreateFlowLogInput) SetResourceType(v string) *CreateFlowLogInput {
	s.ResourceType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateFlowLogInput) SetTags(v []*TagForCreateFlowLogInput) *CreateFlowLogInput {
	s.Tags = v
	return s
}

// SetTrafficType sets the TrafficType field's value.
func (s *CreateFlowLogInput) SetTrafficType(v string) *CreateFlowLogInput {
	s.TrafficType = &v
	return s
}

type CreateFlowLogOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	FlowLogId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateFlowLogOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFlowLogOutput) GoString() string {
	return s.String()
}

// SetFlowLogId sets the FlowLogId field's value.
func (s *CreateFlowLogOutput) SetFlowLogId(v string) *CreateFlowLogOutput {
	s.FlowLogId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateFlowLogOutput) SetRequestId(v string) *CreateFlowLogOutput {
	s.RequestId = &v
	return s
}

type TagForCreateFlowLogInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateFlowLogInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateFlowLogInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateFlowLogInput) SetKey(v string) *TagForCreateFlowLogInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateFlowLogInput) SetValue(v string) *TagForCreateFlowLogInput {
	s.Value = &v
	return s
}
