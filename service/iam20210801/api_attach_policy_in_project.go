// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam20210801

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opAttachPolicyInProjectCommon = "AttachPolicyInProject"

// AttachPolicyInProjectCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the AttachPolicyInProjectCommon operation. The "output" return
// value will be populated with the AttachPolicyInProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachPolicyInProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachPolicyInProjectCommon Send returns without error.
//
// See AttachPolicyInProjectCommon for more information on using the AttachPolicyInProjectCommon
// API call, and error handling.
//
//	// Example sending a request using the AttachPolicyInProjectCommonRequest method.
//	req, resp := client.AttachPolicyInProjectCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM20210801) AttachPolicyInProjectCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachPolicyInProjectCommon,
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

// AttachPolicyInProjectCommon API operation for IAM20210801.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM20210801's
// API operation AttachPolicyInProjectCommon for usage and error information.
func (c *IAM20210801) AttachPolicyInProjectCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachPolicyInProjectCommonRequest(input)
	return out, req.Send()
}

// AttachPolicyInProjectCommonWithContext is the same as AttachPolicyInProjectCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachPolicyInProjectCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) AttachPolicyInProjectCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachPolicyInProjectCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachPolicyInProject = "AttachPolicyInProject"

// AttachPolicyInProjectRequest generates a "byteplus/request.Request" representing the
// client's request for the AttachPolicyInProject operation. The "output" return
// value will be populated with the AttachPolicyInProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachPolicyInProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachPolicyInProjectCommon Send returns without error.
//
// See AttachPolicyInProject for more information on using the AttachPolicyInProject
// API call, and error handling.
//
//	// Example sending a request using the AttachPolicyInProjectRequest method.
//	req, resp := client.AttachPolicyInProjectRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM20210801) AttachPolicyInProjectRequest(input *AttachPolicyInProjectInput) (req *request.Request, output *AttachPolicyInProjectOutput) {
	op := &request.Operation{
		Name:       opAttachPolicyInProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachPolicyInProjectInput{}
	}

	output = &AttachPolicyInProjectOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachPolicyInProject API operation for IAM20210801.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM20210801's
// API operation AttachPolicyInProject for usage and error information.
func (c *IAM20210801) AttachPolicyInProject(input *AttachPolicyInProjectInput) (*AttachPolicyInProjectOutput, error) {
	req, out := c.AttachPolicyInProjectRequest(input)
	return out, req.Send()
}

// AttachPolicyInProjectWithContext is the same as AttachPolicyInProject with the addition of
// the ability to pass a context and additional request options.
//
// See AttachPolicyInProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) AttachPolicyInProjectWithContext(ctx byteplus.Context, input *AttachPolicyInProjectInput, opts ...request.Option) (*AttachPolicyInProjectOutput, error) {
	req, out := c.AttachPolicyInProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachPolicyInProjectInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true" enum:"EnumOfPolicyTypeForAttachPolicyInProjectInput"`

	// PrincipalName is a required field
	PrincipalName *string `type:"string" required:"true"`

	// PrincipalType is a required field
	PrincipalType *string `type:"string" required:"true" enum:"EnumOfPrincipalTypeForAttachPolicyInProjectInput"`

	// ProjectName is a required field
	ProjectName []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s AttachPolicyInProjectInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachPolicyInProjectInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachPolicyInProjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AttachPolicyInProjectInput"}
	if s.PolicyName == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyType == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyType"))
	}
	if s.PrincipalName == nil {
		invalidParams.Add(request.NewErrParamRequired("PrincipalName"))
	}
	if s.PrincipalType == nil {
		invalidParams.Add(request.NewErrParamRequired("PrincipalType"))
	}
	if s.ProjectName == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPolicyName sets the PolicyName field's value.
func (s *AttachPolicyInProjectInput) SetPolicyName(v string) *AttachPolicyInProjectInput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *AttachPolicyInProjectInput) SetPolicyType(v string) *AttachPolicyInProjectInput {
	s.PolicyType = &v
	return s
}

// SetPrincipalName sets the PrincipalName field's value.
func (s *AttachPolicyInProjectInput) SetPrincipalName(v string) *AttachPolicyInProjectInput {
	s.PrincipalName = &v
	return s
}

// SetPrincipalType sets the PrincipalType field's value.
func (s *AttachPolicyInProjectInput) SetPrincipalType(v string) *AttachPolicyInProjectInput {
	s.PrincipalType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *AttachPolicyInProjectInput) SetProjectName(v []*string) *AttachPolicyInProjectInput {
	s.ProjectName = v
	return s
}

type AttachPolicyInProjectOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AttachPolicyInProjectOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachPolicyInProjectOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfPolicyTypeForAttachPolicyInProjectInputSystem is a EnumOfPolicyTypeForAttachPolicyInProjectInput enum value
	EnumOfPolicyTypeForAttachPolicyInProjectInputSystem = "System"

	// EnumOfPolicyTypeForAttachPolicyInProjectInputCustom is a EnumOfPolicyTypeForAttachPolicyInProjectInput enum value
	EnumOfPolicyTypeForAttachPolicyInProjectInputCustom = "Custom"
)

const (
	// EnumOfPrincipalTypeForAttachPolicyInProjectInputUser is a EnumOfPrincipalTypeForAttachPolicyInProjectInput enum value
	EnumOfPrincipalTypeForAttachPolicyInProjectInputUser = "User"

	// EnumOfPrincipalTypeForAttachPolicyInProjectInputRole is a EnumOfPrincipalTypeForAttachPolicyInProjectInput enum value
	EnumOfPrincipalTypeForAttachPolicyInProjectInputRole = "Role"

	// EnumOfPrincipalTypeForAttachPolicyInProjectInputUserGroup is a EnumOfPrincipalTypeForAttachPolicyInProjectInput enum value
	EnumOfPrincipalTypeForAttachPolicyInProjectInputUserGroup = "UserGroup"
)
