// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyReservedInstancesCommon = "ModifyReservedInstances"

// ModifyReservedInstancesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyReservedInstancesCommon operation. The "output" return
// value will be populated with the ModifyReservedInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyReservedInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyReservedInstancesCommon Send returns without error.
//
// See ModifyReservedInstancesCommon for more information on using the ModifyReservedInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyReservedInstancesCommonRequest method.
//    req, resp := client.ModifyReservedInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyReservedInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyReservedInstancesCommon,
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

// ModifyReservedInstancesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyReservedInstancesCommon for usage and error information.
func (c *ECS) ModifyReservedInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyReservedInstancesCommonRequest(input)
	return out, req.Send()
}

// ModifyReservedInstancesCommonWithContext is the same as ModifyReservedInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyReservedInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyReservedInstancesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyReservedInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyReservedInstances = "ModifyReservedInstances"

// ModifyReservedInstancesRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyReservedInstances operation. The "output" return
// value will be populated with the ModifyReservedInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyReservedInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyReservedInstancesCommon Send returns without error.
//
// See ModifyReservedInstances for more information on using the ModifyReservedInstances
// API call, and error handling.
//
//    // Example sending a request using the ModifyReservedInstancesRequest method.
//    req, resp := client.ModifyReservedInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyReservedInstancesRequest(input *ModifyReservedInstancesInput) (req *request.Request, output *ModifyReservedInstancesOutput) {
	op := &request.Operation{
		Name:       opModifyReservedInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyReservedInstancesInput{}
	}

	output = &ModifyReservedInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyReservedInstances API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation ModifyReservedInstances for usage and error information.
func (c *ECS) ModifyReservedInstances(input *ModifyReservedInstancesInput) (*ModifyReservedInstancesOutput, error) {
	req, out := c.ModifyReservedInstancesRequest(input)
	return out, req.Send()
}

// ModifyReservedInstancesWithContext is the same as ModifyReservedInstances with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyReservedInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyReservedInstancesWithContext(ctx byteplus.Context, input *ModifyReservedInstancesInput, opts ...request.Option) (*ModifyReservedInstancesOutput, error) {
	req, out := c.ModifyReservedInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConfigurationForModifyReservedInstancesInput struct {
	_ struct{} `type:"structure"`

	HpcClusterId *string `type:"string"`

	InstanceCount *int32 `type:"int32"`

	InstanceTypeId *string `type:"string"`

	// ReservedInstanceName is a required field
	ReservedInstanceName *string `type:"string" required:"true"`

	Scope *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ConfigurationForModifyReservedInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigurationForModifyReservedInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfigurationForModifyReservedInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ConfigurationForModifyReservedInstancesInput"}
	if s.ReservedInstanceName == nil {
		invalidParams.Add(request.NewErrParamRequired("ReservedInstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *ConfigurationForModifyReservedInstancesInput) SetHpcClusterId(v string) *ConfigurationForModifyReservedInstancesInput {
	s.HpcClusterId = &v
	return s
}

// SetInstanceCount sets the InstanceCount field's value.
func (s *ConfigurationForModifyReservedInstancesInput) SetInstanceCount(v int32) *ConfigurationForModifyReservedInstancesInput {
	s.InstanceCount = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *ConfigurationForModifyReservedInstancesInput) SetInstanceTypeId(v string) *ConfigurationForModifyReservedInstancesInput {
	s.InstanceTypeId = &v
	return s
}

// SetReservedInstanceName sets the ReservedInstanceName field's value.
func (s *ConfigurationForModifyReservedInstancesInput) SetReservedInstanceName(v string) *ConfigurationForModifyReservedInstancesInput {
	s.ReservedInstanceName = &v
	return s
}

// SetScope sets the Scope field's value.
func (s *ConfigurationForModifyReservedInstancesInput) SetScope(v string) *ConfigurationForModifyReservedInstancesInput {
	s.Scope = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ConfigurationForModifyReservedInstancesInput) SetZoneId(v string) *ConfigurationForModifyReservedInstancesInput {
	s.ZoneId = &v
	return s
}

type ModifyReservedInstancesInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	AutoRenewPeriod *int32 `type:"int32"`

	ClientToken *string `type:"string"`

	Configurations []*ConfigurationForModifyReservedInstancesInput `type:"list"`

	Description *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	ReservedInstanceIds []*string `type:"list"`

	Tags []*TagForModifyReservedInstancesInput `type:"list"`
}

// String returns the string representation
func (s ModifyReservedInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyReservedInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyReservedInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyReservedInstancesInput"}
	if s.Configurations != nil {
		for i, v := range s.Configurations {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Configurations", i), err.(request.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ModifyReservedInstancesInput) SetAutoRenew(v bool) *ModifyReservedInstancesInput {
	s.AutoRenew = &v
	return s
}

// SetAutoRenewPeriod sets the AutoRenewPeriod field's value.
func (s *ModifyReservedInstancesInput) SetAutoRenewPeriod(v int32) *ModifyReservedInstancesInput {
	s.AutoRenewPeriod = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyReservedInstancesInput) SetClientToken(v string) *ModifyReservedInstancesInput {
	s.ClientToken = &v
	return s
}

// SetConfigurations sets the Configurations field's value.
func (s *ModifyReservedInstancesInput) SetConfigurations(v []*ConfigurationForModifyReservedInstancesInput) *ModifyReservedInstancesInput {
	s.Configurations = v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyReservedInstancesInput) SetDescription(v string) *ModifyReservedInstancesInput {
	s.Description = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ModifyReservedInstancesInput) SetProjectName(v string) *ModifyReservedInstancesInput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *ModifyReservedInstancesInput) SetRegionId(v string) *ModifyReservedInstancesInput {
	s.RegionId = &v
	return s
}

// SetReservedInstanceIds sets the ReservedInstanceIds field's value.
func (s *ModifyReservedInstancesInput) SetReservedInstanceIds(v []*string) *ModifyReservedInstancesInput {
	s.ReservedInstanceIds = v
	return s
}

// SetTags sets the Tags field's value.
func (s *ModifyReservedInstancesInput) SetTags(v []*TagForModifyReservedInstancesInput) *ModifyReservedInstancesInput {
	s.Tags = v
	return s
}

type ModifyReservedInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ReservedInstanceIds []*string `type:"list"`
}

// String returns the string representation
func (s ModifyReservedInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyReservedInstancesOutput) GoString() string {
	return s.String()
}

// SetReservedInstanceIds sets the ReservedInstanceIds field's value.
func (s *ModifyReservedInstancesOutput) SetReservedInstanceIds(v []*string) *ModifyReservedInstancesOutput {
	s.ReservedInstanceIds = v
	return s
}

type TagForModifyReservedInstancesInput struct {
	_ struct{} `type:"structure"`

	// Key is a required field
	Key *string `type:"string" required:"true"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForModifyReservedInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForModifyReservedInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagForModifyReservedInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TagForModifyReservedInstancesInput"}
	if s.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKey sets the Key field's value.
func (s *TagForModifyReservedInstancesInput) SetKey(v string) *TagForModifyReservedInstancesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForModifyReservedInstancesInput) SetValue(v string) *TagForModifyReservedInstancesInput {
	s.Value = &v
	return s
}