// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeleteRouteTableCommon = "DeleteRouteTable"

// DeleteRouteTableCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteRouteTableCommon operation. The "output" return
// value will be populated with the DeleteRouteTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteRouteTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteRouteTableCommon Send returns without error.
//
// See DeleteRouteTableCommon for more information on using the DeleteRouteTableCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteRouteTableCommonRequest method.
//    req, resp := client.DeleteRouteTableCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteRouteTableCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteRouteTableCommon,
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

// DeleteRouteTableCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteRouteTableCommon for usage and error information.
func (c *VPC) DeleteRouteTableCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteRouteTableCommonRequest(input)
	return out, req.Send()
}

// DeleteRouteTableCommonWithContext is the same as DeleteRouteTableCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteRouteTableCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteRouteTableCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteRouteTableCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteRouteTable = "DeleteRouteTable"

// DeleteRouteTableRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteRouteTable operation. The "output" return
// value will be populated with the DeleteRouteTableCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteRouteTableCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteRouteTableCommon Send returns without error.
//
// See DeleteRouteTable for more information on using the DeleteRouteTable
// API call, and error handling.
//
//    // Example sending a request using the DeleteRouteTableRequest method.
//    req, resp := client.DeleteRouteTableRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteRouteTableRequest(input *DeleteRouteTableInput) (req *request.Request, output *DeleteRouteTableOutput) {
	op := &request.Operation{
		Name:       opDeleteRouteTable,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRouteTableInput{}
	}

	output = &DeleteRouteTableOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteRouteTable API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteRouteTable for usage and error information.
func (c *VPC) DeleteRouteTable(input *DeleteRouteTableInput) (*DeleteRouteTableOutput, error) {
	req, out := c.DeleteRouteTableRequest(input)
	return out, req.Send()
}

// DeleteRouteTableWithContext is the same as DeleteRouteTable with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteRouteTable for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteRouteTableWithContext(ctx byteplus.Context, input *DeleteRouteTableInput, opts ...request.Option) (*DeleteRouteTableOutput, error) {
	req, out := c.DeleteRouteTableRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteRouteTableInput struct {
	_ struct{} `type:"structure"`

	// RouteTableId is a required field
	RouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRouteTableInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRouteTableInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteTableInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteRouteTableInput"}
	if s.RouteTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *DeleteRouteTableInput) SetRouteTableId(v string) *DeleteRouteTableInput {
	s.RouteTableId = &v
	return s
}

type DeleteRouteTableOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteRouteTableOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRouteTableOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteRouteTableOutput) SetRequestId(v string) *DeleteRouteTableOutput {
	s.RequestId = &v
	return s
}
