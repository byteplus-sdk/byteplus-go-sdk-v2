// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opAttachKeyPairCommon = "AttachKeyPair"

// AttachKeyPairCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the AttachKeyPairCommon operation. The "output" return
// value will be populated with the AttachKeyPairCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachKeyPairCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachKeyPairCommon Send returns without error.
//
// See AttachKeyPairCommon for more information on using the AttachKeyPairCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachKeyPairCommonRequest method.
//    req, resp := client.AttachKeyPairCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) AttachKeyPairCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachKeyPairCommon,
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

// AttachKeyPairCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation AttachKeyPairCommon for usage and error information.
func (c *ECS) AttachKeyPairCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachKeyPairCommonRequest(input)
	return out, req.Send()
}

// AttachKeyPairCommonWithContext is the same as AttachKeyPairCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachKeyPairCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) AttachKeyPairCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachKeyPairCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachKeyPair = "AttachKeyPair"

// AttachKeyPairRequest generates a "byteplus/request.Request" representing the
// client's request for the AttachKeyPair operation. The "output" return
// value will be populated with the AttachKeyPairCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachKeyPairCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachKeyPairCommon Send returns without error.
//
// See AttachKeyPair for more information on using the AttachKeyPair
// API call, and error handling.
//
//    // Example sending a request using the AttachKeyPairRequest method.
//    req, resp := client.AttachKeyPairRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) AttachKeyPairRequest(input *AttachKeyPairInput) (req *request.Request, output *AttachKeyPairOutput) {
	op := &request.Operation{
		Name:       opAttachKeyPair,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachKeyPairInput{}
	}

	output = &AttachKeyPairOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachKeyPair API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation AttachKeyPair for usage and error information.
func (c *ECS) AttachKeyPair(input *AttachKeyPairInput) (*AttachKeyPairOutput, error) {
	req, out := c.AttachKeyPairRequest(input)
	return out, req.Send()
}

// AttachKeyPairWithContext is the same as AttachKeyPair with the addition of
// the ability to pass a context and additional request options.
//
// See AttachKeyPair for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) AttachKeyPairWithContext(ctx byteplus.Context, input *AttachKeyPairInput, opts ...request.Option) (*AttachKeyPairOutput, error) {
	req, out := c.AttachKeyPairRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachKeyPairInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	InstanceIds []*string `type:"list"`

	KeyPairId *string `type:"string"`

	KeyPairName *string `type:"string"`
}

// String returns the string representation
func (s AttachKeyPairInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachKeyPairInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *AttachKeyPairInput) SetClientToken(v string) *AttachKeyPairInput {
	s.ClientToken = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *AttachKeyPairInput) SetInstanceIds(v []*string) *AttachKeyPairInput {
	s.InstanceIds = v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *AttachKeyPairInput) SetKeyPairId(v string) *AttachKeyPairInput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *AttachKeyPairInput) SetKeyPairName(v string) *AttachKeyPairInput {
	s.KeyPairName = &v
	return s
}

type AttachKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForAttachKeyPairOutput `type:"list"`
}

// String returns the string representation
func (s AttachKeyPairOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachKeyPairOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *AttachKeyPairOutput) SetOperationDetails(v []*OperationDetailForAttachKeyPairOutput) *AttachKeyPairOutput {
	s.OperationDetails = v
	return s
}

type ErrorForAttachKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForAttachKeyPairOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForAttachKeyPairOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForAttachKeyPairOutput) SetCode(v string) *ErrorForAttachKeyPairOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForAttachKeyPairOutput) SetMessage(v string) *ErrorForAttachKeyPairOutput {
	s.Message = &v
	return s
}

type OperationDetailForAttachKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForAttachKeyPairOutput `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForAttachKeyPairOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForAttachKeyPairOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForAttachKeyPairOutput) SetError(v *ErrorForAttachKeyPairOutput) *OperationDetailForAttachKeyPairOutput {
	s.Error = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *OperationDetailForAttachKeyPairOutput) SetInstanceId(v string) *OperationDetailForAttachKeyPairOutput {
	s.InstanceId = &v
	return s
}
