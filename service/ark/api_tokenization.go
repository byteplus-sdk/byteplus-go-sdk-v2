// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ark

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opTokenizationCommon = "Tokenization"

// TokenizationCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the TokenizationCommon operation. The "output" return
// value will be populated with the TokenizationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned TokenizationCommon Request to send the API call to the service.
// the "output" return value is not valid until after TokenizationCommon Send returns without error.
//
// See TokenizationCommon for more information on using the TokenizationCommon
// API call, and error handling.
//
//    // Example sending a request using the TokenizationCommonRequest method.
//    req, resp := client.TokenizationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) TokenizationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opTokenizationCommon,
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

// TokenizationCommon API operation for ARK.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ARK's
// API operation TokenizationCommon for usage and error information.
func (c *ARK) TokenizationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.TokenizationCommonRequest(input)
	return out, req.Send()
}

// TokenizationCommonWithContext is the same as TokenizationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See TokenizationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) TokenizationCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.TokenizationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opTokenization = "Tokenization"

// TokenizationRequest generates a "byteplus/request.Request" representing the
// client's request for the Tokenization operation. The "output" return
// value will be populated with the TokenizationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned TokenizationCommon Request to send the API call to the service.
// the "output" return value is not valid until after TokenizationCommon Send returns without error.
//
// See Tokenization for more information on using the Tokenization
// API call, and error handling.
//
//    // Example sending a request using the TokenizationRequest method.
//    req, resp := client.TokenizationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) TokenizationRequest(input *TokenizationInput) (req *request.Request, output *TokenizationOutput) {
	op := &request.Operation{
		Name:       opTokenization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TokenizationInput{}
	}

	output = &TokenizationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// Tokenization API operation for ARK.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ARK's
// API operation Tokenization for usage and error information.
func (c *ARK) Tokenization(input *TokenizationInput) (*TokenizationOutput, error) {
	req, out := c.TokenizationRequest(input)
	return out, req.Send()
}

// TokenizationWithContext is the same as Tokenization with the addition of
// the ability to pass a context and additional request options.
//
// See Tokenization for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) TokenizationWithContext(ctx byteplus.Context, input *TokenizationInput, opts ...request.Option) (*TokenizationOutput, error) {
	req, out := c.TokenizationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForTokenizationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Index *int32 `type:"int32" json:"index,omitempty"`

	Object *string `type:"string" json:"object,omitempty"`

	Offset_mapping []*int32 `type:"list" json:"offset_mapping,omitempty"`

	Token_ids []*int32 `type:"list" json:"token_ids,omitempty"`

	Total_tokens *int32 `type:"int32" json:"total_tokens,omitempty"`
}

// String returns the string representation
func (s DataForTokenizationOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForTokenizationOutput) GoString() string {
	return s.String()
}

// SetIndex sets the Index field's value.
func (s *DataForTokenizationOutput) SetIndex(v int32) *DataForTokenizationOutput {
	s.Index = &v
	return s
}

// SetObject sets the Object field's value.
func (s *DataForTokenizationOutput) SetObject(v string) *DataForTokenizationOutput {
	s.Object = &v
	return s
}

// SetOffset_mapping sets the Offset_mapping field's value.
func (s *DataForTokenizationOutput) SetOffset_mapping(v []*int32) *DataForTokenizationOutput {
	s.Offset_mapping = v
	return s
}

// SetToken_ids sets the Token_ids field's value.
func (s *DataForTokenizationOutput) SetToken_ids(v []*int32) *DataForTokenizationOutput {
	s.Token_ids = v
	return s
}

// SetTotal_tokens sets the Total_tokens field's value.
func (s *DataForTokenizationOutput) SetTotal_tokens(v int32) *DataForTokenizationOutput {
	s.Total_tokens = &v
	return s
}

type TokenizationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Model is a required field
	Model *string `type:"string" json:"model,omitempty" required:"true"`

	Text []*string `type:"list" json:"text,omitempty"`
}

// String returns the string representation
func (s TokenizationInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TokenizationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TokenizationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TokenizationInput"}
	if s.Model == nil {
		invalidParams.Add(request.NewErrParamRequired("Model"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetModel sets the Model field's value.
func (s *TokenizationInput) SetModel(v string) *TokenizationInput {
	s.Model = &v
	return s
}

// SetText sets the Text field's value.
func (s *TokenizationInput) SetText(v []*string) *TokenizationInput {
	s.Text = v
	return s
}

type TokenizationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Created *int32 `type:"int32" json:"created,omitempty"`

	Data []*DataForTokenizationOutput `type:"list" json:"data,omitempty"`

	Id *string `type:"string" json:"id,omitempty"`

	Model *string `type:"string" json:"model,omitempty"`

	Object *string `type:"string" json:"object,omitempty"`
}

// String returns the string representation
func (s TokenizationOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TokenizationOutput) GoString() string {
	return s.String()
}

// SetCreated sets the Created field's value.
func (s *TokenizationOutput) SetCreated(v int32) *TokenizationOutput {
	s.Created = &v
	return s
}

// SetData sets the Data field's value.
func (s *TokenizationOutput) SetData(v []*DataForTokenizationOutput) *TokenizationOutput {
	s.Data = v
	return s
}

// SetId sets the Id field's value.
func (s *TokenizationOutput) SetId(v string) *TokenizationOutput {
	s.Id = &v
	return s
}

// SetModel sets the Model field's value.
func (s *TokenizationOutput) SetModel(v string) *TokenizationOutput {
	s.Model = &v
	return s
}

// SetObject sets the Object field's value.
func (s *TokenizationOutput) SetObject(v string) *TokenizationOutput {
	s.Object = &v
	return s
}
