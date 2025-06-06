// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeleteNetworkAclCommon = "DeleteNetworkAcl"

// DeleteNetworkAclCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteNetworkAclCommon operation. The "output" return
// value will be populated with the DeleteNetworkAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNetworkAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNetworkAclCommon Send returns without error.
//
// See DeleteNetworkAclCommon for more information on using the DeleteNetworkAclCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteNetworkAclCommonRequest method.
//    req, resp := client.DeleteNetworkAclCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteNetworkAclCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteNetworkAclCommon,
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

// DeleteNetworkAclCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteNetworkAclCommon for usage and error information.
func (c *VPC) DeleteNetworkAclCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteNetworkAclCommonRequest(input)
	return out, req.Send()
}

// DeleteNetworkAclCommonWithContext is the same as DeleteNetworkAclCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNetworkAclCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteNetworkAclCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteNetworkAclCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteNetworkAcl = "DeleteNetworkAcl"

// DeleteNetworkAclRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteNetworkAcl operation. The "output" return
// value will be populated with the DeleteNetworkAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNetworkAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNetworkAclCommon Send returns without error.
//
// See DeleteNetworkAcl for more information on using the DeleteNetworkAcl
// API call, and error handling.
//
//    // Example sending a request using the DeleteNetworkAclRequest method.
//    req, resp := client.DeleteNetworkAclRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteNetworkAclRequest(input *DeleteNetworkAclInput) (req *request.Request, output *DeleteNetworkAclOutput) {
	op := &request.Operation{
		Name:       opDeleteNetworkAcl,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNetworkAclInput{}
	}

	output = &DeleteNetworkAclOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteNetworkAcl API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteNetworkAcl for usage and error information.
func (c *VPC) DeleteNetworkAcl(input *DeleteNetworkAclInput) (*DeleteNetworkAclOutput, error) {
	req, out := c.DeleteNetworkAclRequest(input)
	return out, req.Send()
}

// DeleteNetworkAclWithContext is the same as DeleteNetworkAcl with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNetworkAcl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteNetworkAclWithContext(ctx byteplus.Context, input *DeleteNetworkAclInput, opts ...request.Option) (*DeleteNetworkAclOutput, error) {
	req, out := c.DeleteNetworkAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteNetworkAclInput struct {
	_ struct{} `type:"structure"`

	// NetworkAclId is a required field
	NetworkAclId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNetworkAclInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNetworkAclInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNetworkAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteNetworkAclInput"}
	if s.NetworkAclId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkAclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNetworkAclId sets the NetworkAclId field's value.
func (s *DeleteNetworkAclInput) SetNetworkAclId(v string) *DeleteNetworkAclInput {
	s.NetworkAclId = &v
	return s
}

type DeleteNetworkAclOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteNetworkAclOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNetworkAclOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteNetworkAclOutput) SetRequestId(v string) *DeleteNetworkAclOutput {
	s.RequestId = &v
	return s
}
