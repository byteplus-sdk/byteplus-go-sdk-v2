// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDisassociateInstancesIamRoleCommon = "DisassociateInstancesIamRole"

// DisassociateInstancesIamRoleCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DisassociateInstancesIamRoleCommon operation. The "output" return
// value will be populated with the DisassociateInstancesIamRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateInstancesIamRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateInstancesIamRoleCommon Send returns without error.
//
// See DisassociateInstancesIamRoleCommon for more information on using the DisassociateInstancesIamRoleCommon
// API call, and error handling.
//
//    // Example sending a request using the DisassociateInstancesIamRoleCommonRequest method.
//    req, resp := client.DisassociateInstancesIamRoleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DisassociateInstancesIamRoleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisassociateInstancesIamRoleCommon,
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

// DisassociateInstancesIamRoleCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DisassociateInstancesIamRoleCommon for usage and error information.
func (c *ECS) DisassociateInstancesIamRoleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisassociateInstancesIamRoleCommonRequest(input)
	return out, req.Send()
}

// DisassociateInstancesIamRoleCommonWithContext is the same as DisassociateInstancesIamRoleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateInstancesIamRoleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DisassociateInstancesIamRoleCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisassociateInstancesIamRoleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisassociateInstancesIamRole = "DisassociateInstancesIamRole"

// DisassociateInstancesIamRoleRequest generates a "byteplus/request.Request" representing the
// client's request for the DisassociateInstancesIamRole operation. The "output" return
// value will be populated with the DisassociateInstancesIamRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateInstancesIamRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateInstancesIamRoleCommon Send returns without error.
//
// See DisassociateInstancesIamRole for more information on using the DisassociateInstancesIamRole
// API call, and error handling.
//
//    // Example sending a request using the DisassociateInstancesIamRoleRequest method.
//    req, resp := client.DisassociateInstancesIamRoleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DisassociateInstancesIamRoleRequest(input *DisassociateInstancesIamRoleInput) (req *request.Request, output *DisassociateInstancesIamRoleOutput) {
	op := &request.Operation{
		Name:       opDisassociateInstancesIamRole,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateInstancesIamRoleInput{}
	}

	output = &DisassociateInstancesIamRoleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisassociateInstancesIamRole API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DisassociateInstancesIamRole for usage and error information.
func (c *ECS) DisassociateInstancesIamRole(input *DisassociateInstancesIamRoleInput) (*DisassociateInstancesIamRoleOutput, error) {
	req, out := c.DisassociateInstancesIamRoleRequest(input)
	return out, req.Send()
}

// DisassociateInstancesIamRoleWithContext is the same as DisassociateInstancesIamRole with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateInstancesIamRole for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DisassociateInstancesIamRoleWithContext(ctx byteplus.Context, input *DisassociateInstancesIamRoleInput, opts ...request.Option) (*DisassociateInstancesIamRoleOutput, error) {
	req, out := c.DisassociateInstancesIamRoleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisassociateInstancesIamRoleInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	IamRoleName *string `type:"string"`

	InstanceIds []*string `type:"list"`
}

// String returns the string representation
func (s DisassociateInstancesIamRoleInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateInstancesIamRoleInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *DisassociateInstancesIamRoleInput) SetClientToken(v string) *DisassociateInstancesIamRoleInput {
	s.ClientToken = &v
	return s
}

// SetIamRoleName sets the IamRoleName field's value.
func (s *DisassociateInstancesIamRoleInput) SetIamRoleName(v string) *DisassociateInstancesIamRoleInput {
	s.IamRoleName = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DisassociateInstancesIamRoleInput) SetInstanceIds(v []*string) *DisassociateInstancesIamRoleInput {
	s.InstanceIds = v
	return s
}

type DisassociateInstancesIamRoleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForDisassociateInstancesIamRoleOutput `type:"list"`
}

// String returns the string representation
func (s DisassociateInstancesIamRoleOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateInstancesIamRoleOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *DisassociateInstancesIamRoleOutput) SetOperationDetails(v []*OperationDetailForDisassociateInstancesIamRoleOutput) *DisassociateInstancesIamRoleOutput {
	s.OperationDetails = v
	return s
}

type ErrorForDisassociateInstancesIamRoleOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForDisassociateInstancesIamRoleOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForDisassociateInstancesIamRoleOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForDisassociateInstancesIamRoleOutput) SetCode(v string) *ErrorForDisassociateInstancesIamRoleOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForDisassociateInstancesIamRoleOutput) SetMessage(v string) *ErrorForDisassociateInstancesIamRoleOutput {
	s.Message = &v
	return s
}

type OperationDetailForDisassociateInstancesIamRoleOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForDisassociateInstancesIamRoleOutput `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForDisassociateInstancesIamRoleOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForDisassociateInstancesIamRoleOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForDisassociateInstancesIamRoleOutput) SetError(v *ErrorForDisassociateInstancesIamRoleOutput) *OperationDetailForDisassociateInstancesIamRoleOutput {
	s.Error = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *OperationDetailForDisassociateInstancesIamRoleOutput) SetInstanceId(v string) *OperationDetailForDisassociateInstancesIamRoleOutput {
	s.InstanceId = &v
	return s
}
