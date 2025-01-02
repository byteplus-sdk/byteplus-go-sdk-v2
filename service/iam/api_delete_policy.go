// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeletePolicyCommon = "DeletePolicy"

// DeletePolicyCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeletePolicyCommon operation. The "output" return
// value will be populated with the DeletePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeletePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeletePolicyCommon Send returns without error.
//
// See DeletePolicyCommon for more information on using the DeletePolicyCommon
// API call, and error handling.
//
//	// Example sending a request using the DeletePolicyCommonRequest method.
//	req, resp := client.DeletePolicyCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) DeletePolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeletePolicyCommon,
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

// DeletePolicyCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation DeletePolicyCommon for usage and error information.
func (c *IAM) DeletePolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeletePolicyCommonRequest(input)
	return out, req.Send()
}

// DeletePolicyCommonWithContext is the same as DeletePolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeletePolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeletePolicyCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeletePolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeletePolicy = "DeletePolicy"

// DeletePolicyRequest generates a "byteplus/request.Request" representing the
// client's request for the DeletePolicy operation. The "output" return
// value will be populated with the DeletePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeletePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeletePolicyCommon Send returns without error.
//
// See DeletePolicy for more information on using the DeletePolicy
// API call, and error handling.
//
//	// Example sending a request using the DeletePolicyRequest method.
//	req, resp := client.DeletePolicyRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) DeletePolicyRequest(input *DeletePolicyInput) (req *request.Request, output *DeletePolicyOutput) {
	op := &request.Operation{
		Name:       opDeletePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePolicyInput{}
	}

	output = &DeletePolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeletePolicy API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation DeletePolicy for usage and error information.
func (c *IAM) DeletePolicy(input *DeletePolicyInput) (*DeletePolicyOutput, error) {
	req, out := c.DeletePolicyRequest(input)
	return out, req.Send()
}

// DeletePolicyWithContext is the same as DeletePolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DeletePolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeletePolicyWithContext(ctx byteplus.Context, input *DeletePolicyInput, opts ...request.Option) (*DeletePolicyOutput, error) {
	req, out := c.DeletePolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeletePolicyInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `min:"1" max:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePolicyInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeletePolicyInput"}
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

// SetPolicyName sets the PolicyName field's value.
func (s *DeletePolicyInput) SetPolicyName(v string) *DeletePolicyInput {
	s.PolicyName = &v
	return s
}

type DeletePolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeletePolicyOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePolicyOutput) GoString() string {
	return s.String()
}
