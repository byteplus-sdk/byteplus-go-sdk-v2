// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeleteAccessKeyCommon = "DeleteAccessKey"

// DeleteAccessKeyCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteAccessKeyCommon operation. The "output" return
// value will be populated with the DeleteAccessKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAccessKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAccessKeyCommon Send returns without error.
//
// See DeleteAccessKeyCommon for more information on using the DeleteAccessKeyCommon
// API call, and error handling.
//
//	// Example sending a request using the DeleteAccessKeyCommonRequest method.
//	req, resp := client.DeleteAccessKeyCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) DeleteAccessKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteAccessKeyCommon,
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

// DeleteAccessKeyCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation DeleteAccessKeyCommon for usage and error information.
func (c *IAM) DeleteAccessKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteAccessKeyCommonRequest(input)
	return out, req.Send()
}

// DeleteAccessKeyCommonWithContext is the same as DeleteAccessKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAccessKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteAccessKeyCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteAccessKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteAccessKey = "DeleteAccessKey"

// DeleteAccessKeyRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteAccessKey operation. The "output" return
// value will be populated with the DeleteAccessKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAccessKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAccessKeyCommon Send returns without error.
//
// See DeleteAccessKey for more information on using the DeleteAccessKey
// API call, and error handling.
//
//	// Example sending a request using the DeleteAccessKeyRequest method.
//	req, resp := client.DeleteAccessKeyRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) DeleteAccessKeyRequest(input *DeleteAccessKeyInput) (req *request.Request, output *DeleteAccessKeyOutput) {
	op := &request.Operation{
		Name:       opDeleteAccessKey,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAccessKeyInput{}
	}

	output = &DeleteAccessKeyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteAccessKey API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation DeleteAccessKey for usage and error information.
func (c *IAM) DeleteAccessKey(input *DeleteAccessKeyInput) (*DeleteAccessKeyOutput, error) {
	req, out := c.DeleteAccessKeyRequest(input)
	return out, req.Send()
}

// DeleteAccessKeyWithContext is the same as DeleteAccessKey with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAccessKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteAccessKeyWithContext(ctx byteplus.Context, input *DeleteAccessKeyInput, opts ...request.Option) (*DeleteAccessKeyOutput, error) {
	req, out := c.DeleteAccessKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteAccessKeyInput struct {
	_ struct{} `type:"structure"`

	// AccessKeyId is a required field
	AccessKeyId *string `type:"string" required:"true"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s DeleteAccessKeyInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAccessKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccessKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteAccessKeyInput"}
	if s.AccessKeyId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccessKeyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *DeleteAccessKeyInput) SetAccessKeyId(v string) *DeleteAccessKeyInput {
	s.AccessKeyId = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *DeleteAccessKeyInput) SetUserName(v string) *DeleteAccessKeyInput {
	s.UserName = &v
	return s
}

type DeleteAccessKeyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteAccessKeyOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAccessKeyOutput) GoString() string {
	return s.String()
}