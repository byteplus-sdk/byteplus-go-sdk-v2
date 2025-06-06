// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDeleteHaVipCommon = "DeleteHaVip"

// DeleteHaVipCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteHaVipCommon operation. The "output" return
// value will be populated with the DeleteHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteHaVipCommon Send returns without error.
//
// See DeleteHaVipCommon for more information on using the DeleteHaVipCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteHaVipCommonRequest method.
//    req, resp := client.DeleteHaVipCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteHaVipCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteHaVipCommon,
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

// DeleteHaVipCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteHaVipCommon for usage and error information.
func (c *VPC) DeleteHaVipCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteHaVipCommonRequest(input)
	return out, req.Send()
}

// DeleteHaVipCommonWithContext is the same as DeleteHaVipCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteHaVipCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteHaVipCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteHaVipCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteHaVip = "DeleteHaVip"

// DeleteHaVipRequest generates a "byteplus/request.Request" representing the
// client's request for the DeleteHaVip operation. The "output" return
// value will be populated with the DeleteHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteHaVipCommon Send returns without error.
//
// See DeleteHaVip for more information on using the DeleteHaVip
// API call, and error handling.
//
//    // Example sending a request using the DeleteHaVipRequest method.
//    req, resp := client.DeleteHaVipRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DeleteHaVipRequest(input *DeleteHaVipInput) (req *request.Request, output *DeleteHaVipOutput) {
	op := &request.Operation{
		Name:       opDeleteHaVip,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHaVipInput{}
	}

	output = &DeleteHaVipOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteHaVip API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DeleteHaVip for usage and error information.
func (c *VPC) DeleteHaVip(input *DeleteHaVipInput) (*DeleteHaVipOutput, error) {
	req, out := c.DeleteHaVipRequest(input)
	return out, req.Send()
}

// DeleteHaVipWithContext is the same as DeleteHaVip with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteHaVip for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteHaVipWithContext(ctx byteplus.Context, input *DeleteHaVipInput, opts ...request.Option) (*DeleteHaVipOutput, error) {
	req, out := c.DeleteHaVipRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteHaVipInput struct {
	_ struct{} `type:"structure"`

	// HaVipId is a required field
	HaVipId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHaVipInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHaVipInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHaVipInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteHaVipInput"}
	if s.HaVipId == nil {
		invalidParams.Add(request.NewErrParamRequired("HaVipId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHaVipId sets the HaVipId field's value.
func (s *DeleteHaVipInput) SetHaVipId(v string) *DeleteHaVipInput {
	s.HaVipId = &v
	return s
}

type DeleteHaVipOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteHaVipOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHaVipOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteHaVipOutput) SetRequestId(v string) *DeleteHaVipOutput {
	s.RequestId = &v
	return s
}
