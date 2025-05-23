// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDetachNetworkInterfaceCommon = "DetachNetworkInterface"

// DetachNetworkInterfaceCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DetachNetworkInterfaceCommon operation. The "output" return
// value will be populated with the DetachNetworkInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachNetworkInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachNetworkInterfaceCommon Send returns without error.
//
// See DetachNetworkInterfaceCommon for more information on using the DetachNetworkInterfaceCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachNetworkInterfaceCommonRequest method.
//    req, resp := client.DetachNetworkInterfaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DetachNetworkInterfaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachNetworkInterfaceCommon,
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

// DetachNetworkInterfaceCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DetachNetworkInterfaceCommon for usage and error information.
func (c *VPC) DetachNetworkInterfaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachNetworkInterfaceCommonRequest(input)
	return out, req.Send()
}

// DetachNetworkInterfaceCommonWithContext is the same as DetachNetworkInterfaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachNetworkInterfaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DetachNetworkInterfaceCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachNetworkInterfaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachNetworkInterface = "DetachNetworkInterface"

// DetachNetworkInterfaceRequest generates a "byteplus/request.Request" representing the
// client's request for the DetachNetworkInterface operation. The "output" return
// value will be populated with the DetachNetworkInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachNetworkInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachNetworkInterfaceCommon Send returns without error.
//
// See DetachNetworkInterface for more information on using the DetachNetworkInterface
// API call, and error handling.
//
//    // Example sending a request using the DetachNetworkInterfaceRequest method.
//    req, resp := client.DetachNetworkInterfaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DetachNetworkInterfaceRequest(input *DetachNetworkInterfaceInput) (req *request.Request, output *DetachNetworkInterfaceOutput) {
	op := &request.Operation{
		Name:       opDetachNetworkInterface,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachNetworkInterfaceInput{}
	}

	output = &DetachNetworkInterfaceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachNetworkInterface API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DetachNetworkInterface for usage and error information.
func (c *VPC) DetachNetworkInterface(input *DetachNetworkInterfaceInput) (*DetachNetworkInterfaceOutput, error) {
	req, out := c.DetachNetworkInterfaceRequest(input)
	return out, req.Send()
}

// DetachNetworkInterfaceWithContext is the same as DetachNetworkInterface with the addition of
// the ability to pass a context and additional request options.
//
// See DetachNetworkInterface for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DetachNetworkInterfaceWithContext(ctx byteplus.Context, input *DetachNetworkInterfaceInput, opts ...request.Option) (*DetachNetworkInterfaceOutput, error) {
	req, out := c.DetachNetworkInterfaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachNetworkInterfaceInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachNetworkInterfaceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachNetworkInterfaceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachNetworkInterfaceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.NetworkInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DetachNetworkInterfaceInput) SetInstanceId(v string) *DetachNetworkInterfaceInput {
	s.InstanceId = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *DetachNetworkInterfaceInput) SetNetworkInterfaceId(v string) *DetachNetworkInterfaceInput {
	s.NetworkInterfaceId = &v
	return s
}

type DetachNetworkInterfaceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DetachNetworkInterfaceOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachNetworkInterfaceOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DetachNetworkInterfaceOutput) SetRequestId(v string) *DetachNetworkInterfaceOutput {
	s.RequestId = &v
	return s
}
