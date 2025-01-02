// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opUpdateUserCommon = "UpdateUser"

// UpdateUserCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the UpdateUserCommon operation. The "output" return
// value will be populated with the UpdateUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateUserCommon Send returns without error.
//
// See UpdateUserCommon for more information on using the UpdateUserCommon
// API call, and error handling.
//
//	// Example sending a request using the UpdateUserCommonRequest method.
//	req, resp := client.UpdateUserCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) UpdateUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateUserCommon,
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

// UpdateUserCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation UpdateUserCommon for usage and error information.
func (c *IAM) UpdateUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateUserCommonRequest(input)
	return out, req.Send()
}

// UpdateUserCommonWithContext is the same as UpdateUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateUserCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateUser = "UpdateUser"

// UpdateUserRequest generates a "byteplus/request.Request" representing the
// client's request for the UpdateUser operation. The "output" return
// value will be populated with the UpdateUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateUserCommon Send returns without error.
//
// See UpdateUser for more information on using the UpdateUser
// API call, and error handling.
//
//	// Example sending a request using the UpdateUserRequest method.
//	req, resp := client.UpdateUserRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) UpdateUserRequest(input *UpdateUserInput) (req *request.Request, output *UpdateUserOutput) {
	op := &request.Operation{
		Name:       opUpdateUser,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateUserInput{}
	}

	output = &UpdateUserOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateUser API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation UpdateUser for usage and error information.
func (c *IAM) UpdateUser(input *UpdateUserInput) (*UpdateUserOutput, error) {
	req, out := c.UpdateUserRequest(input)
	return out, req.Send()
}

// UpdateUserWithContext is the same as UpdateUser with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateUserWithContext(ctx byteplus.Context, input *UpdateUserInput, opts ...request.Option) (*UpdateUserOutput, error) {
	req, out := c.UpdateUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type TagForUpdateUserOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForUpdateUserOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForUpdateUserOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForUpdateUserOutput) SetKey(v string) *TagForUpdateUserOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForUpdateUserOutput) SetValue(v string) *TagForUpdateUserOutput {
	s.Value = &v
	return s
}

type UpdateUserInput struct {
	_ struct{} `type:"structure"`

	NewDescription *string `type:"string"`

	NewDisplayName *string `type:"string"`

	NewEmail *string `type:"string"`

	NewMobilePhone *string `type:"string"`

	NewUserName *string `type:"string"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateUserInput"}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNewDescription sets the NewDescription field's value.
func (s *UpdateUserInput) SetNewDescription(v string) *UpdateUserInput {
	s.NewDescription = &v
	return s
}

// SetNewDisplayName sets the NewDisplayName field's value.
func (s *UpdateUserInput) SetNewDisplayName(v string) *UpdateUserInput {
	s.NewDisplayName = &v
	return s
}

// SetNewEmail sets the NewEmail field's value.
func (s *UpdateUserInput) SetNewEmail(v string) *UpdateUserInput {
	s.NewEmail = &v
	return s
}

// SetNewMobilePhone sets the NewMobilePhone field's value.
func (s *UpdateUserInput) SetNewMobilePhone(v string) *UpdateUserInput {
	s.NewMobilePhone = &v
	return s
}

// SetNewUserName sets the NewUserName field's value.
func (s *UpdateUserInput) SetNewUserName(v string) *UpdateUserInput {
	s.NewUserName = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *UpdateUserInput) SetUserName(v string) *UpdateUserInput {
	s.UserName = &v
	return s
}

type UpdateUserOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	User *UserForUpdateUserOutput `type:"structure"`
}

// String returns the string representation
func (s UpdateUserOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateUserOutput) GoString() string {
	return s.String()
}

// SetUser sets the User field's value.
func (s *UpdateUserOutput) SetUser(v *UserForUpdateUserOutput) *UpdateUserOutput {
	s.User = v
	return s
}

type UserForUpdateUserOutput struct {
	_ struct{} `type:"structure"`

	AccountId *int64 `type:"int64"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	Email *string `type:"string"`

	EmailIsVerify *bool `type:"boolean"`

	Id *int32 `type:"int32"`

	MobilePhone *string `type:"string"`

	MobilePhoneIsVerify *bool `type:"boolean"`

	Tags []*TagForUpdateUserOutput `type:"list"`

	Trn *string `type:"string"`

	UpdateDate *string `type:"string"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s UserForUpdateUserOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UserForUpdateUserOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *UserForUpdateUserOutput) SetAccountId(v int64) *UserForUpdateUserOutput {
	s.AccountId = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *UserForUpdateUserOutput) SetCreateDate(v string) *UserForUpdateUserOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UserForUpdateUserOutput) SetDescription(v string) *UserForUpdateUserOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *UserForUpdateUserOutput) SetDisplayName(v string) *UserForUpdateUserOutput {
	s.DisplayName = &v
	return s
}

// SetEmail sets the Email field's value.
func (s *UserForUpdateUserOutput) SetEmail(v string) *UserForUpdateUserOutput {
	s.Email = &v
	return s
}

// SetEmailIsVerify sets the EmailIsVerify field's value.
func (s *UserForUpdateUserOutput) SetEmailIsVerify(v bool) *UserForUpdateUserOutput {
	s.EmailIsVerify = &v
	return s
}

// SetId sets the Id field's value.
func (s *UserForUpdateUserOutput) SetId(v int32) *UserForUpdateUserOutput {
	s.Id = &v
	return s
}

// SetMobilePhone sets the MobilePhone field's value.
func (s *UserForUpdateUserOutput) SetMobilePhone(v string) *UserForUpdateUserOutput {
	s.MobilePhone = &v
	return s
}

// SetMobilePhoneIsVerify sets the MobilePhoneIsVerify field's value.
func (s *UserForUpdateUserOutput) SetMobilePhoneIsVerify(v bool) *UserForUpdateUserOutput {
	s.MobilePhoneIsVerify = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *UserForUpdateUserOutput) SetTags(v []*TagForUpdateUserOutput) *UserForUpdateUserOutput {
	s.Tags = v
	return s
}

// SetTrn sets the Trn field's value.
func (s *UserForUpdateUserOutput) SetTrn(v string) *UserForUpdateUserOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *UserForUpdateUserOutput) SetUpdateDate(v string) *UserForUpdateUserOutput {
	s.UpdateDate = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *UserForUpdateUserOutput) SetUserName(v string) *UserForUpdateUserOutput {
	s.UserName = &v
	return s
}
