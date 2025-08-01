// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opUpdateSecretCommon = "UpdateSecret"

// UpdateSecretCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the UpdateSecretCommon operation. The "output" return
// value will be populated with the UpdateSecretCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSecretCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSecretCommon Send returns without error.
//
// See UpdateSecretCommon for more information on using the UpdateSecretCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateSecretCommonRequest method.
//    req, resp := client.UpdateSecretCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) UpdateSecretCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateSecretCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateSecretCommon API operation for KMS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for KMS's
// API operation UpdateSecretCommon for usage and error information.
func (c *KMS) UpdateSecretCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateSecretCommonRequest(input)
	return out, req.Send()
}

// UpdateSecretCommonWithContext is the same as UpdateSecretCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSecretCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) UpdateSecretCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateSecretCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateSecret = "UpdateSecret"

// UpdateSecretRequest generates a "byteplus/request.Request" representing the
// client's request for the UpdateSecret operation. The "output" return
// value will be populated with the UpdateSecretCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSecretCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSecretCommon Send returns without error.
//
// See UpdateSecret for more information on using the UpdateSecret
// API call, and error handling.
//
//    // Example sending a request using the UpdateSecretRequest method.
//    req, resp := client.UpdateSecretRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) UpdateSecretRequest(input *UpdateSecretInput) (req *request.Request, output *UpdateSecretOutput) {
	op := &request.Operation{
		Name:       opUpdateSecret,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSecretInput{}
	}

	output = &UpdateSecretOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateSecret API operation for KMS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for KMS's
// API operation UpdateSecret for usage and error information.
func (c *KMS) UpdateSecret(input *UpdateSecretInput) (*UpdateSecretOutput, error) {
	req, out := c.UpdateSecretRequest(input)
	return out, req.Send()
}

// UpdateSecretWithContext is the same as UpdateSecret with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSecret for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) UpdateSecretWithContext(ctx byteplus.Context, input *UpdateSecretInput, opts ...request.Option) (*UpdateSecretOutput, error) {
	req, out := c.UpdateSecretRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateSecretInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `max:"8192" type:"string" json:",omitempty"`

	// SecretName is a required field
	SecretName *string `min:"2" max:"128" type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateSecretInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSecretInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSecretInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateSecretInput"}
	if s.Description != nil && len(*s.Description) > 8192 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 8192, *s.Description))
	}
	if s.SecretName == nil {
		invalidParams.Add(request.NewErrParamRequired("SecretName"))
	}
	if s.SecretName != nil && len(*s.SecretName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("SecretName", 2))
	}
	if s.SecretName != nil && len(*s.SecretName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("SecretName", 128, *s.SecretName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *UpdateSecretInput) SetDescription(v string) *UpdateSecretInput {
	s.Description = &v
	return s
}

// SetSecretName sets the SecretName field's value.
func (s *UpdateSecretInput) SetSecretName(v string) *UpdateSecretInput {
	s.SecretName = &v
	return s
}

type UpdateSecretOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateSecretOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSecretOutput) GoString() string {
	return s.String()
}
