// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyCommandCommon = "ModifyCommand"

// ModifyCommandCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyCommandCommon operation. The "output" return
// value will be populated with the ModifyCommandCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCommandCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCommandCommon Send returns without error.
//
// See ModifyCommandCommon for more information on using the ModifyCommandCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyCommandCommonRequest method.
//    req, resp := client.ModifyCommandCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyCommandCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyCommandCommon,
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

// ModifyCommandCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyCommandCommon for usage and error information.
func (c *ECS) ModifyCommandCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyCommandCommonRequest(input)
	return out, req.Send()
}

// ModifyCommandCommonWithContext is the same as ModifyCommandCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCommandCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyCommandCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyCommandCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyCommand = "ModifyCommand"

// ModifyCommandRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyCommand operation. The "output" return
// value will be populated with the ModifyCommandCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCommandCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCommandCommon Send returns without error.
//
// See ModifyCommand for more information on using the ModifyCommand
// API call, and error handling.
//
//    // Example sending a request using the ModifyCommandRequest method.
//    req, resp := client.ModifyCommandRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyCommandRequest(input *ModifyCommandInput) (req *request.Request, output *ModifyCommandOutput) {
	op := &request.Operation{
		Name:       opModifyCommand,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyCommandInput{}
	}

	output = &ModifyCommandOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyCommand API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyCommand for usage and error information.
func (c *ECS) ModifyCommand(input *ModifyCommandInput) (*ModifyCommandOutput, error) {
	req, out := c.ModifyCommandRequest(input)
	return out, req.Send()
}

// ModifyCommandWithContext is the same as ModifyCommand with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCommand for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyCommandWithContext(ctx byteplus.Context, input *ModifyCommandInput, opts ...request.Option) (*ModifyCommandOutput, error) {
	req, out := c.ModifyCommandRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyCommandInput struct {
	_ struct{} `type:"structure"`

	CommandContent *string `type:"string"`

	CommandId *string `type:"string"`

	ContentEncoding *string `type:"string"`

	Description *string `type:"string"`

	EnableParameter *bool `type:"boolean"`

	Name *string `type:"string"`

	ParameterDefinitions []*ParameterDefinitionForModifyCommandInput `type:"list"`

	Timeout *int32 `type:"int32"`

	Type *string `type:"string"`

	Username *string `type:"string"`

	WorkingDir *string `type:"string"`
}

// String returns the string representation
func (s ModifyCommandInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCommandInput) GoString() string {
	return s.String()
}

// SetCommandContent sets the CommandContent field's value.
func (s *ModifyCommandInput) SetCommandContent(v string) *ModifyCommandInput {
	s.CommandContent = &v
	return s
}

// SetCommandId sets the CommandId field's value.
func (s *ModifyCommandInput) SetCommandId(v string) *ModifyCommandInput {
	s.CommandId = &v
	return s
}

// SetContentEncoding sets the ContentEncoding field's value.
func (s *ModifyCommandInput) SetContentEncoding(v string) *ModifyCommandInput {
	s.ContentEncoding = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyCommandInput) SetDescription(v string) *ModifyCommandInput {
	s.Description = &v
	return s
}

// SetEnableParameter sets the EnableParameter field's value.
func (s *ModifyCommandInput) SetEnableParameter(v bool) *ModifyCommandInput {
	s.EnableParameter = &v
	return s
}

// SetName sets the Name field's value.
func (s *ModifyCommandInput) SetName(v string) *ModifyCommandInput {
	s.Name = &v
	return s
}

// SetParameterDefinitions sets the ParameterDefinitions field's value.
func (s *ModifyCommandInput) SetParameterDefinitions(v []*ParameterDefinitionForModifyCommandInput) *ModifyCommandInput {
	s.ParameterDefinitions = v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *ModifyCommandInput) SetTimeout(v int32) *ModifyCommandInput {
	s.Timeout = &v
	return s
}

// SetType sets the Type field's value.
func (s *ModifyCommandInput) SetType(v string) *ModifyCommandInput {
	s.Type = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *ModifyCommandInput) SetUsername(v string) *ModifyCommandInput {
	s.Username = &v
	return s
}

// SetWorkingDir sets the WorkingDir field's value.
func (s *ModifyCommandInput) SetWorkingDir(v string) *ModifyCommandInput {
	s.WorkingDir = &v
	return s
}

type ModifyCommandOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CommandId *string `type:"string"`
}

// String returns the string representation
func (s ModifyCommandOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCommandOutput) GoString() string {
	return s.String()
}

// SetCommandId sets the CommandId field's value.
func (s *ModifyCommandOutput) SetCommandId(v string) *ModifyCommandOutput {
	s.CommandId = &v
	return s
}

type ParameterDefinitionForModifyCommandInput struct {
	_ struct{} `type:"structure"`

	DecimalPrecision *int32 `type:"int32"`

	DefaultValue *string `type:"string"`

	MaxLength *int32 `type:"int32"`

	MaxValue *string `type:"string"`

	MinLength *int32 `type:"int32"`

	MinValue *string `type:"string"`

	Name *string `type:"string"`

	Required *bool `type:"boolean"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ParameterDefinitionForModifyCommandInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ParameterDefinitionForModifyCommandInput) GoString() string {
	return s.String()
}

// SetDecimalPrecision sets the DecimalPrecision field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetDecimalPrecision(v int32) *ParameterDefinitionForModifyCommandInput {
	s.DecimalPrecision = &v
	return s
}

// SetDefaultValue sets the DefaultValue field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetDefaultValue(v string) *ParameterDefinitionForModifyCommandInput {
	s.DefaultValue = &v
	return s
}

// SetMaxLength sets the MaxLength field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetMaxLength(v int32) *ParameterDefinitionForModifyCommandInput {
	s.MaxLength = &v
	return s
}

// SetMaxValue sets the MaxValue field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetMaxValue(v string) *ParameterDefinitionForModifyCommandInput {
	s.MaxValue = &v
	return s
}

// SetMinLength sets the MinLength field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetMinLength(v int32) *ParameterDefinitionForModifyCommandInput {
	s.MinLength = &v
	return s
}

// SetMinValue sets the MinValue field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetMinValue(v string) *ParameterDefinitionForModifyCommandInput {
	s.MinValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetName(v string) *ParameterDefinitionForModifyCommandInput {
	s.Name = &v
	return s
}

// SetRequired sets the Required field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetRequired(v bool) *ParameterDefinitionForModifyCommandInput {
	s.Required = &v
	return s
}

// SetType sets the Type field's value.
func (s *ParameterDefinitionForModifyCommandInput) SetType(v string) *ParameterDefinitionForModifyCommandInput {
	s.Type = &v
	return s
}
