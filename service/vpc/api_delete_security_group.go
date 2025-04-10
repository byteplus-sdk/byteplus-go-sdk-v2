// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeleteSecurityGroupCommon = "DeleteSecurityGroup"

// DeleteSecurityGroupCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteSecurityGroupCommon operation. The "output" return
// value will be populated with the DeleteSecurityGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteSecurityGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteSecurityGroupCommon Send returns without error.
//
// See DeleteSecurityGroupCommon for more information on using the DeleteSecurityGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteSecurityGroupCommonRequest method.
//    req, resp := client.DeleteSecurityGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteSecurityGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteSecurityGroupCommon,
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

// DeleteSecurityGroupCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteSecurityGroupCommon for usage and error information.
func (c *VPC) DeleteSecurityGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteSecurityGroupCommonRequest(input)
	return out, req.Send()
}

// DeleteSecurityGroupCommonWithContext is the same as DeleteSecurityGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteSecurityGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteSecurityGroupCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteSecurityGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteSecurityGroup = "DeleteSecurityGroup"

// DeleteSecurityGroupRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteSecurityGroup operation. The "output" return
// value will be populated with the DeleteSecurityGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteSecurityGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteSecurityGroupCommon Send returns without error.
//
// See DeleteSecurityGroup for more information on using the DeleteSecurityGroup
// API call, and error handling.
//
//    // Example sending a request using the DeleteSecurityGroupRequest method.
//    req, resp := client.DeleteSecurityGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteSecurityGroupRequest(input *DeleteSecurityGroupInput) (req *request.Request, output *DeleteSecurityGroupOutput) {
	op := &request.Operation{
		Name:       opDeleteSecurityGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSecurityGroupInput{}
	}

	output = &DeleteSecurityGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteSecurityGroup API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteSecurityGroup for usage and error information.
func (c *VPC) DeleteSecurityGroup(input *DeleteSecurityGroupInput) (*DeleteSecurityGroupOutput, error) {
	req, out := c.DeleteSecurityGroupRequest(input)
	return out, req.Send()
}

// DeleteSecurityGroupWithContext is the same as DeleteSecurityGroup with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteSecurityGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteSecurityGroupWithContext(ctx byteplus.Context, input *DeleteSecurityGroupInput, opts ...request.Option) (*DeleteSecurityGroupOutput, error) {
	req, out := c.DeleteSecurityGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	// SecurityGroupId is a required field
	SecurityGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSecurityGroupInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteSecurityGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSecurityGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteSecurityGroupInput"}
	if s.SecurityGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *DeleteSecurityGroupInput) SetSecurityGroupId(v string) *DeleteSecurityGroupInput {
	s.SecurityGroupId = &v
	return s
}

type DeleteSecurityGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteSecurityGroupOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteSecurityGroupOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteSecurityGroupOutput) SetRequestId(v string) *DeleteSecurityGroupOutput {
	s.RequestId = &v
	return s
}
