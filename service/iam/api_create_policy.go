// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreatePolicyCommon = "CreatePolicy"

// CreatePolicyCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreatePolicyCommon operation. The "output" return
// value will be populated with the CreatePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreatePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreatePolicyCommon Send returns without error.
//
// See CreatePolicyCommon for more information on using the CreatePolicyCommon
// API call, and error handling.
//
//	// Example sending a request using the CreatePolicyCommonRequest method.
//	req, resp := client.CreatePolicyCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) CreatePolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreatePolicyCommon,
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

// CreatePolicyCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation CreatePolicyCommon for usage and error information.
func (c *IAM) CreatePolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreatePolicyCommonRequest(input)
	return out, req.Send()
}

// CreatePolicyCommonWithContext is the same as CreatePolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreatePolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreatePolicyCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreatePolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreatePolicy = "CreatePolicy"

// CreatePolicyRequest generates a "byteplus/request.Request" representing the
// client's request for the CreatePolicy operation. The "output" return
// value will be populated with the CreatePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreatePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreatePolicyCommon Send returns without error.
//
// See CreatePolicy for more information on using the CreatePolicy
// API call, and error handling.
//
//	// Example sending a request using the CreatePolicyRequest method.
//	req, resp := client.CreatePolicyRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) CreatePolicyRequest(input *CreatePolicyInput) (req *request.Request, output *CreatePolicyOutput) {
	op := &request.Operation{
		Name:       opCreatePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePolicyInput{}
	}

	output = &CreatePolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreatePolicy API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation CreatePolicy for usage and error information.
func (c *IAM) CreatePolicy(input *CreatePolicyInput) (*CreatePolicyOutput, error) {
	req, out := c.CreatePolicyRequest(input)
	return out, req.Send()
}

// CreatePolicyWithContext is the same as CreatePolicy with the addition of
// the ability to pass a context and additional request options.
//
// See CreatePolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreatePolicyWithContext(ctx byteplus.Context, input *CreatePolicyInput, opts ...request.Option) (*CreatePolicyOutput, error) {
	req, out := c.CreatePolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreatePolicyInput struct {
	_ struct{} `type:"structure"`

	Description *string `max:"255" type:"string"`

	// PolicyDocument is a required field
	PolicyDocument *string `min:"1" max:"6144" type:"string" required:"true"`

	// PolicyName is a required field
	PolicyName *string `min:"1" max:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePolicyInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatePolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreatePolicyInput"}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.PolicyDocument == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyDocument"))
	}
	if s.PolicyDocument != nil && len(*s.PolicyDocument) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("PolicyDocument", 1))
	}
	if s.PolicyDocument != nil && len(*s.PolicyDocument) > 6144 {
		invalidParams.Add(request.NewErrParamMaxLen("PolicyDocument", 6144, *s.PolicyDocument))
	}
	if s.PolicyName == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("PolicyName", 1))
	}
	if s.PolicyName != nil && len(*s.PolicyName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("PolicyName", 64, *s.PolicyName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreatePolicyInput) SetDescription(v string) *CreatePolicyInput {
	s.Description = &v
	return s
}

// SetPolicyDocument sets the PolicyDocument field's value.
func (s *CreatePolicyInput) SetPolicyDocument(v string) *CreatePolicyInput {
	s.PolicyDocument = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *CreatePolicyInput) SetPolicyName(v string) *CreatePolicyInput {
	s.PolicyName = &v
	return s
}

type CreatePolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Policy *PolicyForCreatePolicyOutput `type:"structure"`
}

// String returns the string representation
func (s CreatePolicyOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatePolicyOutput) GoString() string {
	return s.String()
}

// SetPolicy sets the Policy field's value.
func (s *CreatePolicyOutput) SetPolicy(v *PolicyForCreatePolicyOutput) *CreatePolicyOutput {
	s.Policy = v
	return s
}

type PolicyForCreatePolicyOutput struct {
	_ struct{} `type:"structure"`

	AttachmentCount *int32 `type:"int32"`

	Category *string `type:"string"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	IsServiceRolePolicy *int32 `type:"int32"`

	PolicyDocument *string `type:"string"`

	PolicyName *string `type:"string"`

	PolicyTrn *string `type:"string"`

	PolicyType *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s PolicyForCreatePolicyOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s PolicyForCreatePolicyOutput) GoString() string {
	return s.String()
}

// SetAttachmentCount sets the AttachmentCount field's value.
func (s *PolicyForCreatePolicyOutput) SetAttachmentCount(v int32) *PolicyForCreatePolicyOutput {
	s.AttachmentCount = &v
	return s
}

// SetCategory sets the Category field's value.
func (s *PolicyForCreatePolicyOutput) SetCategory(v string) *PolicyForCreatePolicyOutput {
	s.Category = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *PolicyForCreatePolicyOutput) SetCreateDate(v string) *PolicyForCreatePolicyOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *PolicyForCreatePolicyOutput) SetDescription(v string) *PolicyForCreatePolicyOutput {
	s.Description = &v
	return s
}

// SetIsServiceRolePolicy sets the IsServiceRolePolicy field's value.
func (s *PolicyForCreatePolicyOutput) SetIsServiceRolePolicy(v int32) *PolicyForCreatePolicyOutput {
	s.IsServiceRolePolicy = &v
	return s
}

// SetPolicyDocument sets the PolicyDocument field's value.
func (s *PolicyForCreatePolicyOutput) SetPolicyDocument(v string) *PolicyForCreatePolicyOutput {
	s.PolicyDocument = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *PolicyForCreatePolicyOutput) SetPolicyName(v string) *PolicyForCreatePolicyOutput {
	s.PolicyName = &v
	return s
}

// SetPolicyTrn sets the PolicyTrn field's value.
func (s *PolicyForCreatePolicyOutput) SetPolicyTrn(v string) *PolicyForCreatePolicyOutput {
	s.PolicyTrn = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *PolicyForCreatePolicyOutput) SetPolicyType(v string) *PolicyForCreatePolicyOutput {
	s.PolicyType = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *PolicyForCreatePolicyOutput) SetUpdateDate(v string) *PolicyForCreatePolicyOutput {
	s.UpdateDate = &v
	return s
}
