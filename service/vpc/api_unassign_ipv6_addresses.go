// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opUnassignIpv6AddressesCommon = "UnassignIpv6Addresses"

// UnassignIpv6AddressesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the UnassignIpv6AddressesCommon operation. The "output" return
// value will be populated with the UnassignIpv6AddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnassignIpv6AddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnassignIpv6AddressesCommon Send returns without error.
//
// See UnassignIpv6AddressesCommon for more information on using the UnassignIpv6AddressesCommon
// API call, and error handling.
//
//    // Example sending a request using the UnassignIpv6AddressesCommonRequest method.
//    req, resp := client.UnassignIpv6AddressesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) UnassignIpv6AddressesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUnassignIpv6AddressesCommon,
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

// UnassignIpv6AddressesCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation UnassignIpv6AddressesCommon for usage and error information.
func (c *VPC) UnassignIpv6AddressesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UnassignIpv6AddressesCommonRequest(input)
	return out, req.Send()
}

// UnassignIpv6AddressesCommonWithContext is the same as UnassignIpv6AddressesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UnassignIpv6AddressesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) UnassignIpv6AddressesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UnassignIpv6AddressesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUnassignIpv6Addresses = "UnassignIpv6Addresses"

// UnassignIpv6AddressesRequest generates a "byteplus/request.Request" representing the
// client's request for the UnassignIpv6Addresses operation. The "output" return
// value will be populated with the UnassignIpv6AddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnassignIpv6AddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnassignIpv6AddressesCommon Send returns without error.
//
// See UnassignIpv6Addresses for more information on using the UnassignIpv6Addresses
// API call, and error handling.
//
//    // Example sending a request using the UnassignIpv6AddressesRequest method.
//    req, resp := client.UnassignIpv6AddressesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) UnassignIpv6AddressesRequest(input *UnassignIpv6AddressesInput) (req *request.Request, output *UnassignIpv6AddressesOutput) {
	op := &request.Operation{
		Name:       opUnassignIpv6Addresses,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UnassignIpv6AddressesInput{}
	}

	output = &UnassignIpv6AddressesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UnassignIpv6Addresses API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation UnassignIpv6Addresses for usage and error information.
func (c *VPC) UnassignIpv6Addresses(input *UnassignIpv6AddressesInput) (*UnassignIpv6AddressesOutput, error) {
	req, out := c.UnassignIpv6AddressesRequest(input)
	return out, req.Send()
}

// UnassignIpv6AddressesWithContext is the same as UnassignIpv6Addresses with the addition of
// the ability to pass a context and additional request options.
//
// See UnassignIpv6Addresses for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) UnassignIpv6AddressesWithContext(ctx byteplus.Context, input *UnassignIpv6AddressesInput, opts ...request.Option) (*UnassignIpv6AddressesOutput, error) {
	req, out := c.UnassignIpv6AddressesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UnassignIpv6AddressesInput struct {
	_ struct{} `type:"structure"`

	// Ipv6Address is a required field
	Ipv6Address []*string `type:"list" required:"true"`

	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UnassignIpv6AddressesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UnassignIpv6AddressesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnassignIpv6AddressesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UnassignIpv6AddressesInput"}
	if s.Ipv6Address == nil {
		invalidParams.Add(request.NewErrParamRequired("Ipv6Address"))
	}
	if s.NetworkInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIpv6Address sets the Ipv6Address field's value.
func (s *UnassignIpv6AddressesInput) SetIpv6Address(v []*string) *UnassignIpv6AddressesInput {
	s.Ipv6Address = v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *UnassignIpv6AddressesInput) SetNetworkInterfaceId(v string) *UnassignIpv6AddressesInput {
	s.NetworkInterfaceId = &v
	return s
}

type UnassignIpv6AddressesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s UnassignIpv6AddressesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UnassignIpv6AddressesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *UnassignIpv6AddressesOutput) SetRequestId(v string) *UnassignIpv6AddressesOutput {
	s.RequestId = &v
	return s
}
