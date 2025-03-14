// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opListAttachedUserGroupPoliciesCommon = "ListAttachedUserGroupPolicies"

// ListAttachedUserGroupPoliciesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ListAttachedUserGroupPoliciesCommon operation. The "output" return
// value will be populated with the ListAttachedUserGroupPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAttachedUserGroupPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAttachedUserGroupPoliciesCommon Send returns without error.
//
// See ListAttachedUserGroupPoliciesCommon for more information on using the ListAttachedUserGroupPoliciesCommon
// API call, and error handling.
//
//	// Example sending a request using the ListAttachedUserGroupPoliciesCommonRequest method.
//	req, resp := client.ListAttachedUserGroupPoliciesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) ListAttachedUserGroupPoliciesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAttachedUserGroupPoliciesCommon,
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

// ListAttachedUserGroupPoliciesCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation ListAttachedUserGroupPoliciesCommon for usage and error information.
func (c *IAM) ListAttachedUserGroupPoliciesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAttachedUserGroupPoliciesCommonRequest(input)
	return out, req.Send()
}

// ListAttachedUserGroupPoliciesCommonWithContext is the same as ListAttachedUserGroupPoliciesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAttachedUserGroupPoliciesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListAttachedUserGroupPoliciesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAttachedUserGroupPoliciesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAttachedUserGroupPolicies = "ListAttachedUserGroupPolicies"

// ListAttachedUserGroupPoliciesRequest generates a "byteplus/request.Request" representing the
// client's request for the ListAttachedUserGroupPolicies operation. The "output" return
// value will be populated with the ListAttachedUserGroupPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAttachedUserGroupPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAttachedUserGroupPoliciesCommon Send returns without error.
//
// See ListAttachedUserGroupPolicies for more information on using the ListAttachedUserGroupPolicies
// API call, and error handling.
//
//	// Example sending a request using the ListAttachedUserGroupPoliciesRequest method.
//	req, resp := client.ListAttachedUserGroupPoliciesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) ListAttachedUserGroupPoliciesRequest(input *ListAttachedUserGroupPoliciesInput) (req *request.Request, output *ListAttachedUserGroupPoliciesOutput) {
	op := &request.Operation{
		Name:       opListAttachedUserGroupPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAttachedUserGroupPoliciesInput{}
	}

	output = &ListAttachedUserGroupPoliciesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListAttachedUserGroupPolicies API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation ListAttachedUserGroupPolicies for usage and error information.
func (c *IAM) ListAttachedUserGroupPolicies(input *ListAttachedUserGroupPoliciesInput) (*ListAttachedUserGroupPoliciesOutput, error) {
	req, out := c.ListAttachedUserGroupPoliciesRequest(input)
	return out, req.Send()
}

// ListAttachedUserGroupPoliciesWithContext is the same as ListAttachedUserGroupPolicies with the addition of
// the ability to pass a context and additional request options.
//
// See ListAttachedUserGroupPolicies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListAttachedUserGroupPoliciesWithContext(ctx byteplus.Context, input *ListAttachedUserGroupPoliciesInput, opts ...request.Option) (*ListAttachedUserGroupPoliciesOutput, error) {
	req, out := c.ListAttachedUserGroupPoliciesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput struct {
	_ struct{} `type:"structure"`

	AttachDate *string `type:"string"`

	Description *string `type:"string"`

	PolicyName *string `type:"string"`

	PolicyScope []*PolicyScopeForListAttachedUserGroupPoliciesOutput `type:"list"`

	PolicyTrn *string `type:"string"`

	PolicyType *string `type:"string"`
}

// String returns the string representation
func (s AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) GoString() string {
	return s.String()
}

// SetAttachDate sets the AttachDate field's value.
func (s *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) SetAttachDate(v string) *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput {
	s.AttachDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) SetDescription(v string) *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput {
	s.Description = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) SetPolicyName(v string) *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput {
	s.PolicyName = &v
	return s
}

// SetPolicyScope sets the PolicyScope field's value.
func (s *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) SetPolicyScope(v []*PolicyScopeForListAttachedUserGroupPoliciesOutput) *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput {
	s.PolicyScope = v
	return s
}

// SetPolicyTrn sets the PolicyTrn field's value.
func (s *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) SetPolicyTrn(v string) *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput {
	s.PolicyTrn = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) SetPolicyType(v string) *AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput {
	s.PolicyType = &v
	return s
}

type ListAttachedUserGroupPoliciesInput struct {
	_ struct{} `type:"structure"`

	// UserGroupName is a required field
	UserGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListAttachedUserGroupPoliciesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAttachedUserGroupPoliciesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttachedUserGroupPoliciesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAttachedUserGroupPoliciesInput"}
	if s.UserGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *ListAttachedUserGroupPoliciesInput) SetUserGroupName(v string) *ListAttachedUserGroupPoliciesInput {
	s.UserGroupName = &v
	return s
}

type ListAttachedUserGroupPoliciesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AttachedPolicyMetadata []*AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput `type:"list"`
}

// String returns the string representation
func (s ListAttachedUserGroupPoliciesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAttachedUserGroupPoliciesOutput) GoString() string {
	return s.String()
}

// SetAttachedPolicyMetadata sets the AttachedPolicyMetadata field's value.
func (s *ListAttachedUserGroupPoliciesOutput) SetAttachedPolicyMetadata(v []*AttachedPolicyMetadataForListAttachedUserGroupPoliciesOutput) *ListAttachedUserGroupPoliciesOutput {
	s.AttachedPolicyMetadata = v
	return s
}

type PolicyScopeForListAttachedUserGroupPoliciesOutput struct {
	_ struct{} `type:"structure"`

	AttachDate *string `type:"string"`

	PolicyScopeType *string `type:"string"`

	ProjectDisplayName *string `type:"string"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s PolicyScopeForListAttachedUserGroupPoliciesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s PolicyScopeForListAttachedUserGroupPoliciesOutput) GoString() string {
	return s.String()
}

// SetAttachDate sets the AttachDate field's value.
func (s *PolicyScopeForListAttachedUserGroupPoliciesOutput) SetAttachDate(v string) *PolicyScopeForListAttachedUserGroupPoliciesOutput {
	s.AttachDate = &v
	return s
}

// SetPolicyScopeType sets the PolicyScopeType field's value.
func (s *PolicyScopeForListAttachedUserGroupPoliciesOutput) SetPolicyScopeType(v string) *PolicyScopeForListAttachedUserGroupPoliciesOutput {
	s.PolicyScopeType = &v
	return s
}

// SetProjectDisplayName sets the ProjectDisplayName field's value.
func (s *PolicyScopeForListAttachedUserGroupPoliciesOutput) SetProjectDisplayName(v string) *PolicyScopeForListAttachedUserGroupPoliciesOutput {
	s.ProjectDisplayName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *PolicyScopeForListAttachedUserGroupPoliciesOutput) SetProjectName(v string) *PolicyScopeForListAttachedUserGroupPoliciesOutput {
	s.ProjectName = &v
	return s
}
