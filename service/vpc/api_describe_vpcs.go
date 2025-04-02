// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeVpcsCommon = "DescribeVpcs"

// DescribeVpcsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeVpcsCommon operation. The "output" return
// value will be populated with the DescribeVpcsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcsCommon Send returns without error.
//
// See DescribeVpcsCommon for more information on using the DescribeVpcsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcsCommonRequest method.
//    req, resp := client.DescribeVpcsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeVpcsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcsCommon,
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

// DescribeVpcsCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeVpcsCommon for usage and error information.
func (c *VPC) DescribeVpcsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcsCommonRequest(input)
	return out, req.Send()
}

// DescribeVpcsCommonWithContext is the same as DescribeVpcsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeVpcsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcs = "DescribeVpcs"

// DescribeVpcsRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeVpcs operation. The "output" return
// value will be populated with the DescribeVpcsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcsCommon Send returns without error.
//
// See DescribeVpcs for more information on using the DescribeVpcs
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcsRequest method.
//    req, resp := client.DescribeVpcsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeVpcsRequest(input *DescribeVpcsInput) (req *request.Request, output *DescribeVpcsOutput) {
	op := &request.Operation{
		Name:       opDescribeVpcs,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcsInput{}
	}

	output = &DescribeVpcsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpcs API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeVpcs for usage and error information.
func (c *VPC) DescribeVpcs(input *DescribeVpcsInput) (*DescribeVpcsOutput, error) {
	req, out := c.DescribeVpcsRequest(input)
	return out, req.Send()
}

// DescribeVpcsWithContext is the same as DescribeVpcs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeVpcsWithContext(ctx byteplus.Context, input *DescribeVpcsInput, opts ...request.Option) (*DescribeVpcsOutput, error) {
	req, out := c.DescribeVpcsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateCenForDescribeVpcsOutput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	CenOwnerId *string `type:"string"`

	CenStatus *string `type:"string"`
}

// String returns the string representation
func (s AssociateCenForDescribeVpcsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateCenForDescribeVpcsOutput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *AssociateCenForDescribeVpcsOutput) SetCenId(v string) *AssociateCenForDescribeVpcsOutput {
	s.CenId = &v
	return s
}

// SetCenOwnerId sets the CenOwnerId field's value.
func (s *AssociateCenForDescribeVpcsOutput) SetCenOwnerId(v string) *AssociateCenForDescribeVpcsOutput {
	s.CenOwnerId = &v
	return s
}

// SetCenStatus sets the CenStatus field's value.
func (s *AssociateCenForDescribeVpcsOutput) SetCenStatus(v string) *AssociateCenForDescribeVpcsOutput {
	s.CenStatus = &v
	return s
}

type DescribeVpcsInput struct {
	_ struct{} `type:"structure"`

	IsDefault *bool `type:"boolean"`

	MaxResults *int64 `min:"1" max:"100" type:"integer"`

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	ProjectName *string `type:"string"`

	TagFilters []*TagFilterForDescribeVpcsInput `type:"list"`

	VpcIds []*string `type:"list"`

	VpcName *string `type:"string"`

	VpcOwnerId *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeVpcsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVpcsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MaxResults != nil && *s.MaxResults > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("MaxResults", 100))
	}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIsDefault sets the IsDefault field's value.
func (s *DescribeVpcsInput) SetIsDefault(v bool) *DescribeVpcsInput {
	s.IsDefault = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeVpcsInput) SetMaxResults(v int64) *DescribeVpcsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeVpcsInput) SetNextToken(v string) *DescribeVpcsInput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcsInput) SetPageNumber(v int64) *DescribeVpcsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcsInput) SetPageSize(v int64) *DescribeVpcsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeVpcsInput) SetProjectName(v string) *DescribeVpcsInput {
	s.ProjectName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeVpcsInput) SetTagFilters(v []*TagFilterForDescribeVpcsInput) *DescribeVpcsInput {
	s.TagFilters = v
	return s
}

// SetVpcIds sets the VpcIds field's value.
func (s *DescribeVpcsInput) SetVpcIds(v []*string) *DescribeVpcsInput {
	s.VpcIds = v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *DescribeVpcsInput) SetVpcName(v string) *DescribeVpcsInput {
	s.VpcName = &v
	return s
}

// SetVpcOwnerId sets the VpcOwnerId field's value.
func (s *DescribeVpcsInput) SetVpcOwnerId(v int64) *DescribeVpcsInput {
	s.VpcOwnerId = &v
	return s
}

type DescribeVpcsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`

	Vpcs []*VpcForDescribeVpcsOutput `type:"list"`
}

// String returns the string representation
func (s DescribeVpcsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeVpcsOutput) SetNextToken(v string) *DescribeVpcsOutput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcsOutput) SetPageNumber(v int64) *DescribeVpcsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcsOutput) SetPageSize(v int64) *DescribeVpcsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpcsOutput) SetRequestId(v string) *DescribeVpcsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVpcsOutput) SetTotalCount(v int64) *DescribeVpcsOutput {
	s.TotalCount = &v
	return s
}

// SetVpcs sets the Vpcs field's value.
func (s *DescribeVpcsOutput) SetVpcs(v []*VpcForDescribeVpcsOutput) *DescribeVpcsOutput {
	s.Vpcs = v
	return s
}

type TagFilterForDescribeVpcsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeVpcsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeVpcsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeVpcsInput) SetKey(v string) *TagFilterForDescribeVpcsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeVpcsInput) SetValues(v []*string) *TagFilterForDescribeVpcsInput {
	s.Values = v
	return s
}

type TagForDescribeVpcsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeVpcsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeVpcsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeVpcsOutput) SetKey(v string) *TagForDescribeVpcsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeVpcsOutput) SetValue(v string) *TagForDescribeVpcsOutput {
	s.Value = &v
	return s
}

type VpcForDescribeVpcsOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	AssociateCens []*AssociateCenForDescribeVpcsOutput `type:"list"`

	CidrBlock *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	DnsServers []*string `type:"list"`

	Ipv4GatewayId *string `type:"string"`

	IsDefault *bool `type:"boolean"`

	NatGatewayIds []*string `type:"list"`

	NetworkAclNum *string `type:"string"`

	ProjectName *string `type:"string"`

	RouteTableIds []*string `type:"list"`

	SecondaryCidrBlocks []*string `type:"list"`

	SecurityGroupIds []*string `type:"list"`

	Status *string `type:"string"`

	SubnetIds []*string `type:"list"`

	SupportIpv4Gateway *bool `type:"boolean"`

	Tags []*TagForDescribeVpcsOutput `type:"list"`

	UpdateTime *string `type:"string"`

	UserCidrBlocks []*string `type:"list"`

	VpcId *string `type:"string"`

	VpcName *string `type:"string"`
}

// String returns the string representation
func (s VpcForDescribeVpcsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcForDescribeVpcsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *VpcForDescribeVpcsOutput) SetAccountId(v string) *VpcForDescribeVpcsOutput {
	s.AccountId = &v
	return s
}

// SetAssociateCens sets the AssociateCens field's value.
func (s *VpcForDescribeVpcsOutput) SetAssociateCens(v []*AssociateCenForDescribeVpcsOutput) *VpcForDescribeVpcsOutput {
	s.AssociateCens = v
	return s
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *VpcForDescribeVpcsOutput) SetCidrBlock(v string) *VpcForDescribeVpcsOutput {
	s.CidrBlock = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *VpcForDescribeVpcsOutput) SetCreationTime(v string) *VpcForDescribeVpcsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *VpcForDescribeVpcsOutput) SetDescription(v string) *VpcForDescribeVpcsOutput {
	s.Description = &v
	return s
}

// SetDnsServers sets the DnsServers field's value.
func (s *VpcForDescribeVpcsOutput) SetDnsServers(v []*string) *VpcForDescribeVpcsOutput {
	s.DnsServers = v
	return s
}

// SetIpv4GatewayId sets the Ipv4GatewayId field's value.
func (s *VpcForDescribeVpcsOutput) SetIpv4GatewayId(v string) *VpcForDescribeVpcsOutput {
	s.Ipv4GatewayId = &v
	return s
}

// SetIsDefault sets the IsDefault field's value.
func (s *VpcForDescribeVpcsOutput) SetIsDefault(v bool) *VpcForDescribeVpcsOutput {
	s.IsDefault = &v
	return s
}

// SetNatGatewayIds sets the NatGatewayIds field's value.
func (s *VpcForDescribeVpcsOutput) SetNatGatewayIds(v []*string) *VpcForDescribeVpcsOutput {
	s.NatGatewayIds = v
	return s
}

// SetNetworkAclNum sets the NetworkAclNum field's value.
func (s *VpcForDescribeVpcsOutput) SetNetworkAclNum(v string) *VpcForDescribeVpcsOutput {
	s.NetworkAclNum = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *VpcForDescribeVpcsOutput) SetProjectName(v string) *VpcForDescribeVpcsOutput {
	s.ProjectName = &v
	return s
}

// SetRouteTableIds sets the RouteTableIds field's value.
func (s *VpcForDescribeVpcsOutput) SetRouteTableIds(v []*string) *VpcForDescribeVpcsOutput {
	s.RouteTableIds = v
	return s
}

// SetSecondaryCidrBlocks sets the SecondaryCidrBlocks field's value.
func (s *VpcForDescribeVpcsOutput) SetSecondaryCidrBlocks(v []*string) *VpcForDescribeVpcsOutput {
	s.SecondaryCidrBlocks = v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *VpcForDescribeVpcsOutput) SetSecurityGroupIds(v []*string) *VpcForDescribeVpcsOutput {
	s.SecurityGroupIds = v
	return s
}

// SetStatus sets the Status field's value.
func (s *VpcForDescribeVpcsOutput) SetStatus(v string) *VpcForDescribeVpcsOutput {
	s.Status = &v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *VpcForDescribeVpcsOutput) SetSubnetIds(v []*string) *VpcForDescribeVpcsOutput {
	s.SubnetIds = v
	return s
}

// SetSupportIpv4Gateway sets the SupportIpv4Gateway field's value.
func (s *VpcForDescribeVpcsOutput) SetSupportIpv4Gateway(v bool) *VpcForDescribeVpcsOutput {
	s.SupportIpv4Gateway = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *VpcForDescribeVpcsOutput) SetTags(v []*TagForDescribeVpcsOutput) *VpcForDescribeVpcsOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *VpcForDescribeVpcsOutput) SetUpdateTime(v string) *VpcForDescribeVpcsOutput {
	s.UpdateTime = &v
	return s
}

// SetUserCidrBlocks sets the UserCidrBlocks field's value.
func (s *VpcForDescribeVpcsOutput) SetUserCidrBlocks(v []*string) *VpcForDescribeVpcsOutput {
	s.UserCidrBlocks = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *VpcForDescribeVpcsOutput) SetVpcId(v string) *VpcForDescribeVpcsOutput {
	s.VpcId = &v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *VpcForDescribeVpcsOutput) SetVpcName(v string) *VpcForDescribeVpcsOutput {
	s.VpcName = &v
	return s
}
