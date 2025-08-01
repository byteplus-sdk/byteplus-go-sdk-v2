// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeleteTagsCommon = "DeleteTags"

// DeleteTagsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteTagsCommon operation. The "output" return
// value will be populated with the DeleteTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTagsCommon Send returns without error.
//
// See DeleteTagsCommon for more information on using the DeleteTagsCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTagsCommonRequest method.
//    req, resp := client.DeleteTagsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) DeleteTagsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTagsCommon,
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

// DeleteTagsCommon API operation for CR.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for CR's
// API operation DeleteTagsCommon for usage and error information.
func (c *CR) DeleteTagsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTagsCommonRequest(input)
	return out, req.Send()
}

// DeleteTagsCommonWithContext is the same as DeleteTagsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTagsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) DeleteTagsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTagsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTags = "DeleteTags"

// DeleteTagsRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteTags operation. The "output" return
// value will be populated with the DeleteTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTagsCommon Send returns without error.
//
// See DeleteTags for more information on using the DeleteTags
// API call, and error handling.
//
//    // Example sending a request using the DeleteTagsRequest method.
//    req, resp := client.DeleteTagsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) DeleteTagsRequest(input *DeleteTagsInput) (req *request.Request, output *DeleteTagsOutput) {
	op := &request.Operation{
		Name:       opDeleteTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTagsInput{}
	}

	output = &DeleteTagsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteTags API operation for CR.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for CR's
// API operation DeleteTags for usage and error information.
func (c *CR) DeleteTags(input *DeleteTagsInput) (*DeleteTagsOutput, error) {
	req, out := c.DeleteTagsRequest(input)
	return out, req.Send()
}

// DeleteTagsWithContext is the same as DeleteTags with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTags for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) DeleteTagsWithContext(ctx byteplus.Context, input *DeleteTagsInput, opts ...request.Option) (*DeleteTagsOutput, error) {
	req, out := c.DeleteTagsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTagsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Names []*string `type:"list" json:",omitempty"`

	// Namespace is a required field
	Namespace *string `min:"2" max:"90" type:"string" json:",omitempty" required:"true"`

	// Registry is a required field
	Registry *string `min:"3" max:"30" type:"string" json:",omitempty" required:"true"`

	// Repository is a required field
	Repository *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteTagsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTagsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteTagsInput"}
	if s.Namespace == nil {
		invalidParams.Add(request.NewErrParamRequired("Namespace"))
	}
	if s.Namespace != nil && len(*s.Namespace) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("Namespace", 2))
	}
	if s.Namespace != nil && len(*s.Namespace) > 90 {
		invalidParams.Add(request.NewErrParamMaxLen("Namespace", 90, *s.Namespace))
	}
	if s.Registry == nil {
		invalidParams.Add(request.NewErrParamRequired("Registry"))
	}
	if s.Registry != nil && len(*s.Registry) < 3 {
		invalidParams.Add(request.NewErrParamMinLen("Registry", 3))
	}
	if s.Registry != nil && len(*s.Registry) > 30 {
		invalidParams.Add(request.NewErrParamMaxLen("Registry", 30, *s.Registry))
	}
	if s.Repository == nil {
		invalidParams.Add(request.NewErrParamRequired("Repository"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNames sets the Names field's value.
func (s *DeleteTagsInput) SetNames(v []*string) *DeleteTagsInput {
	s.Names = v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DeleteTagsInput) SetNamespace(v string) *DeleteTagsInput {
	s.Namespace = &v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *DeleteTagsInput) SetRegistry(v string) *DeleteTagsInput {
	s.Registry = &v
	return s
}

// SetRepository sets the Repository field's value.
func (s *DeleteTagsInput) SetRepository(v string) *DeleteTagsInput {
	s.Repository = &v
	return s
}

type DeleteTagsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Failures []*FailureForDeleteTagsOutput `type:"list" json:",omitempty"`

	Successes []*SuccessForDeleteTagsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteTagsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTagsOutput) GoString() string {
	return s.String()
}

// SetFailures sets the Failures field's value.
func (s *DeleteTagsOutput) SetFailures(v []*FailureForDeleteTagsOutput) *DeleteTagsOutput {
	s.Failures = v
	return s
}

// SetSuccesses sets the Successes field's value.
func (s *DeleteTagsOutput) SetSuccesses(v []*SuccessForDeleteTagsOutput) *DeleteTagsOutput {
	s.Successes = v
	return s
}

type FailureForDeleteTagsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Reason *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FailureForDeleteTagsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s FailureForDeleteTagsOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *FailureForDeleteTagsOutput) SetName(v string) *FailureForDeleteTagsOutput {
	s.Name = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *FailureForDeleteTagsOutput) SetReason(v string) *FailureForDeleteTagsOutput {
	s.Reason = &v
	return s
}

type SuccessForDeleteTagsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SuccessForDeleteTagsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s SuccessForDeleteTagsOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *SuccessForDeleteTagsOutput) SetName(v string) *SuccessForDeleteTagsOutput {
	s.Name = &v
	return s
}
