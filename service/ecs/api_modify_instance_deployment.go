// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyInstanceDeploymentCommon = "ModifyInstanceDeployment"

// ModifyInstanceDeploymentCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyInstanceDeploymentCommon operation. The "output" return
// value will be populated with the ModifyInstanceDeploymentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceDeploymentCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceDeploymentCommon Send returns without error.
//
// See ModifyInstanceDeploymentCommon for more information on using the ModifyInstanceDeploymentCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceDeploymentCommonRequest method.
//    req, resp := client.ModifyInstanceDeploymentCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceDeploymentCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyInstanceDeploymentCommon,
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

// ModifyInstanceDeploymentCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyInstanceDeploymentCommon for usage and error information.
func (c *ECS) ModifyInstanceDeploymentCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceDeploymentCommonRequest(input)
	return out, req.Send()
}

// ModifyInstanceDeploymentCommonWithContext is the same as ModifyInstanceDeploymentCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceDeploymentCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceDeploymentCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceDeploymentCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyInstanceDeployment = "ModifyInstanceDeployment"

// ModifyInstanceDeploymentRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyInstanceDeployment operation. The "output" return
// value will be populated with the ModifyInstanceDeploymentCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceDeploymentCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceDeploymentCommon Send returns without error.
//
// See ModifyInstanceDeployment for more information on using the ModifyInstanceDeployment
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceDeploymentRequest method.
//    req, resp := client.ModifyInstanceDeploymentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceDeploymentRequest(input *ModifyInstanceDeploymentInput) (req *request.Request, output *ModifyInstanceDeploymentOutput) {
	op := &request.Operation{
		Name:       opModifyInstanceDeployment,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceDeploymentInput{}
	}

	output = &ModifyInstanceDeploymentOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyInstanceDeployment API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyInstanceDeployment for usage and error information.
func (c *ECS) ModifyInstanceDeployment(input *ModifyInstanceDeploymentInput) (*ModifyInstanceDeploymentOutput, error) {
	req, out := c.ModifyInstanceDeploymentRequest(input)
	return out, req.Send()
}

// ModifyInstanceDeploymentWithContext is the same as ModifyInstanceDeployment with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceDeployment for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceDeploymentWithContext(ctx byteplus.Context, input *ModifyInstanceDeploymentInput, opts ...request.Option) (*ModifyInstanceDeploymentOutput, error) {
	req, out := c.ModifyInstanceDeploymentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyInstanceDeploymentInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	DeploymentSetGroupNumber *int32 `type:"int32"`

	// DeploymentSetId is a required field
	DeploymentSetId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyInstanceDeploymentInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceDeploymentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstanceDeploymentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyInstanceDeploymentInput"}
	if s.DeploymentSetId == nil {
		invalidParams.Add(request.NewErrParamRequired("DeploymentSetId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyInstanceDeploymentInput) SetClientToken(v string) *ModifyInstanceDeploymentInput {
	s.ClientToken = &v
	return s
}

// SetDeploymentSetGroupNumber sets the DeploymentSetGroupNumber field's value.
func (s *ModifyInstanceDeploymentInput) SetDeploymentSetGroupNumber(v int32) *ModifyInstanceDeploymentInput {
	s.DeploymentSetGroupNumber = &v
	return s
}

// SetDeploymentSetId sets the DeploymentSetId field's value.
func (s *ModifyInstanceDeploymentInput) SetDeploymentSetId(v string) *ModifyInstanceDeploymentInput {
	s.DeploymentSetId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyInstanceDeploymentInput) SetInstanceId(v string) *ModifyInstanceDeploymentInput {
	s.InstanceId = &v
	return s
}

type ModifyInstanceDeploymentOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyInstanceDeploymentOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceDeploymentOutput) GoString() string {
	return s.String()
}
