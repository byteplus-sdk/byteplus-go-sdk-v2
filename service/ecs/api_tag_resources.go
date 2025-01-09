// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opTagResourcesCommon = "TagResources"

// TagResourcesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the TagResourcesCommon operation. The "output" return
// value will be populated with the TagResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned TagResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after TagResourcesCommon Send returns without error.
//
// See TagResourcesCommon for more information on using the TagResourcesCommon
// API call, and error handling.
//
//    // Example sending a request using the TagResourcesCommonRequest method.
//    req, resp := client.TagResourcesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) TagResourcesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opTagResourcesCommon,
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

// TagResourcesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation TagResourcesCommon for usage and error information.
func (c *ECS) TagResourcesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.TagResourcesCommonRequest(input)
	return out, req.Send()
}

// TagResourcesCommonWithContext is the same as TagResourcesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See TagResourcesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) TagResourcesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.TagResourcesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opTagResources = "TagResources"

// TagResourcesRequest generates a "byteplus/request.Request" representing the
// client's request for the TagResources operation. The "output" return
// value will be populated with the TagResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned TagResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after TagResourcesCommon Send returns without error.
//
// See TagResources for more information on using the TagResources
// API call, and error handling.
//
//    // Example sending a request using the TagResourcesRequest method.
//    req, resp := client.TagResourcesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) TagResourcesRequest(input *TagResourcesInput) (req *request.Request, output *TagResourcesOutput) {
	op := &request.Operation{
		Name:       opTagResources,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagResourcesInput{}
	}

	output = &TagResourcesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// TagResources API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation TagResources for usage and error information.
func (c *ECS) TagResources(input *TagResourcesInput) (*TagResourcesOutput, error) {
	req, out := c.TagResourcesRequest(input)
	return out, req.Send()
}

// TagResourcesWithContext is the same as TagResources with the addition of
// the ability to pass a context and additional request options.
//
// See TagResources for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) TagResourcesWithContext(ctx byteplus.Context, input *TagResourcesInput, opts ...request.Option) (*TagResourcesOutput, error) {
	req, out := c.TagResourcesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type TagForTagResourcesInput struct {
	_ struct{} `type:"structure"`

	// Key is a required field
	Key *string `type:"string" required:"true"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForTagResourcesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForTagResourcesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagForTagResourcesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TagForTagResourcesInput"}
	if s.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKey sets the Key field's value.
func (s *TagForTagResourcesInput) SetKey(v string) *TagForTagResourcesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForTagResourcesInput) SetValue(v string) *TagForTagResourcesInput {
	s.Value = &v
	return s
}

type TagResourcesInput struct {
	_ struct{} `type:"structure"`

	ResourceIds []*string `type:"list"`

	ResourceType *string `type:"string"`

	Tags []*TagForTagResourcesInput `type:"list"`
}

// String returns the string representation
func (s TagResourcesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagResourcesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourcesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TagResourcesInput"}
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

// SetResourceIds sets the ResourceIds field's value.
func (s *TagResourcesInput) SetResourceIds(v []*string) *TagResourcesInput {
	s.ResourceIds = v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *TagResourcesInput) SetResourceType(v string) *TagResourcesInput {
	s.ResourceType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *TagResourcesInput) SetTags(v []*TagForTagResourcesInput) *TagResourcesInput {
	s.Tags = v
	return s
}

type TagResourcesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s TagResourcesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagResourcesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *TagResourcesOutput) SetRequestId(v string) *TagResourcesOutput {
	s.RequestId = &v
	return s
}
