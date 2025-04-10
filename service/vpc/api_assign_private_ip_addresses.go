// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opAssignPrivateIpAddressesCommon = "AssignPrivateIpAddresses"

// AssignPrivateIpAddressesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the AssignPrivateIpAddressesCommon operation. The "output" return
// value will be populated with the AssignPrivateIpAddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssignPrivateIpAddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssignPrivateIpAddressesCommon Send returns without error.
//
// See AssignPrivateIpAddressesCommon for more information on using the AssignPrivateIpAddressesCommon
// API call, and error handling.
//
//    // Example sending a request using the AssignPrivateIpAddressesCommonRequest method.
//    req, resp := client.AssignPrivateIpAddressesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AssignPrivateIpAddressesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssignPrivateIpAddressesCommon,
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

// AssignPrivateIpAddressesCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation AssignPrivateIpAddressesCommon for usage and error information.
func (c *VPC) AssignPrivateIpAddressesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssignPrivateIpAddressesCommonRequest(input)
	return out, req.Send()
}

// AssignPrivateIpAddressesCommonWithContext is the same as AssignPrivateIpAddressesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssignPrivateIpAddressesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssignPrivateIpAddressesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssignPrivateIpAddressesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssignPrivateIpAddresses = "AssignPrivateIpAddresses"

// AssignPrivateIpAddressesRequest generates a "byteplus/request.Request" representing the
// client's request for the AssignPrivateIpAddresses operation. The "output" return
// value will be populated with the AssignPrivateIpAddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssignPrivateIpAddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssignPrivateIpAddressesCommon Send returns without error.
//
// See AssignPrivateIpAddresses for more information on using the AssignPrivateIpAddresses
// API call, and error handling.
//
//    // Example sending a request using the AssignPrivateIpAddressesRequest method.
//    req, resp := client.AssignPrivateIpAddressesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AssignPrivateIpAddressesRequest(input *AssignPrivateIpAddressesInput) (req *request.Request, output *AssignPrivateIpAddressesOutput) {
	op := &request.Operation{
		Name:       opAssignPrivateIpAddresses,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssignPrivateIpAddressesInput{}
	}

	output = &AssignPrivateIpAddressesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssignPrivateIpAddresses API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation AssignPrivateIpAddresses for usage and error information.
func (c *VPC) AssignPrivateIpAddresses(input *AssignPrivateIpAddressesInput) (*AssignPrivateIpAddressesOutput, error) {
	req, out := c.AssignPrivateIpAddressesRequest(input)
	return out, req.Send()
}

// AssignPrivateIpAddressesWithContext is the same as AssignPrivateIpAddresses with the addition of
// the ability to pass a context and additional request options.
//
// See AssignPrivateIpAddresses for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssignPrivateIpAddressesWithContext(ctx byteplus.Context, input *AssignPrivateIpAddressesInput, opts ...request.Option) (*AssignPrivateIpAddressesOutput, error) {
	req, out := c.AssignPrivateIpAddressesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssignPrivateIpAddressesInput struct {
	_ struct{} `type:"structure"`

	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`

	PrivateIpAddress []*string `type:"list"`

	SecondaryPrivateIpAddressCount *int64 `type:"integer"`
}

// String returns the string representation
func (s AssignPrivateIpAddressesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AssignPrivateIpAddressesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssignPrivateIpAddressesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssignPrivateIpAddressesInput"}
	if s.NetworkInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *AssignPrivateIpAddressesInput) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesInput {
	s.NetworkInterfaceId = &v
	return s
}

// SetPrivateIpAddress sets the PrivateIpAddress field's value.
func (s *AssignPrivateIpAddressesInput) SetPrivateIpAddress(v []*string) *AssignPrivateIpAddressesInput {
	s.PrivateIpAddress = v
	return s
}

// SetSecondaryPrivateIpAddressCount sets the SecondaryPrivateIpAddressCount field's value.
func (s *AssignPrivateIpAddressesInput) SetSecondaryPrivateIpAddressCount(v int64) *AssignPrivateIpAddressesInput {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

type AssignPrivateIpAddressesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NetworkInterfaceId *string `type:"string"`

	PrivateIpSet []*string `type:"list"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AssignPrivateIpAddressesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AssignPrivateIpAddressesOutput) GoString() string {
	return s.String()
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *AssignPrivateIpAddressesOutput) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetPrivateIpSet sets the PrivateIpSet field's value.
func (s *AssignPrivateIpAddressesOutput) SetPrivateIpSet(v []*string) *AssignPrivateIpAddressesOutput {
	s.PrivateIpSet = v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *AssignPrivateIpAddressesOutput) SetRequestId(v string) *AssignPrivateIpAddressesOutput {
	s.RequestId = &v
	return s
}
