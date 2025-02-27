// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ark

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opEmbeddingsCommon = "Embeddings"

// EmbeddingsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the EmbeddingsCommon operation. The "output" return
// value will be populated with the EmbeddingsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EmbeddingsCommon Request to send the API call to the service.
// the "output" return value is not valid until after EmbeddingsCommon Send returns without error.
//
// See EmbeddingsCommon for more information on using the EmbeddingsCommon
// API call, and error handling.
//
//    // Example sending a request using the EmbeddingsCommonRequest method.
//    req, resp := client.EmbeddingsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) EmbeddingsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEmbeddingsCommon,
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

// EmbeddingsCommon API operation for ARK.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ARK's
// API operation EmbeddingsCommon for usage and error information.
func (c *ARK) EmbeddingsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EmbeddingsCommonRequest(input)
	return out, req.Send()
}

// EmbeddingsCommonWithContext is the same as EmbeddingsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EmbeddingsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) EmbeddingsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EmbeddingsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEmbeddings = "Embeddings"

// EmbeddingsRequest generates a "byteplus/request.Request" representing the
// client's request for the Embeddings operation. The "output" return
// value will be populated with the EmbeddingsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EmbeddingsCommon Request to send the API call to the service.
// the "output" return value is not valid until after EmbeddingsCommon Send returns without error.
//
// See Embeddings for more information on using the Embeddings
// API call, and error handling.
//
//    // Example sending a request using the EmbeddingsRequest method.
//    req, resp := client.EmbeddingsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) EmbeddingsRequest(input *EmbeddingsInput) (req *request.Request, output *EmbeddingsOutput) {
	op := &request.Operation{
		Name:       opEmbeddings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EmbeddingsInput{}
	}

	output = &EmbeddingsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// Embeddings API operation for ARK.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ARK's
// API operation Embeddings for usage and error information.
func (c *ARK) Embeddings(input *EmbeddingsInput) (*EmbeddingsOutput, error) {
	req, out := c.EmbeddingsRequest(input)
	return out, req.Send()
}

// EmbeddingsWithContext is the same as Embeddings with the addition of
// the ability to pass a context and additional request options.
//
// See Embeddings for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) EmbeddingsWithContext(ctx byteplus.Context, input *EmbeddingsInput, opts ...request.Option) (*EmbeddingsOutput, error) {
	req, out := c.EmbeddingsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForEmbeddingsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Embedding []*string `type:"list" json:"embedding,omitempty"`

	Index *int32 `type:"int32" json:"index,omitempty"`

	Object *string `type:"string" json:"object,omitempty"`
}

// String returns the string representation
func (s DataForEmbeddingsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForEmbeddingsOutput) GoString() string {
	return s.String()
}

// SetEmbedding sets the Embedding field's value.
func (s *DataForEmbeddingsOutput) SetEmbedding(v []*string) *DataForEmbeddingsOutput {
	s.Embedding = v
	return s
}

// SetIndex sets the Index field's value.
func (s *DataForEmbeddingsOutput) SetIndex(v int32) *DataForEmbeddingsOutput {
	s.Index = &v
	return s
}

// SetObject sets the Object field's value.
func (s *DataForEmbeddingsOutput) SetObject(v string) *DataForEmbeddingsOutput {
	s.Object = &v
	return s
}

type EmbeddingsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Encoding_format *string `type:"string" json:"encoding_format,omitempty"`

	Input []*string `type:"list" json:"input,omitempty"`

	// Model is a required field
	Model *string `type:"string" json:"model,omitempty" required:"true"`
}

// String returns the string representation
func (s EmbeddingsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s EmbeddingsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EmbeddingsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EmbeddingsInput"}
	if s.Model == nil {
		invalidParams.Add(request.NewErrParamRequired("Model"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEncoding_format sets the Encoding_format field's value.
func (s *EmbeddingsInput) SetEncoding_format(v string) *EmbeddingsInput {
	s.Encoding_format = &v
	return s
}

// SetInput sets the Input field's value.
func (s *EmbeddingsInput) SetInput(v []*string) *EmbeddingsInput {
	s.Input = v
	return s
}

// SetModel sets the Model field's value.
func (s *EmbeddingsInput) SetModel(v string) *EmbeddingsInput {
	s.Model = &v
	return s
}

type EmbeddingsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Created *int32 `type:"int32" json:"created,omitempty"`

	Data []*DataForEmbeddingsOutput `type:"list" json:"data,omitempty"`

	Id *string `type:"string" json:"id,omitempty"`

	Model *string `type:"string" json:"model,omitempty"`

	Object *string `type:"string" json:"object,omitempty"`

	Usage *UsageForEmbeddingsOutput `type:"structure" json:"usage,omitempty"`
}

// String returns the string representation
func (s EmbeddingsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s EmbeddingsOutput) GoString() string {
	return s.String()
}

// SetCreated sets the Created field's value.
func (s *EmbeddingsOutput) SetCreated(v int32) *EmbeddingsOutput {
	s.Created = &v
	return s
}

// SetData sets the Data field's value.
func (s *EmbeddingsOutput) SetData(v []*DataForEmbeddingsOutput) *EmbeddingsOutput {
	s.Data = v
	return s
}

// SetId sets the Id field's value.
func (s *EmbeddingsOutput) SetId(v string) *EmbeddingsOutput {
	s.Id = &v
	return s
}

// SetModel sets the Model field's value.
func (s *EmbeddingsOutput) SetModel(v string) *EmbeddingsOutput {
	s.Model = &v
	return s
}

// SetObject sets the Object field's value.
func (s *EmbeddingsOutput) SetObject(v string) *EmbeddingsOutput {
	s.Object = &v
	return s
}

// SetUsage sets the Usage field's value.
func (s *EmbeddingsOutput) SetUsage(v *UsageForEmbeddingsOutput) *EmbeddingsOutput {
	s.Usage = v
	return s
}

type UsageForEmbeddingsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Prompt_tokens *int32 `type:"int32" json:"prompt_tokens,omitempty"`

	Total_tokens *int32 `type:"int32" json:"total_tokens,omitempty"`
}

// String returns the string representation
func (s UsageForEmbeddingsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageForEmbeddingsOutput) GoString() string {
	return s.String()
}

// SetPrompt_tokens sets the Prompt_tokens field's value.
func (s *UsageForEmbeddingsOutput) SetPrompt_tokens(v int32) *UsageForEmbeddingsOutput {
	s.Prompt_tokens = &v
	return s
}

// SetTotal_tokens sets the Total_tokens field's value.
func (s *UsageForEmbeddingsOutput) SetTotal_tokens(v int32) *UsageForEmbeddingsOutput {
	s.Total_tokens = &v
	return s
}
