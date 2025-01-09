// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeCommandsCommon = "DescribeCommands"

// DescribeCommandsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeCommandsCommon operation. The "output" return
// value will be populated with the DescribeCommandsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCommandsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCommandsCommon Send returns without error.
//
// See DescribeCommandsCommon for more information on using the DescribeCommandsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCommandsCommonRequest method.
//    req, resp := client.DescribeCommandsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeCommandsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCommandsCommon,
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

// DescribeCommandsCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeCommandsCommon for usage and error information.
func (c *ECS) DescribeCommandsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCommandsCommonRequest(input)
	return out, req.Send()
}

// DescribeCommandsCommonWithContext is the same as DescribeCommandsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCommandsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeCommandsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCommandsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCommands = "DescribeCommands"

// DescribeCommandsRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeCommands operation. The "output" return
// value will be populated with the DescribeCommandsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCommandsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCommandsCommon Send returns without error.
//
// See DescribeCommands for more information on using the DescribeCommands
// API call, and error handling.
//
//    // Example sending a request using the DescribeCommandsRequest method.
//    req, resp := client.DescribeCommandsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeCommandsRequest(input *DescribeCommandsInput) (req *request.Request, output *DescribeCommandsOutput) {
	op := &request.Operation{
		Name:       opDescribeCommands,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCommandsInput{}
	}

	output = &DescribeCommandsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCommands API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeCommands for usage and error information.
func (c *ECS) DescribeCommands(input *DescribeCommandsInput) (*DescribeCommandsOutput, error) {
	req, out := c.DescribeCommandsRequest(input)
	return out, req.Send()
}

// DescribeCommandsWithContext is the same as DescribeCommands with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCommands for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeCommandsWithContext(ctx byteplus.Context, input *DescribeCommandsInput, opts ...request.Option) (*DescribeCommandsOutput, error) {
	req, out := c.DescribeCommandsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CommandForDescribeCommandsOutput struct {
	_ struct{} `type:"structure"`

	CommandContent *string `type:"string"`

	CommandId *string `type:"string"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	EnableParameter *bool `type:"boolean"`

	InvocationTimes *int32 `type:"int32"`

	Name *string `type:"string"`

	ParameterDefinitions []*ParameterDefinitionForDescribeCommandsOutput `type:"list"`

	ProjectName *string `type:"string"`

	Provider *string `type:"string"`

	Tags []*TagForDescribeCommandsOutput `type:"list"`

	Timeout *int32 `type:"int32"`

	Type *string `type:"string"`

	UpdatedAt *string `type:"string"`

	Username *string `type:"string"`

	WorkingDir *string `type:"string"`
}

// String returns the string representation
func (s CommandForDescribeCommandsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CommandForDescribeCommandsOutput) GoString() string {
	return s.String()
}

// SetCommandContent sets the CommandContent field's value.
func (s *CommandForDescribeCommandsOutput) SetCommandContent(v string) *CommandForDescribeCommandsOutput {
	s.CommandContent = &v
	return s
}

// SetCommandId sets the CommandId field's value.
func (s *CommandForDescribeCommandsOutput) SetCommandId(v string) *CommandForDescribeCommandsOutput {
	s.CommandId = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *CommandForDescribeCommandsOutput) SetCreatedAt(v string) *CommandForDescribeCommandsOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CommandForDescribeCommandsOutput) SetDescription(v string) *CommandForDescribeCommandsOutput {
	s.Description = &v
	return s
}

// SetEnableParameter sets the EnableParameter field's value.
func (s *CommandForDescribeCommandsOutput) SetEnableParameter(v bool) *CommandForDescribeCommandsOutput {
	s.EnableParameter = &v
	return s
}

// SetInvocationTimes sets the InvocationTimes field's value.
func (s *CommandForDescribeCommandsOutput) SetInvocationTimes(v int32) *CommandForDescribeCommandsOutput {
	s.InvocationTimes = &v
	return s
}

// SetName sets the Name field's value.
func (s *CommandForDescribeCommandsOutput) SetName(v string) *CommandForDescribeCommandsOutput {
	s.Name = &v
	return s
}

// SetParameterDefinitions sets the ParameterDefinitions field's value.
func (s *CommandForDescribeCommandsOutput) SetParameterDefinitions(v []*ParameterDefinitionForDescribeCommandsOutput) *CommandForDescribeCommandsOutput {
	s.ParameterDefinitions = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CommandForDescribeCommandsOutput) SetProjectName(v string) *CommandForDescribeCommandsOutput {
	s.ProjectName = &v
	return s
}

// SetProvider sets the Provider field's value.
func (s *CommandForDescribeCommandsOutput) SetProvider(v string) *CommandForDescribeCommandsOutput {
	s.Provider = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CommandForDescribeCommandsOutput) SetTags(v []*TagForDescribeCommandsOutput) *CommandForDescribeCommandsOutput {
	s.Tags = v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *CommandForDescribeCommandsOutput) SetTimeout(v int32) *CommandForDescribeCommandsOutput {
	s.Timeout = &v
	return s
}

// SetType sets the Type field's value.
func (s *CommandForDescribeCommandsOutput) SetType(v string) *CommandForDescribeCommandsOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *CommandForDescribeCommandsOutput) SetUpdatedAt(v string) *CommandForDescribeCommandsOutput {
	s.UpdatedAt = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *CommandForDescribeCommandsOutput) SetUsername(v string) *CommandForDescribeCommandsOutput {
	s.Username = &v
	return s
}

// SetWorkingDir sets the WorkingDir field's value.
func (s *CommandForDescribeCommandsOutput) SetWorkingDir(v string) *CommandForDescribeCommandsOutput {
	s.WorkingDir = &v
	return s
}

type DescribeCommandsInput struct {
	_ struct{} `type:"structure"`

	CommandId *string `type:"string"`

	ContentEncoding *string `type:"string"`

	Name *string `type:"string"`

	Order *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	Provider *string `type:"string"`

	TagFilters []*TagFilterForDescribeCommandsInput `type:"list"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s DescribeCommandsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCommandsInput) GoString() string {
	return s.String()
}

// SetCommandId sets the CommandId field's value.
func (s *DescribeCommandsInput) SetCommandId(v string) *DescribeCommandsInput {
	s.CommandId = &v
	return s
}

// SetContentEncoding sets the ContentEncoding field's value.
func (s *DescribeCommandsInput) SetContentEncoding(v string) *DescribeCommandsInput {
	s.ContentEncoding = &v
	return s
}

// SetName sets the Name field's value.
func (s *DescribeCommandsInput) SetName(v string) *DescribeCommandsInput {
	s.Name = &v
	return s
}

// SetOrder sets the Order field's value.
func (s *DescribeCommandsInput) SetOrder(v string) *DescribeCommandsInput {
	s.Order = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCommandsInput) SetPageNumber(v int32) *DescribeCommandsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCommandsInput) SetPageSize(v int32) *DescribeCommandsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeCommandsInput) SetProjectName(v string) *DescribeCommandsInput {
	s.ProjectName = &v
	return s
}

// SetProvider sets the Provider field's value.
func (s *DescribeCommandsInput) SetProvider(v string) *DescribeCommandsInput {
	s.Provider = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeCommandsInput) SetTagFilters(v []*TagFilterForDescribeCommandsInput) *DescribeCommandsInput {
	s.TagFilters = v
	return s
}

// SetType sets the Type field's value.
func (s *DescribeCommandsInput) SetType(v string) *DescribeCommandsInput {
	s.Type = &v
	return s
}

type DescribeCommandsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Commands []*CommandForDescribeCommandsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeCommandsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCommandsOutput) GoString() string {
	return s.String()
}

// SetCommands sets the Commands field's value.
func (s *DescribeCommandsOutput) SetCommands(v []*CommandForDescribeCommandsOutput) *DescribeCommandsOutput {
	s.Commands = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCommandsOutput) SetPageNumber(v int32) *DescribeCommandsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCommandsOutput) SetPageSize(v int32) *DescribeCommandsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCommandsOutput) SetTotalCount(v int32) *DescribeCommandsOutput {
	s.TotalCount = &v
	return s
}

type ParameterDefinitionForDescribeCommandsOutput struct {
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
func (s ParameterDefinitionForDescribeCommandsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ParameterDefinitionForDescribeCommandsOutput) GoString() string {
	return s.String()
}

// SetDecimalPrecision sets the DecimalPrecision field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetDecimalPrecision(v int32) *ParameterDefinitionForDescribeCommandsOutput {
	s.DecimalPrecision = &v
	return s
}

// SetDefaultValue sets the DefaultValue field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetDefaultValue(v string) *ParameterDefinitionForDescribeCommandsOutput {
	s.DefaultValue = &v
	return s
}

// SetMaxLength sets the MaxLength field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetMaxLength(v int32) *ParameterDefinitionForDescribeCommandsOutput {
	s.MaxLength = &v
	return s
}

// SetMaxValue sets the MaxValue field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetMaxValue(v string) *ParameterDefinitionForDescribeCommandsOutput {
	s.MaxValue = &v
	return s
}

// SetMinLength sets the MinLength field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetMinLength(v int32) *ParameterDefinitionForDescribeCommandsOutput {
	s.MinLength = &v
	return s
}

// SetMinValue sets the MinValue field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetMinValue(v string) *ParameterDefinitionForDescribeCommandsOutput {
	s.MinValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetName(v string) *ParameterDefinitionForDescribeCommandsOutput {
	s.Name = &v
	return s
}

// SetRequired sets the Required field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetRequired(v bool) *ParameterDefinitionForDescribeCommandsOutput {
	s.Required = &v
	return s
}

// SetType sets the Type field's value.
func (s *ParameterDefinitionForDescribeCommandsOutput) SetType(v string) *ParameterDefinitionForDescribeCommandsOutput {
	s.Type = &v
	return s
}

type TagFilterForDescribeCommandsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeCommandsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeCommandsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeCommandsInput) SetKey(v string) *TagFilterForDescribeCommandsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeCommandsInput) SetValues(v []*string) *TagFilterForDescribeCommandsInput {
	s.Values = v
	return s
}

type TagForDescribeCommandsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeCommandsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeCommandsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeCommandsOutput) SetKey(v string) *TagForDescribeCommandsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeCommandsOutput) SetValue(v string) *TagForDescribeCommandsOutput {
	s.Value = &v
	return s
}
