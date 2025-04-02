// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeIpv6AddressBandwidthAttributesCommon = "DescribeIpv6AddressBandwidthAttributes"

// DescribeIpv6AddressBandwidthAttributesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeIpv6AddressBandwidthAttributesCommon operation. The "output" return
// value will be populated with the DescribeIpv6AddressBandwidthAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpv6AddressBandwidthAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpv6AddressBandwidthAttributesCommon Send returns without error.
//
// See DescribeIpv6AddressBandwidthAttributesCommon for more information on using the DescribeIpv6AddressBandwidthAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpv6AddressBandwidthAttributesCommonRequest method.
//    req, resp := client.DescribeIpv6AddressBandwidthAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpv6AddressBandwidthAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeIpv6AddressBandwidthAttributesCommon,
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

// DescribeIpv6AddressBandwidthAttributesCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeIpv6AddressBandwidthAttributesCommon for usage and error information.
func (c *VPC) DescribeIpv6AddressBandwidthAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeIpv6AddressBandwidthAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeIpv6AddressBandwidthAttributesCommonWithContext is the same as DescribeIpv6AddressBandwidthAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpv6AddressBandwidthAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpv6AddressBandwidthAttributesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeIpv6AddressBandwidthAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeIpv6AddressBandwidthAttributes = "DescribeIpv6AddressBandwidthAttributes"

// DescribeIpv6AddressBandwidthAttributesRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeIpv6AddressBandwidthAttributes operation. The "output" return
// value will be populated with the DescribeIpv6AddressBandwidthAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpv6AddressBandwidthAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpv6AddressBandwidthAttributesCommon Send returns without error.
//
// See DescribeIpv6AddressBandwidthAttributes for more information on using the DescribeIpv6AddressBandwidthAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpv6AddressBandwidthAttributesRequest method.
//    req, resp := client.DescribeIpv6AddressBandwidthAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpv6AddressBandwidthAttributesRequest(input *DescribeIpv6AddressBandwidthAttributesInput) (req *request.Request, output *DescribeIpv6AddressBandwidthAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeIpv6AddressBandwidthAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeIpv6AddressBandwidthAttributesInput{}
	}

	output = &DescribeIpv6AddressBandwidthAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeIpv6AddressBandwidthAttributes API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeIpv6AddressBandwidthAttributes for usage and error information.
func (c *VPC) DescribeIpv6AddressBandwidthAttributes(input *DescribeIpv6AddressBandwidthAttributesInput) (*DescribeIpv6AddressBandwidthAttributesOutput, error) {
	req, out := c.DescribeIpv6AddressBandwidthAttributesRequest(input)
	return out, req.Send()
}

// DescribeIpv6AddressBandwidthAttributesWithContext is the same as DescribeIpv6AddressBandwidthAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpv6AddressBandwidthAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpv6AddressBandwidthAttributesWithContext(ctx byteplus.Context, input *DescribeIpv6AddressBandwidthAttributesInput, opts ...request.Option) (*DescribeIpv6AddressBandwidthAttributesOutput, error) {
	req, out := c.DescribeIpv6AddressBandwidthAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeIpv6AddressBandwidthAttributesInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeIpv6AddressBandwidthAttributesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpv6AddressBandwidthAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIpv6AddressBandwidthAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeIpv6AddressBandwidthAttributesInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *DescribeIpv6AddressBandwidthAttributesInput) SetAllocationId(v string) *DescribeIpv6AddressBandwidthAttributesInput {
	s.AllocationId = &v
	return s
}

type DescribeIpv6AddressBandwidthAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AllocationId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BandwidthPackageId *string `type:"string"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeleteTime *string `type:"string"`

	ISP *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceType *string `type:"string"`

	Ipv6Address *string `type:"string"`

	Ipv6GatewayId *string `type:"string"`

	LockReason *string `type:"string"`

	NetworkType *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	ServiceManaged *bool `type:"boolean"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s DescribeIpv6AddressBandwidthAttributesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpv6AddressBandwidthAttributesOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetAllocationId(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.AllocationId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetBandwidth(v int64) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetBandwidthPackageId(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetBillingType(v int64) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetBusinessStatus(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetCreationTime(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDeleteTime sets the DeleteTime field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetDeleteTime(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.DeleteTime = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetISP(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.ISP = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetInstanceId(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetInstanceType(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.InstanceType = &v
	return s
}

// SetIpv6Address sets the Ipv6Address field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetIpv6Address(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.Ipv6Address = &v
	return s
}

// SetIpv6GatewayId sets the Ipv6GatewayId field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetIpv6GatewayId(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.Ipv6GatewayId = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetLockReason(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.LockReason = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetNetworkType(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.NetworkType = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetOverdueTime(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetProjectName(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetRequestId(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.RequestId = &v
	return s
}

// SetServiceManaged sets the ServiceManaged field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetServiceManaged(v bool) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.ServiceManaged = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetStatus(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeIpv6AddressBandwidthAttributesOutput) SetUpdateTime(v string) *DescribeIpv6AddressBandwidthAttributesOutput {
	s.UpdateTime = &v
	return s
}
