// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opListAttachedUserPoliciesCommon = "ListAttachedUserPolicies"

// ListAttachedUserPoliciesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ListAttachedUserPoliciesCommon operation. The "output" return
// value will be populated with the ListAttachedUserPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAttachedUserPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAttachedUserPoliciesCommon Send returns without error.
//
// See ListAttachedUserPoliciesCommon for more information on using the ListAttachedUserPoliciesCommon
// API call, and error handling.
//
//	// Example sending a request using the ListAttachedUserPoliciesCommonRequest method.
//	req, resp := client.ListAttachedUserPoliciesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) ListAttachedUserPoliciesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAttachedUserPoliciesCommon,
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

// ListAttachedUserPoliciesCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation ListAttachedUserPoliciesCommon for usage and error information.
func (c *IAM) ListAttachedUserPoliciesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAttachedUserPoliciesCommonRequest(input)
	return out, req.Send()
}

// ListAttachedUserPoliciesCommonWithContext is the same as ListAttachedUserPoliciesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAttachedUserPoliciesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListAttachedUserPoliciesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAttachedUserPoliciesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAttachedUserPolicies = "ListAttachedUserPolicies"

// ListAttachedUserPoliciesRequest generates a "byteplus/request.Request" representing the
// client's request for the ListAttachedUserPolicies operation. The "output" return
// value will be populated with the ListAttachedUserPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAttachedUserPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAttachedUserPoliciesCommon Send returns without error.
//
// See ListAttachedUserPolicies for more information on using the ListAttachedUserPolicies
// API call, and error handling.
//
//	// Example sending a request using the ListAttachedUserPoliciesRequest method.
//	req, resp := client.ListAttachedUserPoliciesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) ListAttachedUserPoliciesRequest(input *ListAttachedUserPoliciesInput) (req *request.Request, output *ListAttachedUserPoliciesOutput) {
	op := &request.Operation{
		Name:       opListAttachedUserPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAttachedUserPoliciesInput{}
	}

	output = &ListAttachedUserPoliciesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListAttachedUserPolicies API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation ListAttachedUserPolicies for usage and error information.
func (c *IAM) ListAttachedUserPolicies(input *ListAttachedUserPoliciesInput) (*ListAttachedUserPoliciesOutput, error) {
	req, out := c.ListAttachedUserPoliciesRequest(input)
	return out, req.Send()
}

// ListAttachedUserPoliciesWithContext is the same as ListAttachedUserPolicies with the addition of
// the ability to pass a context and additional request options.
//
// See ListAttachedUserPolicies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListAttachedUserPoliciesWithContext(ctx byteplus.Context, input *ListAttachedUserPoliciesInput, opts ...request.Option) (*ListAttachedUserPoliciesOutput, error) {
	req, out := c.ListAttachedUserPoliciesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachedPolicyMetadataForListAttachedUserPoliciesOutput struct {
	_ struct{} `type:"structure"`

	AttachDate *string `type:"string"`

	Description *string `type:"string"`

	PolicyName *string `type:"string"`

	PolicyScope []*PolicyScopeForListAttachedUserPoliciesOutput `type:"list"`

	PolicyTrn *string `type:"string"`

	PolicyType *string `type:"string"`
}

// String returns the string representation
func (s AttachedPolicyMetadataForListAttachedUserPoliciesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachedPolicyMetadataForListAttachedUserPoliciesOutput) GoString() string {
	return s.String()
}

// SetAttachDate sets the AttachDate field's value.
func (s *AttachedPolicyMetadataForListAttachedUserPoliciesOutput) SetAttachDate(v string) *AttachedPolicyMetadataForListAttachedUserPoliciesOutput {
	s.AttachDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AttachedPolicyMetadataForListAttachedUserPoliciesOutput) SetDescription(v string) *AttachedPolicyMetadataForListAttachedUserPoliciesOutput {
	s.Description = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *AttachedPolicyMetadataForListAttachedUserPoliciesOutput) SetPolicyName(v string) *AttachedPolicyMetadataForListAttachedUserPoliciesOutput {
	s.PolicyName = &v
	return s
}

// SetPolicyScope sets the PolicyScope field's value.
func (s *AttachedPolicyMetadataForListAttachedUserPoliciesOutput) SetPolicyScope(v []*PolicyScopeForListAttachedUserPoliciesOutput) *AttachedPolicyMetadataForListAttachedUserPoliciesOutput {
	s.PolicyScope = v
	return s
}

// SetPolicyTrn sets the PolicyTrn field's value.
func (s *AttachedPolicyMetadataForListAttachedUserPoliciesOutput) SetPolicyTrn(v string) *AttachedPolicyMetadataForListAttachedUserPoliciesOutput {
	s.PolicyTrn = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *AttachedPolicyMetadataForListAttachedUserPoliciesOutput) SetPolicyType(v string) *AttachedPolicyMetadataForListAttachedUserPoliciesOutput {
	s.PolicyType = &v
	return s
}

type ListAttachedUserPoliciesInput struct {
	_ struct{} `type:"structure"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListAttachedUserPoliciesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAttachedUserPoliciesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttachedUserPoliciesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAttachedUserPoliciesInput"}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserName sets the UserName field's value.
func (s *ListAttachedUserPoliciesInput) SetUserName(v string) *ListAttachedUserPoliciesInput {
	s.UserName = &v
	return s
}

type ListAttachedUserPoliciesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AttachedPolicyMetadata []*AttachedPolicyMetadataForListAttachedUserPoliciesOutput `type:"list"`
}

// String returns the string representation
func (s ListAttachedUserPoliciesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAttachedUserPoliciesOutput) GoString() string {
	return s.String()
}

// SetAttachedPolicyMetadata sets the AttachedPolicyMetadata field's value.
func (s *ListAttachedUserPoliciesOutput) SetAttachedPolicyMetadata(v []*AttachedPolicyMetadataForListAttachedUserPoliciesOutput) *ListAttachedUserPoliciesOutput {
	s.AttachedPolicyMetadata = v
	return s
}

type PolicyScopeForListAttachedUserPoliciesOutput struct {
	_ struct{} `type:"structure"`

	AttachDate *string `type:"string"`

	PolicyScopeType *string `type:"string"`

	ProjectDisplayName *string `type:"string"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s PolicyScopeForListAttachedUserPoliciesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s PolicyScopeForListAttachedUserPoliciesOutput) GoString() string {
	return s.String()
}

// SetAttachDate sets the AttachDate field's value.
func (s *PolicyScopeForListAttachedUserPoliciesOutput) SetAttachDate(v string) *PolicyScopeForListAttachedUserPoliciesOutput {
	s.AttachDate = &v
	return s
}

// SetPolicyScopeType sets the PolicyScopeType field's value.
func (s *PolicyScopeForListAttachedUserPoliciesOutput) SetPolicyScopeType(v string) *PolicyScopeForListAttachedUserPoliciesOutput {
	s.PolicyScopeType = &v
	return s
}

// SetProjectDisplayName sets the ProjectDisplayName field's value.
func (s *PolicyScopeForListAttachedUserPoliciesOutput) SetProjectDisplayName(v string) *PolicyScopeForListAttachedUserPoliciesOutput {
	s.ProjectDisplayName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *PolicyScopeForListAttachedUserPoliciesOutput) SetProjectName(v string) *PolicyScopeForListAttachedUserPoliciesOutput {
	s.ProjectName = &v
	return s
}
