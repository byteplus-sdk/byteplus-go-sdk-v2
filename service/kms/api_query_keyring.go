// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opQueryKeyringCommon = "QueryKeyring"

// QueryKeyringCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the QueryKeyringCommon operation. The "output" return
// value will be populated with the QueryKeyringCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryKeyringCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryKeyringCommon Send returns without error.
//
// See QueryKeyringCommon for more information on using the QueryKeyringCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryKeyringCommonRequest method.
//    req, resp := client.QueryKeyringCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) QueryKeyringCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryKeyringCommon,
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

// QueryKeyringCommon API operation for KMS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for KMS's
// API operation QueryKeyringCommon for usage and error information.
func (c *KMS) QueryKeyringCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryKeyringCommonRequest(input)
	return out, req.Send()
}

// QueryKeyringCommonWithContext is the same as QueryKeyringCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryKeyringCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) QueryKeyringCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryKeyringCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryKeyring = "QueryKeyring"

// QueryKeyringRequest generates a "byteplus/request.Request" representing the
// client's request for the QueryKeyring operation. The "output" return
// value will be populated with the QueryKeyringCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryKeyringCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryKeyringCommon Send returns without error.
//
// See QueryKeyring for more information on using the QueryKeyring
// API call, and error handling.
//
//    // Example sending a request using the QueryKeyringRequest method.
//    req, resp := client.QueryKeyringRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) QueryKeyringRequest(input *QueryKeyringInput) (req *request.Request, output *QueryKeyringOutput) {
	op := &request.Operation{
		Name:       opQueryKeyring,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryKeyringInput{}
	}

	output = &QueryKeyringOutput{}
	req = c.newRequest(op, input, output)

	return
}

// QueryKeyring API operation for KMS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for KMS's
// API operation QueryKeyring for usage and error information.
func (c *KMS) QueryKeyring(input *QueryKeyringInput) (*QueryKeyringOutput, error) {
	req, out := c.QueryKeyringRequest(input)
	return out, req.Send()
}

// QueryKeyringWithContext is the same as QueryKeyring with the addition of
// the ability to pass a context and additional request options.
//
// See QueryKeyring for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) QueryKeyringWithContext(ctx byteplus.Context, input *QueryKeyringInput, opts ...request.Option) (*QueryKeyringOutput, error) {
	req, out := c.QueryKeyringRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type KeyringForQueryKeyringOutput struct {
	_ struct{} `type:"structure"`

	CreationDate *int64 `type:"int64"`

	Description *string `type:"string"`

	ID *string `type:"string"`

	KeyringName *string `type:"string"`

	KeyringType *string `type:"string"`

	ProjectName *string `type:"string"`

	TRN *string `type:"string"`

	UID *string `type:"string"`

	UpdateDate *int64 `type:"int64"`
}

// String returns the string representation
func (s KeyringForQueryKeyringOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s KeyringForQueryKeyringOutput) GoString() string {
	return s.String()
}

// SetCreationDate sets the CreationDate field's value.
func (s *KeyringForQueryKeyringOutput) SetCreationDate(v int64) *KeyringForQueryKeyringOutput {
	s.CreationDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *KeyringForQueryKeyringOutput) SetDescription(v string) *KeyringForQueryKeyringOutput {
	s.Description = &v
	return s
}

// SetID sets the ID field's value.
func (s *KeyringForQueryKeyringOutput) SetID(v string) *KeyringForQueryKeyringOutput {
	s.ID = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *KeyringForQueryKeyringOutput) SetKeyringName(v string) *KeyringForQueryKeyringOutput {
	s.KeyringName = &v
	return s
}

// SetKeyringType sets the KeyringType field's value.
func (s *KeyringForQueryKeyringOutput) SetKeyringType(v string) *KeyringForQueryKeyringOutput {
	s.KeyringType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *KeyringForQueryKeyringOutput) SetProjectName(v string) *KeyringForQueryKeyringOutput {
	s.ProjectName = &v
	return s
}

// SetTRN sets the TRN field's value.
func (s *KeyringForQueryKeyringOutput) SetTRN(v string) *KeyringForQueryKeyringOutput {
	s.TRN = &v
	return s
}

// SetUID sets the UID field's value.
func (s *KeyringForQueryKeyringOutput) SetUID(v string) *KeyringForQueryKeyringOutput {
	s.UID = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *KeyringForQueryKeyringOutput) SetUpdateDate(v int64) *KeyringForQueryKeyringOutput {
	s.UpdateDate = &v
	return s
}

type QueryKeyringInput struct {
	_ struct{} `type:"structure"`

	KeyringID *string `type:"string"`

	KeyringName *string `min:"2" max:"31" type:"string"`
}

// String returns the string representation
func (s QueryKeyringInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryKeyringInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryKeyringInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryKeyringInput"}
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
func (s *QueryKeyringInput) SetKeyringID(v string) *QueryKeyringInput {
	s.KeyringID = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *QueryKeyringInput) SetKeyringName(v string) *QueryKeyringInput {
	s.KeyringName = &v
	return s
}

type QueryKeyringOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Keyring *KeyringForQueryKeyringOutput `type:"structure"`
}

// String returns the string representation
func (s QueryKeyringOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryKeyringOutput) GoString() string {
	return s.String()
}

// SetKeyring sets the Keyring field's value.
func (s *QueryKeyringOutput) SetKeyring(v *KeyringForQueryKeyringOutput) *QueryKeyringOutput {
	s.Keyring = v
	return s
}
