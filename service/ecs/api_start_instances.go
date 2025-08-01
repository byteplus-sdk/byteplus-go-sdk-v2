// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opStartInstancesCommon = "StartInstances"

// StartInstancesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the StartInstancesCommon operation. The "output" return
// value will be populated with the StartInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartInstancesCommon Send returns without error.
//
// See StartInstancesCommon for more information on using the StartInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the StartInstancesCommonRequest method.
//    req, resp := client.StartInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) StartInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartInstancesCommon,
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

// StartInstancesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation StartInstancesCommon for usage and error information.
func (c *ECS) StartInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartInstancesCommonRequest(input)
	return out, req.Send()
}

// StartInstancesCommonWithContext is the same as StartInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) StartInstancesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartInstances = "StartInstances"

// StartInstancesRequest generates a "byteplus/request.Request" representing the
// client's request for the StartInstances operation. The "output" return
// value will be populated with the StartInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartInstancesCommon Send returns without error.
//
// See StartInstances for more information on using the StartInstances
// API call, and error handling.
//
//    // Example sending a request using the StartInstancesRequest method.
//    req, resp := client.StartInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) StartInstancesRequest(input *StartInstancesInput) (req *request.Request, output *StartInstancesOutput) {
	op := &request.Operation{
		Name:       opStartInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartInstancesInput{}
	}

	output = &StartInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// StartInstances API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation StartInstances for usage and error information.
func (c *ECS) StartInstances(input *StartInstancesInput) (*StartInstancesOutput, error) {
	req, out := c.StartInstancesRequest(input)
	return out, req.Send()
}

// StartInstancesWithContext is the same as StartInstances with the addition of
// the ability to pass a context and additional request options.
//
// See StartInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) StartInstancesWithContext(ctx byteplus.Context, input *StartInstancesInput, opts ...request.Option) (*StartInstancesOutput, error) {
	req, out := c.StartInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ErrorForStartInstancesOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForStartInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForStartInstancesOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForStartInstancesOutput) SetCode(v string) *ErrorForStartInstancesOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForStartInstancesOutput) SetMessage(v string) *ErrorForStartInstancesOutput {
	s.Message = &v
	return s
}

type OperationDetailForStartInstancesOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForStartInstancesOutput `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForStartInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForStartInstancesOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForStartInstancesOutput) SetError(v *ErrorForStartInstancesOutput) *OperationDetailForStartInstancesOutput {
	s.Error = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *OperationDetailForStartInstancesOutput) SetInstanceId(v string) *OperationDetailForStartInstancesOutput {
	s.InstanceId = &v
	return s
}

type StartInstancesInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// InstanceIds is a required field
	InstanceIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s StartInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s StartInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartInstancesInput"}
	if s.InstanceIds == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *StartInstancesInput) SetClientToken(v string) *StartInstancesInput {
	s.ClientToken = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *StartInstancesInput) SetInstanceIds(v []*string) *StartInstancesInput {
	s.InstanceIds = v
	return s
}

type StartInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForStartInstancesOutput `type:"list"`
}

// String returns the string representation
func (s StartInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s StartInstancesOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *StartInstancesOutput) SetOperationDetails(v []*OperationDetailForStartInstancesOutput) *StartInstancesOutput {
	s.OperationDetails = v
	return s
}
