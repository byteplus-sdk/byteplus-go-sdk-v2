// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreateTagsCommon = "CreateTags"

// CreateTagsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateTagsCommon operation. The "output" return
// value will be populated with the CreateTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTagsCommon Send returns without error.
//
// See CreateTagsCommon for more information on using the CreateTagsCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTagsCommonRequest method.
//    req, resp := client.CreateTagsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateTagsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTagsCommon,
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

// CreateTagsCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation CreateTagsCommon for usage and error information.
func (c *ECS) CreateTagsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTagsCommonRequest(input)
	return out, req.Send()
}

// CreateTagsCommonWithContext is the same as CreateTagsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTagsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateTagsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTagsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTags = "CreateTags"

// CreateTagsRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateTags operation. The "output" return
// value will be populated with the CreateTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTagsCommon Send returns without error.
//
// See CreateTags for more information on using the CreateTags
// API call, and error handling.
//
//    // Example sending a request using the CreateTagsRequest method.
//    req, resp := client.CreateTagsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateTagsRequest(input *CreateTagsInput) (req *request.Request, output *CreateTagsOutput) {
	op := &request.Operation{
		Name:       opCreateTags,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTagsInput{}
	}

	output = &CreateTagsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTags API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation CreateTags for usage and error information.
func (c *ECS) CreateTags(input *CreateTagsInput) (*CreateTagsOutput, error) {
	req, out := c.CreateTagsRequest(input)
	return out, req.Send()
}

// CreateTagsWithContext is the same as CreateTags with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTags for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateTagsWithContext(ctx byteplus.Context, input *CreateTagsInput, opts ...request.Option) (*CreateTagsOutput, error) {
	req, out := c.CreateTagsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTagsInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	ResourceIds []*string `type:"list"`

	// ResourceType is a required field
	ResourceType *string `type:"string" required:"true"`

	Tags []*TagForCreateTagsInput `type:"list"`
}

// String returns the string representation
func (s CreateTagsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTagsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTagsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTagsInput"}
	if s.ResourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceType"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateTagsInput) SetClientToken(v string) *CreateTagsInput {
	s.ClientToken = &v
	return s
}

// SetResourceIds sets the ResourceIds field's value.
func (s *CreateTagsInput) SetResourceIds(v []*string) *CreateTagsInput {
	s.ResourceIds = v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *CreateTagsInput) SetResourceType(v string) *CreateTagsInput {
	s.ResourceType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateTagsInput) SetTags(v []*TagForCreateTagsInput) *CreateTagsInput {
	s.Tags = v
	return s
}

type CreateTagsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForCreateTagsOutput `type:"list"`
}

// String returns the string representation
func (s CreateTagsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTagsOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *CreateTagsOutput) SetOperationDetails(v []*OperationDetailForCreateTagsOutput) *CreateTagsOutput {
	s.OperationDetails = v
	return s
}

type ErrorForCreateTagsOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForCreateTagsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForCreateTagsOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForCreateTagsOutput) SetCode(v string) *ErrorForCreateTagsOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForCreateTagsOutput) SetMessage(v string) *ErrorForCreateTagsOutput {
	s.Message = &v
	return s
}

type OperationDetailForCreateTagsOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForCreateTagsOutput `type:"structure"`

	ResourceId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForCreateTagsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForCreateTagsOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForCreateTagsOutput) SetError(v *ErrorForCreateTagsOutput) *OperationDetailForCreateTagsOutput {
	s.Error = v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *OperationDetailForCreateTagsOutput) SetResourceId(v string) *OperationDetailForCreateTagsOutput {
	s.ResourceId = &v
	return s
}

type TagForCreateTagsInput struct {
	_ struct{} `type:"structure"`

	// Key is a required field
	Key *string `type:"string" required:"true"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateTagsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateTagsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagForCreateTagsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TagForCreateTagsInput"}
	if s.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKey sets the Key field's value.
func (s *TagForCreateTagsInput) SetKey(v string) *TagForCreateTagsInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateTagsInput) SetValue(v string) *TagForCreateTagsInput {
	s.Value = &v
	return s
}
