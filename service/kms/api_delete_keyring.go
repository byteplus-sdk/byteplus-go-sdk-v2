// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeleteKeyringCommon = "DeleteKeyring"

// DeleteKeyringCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteKeyringCommon operation. The "output" return
// value will be populated with the DeleteKeyringCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteKeyringCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteKeyringCommon Send returns without error.
//
// See DeleteKeyringCommon for more information on using the DeleteKeyringCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteKeyringCommonRequest method.
//    req, resp := client.DeleteKeyringCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DeleteKeyringCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteKeyringCommon,
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

// DeleteKeyringCommon API operation for KMS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for KMS's
// API operation DeleteKeyringCommon for usage and error information.
func (c *KMS) DeleteKeyringCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteKeyringCommonRequest(input)
	return out, req.Send()
}

// DeleteKeyringCommonWithContext is the same as DeleteKeyringCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteKeyringCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DeleteKeyringCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteKeyringCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteKeyring = "DeleteKeyring"

// DeleteKeyringRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteKeyring operation. The "output" return
// value will be populated with the DeleteKeyringCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteKeyringCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteKeyringCommon Send returns without error.
//
// See DeleteKeyring for more information on using the DeleteKeyring
// API call, and error handling.
//
//    // Example sending a request using the DeleteKeyringRequest method.
//    req, resp := client.DeleteKeyringRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DeleteKeyringRequest(input *DeleteKeyringInput) (req *request.Request, output *DeleteKeyringOutput) {
	op := &request.Operation{
		Name:       opDeleteKeyring,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteKeyringInput{}
	}

	output = &DeleteKeyringOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteKeyring API operation for KMS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for KMS's
// API operation DeleteKeyring for usage and error information.
func (c *KMS) DeleteKeyring(input *DeleteKeyringInput) (*DeleteKeyringOutput, error) {
	req, out := c.DeleteKeyringRequest(input)
	return out, req.Send()
}

// DeleteKeyringWithContext is the same as DeleteKeyring with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteKeyring for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DeleteKeyringWithContext(ctx byteplus.Context, input *DeleteKeyringInput, opts ...request.Option) (*DeleteKeyringOutput, error) {
	req, out := c.DeleteKeyringRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteKeyringInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	KeyringID *string `type:"string" json:",omitempty"`

	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteKeyringInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteKeyringInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteKeyringInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteKeyringInput"}
	if s.KeyringName != nil && len(*s.KeyringName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyringName", 2))
	}
	if s.KeyringName != nil && len(*s.KeyringName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyringName", 31, *s.KeyringName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKeyringID sets the KeyringID field's value.
func (s *DeleteKeyringInput) SetKeyringID(v string) *DeleteKeyringInput {
	s.KeyringID = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *DeleteKeyringInput) SetKeyringName(v string) *DeleteKeyringInput {
	s.KeyringName = &v
	return s
}

type DeleteKeyringOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteKeyringOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteKeyringOutput) GoString() string {
	return s.String()
}
