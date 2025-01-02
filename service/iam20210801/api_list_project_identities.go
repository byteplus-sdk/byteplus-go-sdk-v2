// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam20210801

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opListProjectIdentitiesCommon = "ListProjectIdentities"

// ListProjectIdentitiesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ListProjectIdentitiesCommon operation. The "output" return
// value will be populated with the ListProjectIdentitiesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProjectIdentitiesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProjectIdentitiesCommon Send returns without error.
//
// See ListProjectIdentitiesCommon for more information on using the ListProjectIdentitiesCommon
// API call, and error handling.
//
//	// Example sending a request using the ListProjectIdentitiesCommonRequest method.
//	req, resp := client.ListProjectIdentitiesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM20210801) ListProjectIdentitiesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListProjectIdentitiesCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ListProjectIdentitiesCommon API operation for IAM20210801.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM20210801's
// API operation ListProjectIdentitiesCommon for usage and error information.
func (c *IAM20210801) ListProjectIdentitiesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListProjectIdentitiesCommonRequest(input)
	return out, req.Send()
}

// ListProjectIdentitiesCommonWithContext is the same as ListProjectIdentitiesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListProjectIdentitiesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) ListProjectIdentitiesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListProjectIdentitiesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListProjectIdentities = "ListProjectIdentities"

// ListProjectIdentitiesRequest generates a "byteplus/request.Request" representing the
// client's request for the ListProjectIdentities operation. The "output" return
// value will be populated with the ListProjectIdentitiesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProjectIdentitiesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProjectIdentitiesCommon Send returns without error.
//
// See ListProjectIdentities for more information on using the ListProjectIdentities
// API call, and error handling.
//
//	// Example sending a request using the ListProjectIdentitiesRequest method.
//	req, resp := client.ListProjectIdentitiesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM20210801) ListProjectIdentitiesRequest(input *ListProjectIdentitiesInput) (req *request.Request, output *ListProjectIdentitiesOutput) {
	op := &request.Operation{
		Name:       opListProjectIdentities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProjectIdentitiesInput{}
	}

	output = &ListProjectIdentitiesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListProjectIdentities API operation for IAM20210801.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM20210801's
// API operation ListProjectIdentities for usage and error information.
func (c *IAM20210801) ListProjectIdentities(input *ListProjectIdentitiesInput) (*ListProjectIdentitiesOutput, error) {
	req, out := c.ListProjectIdentitiesRequest(input)
	return out, req.Send()
}

// ListProjectIdentitiesWithContext is the same as ListProjectIdentities with the addition of
// the ability to pass a context and additional request options.
//
// See ListProjectIdentities for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) ListProjectIdentitiesWithContext(ctx byteplus.Context, input *ListProjectIdentitiesInput, opts ...request.Option) (*ListProjectIdentitiesOutput, error) {
	req, out := c.ListProjectIdentitiesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListProjectIdentitiesInput struct {
	_ struct{} `type:"structure"`

	// IdentityType is a required field
	IdentityType *string `type:"string" required:"true" enum:"EnumOfIdentityTypeForListProjectIdentitiesInput"`

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	// ProjectName is a required field
	ProjectName *string `type:"string" required:"true"`

	Query *string `type:"string"`
}

// String returns the string representation
func (s ListProjectIdentitiesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProjectIdentitiesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProjectIdentitiesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListProjectIdentitiesInput"}
	if s.IdentityType == nil {
		invalidParams.Add(request.NewErrParamRequired("IdentityType"))
	}
	if s.ProjectName == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIdentityType sets the IdentityType field's value.
func (s *ListProjectIdentitiesInput) SetIdentityType(v string) *ListProjectIdentitiesInput {
	s.IdentityType = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListProjectIdentitiesInput) SetLimit(v int32) *ListProjectIdentitiesInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListProjectIdentitiesInput) SetOffset(v int32) *ListProjectIdentitiesInput {
	s.Offset = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListProjectIdentitiesInput) SetProjectName(v string) *ListProjectIdentitiesInput {
	s.ProjectName = &v
	return s
}

// SetQuery sets the Query field's value.
func (s *ListProjectIdentitiesInput) SetQuery(v string) *ListProjectIdentitiesInput {
	s.Query = &v
	return s
}

type ListProjectIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	ProjectRoles []*ProjectRoleForListProjectIdentitiesOutput `type:"list"`

	ProjectUserGroups []*ProjectUserGroupForListProjectIdentitiesOutput `type:"list"`

	ProjectUsers []*ProjectUserForListProjectIdentitiesOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListProjectIdentitiesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProjectIdentitiesOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListProjectIdentitiesOutput) SetLimit(v int32) *ListProjectIdentitiesOutput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListProjectIdentitiesOutput) SetOffset(v int32) *ListProjectIdentitiesOutput {
	s.Offset = &v
	return s
}

// SetProjectRoles sets the ProjectRoles field's value.
func (s *ListProjectIdentitiesOutput) SetProjectRoles(v []*ProjectRoleForListProjectIdentitiesOutput) *ListProjectIdentitiesOutput {
	s.ProjectRoles = v
	return s
}

// SetProjectUserGroups sets the ProjectUserGroups field's value.
func (s *ListProjectIdentitiesOutput) SetProjectUserGroups(v []*ProjectUserGroupForListProjectIdentitiesOutput) *ListProjectIdentitiesOutput {
	s.ProjectUserGroups = v
	return s
}

// SetProjectUsers sets the ProjectUsers field's value.
func (s *ListProjectIdentitiesOutput) SetProjectUsers(v []*ProjectUserForListProjectIdentitiesOutput) *ListProjectIdentitiesOutput {
	s.ProjectUsers = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListProjectIdentitiesOutput) SetTotal(v int32) *ListProjectIdentitiesOutput {
	s.Total = &v
	return s
}

type PolicyForListProjectIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	AttachDate *string `type:"string"`

	Description *string `type:"string"`

	IsServiceRolePolicy *int32 `type:"int32"`

	PolicyName *string `type:"string"`

	PolicyTrn *string `type:"string"`

	PolicyType *string `type:"string"`
}

// String returns the string representation
func (s PolicyForListProjectIdentitiesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s PolicyForListProjectIdentitiesOutput) GoString() string {
	return s.String()
}

// SetAttachDate sets the AttachDate field's value.
func (s *PolicyForListProjectIdentitiesOutput) SetAttachDate(v string) *PolicyForListProjectIdentitiesOutput {
	s.AttachDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *PolicyForListProjectIdentitiesOutput) SetDescription(v string) *PolicyForListProjectIdentitiesOutput {
	s.Description = &v
	return s
}

// SetIsServiceRolePolicy sets the IsServiceRolePolicy field's value.
func (s *PolicyForListProjectIdentitiesOutput) SetIsServiceRolePolicy(v int32) *PolicyForListProjectIdentitiesOutput {
	s.IsServiceRolePolicy = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *PolicyForListProjectIdentitiesOutput) SetPolicyName(v string) *PolicyForListProjectIdentitiesOutput {
	s.PolicyName = &v
	return s
}

// SetPolicyTrn sets the PolicyTrn field's value.
func (s *PolicyForListProjectIdentitiesOutput) SetPolicyTrn(v string) *PolicyForListProjectIdentitiesOutput {
	s.PolicyTrn = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *PolicyForListProjectIdentitiesOutput) SetPolicyType(v string) *PolicyForListProjectIdentitiesOutput {
	s.PolicyType = &v
	return s
}

type ProjectRoleForListProjectIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	Policy []*PolicyForListProjectIdentitiesOutput `type:"list"`

	RoleName *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s ProjectRoleForListProjectIdentitiesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ProjectRoleForListProjectIdentitiesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ProjectRoleForListProjectIdentitiesOutput) SetDescription(v string) *ProjectRoleForListProjectIdentitiesOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *ProjectRoleForListProjectIdentitiesOutput) SetDisplayName(v string) *ProjectRoleForListProjectIdentitiesOutput {
	s.DisplayName = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *ProjectRoleForListProjectIdentitiesOutput) SetPolicy(v []*PolicyForListProjectIdentitiesOutput) *ProjectRoleForListProjectIdentitiesOutput {
	s.Policy = v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *ProjectRoleForListProjectIdentitiesOutput) SetRoleName(v string) *ProjectRoleForListProjectIdentitiesOutput {
	s.RoleName = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *ProjectRoleForListProjectIdentitiesOutput) SetUpdateDate(v string) *ProjectRoleForListProjectIdentitiesOutput {
	s.UpdateDate = &v
	return s
}

type ProjectUserForListProjectIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	Policy []*PolicyForListProjectIdentitiesOutput `type:"list"`

	UpdateDate *string `type:"string"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s ProjectUserForListProjectIdentitiesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ProjectUserForListProjectIdentitiesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ProjectUserForListProjectIdentitiesOutput) SetDescription(v string) *ProjectUserForListProjectIdentitiesOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *ProjectUserForListProjectIdentitiesOutput) SetDisplayName(v string) *ProjectUserForListProjectIdentitiesOutput {
	s.DisplayName = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *ProjectUserForListProjectIdentitiesOutput) SetPolicy(v []*PolicyForListProjectIdentitiesOutput) *ProjectUserForListProjectIdentitiesOutput {
	s.Policy = v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *ProjectUserForListProjectIdentitiesOutput) SetUpdateDate(v string) *ProjectUserForListProjectIdentitiesOutput {
	s.UpdateDate = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *ProjectUserForListProjectIdentitiesOutput) SetUserName(v string) *ProjectUserForListProjectIdentitiesOutput {
	s.UserName = &v
	return s
}

type ProjectUserGroupForListProjectIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	Policy []*PolicyForListProjectIdentitiesOutput `type:"list"`

	UpdateDate *string `type:"string"`

	UserGroupName *string `type:"string"`
}

// String returns the string representation
func (s ProjectUserGroupForListProjectIdentitiesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ProjectUserGroupForListProjectIdentitiesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ProjectUserGroupForListProjectIdentitiesOutput) SetDescription(v string) *ProjectUserGroupForListProjectIdentitiesOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *ProjectUserGroupForListProjectIdentitiesOutput) SetDisplayName(v string) *ProjectUserGroupForListProjectIdentitiesOutput {
	s.DisplayName = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *ProjectUserGroupForListProjectIdentitiesOutput) SetPolicy(v []*PolicyForListProjectIdentitiesOutput) *ProjectUserGroupForListProjectIdentitiesOutput {
	s.Policy = v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *ProjectUserGroupForListProjectIdentitiesOutput) SetUpdateDate(v string) *ProjectUserGroupForListProjectIdentitiesOutput {
	s.UpdateDate = &v
	return s
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *ProjectUserGroupForListProjectIdentitiesOutput) SetUserGroupName(v string) *ProjectUserGroupForListProjectIdentitiesOutput {
	s.UserGroupName = &v
	return s
}

const (
	// EnumOfIdentityTypeForListProjectIdentitiesInputUser is a EnumOfIdentityTypeForListProjectIdentitiesInput enum value
	EnumOfIdentityTypeForListProjectIdentitiesInputUser = "User"

	// EnumOfIdentityTypeForListProjectIdentitiesInputRole is a EnumOfIdentityTypeForListProjectIdentitiesInput enum value
	EnumOfIdentityTypeForListProjectIdentitiesInputRole = "Role"

	// EnumOfIdentityTypeForListProjectIdentitiesInputUserGroup is a EnumOfIdentityTypeForListProjectIdentitiesInput enum value
	EnumOfIdentityTypeForListProjectIdentitiesInputUserGroup = "UserGroup"
)