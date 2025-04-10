// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeIpv6AddressBandwidthsCommon = "DescribeIpv6AddressBandwidths"

// DescribeIpv6AddressBandwidthsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeIpv6AddressBandwidthsCommon operation. The "output" return
// value will be populated with the DescribeIpv6AddressBandwidthsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpv6AddressBandwidthsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpv6AddressBandwidthsCommon Send returns without error.
//
// See DescribeIpv6AddressBandwidthsCommon for more information on using the DescribeIpv6AddressBandwidthsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpv6AddressBandwidthsCommonRequest method.
//    req, resp := client.DescribeIpv6AddressBandwidthsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpv6AddressBandwidthsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeIpv6AddressBandwidthsCommon,
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

// DescribeIpv6AddressBandwidthsCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeIpv6AddressBandwidthsCommon for usage and error information.
func (c *VPC) DescribeIpv6AddressBandwidthsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeIpv6AddressBandwidthsCommonRequest(input)
	return out, req.Send()
}

// DescribeIpv6AddressBandwidthsCommonWithContext is the same as DescribeIpv6AddressBandwidthsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpv6AddressBandwidthsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpv6AddressBandwidthsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeIpv6AddressBandwidthsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeIpv6AddressBandwidths = "DescribeIpv6AddressBandwidths"

// DescribeIpv6AddressBandwidthsRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeIpv6AddressBandwidths operation. The "output" return
// value will be populated with the DescribeIpv6AddressBandwidthsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpv6AddressBandwidthsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpv6AddressBandwidthsCommon Send returns without error.
//
// See DescribeIpv6AddressBandwidths for more information on using the DescribeIpv6AddressBandwidths
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpv6AddressBandwidthsRequest method.
//    req, resp := client.DescribeIpv6AddressBandwidthsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpv6AddressBandwidthsRequest(input *DescribeIpv6AddressBandwidthsInput) (req *request.Request, output *DescribeIpv6AddressBandwidthsOutput) {
	op := &request.Operation{
		Name:       opDescribeIpv6AddressBandwidths,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeIpv6AddressBandwidthsInput{}
	}

	output = &DescribeIpv6AddressBandwidthsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeIpv6AddressBandwidths API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeIpv6AddressBandwidths for usage and error information.
func (c *VPC) DescribeIpv6AddressBandwidths(input *DescribeIpv6AddressBandwidthsInput) (*DescribeIpv6AddressBandwidthsOutput, error) {
	req, out := c.DescribeIpv6AddressBandwidthsRequest(input)
	return out, req.Send()
}

// DescribeIpv6AddressBandwidthsWithContext is the same as DescribeIpv6AddressBandwidths with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpv6AddressBandwidths for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpv6AddressBandwidthsWithContext(ctx byteplus.Context, input *DescribeIpv6AddressBandwidthsInput, opts ...request.Option) (*DescribeIpv6AddressBandwidthsOutput, error) {
	req, out := c.DescribeIpv6AddressBandwidthsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeIpv6AddressBandwidthsInput struct {
	_ struct{} `type:"structure"`

	AllocationIds *string `type:"string"`

	AssociatedInstanceId *string `type:"string"`

	AssociatedInstanceType *string `type:"string"`

	BandwidthPackageId *string `type:"string"`

	ISP *string `type:"string"`

	Ipv6Addresses *string `type:"string"`

	MaxResults *int64 `type:"integer"`

	NetworkType *string `type:"string"`

	NextToken *string `type:"string"`

	ProjectName *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeIpv6AddressBandwidthsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpv6AddressBandwidthsInput) GoString() string {
	return s.String()
}

// SetAllocationIds sets the AllocationIds field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetAllocationIds(v string) *DescribeIpv6AddressBandwidthsInput {
	s.AllocationIds = &v
	return s
}

// SetAssociatedInstanceId sets the AssociatedInstanceId field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetAssociatedInstanceId(v string) *DescribeIpv6AddressBandwidthsInput {
	s.AssociatedInstanceId = &v
	return s
}

// SetAssociatedInstanceType sets the AssociatedInstanceType field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetAssociatedInstanceType(v string) *DescribeIpv6AddressBandwidthsInput {
	s.AssociatedInstanceType = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetBandwidthPackageId(v string) *DescribeIpv6AddressBandwidthsInput {
	s.BandwidthPackageId = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetISP(v string) *DescribeIpv6AddressBandwidthsInput {
	s.ISP = &v
	return s
}

// SetIpv6Addresses sets the Ipv6Addresses field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetIpv6Addresses(v string) *DescribeIpv6AddressBandwidthsInput {
	s.Ipv6Addresses = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetMaxResults(v int64) *DescribeIpv6AddressBandwidthsInput {
	s.MaxResults = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetNetworkType(v string) *DescribeIpv6AddressBandwidthsInput {
	s.NetworkType = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetNextToken(v string) *DescribeIpv6AddressBandwidthsInput {
	s.NextToken = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetProjectName(v string) *DescribeIpv6AddressBandwidthsInput {
	s.ProjectName = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeIpv6AddressBandwidthsInput) SetVpcId(v string) *DescribeIpv6AddressBandwidthsInput {
	s.VpcId = &v
	return s
}

type DescribeIpv6AddressBandwidthsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Ipv6AddressBandwidths []*Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput `type:"list"`

	NextToken *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DescribeIpv6AddressBandwidthsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpv6AddressBandwidthsOutput) GoString() string {
	return s.String()
}

// SetIpv6AddressBandwidths sets the Ipv6AddressBandwidths field's value.
func (s *DescribeIpv6AddressBandwidthsOutput) SetIpv6AddressBandwidths(v []*Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) *DescribeIpv6AddressBandwidthsOutput {
	s.Ipv6AddressBandwidths = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeIpv6AddressBandwidthsOutput) SetNextToken(v string) *DescribeIpv6AddressBandwidthsOutput {
	s.NextToken = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeIpv6AddressBandwidthsOutput) SetRequestId(v string) *DescribeIpv6AddressBandwidthsOutput {
	s.RequestId = &v
	return s
}

type Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	Bandwidth *string `type:"string"`

	BandwidthPackageId *string `type:"string"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	ISP *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceType *string `type:"string"`

	Ipv6Address *string `type:"string"`

	Ipv6GatewayId *string `type:"string"`

	LockReason *string `type:"string"`

	NetworkType *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetAllocationId(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.AllocationId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetBandwidth(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetBandwidthPackageId(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetBillingType(v int64) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetBusinessStatus(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetCreationTime(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetDeletedTime(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.DeletedTime = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetISP(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.ISP = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetInstanceId(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetInstanceType(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.InstanceType = &v
	return s
}

// SetIpv6Address sets the Ipv6Address field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetIpv6Address(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.Ipv6Address = &v
	return s
}

// SetIpv6GatewayId sets the Ipv6GatewayId field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetIpv6GatewayId(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.Ipv6GatewayId = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetLockReason(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.LockReason = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetNetworkType(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.NetworkType = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetOverdueTime(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetProjectName(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetStatus(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.Status = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput) SetUpdatedAt(v string) *Ipv6AddressBandwidthForDescribeIpv6AddressBandwidthsOutput {
	s.UpdatedAt = &v
	return s
}
