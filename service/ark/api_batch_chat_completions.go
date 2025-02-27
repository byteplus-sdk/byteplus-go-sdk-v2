// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ark

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opBatchChatCompletionsCommon = "BatchChatCompletions"

// BatchChatCompletionsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the BatchChatCompletionsCommon operation. The "output" return
// value will be populated with the BatchChatCompletionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchChatCompletionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchChatCompletionsCommon Send returns without error.
//
// See BatchChatCompletionsCommon for more information on using the BatchChatCompletionsCommon
// API call, and error handling.
//
//    // Example sending a request using the BatchChatCompletionsCommonRequest method.
//    req, resp := client.BatchChatCompletionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) BatchChatCompletionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opBatchChatCompletionsCommon,
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

// BatchChatCompletionsCommon API operation for ARK.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ARK's
// API operation BatchChatCompletionsCommon for usage and error information.
func (c *ARK) BatchChatCompletionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.BatchChatCompletionsCommonRequest(input)
	return out, req.Send()
}

// BatchChatCompletionsCommonWithContext is the same as BatchChatCompletionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See BatchChatCompletionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) BatchChatCompletionsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.BatchChatCompletionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opBatchChatCompletions = "BatchChatCompletions"

// BatchChatCompletionsRequest generates a "byteplus/request.Request" representing the
// client's request for the BatchChatCompletions operation. The "output" return
// value will be populated with the BatchChatCompletionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BatchChatCompletionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after BatchChatCompletionsCommon Send returns without error.
//
// See BatchChatCompletions for more information on using the BatchChatCompletions
// API call, and error handling.
//
//    // Example sending a request using the BatchChatCompletionsRequest method.
//    req, resp := client.BatchChatCompletionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) BatchChatCompletionsRequest(input *BatchChatCompletionsInput) (req *request.Request, output *BatchChatCompletionsOutput) {
	op := &request.Operation{
		Name:       opBatchChatCompletions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchChatCompletionsInput{}
	}

	output = &BatchChatCompletionsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// BatchChatCompletions API operation for ARK.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ARK's
// API operation BatchChatCompletions for usage and error information.
func (c *ARK) BatchChatCompletions(input *BatchChatCompletionsInput) (*BatchChatCompletionsOutput, error) {
	req, out := c.BatchChatCompletionsRequest(input)
	return out, req.Send()
}

// BatchChatCompletionsWithContext is the same as BatchChatCompletions with the addition of
// the ability to pass a context and additional request options.
//
// See BatchChatCompletions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) BatchChatCompletionsWithContext(ctx byteplus.Context, input *BatchChatCompletionsInput, opts ...request.Option) (*BatchChatCompletionsOutput, error) {
	req, out := c.BatchChatCompletionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Frequency_penalty *float64 `type:"float" json:"frequency_penalty,omitempty"`

	Logit_bias *Logit_biasForBatchChatCompletionsInput `type:"structure" json:"logit_bias,omitempty"`

	Logprobs *bool `type:"boolean" json:"logprobs,omitempty"`

	Max_tokens *int32 `type:"int32" json:"max_tokens,omitempty"`

	Messages []*MessageForBatchChatCompletionsInput `type:"list" json:"messages,omitempty"`

	// Model is a required field
	Model *string `type:"string" json:"model,omitempty" required:"true"`

	Presence_penalty *float64 `type:"float" json:"presence_penalty,omitempty"`

	Stop []*string `type:"list" json:"stop,omitempty"`

	Stream *bool `type:"boolean" json:"stream,omitempty"`

	Stream_options *Stream_optionsForBatchChatCompletionsInput `type:"structure" json:"stream_options,omitempty"`

	Temperature *float64 `type:"float" json:"temperature,omitempty"`

	Tools []*ToolForBatchChatCompletionsInput `type:"list" json:"tools,omitempty"`

	Top_logprobs *int32 `type:"int32" json:"top_logprobs,omitempty"`

	Top_p *float64 `type:"float" json:"top_p,omitempty"`
}

// String returns the string representation
func (s BatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchChatCompletionsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchChatCompletionsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "BatchChatCompletionsInput"}
	if s.Model == nil {
		invalidParams.Add(request.NewErrParamRequired("Model"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFrequency_penalty sets the Frequency_penalty field's value.
func (s *BatchChatCompletionsInput) SetFrequency_penalty(v float64) *BatchChatCompletionsInput {
	s.Frequency_penalty = &v
	return s
}

// SetLogit_bias sets the Logit_bias field's value.
func (s *BatchChatCompletionsInput) SetLogit_bias(v *Logit_biasForBatchChatCompletionsInput) *BatchChatCompletionsInput {
	s.Logit_bias = v
	return s
}

// SetLogprobs sets the Logprobs field's value.
func (s *BatchChatCompletionsInput) SetLogprobs(v bool) *BatchChatCompletionsInput {
	s.Logprobs = &v
	return s
}

// SetMax_tokens sets the Max_tokens field's value.
func (s *BatchChatCompletionsInput) SetMax_tokens(v int32) *BatchChatCompletionsInput {
	s.Max_tokens = &v
	return s
}

// SetMessages sets the Messages field's value.
func (s *BatchChatCompletionsInput) SetMessages(v []*MessageForBatchChatCompletionsInput) *BatchChatCompletionsInput {
	s.Messages = v
	return s
}

// SetModel sets the Model field's value.
func (s *BatchChatCompletionsInput) SetModel(v string) *BatchChatCompletionsInput {
	s.Model = &v
	return s
}

// SetPresence_penalty sets the Presence_penalty field's value.
func (s *BatchChatCompletionsInput) SetPresence_penalty(v float64) *BatchChatCompletionsInput {
	s.Presence_penalty = &v
	return s
}

// SetStop sets the Stop field's value.
func (s *BatchChatCompletionsInput) SetStop(v []*string) *BatchChatCompletionsInput {
	s.Stop = v
	return s
}

// SetStream sets the Stream field's value.
func (s *BatchChatCompletionsInput) SetStream(v bool) *BatchChatCompletionsInput {
	s.Stream = &v
	return s
}

// SetStream_options sets the Stream_options field's value.
func (s *BatchChatCompletionsInput) SetStream_options(v *Stream_optionsForBatchChatCompletionsInput) *BatchChatCompletionsInput {
	s.Stream_options = v
	return s
}

// SetTemperature sets the Temperature field's value.
func (s *BatchChatCompletionsInput) SetTemperature(v float64) *BatchChatCompletionsInput {
	s.Temperature = &v
	return s
}

// SetTools sets the Tools field's value.
func (s *BatchChatCompletionsInput) SetTools(v []*ToolForBatchChatCompletionsInput) *BatchChatCompletionsInput {
	s.Tools = v
	return s
}

// SetTop_logprobs sets the Top_logprobs field's value.
func (s *BatchChatCompletionsInput) SetTop_logprobs(v int32) *BatchChatCompletionsInput {
	s.Top_logprobs = &v
	return s
}

// SetTop_p sets the Top_p field's value.
func (s *BatchChatCompletionsInput) SetTop_p(v float64) *BatchChatCompletionsInput {
	s.Top_p = &v
	return s
}

type BatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Choices []*ChoiceForBatchChatCompletionsOutput `type:"list" json:"choices,omitempty"`

	Created *int32 `type:"int32" json:"created,omitempty"`

	Id *string `type:"string" json:"id,omitempty"`

	Model *string `type:"string" json:"model,omitempty"`

	Object *string `type:"string" json:"object,omitempty"`

	Usage *UsageForBatchChatCompletionsOutput `type:"structure" json:"usage,omitempty"`
}

// String returns the string representation
func (s BatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetChoices sets the Choices field's value.
func (s *BatchChatCompletionsOutput) SetChoices(v []*ChoiceForBatchChatCompletionsOutput) *BatchChatCompletionsOutput {
	s.Choices = v
	return s
}

// SetCreated sets the Created field's value.
func (s *BatchChatCompletionsOutput) SetCreated(v int32) *BatchChatCompletionsOutput {
	s.Created = &v
	return s
}

// SetId sets the Id field's value.
func (s *BatchChatCompletionsOutput) SetId(v string) *BatchChatCompletionsOutput {
	s.Id = &v
	return s
}

// SetModel sets the Model field's value.
func (s *BatchChatCompletionsOutput) SetModel(v string) *BatchChatCompletionsOutput {
	s.Model = &v
	return s
}

// SetObject sets the Object field's value.
func (s *BatchChatCompletionsOutput) SetObject(v string) *BatchChatCompletionsOutput {
	s.Object = &v
	return s
}

// SetUsage sets the Usage field's value.
func (s *BatchChatCompletionsOutput) SetUsage(v *UsageForBatchChatCompletionsOutput) *BatchChatCompletionsOutput {
	s.Usage = v
	return s
}

type ChoiceForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Delta *DeltaForBatchChatCompletionsOutput `type:"structure" json:"delta,omitempty"`

	Finish_reason *string `type:"string" json:"finish_reason,omitempty"`

	Index *int32 `type:"int32" json:"index,omitempty"`

	Logprobs *LogprobsForBatchChatCompletionsOutput `type:"structure" json:"logprobs,omitempty"`

	Message *MessageForBatchChatCompletionsOutput `type:"structure" json:"message,omitempty"`
}

// String returns the string representation
func (s ChoiceForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ChoiceForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetDelta sets the Delta field's value.
func (s *ChoiceForBatchChatCompletionsOutput) SetDelta(v *DeltaForBatchChatCompletionsOutput) *ChoiceForBatchChatCompletionsOutput {
	s.Delta = v
	return s
}

// SetFinish_reason sets the Finish_reason field's value.
func (s *ChoiceForBatchChatCompletionsOutput) SetFinish_reason(v string) *ChoiceForBatchChatCompletionsOutput {
	s.Finish_reason = &v
	return s
}

// SetIndex sets the Index field's value.
func (s *ChoiceForBatchChatCompletionsOutput) SetIndex(v int32) *ChoiceForBatchChatCompletionsOutput {
	s.Index = &v
	return s
}

// SetLogprobs sets the Logprobs field's value.
func (s *ChoiceForBatchChatCompletionsOutput) SetLogprobs(v *LogprobsForBatchChatCompletionsOutput) *ChoiceForBatchChatCompletionsOutput {
	s.Logprobs = v
	return s
}

// SetMessage sets the Message field's value.
func (s *ChoiceForBatchChatCompletionsOutput) SetMessage(v *MessageForBatchChatCompletionsOutput) *ChoiceForBatchChatCompletionsOutput {
	s.Message = v
	return s
}

type Completion_tokens_detailsForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s Completion_tokens_detailsForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Completion_tokens_detailsForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

type ContentForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Bytes []*int32 `type:"list" json:"bytes,omitempty"`

	Logprob *float64 `type:"float" json:"logprob,omitempty"`

	Token *string `type:"string" json:"token,omitempty"`

	Top_logprobs []*Top_logprobForBatchChatCompletionsOutput `type:"list" json:"top_logprobs,omitempty"`
}

// String returns the string representation
func (s ContentForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ContentForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetBytes sets the Bytes field's value.
func (s *ContentForBatchChatCompletionsOutput) SetBytes(v []*int32) *ContentForBatchChatCompletionsOutput {
	s.Bytes = v
	return s
}

// SetLogprob sets the Logprob field's value.
func (s *ContentForBatchChatCompletionsOutput) SetLogprob(v float64) *ContentForBatchChatCompletionsOutput {
	s.Logprob = &v
	return s
}

// SetToken sets the Token field's value.
func (s *ContentForBatchChatCompletionsOutput) SetToken(v string) *ContentForBatchChatCompletionsOutput {
	s.Token = &v
	return s
}

// SetTop_logprobs sets the Top_logprobs field's value.
func (s *ContentForBatchChatCompletionsOutput) SetTop_logprobs(v []*Top_logprobForBatchChatCompletionsOutput) *ContentForBatchChatCompletionsOutput {
	s.Top_logprobs = v
	return s
}

type ConvertfunctionForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:"description,omitempty"`

	Name *string `type:"string" json:"name,omitempty"`

	Parameters *ParametersForBatchChatCompletionsInput `type:"structure" json:"parameters,omitempty"`
}

// String returns the string representation
func (s ConvertfunctionForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertfunctionForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ConvertfunctionForBatchChatCompletionsInput) SetDescription(v string) *ConvertfunctionForBatchChatCompletionsInput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *ConvertfunctionForBatchChatCompletionsInput) SetName(v string) *ConvertfunctionForBatchChatCompletionsInput {
	s.Name = &v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *ConvertfunctionForBatchChatCompletionsInput) SetParameters(v *ParametersForBatchChatCompletionsInput) *ConvertfunctionForBatchChatCompletionsInput {
	s.Parameters = v
	return s
}

type Converttool_callForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Function *FunctionForBatchChatCompletionsOutput `type:"structure" json:"function,omitempty"`

	Id *string `type:"string" json:"id,omitempty"`

	Type *string `type:"string" json:"type,omitempty"`
}

// String returns the string representation
func (s Converttool_callForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Converttool_callForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetFunction sets the Function field's value.
func (s *Converttool_callForBatchChatCompletionsOutput) SetFunction(v *FunctionForBatchChatCompletionsOutput) *Converttool_callForBatchChatCompletionsOutput {
	s.Function = v
	return s
}

// SetId sets the Id field's value.
func (s *Converttool_callForBatchChatCompletionsOutput) SetId(v string) *Converttool_callForBatchChatCompletionsOutput {
	s.Id = &v
	return s
}

// SetType sets the Type field's value.
func (s *Converttool_callForBatchChatCompletionsOutput) SetType(v string) *Converttool_callForBatchChatCompletionsOutput {
	s.Type = &v
	return s
}

type DeltaForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Content *string `type:"string" json:"content,omitempty"`

	Role *string `type:"string" json:"role,omitempty"`

	Tool_calls []*Tool_callForBatchChatCompletionsOutput `type:"list" json:"tool_calls,omitempty"`
}

// String returns the string representation
func (s DeltaForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeltaForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetContent sets the Content field's value.
func (s *DeltaForBatchChatCompletionsOutput) SetContent(v string) *DeltaForBatchChatCompletionsOutput {
	s.Content = &v
	return s
}

// SetRole sets the Role field's value.
func (s *DeltaForBatchChatCompletionsOutput) SetRole(v string) *DeltaForBatchChatCompletionsOutput {
	s.Role = &v
	return s
}

// SetTool_calls sets the Tool_calls field's value.
func (s *DeltaForBatchChatCompletionsOutput) SetTool_calls(v []*Tool_callForBatchChatCompletionsOutput) *DeltaForBatchChatCompletionsOutput {
	s.Tool_calls = v
	return s
}

type FunctionForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Arguments *string `type:"string" json:"arguments,omitempty"`

	Name *string `type:"string" json:"name,omitempty"`
}

// String returns the string representation
func (s FunctionForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s FunctionForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

// SetArguments sets the Arguments field's value.
func (s *FunctionForBatchChatCompletionsInput) SetArguments(v string) *FunctionForBatchChatCompletionsInput {
	s.Arguments = &v
	return s
}

// SetName sets the Name field's value.
func (s *FunctionForBatchChatCompletionsInput) SetName(v string) *FunctionForBatchChatCompletionsInput {
	s.Name = &v
	return s
}

type FunctionForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Arguments *string `type:"string" json:"arguments,omitempty"`

	Name *string `type:"string" json:"name,omitempty"`
}

// String returns the string representation
func (s FunctionForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s FunctionForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetArguments sets the Arguments field's value.
func (s *FunctionForBatchChatCompletionsOutput) SetArguments(v string) *FunctionForBatchChatCompletionsOutput {
	s.Arguments = &v
	return s
}

// SetName sets the Name field's value.
func (s *FunctionForBatchChatCompletionsOutput) SetName(v string) *FunctionForBatchChatCompletionsOutput {
	s.Name = &v
	return s
}

type Logit_biasForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s Logit_biasForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Logit_biasForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

type LogprobsForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Content []*ContentForBatchChatCompletionsOutput `type:"list" json:"content,omitempty"`
}

// String returns the string representation
func (s LogprobsForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s LogprobsForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetContent sets the Content field's value.
func (s *LogprobsForBatchChatCompletionsOutput) SetContent(v []*ContentForBatchChatCompletionsOutput) *LogprobsForBatchChatCompletionsOutput {
	s.Content = v
	return s
}

type MessageForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Content *string `type:"string" json:"content,omitempty"`

	Role *string `type:"string" json:"role,omitempty" enum:"EnumOfroleForBatchChatCompletionsInput"`

	Tool_call_id *string `type:"string" json:"tool_call_id,omitempty"`

	Tool_calls []*Tool_callForBatchChatCompletionsInput `type:"list" json:"tool_calls,omitempty"`
}

// String returns the string representation
func (s MessageForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s MessageForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

// SetContent sets the Content field's value.
func (s *MessageForBatchChatCompletionsInput) SetContent(v string) *MessageForBatchChatCompletionsInput {
	s.Content = &v
	return s
}

// SetRole sets the Role field's value.
func (s *MessageForBatchChatCompletionsInput) SetRole(v string) *MessageForBatchChatCompletionsInput {
	s.Role = &v
	return s
}

// SetTool_call_id sets the Tool_call_id field's value.
func (s *MessageForBatchChatCompletionsInput) SetTool_call_id(v string) *MessageForBatchChatCompletionsInput {
	s.Tool_call_id = &v
	return s
}

// SetTool_calls sets the Tool_calls field's value.
func (s *MessageForBatchChatCompletionsInput) SetTool_calls(v []*Tool_callForBatchChatCompletionsInput) *MessageForBatchChatCompletionsInput {
	s.Tool_calls = v
	return s
}

type MessageForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Content *string `type:"string" json:"content,omitempty"`

	Reasoning_content *string `type:"string" json:"reasoning_content,omitempty"`

	Role *string `type:"string" json:"role,omitempty"`

	Tool_calls []*Converttool_callForBatchChatCompletionsOutput `type:"list" json:"tool_calls,omitempty"`
}

// String returns the string representation
func (s MessageForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s MessageForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetContent sets the Content field's value.
func (s *MessageForBatchChatCompletionsOutput) SetContent(v string) *MessageForBatchChatCompletionsOutput {
	s.Content = &v
	return s
}

// SetReasoning_content sets the Reasoning_content field's value.
func (s *MessageForBatchChatCompletionsOutput) SetReasoning_content(v string) *MessageForBatchChatCompletionsOutput {
	s.Reasoning_content = &v
	return s
}

// SetRole sets the Role field's value.
func (s *MessageForBatchChatCompletionsOutput) SetRole(v string) *MessageForBatchChatCompletionsOutput {
	s.Role = &v
	return s
}

// SetTool_calls sets the Tool_calls field's value.
func (s *MessageForBatchChatCompletionsOutput) SetTool_calls(v []*Converttool_callForBatchChatCompletionsOutput) *MessageForBatchChatCompletionsOutput {
	s.Tool_calls = v
	return s
}

type ParametersForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ParametersForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ParametersForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

type Prompt_tokens_detailsForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s Prompt_tokens_detailsForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Prompt_tokens_detailsForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

type Stream_optionsForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Include_usage *bool `type:"boolean" json:"include_usage,omitempty"`
}

// String returns the string representation
func (s Stream_optionsForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Stream_optionsForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

// SetInclude_usage sets the Include_usage field's value.
func (s *Stream_optionsForBatchChatCompletionsInput) SetInclude_usage(v bool) *Stream_optionsForBatchChatCompletionsInput {
	s.Include_usage = &v
	return s
}

type ToolForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Function *ConvertfunctionForBatchChatCompletionsInput `type:"structure" json:"function,omitempty"`

	Type *string `type:"string" json:"type,omitempty"`
}

// String returns the string representation
func (s ToolForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ToolForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

// SetFunction sets the Function field's value.
func (s *ToolForBatchChatCompletionsInput) SetFunction(v *ConvertfunctionForBatchChatCompletionsInput) *ToolForBatchChatCompletionsInput {
	s.Function = v
	return s
}

// SetType sets the Type field's value.
func (s *ToolForBatchChatCompletionsInput) SetType(v string) *ToolForBatchChatCompletionsInput {
	s.Type = &v
	return s
}

type Tool_callForBatchChatCompletionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Function *FunctionForBatchChatCompletionsInput `type:"structure" json:"function,omitempty"`

	Id *string `type:"string" json:"id,omitempty"`

	Type *string `type:"string" json:"type,omitempty"`
}

// String returns the string representation
func (s Tool_callForBatchChatCompletionsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Tool_callForBatchChatCompletionsInput) GoString() string {
	return s.String()
}

// SetFunction sets the Function field's value.
func (s *Tool_callForBatchChatCompletionsInput) SetFunction(v *FunctionForBatchChatCompletionsInput) *Tool_callForBatchChatCompletionsInput {
	s.Function = v
	return s
}

// SetId sets the Id field's value.
func (s *Tool_callForBatchChatCompletionsInput) SetId(v string) *Tool_callForBatchChatCompletionsInput {
	s.Id = &v
	return s
}

// SetType sets the Type field's value.
func (s *Tool_callForBatchChatCompletionsInput) SetType(v string) *Tool_callForBatchChatCompletionsInput {
	s.Type = &v
	return s
}

type Tool_callForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Function *FunctionForBatchChatCompletionsOutput `type:"structure" json:"function,omitempty"`

	Id *string `type:"string" json:"id,omitempty"`

	Index *int32 `type:"int32" json:"index,omitempty"`

	Type *string `type:"string" json:"type,omitempty"`
}

// String returns the string representation
func (s Tool_callForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Tool_callForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetFunction sets the Function field's value.
func (s *Tool_callForBatchChatCompletionsOutput) SetFunction(v *FunctionForBatchChatCompletionsOutput) *Tool_callForBatchChatCompletionsOutput {
	s.Function = v
	return s
}

// SetId sets the Id field's value.
func (s *Tool_callForBatchChatCompletionsOutput) SetId(v string) *Tool_callForBatchChatCompletionsOutput {
	s.Id = &v
	return s
}

// SetIndex sets the Index field's value.
func (s *Tool_callForBatchChatCompletionsOutput) SetIndex(v int32) *Tool_callForBatchChatCompletionsOutput {
	s.Index = &v
	return s
}

// SetType sets the Type field's value.
func (s *Tool_callForBatchChatCompletionsOutput) SetType(v string) *Tool_callForBatchChatCompletionsOutput {
	s.Type = &v
	return s
}

type Top_logprobForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Bytes []*int32 `type:"list" json:"bytes,omitempty"`

	Logprob *float64 `type:"float" json:"logprob,omitempty"`

	Token *string `type:"string" json:"token,omitempty"`
}

// String returns the string representation
func (s Top_logprobForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Top_logprobForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetBytes sets the Bytes field's value.
func (s *Top_logprobForBatchChatCompletionsOutput) SetBytes(v []*int32) *Top_logprobForBatchChatCompletionsOutput {
	s.Bytes = v
	return s
}

// SetLogprob sets the Logprob field's value.
func (s *Top_logprobForBatchChatCompletionsOutput) SetLogprob(v float64) *Top_logprobForBatchChatCompletionsOutput {
	s.Logprob = &v
	return s
}

// SetToken sets the Token field's value.
func (s *Top_logprobForBatchChatCompletionsOutput) SetToken(v string) *Top_logprobForBatchChatCompletionsOutput {
	s.Token = &v
	return s
}

type UsageForBatchChatCompletionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Completion_tokens *int32 `type:"int32" json:"completion_tokens,omitempty"`

	Completion_tokens_details *Completion_tokens_detailsForBatchChatCompletionsOutput `type:"structure" json:"completion_tokens_details,omitempty"`

	Prompt_tokens *int32 `type:"int32" json:"prompt_tokens,omitempty"`

	Prompt_tokens_details *Prompt_tokens_detailsForBatchChatCompletionsOutput `type:"structure" json:"prompt_tokens_details,omitempty"`

	Total_tokens *int32 `type:"int32" json:"total_tokens,omitempty"`
}

// String returns the string representation
func (s UsageForBatchChatCompletionsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageForBatchChatCompletionsOutput) GoString() string {
	return s.String()
}

// SetCompletion_tokens sets the Completion_tokens field's value.
func (s *UsageForBatchChatCompletionsOutput) SetCompletion_tokens(v int32) *UsageForBatchChatCompletionsOutput {
	s.Completion_tokens = &v
	return s
}

// SetCompletion_tokens_details sets the Completion_tokens_details field's value.
func (s *UsageForBatchChatCompletionsOutput) SetCompletion_tokens_details(v *Completion_tokens_detailsForBatchChatCompletionsOutput) *UsageForBatchChatCompletionsOutput {
	s.Completion_tokens_details = v
	return s
}

// SetPrompt_tokens sets the Prompt_tokens field's value.
func (s *UsageForBatchChatCompletionsOutput) SetPrompt_tokens(v int32) *UsageForBatchChatCompletionsOutput {
	s.Prompt_tokens = &v
	return s
}

// SetPrompt_tokens_details sets the Prompt_tokens_details field's value.
func (s *UsageForBatchChatCompletionsOutput) SetPrompt_tokens_details(v *Prompt_tokens_detailsForBatchChatCompletionsOutput) *UsageForBatchChatCompletionsOutput {
	s.Prompt_tokens_details = v
	return s
}

// SetTotal_tokens sets the Total_tokens field's value.
func (s *UsageForBatchChatCompletionsOutput) SetTotal_tokens(v int32) *UsageForBatchChatCompletionsOutput {
	s.Total_tokens = &v
	return s
}

const (
	// EnumOfroleForBatchChatCompletionsInputSystem is a EnumOfroleForBatchChatCompletionsInput enum value
	EnumOfroleForBatchChatCompletionsInputSystem = "system"

	// EnumOfroleForBatchChatCompletionsInputUser is a EnumOfroleForBatchChatCompletionsInput enum value
	EnumOfroleForBatchChatCompletionsInputUser = "user"

	// EnumOfroleForBatchChatCompletionsInputAssistant is a EnumOfroleForBatchChatCompletionsInput enum value
	EnumOfroleForBatchChatCompletionsInputAssistant = "assistant"

	// EnumOfroleForBatchChatCompletionsInputTool is a EnumOfroleForBatchChatCompletionsInput enum value
	EnumOfroleForBatchChatCompletionsInputTool = "tool"
)
