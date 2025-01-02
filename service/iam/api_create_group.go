// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreateGroupCommon = "CreateGroup"

// CreateGroupCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateGroupCommon operation. The "output" return
// value will be populated with the CreateGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateGroupCommon Send returns without error.
//
// See CreateGroupCommon for more information on using the CreateGroupCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateGroupCommonRequest method.
//	req, resp := client.CreateGroupCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) CreateGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateGroupCommon,
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

// CreateGroupCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation CreateGroupCommon for usage and error information.
func (c *IAM) CreateGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateGroupCommonRequest(input)
	return out, req.Send()
}

// CreateGroupCommonWithContext is the same as CreateGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateGroupCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateGroup = "CreateGroup"

// CreateGroupRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateGroup operation. The "output" return
// value will be populated with the CreateGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateGroupCommon Send returns without error.
//
// See CreateGroup for more information on using the CreateGroup
// API call, and error handling.
//
//	// Example sending a request using the CreateGroupRequest method.
//	req, resp := client.CreateGroupRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) CreateGroupRequest(input *CreateGroupInput) (req *request.Request, output *CreateGroupOutput) {
	op := &request.Operation{
		Name:       opCreateGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGroupInput{}
	}

	output = &CreateGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateGroup API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation CreateGroup for usage and error information.
func (c *IAM) CreateGroup(input *CreateGroupInput) (*CreateGroupOutput, error) {
	req, out := c.CreateGroupRequest(input)
	return out, req.Send()
}

// CreateGroupWithContext is the same as CreateGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CreateGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateGroupWithContext(ctx byteplus.Context, input *CreateGroupInput, opts ...request.Option) (*CreateGroupOutput, error) {
	req, out := c.CreateGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateGroupInput struct {
	_ struct{} `type:"structure"`

	Description *string `max:"128" type:"string"`

	DisplayName *string `max:"64" type:"string"`

	// UserGroupName is a required field
	UserGroupName *string `max:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGroupInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateGroupInput"}
	if s.Description != nil && len(*s.Description) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 128, *s.Description))
	}
	if s.DisplayName != nil && len(*s.DisplayName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("DisplayName", 64, *s.DisplayName))
	}
	if s.UserGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserGroupName"))
	}
	if s.UserGroupName != nil && len(*s.UserGroupName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("UserGroupName", 64, *s.UserGroupName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateGroupInput) SetDescription(v string) *CreateGroupInput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *CreateGroupInput) SetDisplayName(v string) *CreateGroupInput {
	s.DisplayName = &v
	return s
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *CreateGroupInput) SetUserGroupName(v string) *CreateGroupInput {
	s.UserGroupName = &v
	return s
}

type CreateGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	UserGroup *UserGroupForCreateGroupOutput `type:"structure"`
}

// String returns the string representation
func (s CreateGroupOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateGroupOutput) GoString() string {
	return s.String()
}

// SetUserGroup sets the UserGroup field's value.
func (s *CreateGroupOutput) SetUserGroup(v *UserGroupForCreateGroupOutput) *CreateGroupOutput {
	s.UserGroup = v
	return s
}

type UserGroupForCreateGroupOutput struct {
	_ struct{} `type:"structure"`

	AccountID *int64 `type:"int64"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	UpdateDate *string `type:"string"`

	UserGroupID *int32 `type:"int32"`

	UserGroupName *string `type:"string"`
}

// String returns the string representation
func (s UserGroupForCreateGroupOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UserGroupForCreateGroupOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *UserGroupForCreateGroupOutput) SetAccountID(v int64) *UserGroupForCreateGroupOutput {
	s.AccountID = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *UserGroupForCreateGroupOutput) SetCreateDate(v string) *UserGroupForCreateGroupOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UserGroupForCreateGroupOutput) SetDescription(v string) *UserGroupForCreateGroupOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *UserGroupForCreateGroupOutput) SetDisplayName(v string) *UserGroupForCreateGroupOutput {
	s.DisplayName = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *UserGroupForCreateGroupOutput) SetUpdateDate(v string) *UserGroupForCreateGroupOutput {
	s.UpdateDate = &v
	return s
}

// SetUserGroupID sets the UserGroupID field's value.
func (s *UserGroupForCreateGroupOutput) SetUserGroupID(v int32) *UserGroupForCreateGroupOutput {
	s.UserGroupID = &v
	return s
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *UserGroupForCreateGroupOutput) SetUserGroupName(v string) *UserGroupForCreateGroupOutput {
	s.UserGroupName = &v
	return s
}
