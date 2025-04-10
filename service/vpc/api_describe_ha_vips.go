// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeHaVipsCommon = "DescribeHaVips"

// DescribeHaVipsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeHaVipsCommon operation. The "output" return
// value will be populated with the DescribeHaVipsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeHaVipsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeHaVipsCommon Send returns without error.
//
// See DescribeHaVipsCommon for more information on using the DescribeHaVipsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeHaVipsCommonRequest method.
//    req, resp := client.DescribeHaVipsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeHaVipsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeHaVipsCommon,
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

// DescribeHaVipsCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeHaVipsCommon for usage and error information.
func (c *VPC) DescribeHaVipsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeHaVipsCommonRequest(input)
	return out, req.Send()
}

// DescribeHaVipsCommonWithContext is the same as DescribeHaVipsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeHaVipsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeHaVipsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeHaVipsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeHaVips = "DescribeHaVips"

// DescribeHaVipsRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeHaVips operation. The "output" return
// value will be populated with the DescribeHaVipsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeHaVipsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeHaVipsCommon Send returns without error.
//
// See DescribeHaVips for more information on using the DescribeHaVips
// API call, and error handling.
//
//    // Example sending a request using the DescribeHaVipsRequest method.
//    req, resp := client.DescribeHaVipsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeHaVipsRequest(input *DescribeHaVipsInput) (req *request.Request, output *DescribeHaVipsOutput) {
	op := &request.Operation{
		Name:       opDescribeHaVips,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHaVipsInput{}
	}

	output = &DescribeHaVipsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeHaVips API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation DescribeHaVips for usage and error information.
func (c *VPC) DescribeHaVips(input *DescribeHaVipsInput) (*DescribeHaVipsOutput, error) {
	req, out := c.DescribeHaVipsRequest(input)
	return out, req.Send()
}

// DescribeHaVipsWithContext is the same as DescribeHaVips with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeHaVips for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeHaVipsWithContext(ctx byteplus.Context, input *DescribeHaVipsInput, opts ...request.Option) (*DescribeHaVipsOutput, error) {
	req, out := c.DescribeHaVipsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeHaVipsInput struct {
	_ struct{} `type:"structure"`

	HaVipIds []*string `type:"list"`

	HaVipName *string `type:"string"`

	IpAddress *string `type:"string"`

	MaxResults *int64 `min:"1" max:"100" type:"integer"`

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	TagFilters []*TagFilterForDescribeHaVipsInput `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeHaVipsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHaVipsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeHaVipsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeHaVipsInput"}
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

// SetHaVipIds sets the HaVipIds field's value.
func (s *DescribeHaVipsInput) SetHaVipIds(v []*string) *DescribeHaVipsInput {
	s.HaVipIds = v
	return s
}

// SetHaVipName sets the HaVipName field's value.
func (s *DescribeHaVipsInput) SetHaVipName(v string) *DescribeHaVipsInput {
	s.HaVipName = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *DescribeHaVipsInput) SetIpAddress(v string) *DescribeHaVipsInput {
	s.IpAddress = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeHaVipsInput) SetMaxResults(v int64) *DescribeHaVipsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeHaVipsInput) SetNextToken(v string) *DescribeHaVipsInput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeHaVipsInput) SetPageNumber(v int64) *DescribeHaVipsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeHaVipsInput) SetPageSize(v int64) *DescribeHaVipsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeHaVipsInput) SetProjectName(v string) *DescribeHaVipsInput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeHaVipsInput) SetStatus(v string) *DescribeHaVipsInput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeHaVipsInput) SetSubnetId(v string) *DescribeHaVipsInput {
	s.SubnetId = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeHaVipsInput) SetTagFilters(v []*TagFilterForDescribeHaVipsInput) *DescribeHaVipsInput {
	s.TagFilters = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeHaVipsInput) SetVpcId(v string) *DescribeHaVipsInput {
	s.VpcId = &v
	return s
}

type DescribeHaVipsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	HaVips []*HaVipForDescribeHaVipsOutput `type:"list"`

	NextToken *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeHaVipsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHaVipsOutput) GoString() string {
	return s.String()
}

// SetHaVips sets the HaVips field's value.
func (s *DescribeHaVipsOutput) SetHaVips(v []*HaVipForDescribeHaVipsOutput) *DescribeHaVipsOutput {
	s.HaVips = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeHaVipsOutput) SetNextToken(v string) *DescribeHaVipsOutput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeHaVipsOutput) SetPageNumber(v int64) *DescribeHaVipsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeHaVipsOutput) SetPageSize(v int64) *DescribeHaVipsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeHaVipsOutput) SetRequestId(v string) *DescribeHaVipsOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeHaVipsOutput) SetTotalCount(v int64) *DescribeHaVipsOutput {
	s.TotalCount = &v
	return s
}

type HaVipForDescribeHaVipsOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	AssociatedEipAddress *string `type:"string"`

	AssociatedEipId *string `type:"string"`

	AssociatedInstanceIds []*string `type:"list"`

	AssociatedInstanceType *string `type:"string"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	HaVipId *string `type:"string"`

	HaVipName *string `type:"string"`

	IpAddress *string `type:"string"`

	MasterInstanceId *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags []*TagForDescribeHaVipsOutput `type:"list"`

	UpdatedAt *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s HaVipForDescribeHaVipsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s HaVipForDescribeHaVipsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *HaVipForDescribeHaVipsOutput) SetAccountId(v string) *HaVipForDescribeHaVipsOutput {
	s.AccountId = &v
	return s
}

// SetAssociatedEipAddress sets the AssociatedEipAddress field's value.
func (s *HaVipForDescribeHaVipsOutput) SetAssociatedEipAddress(v string) *HaVipForDescribeHaVipsOutput {
	s.AssociatedEipAddress = &v
	return s
}

// SetAssociatedEipId sets the AssociatedEipId field's value.
func (s *HaVipForDescribeHaVipsOutput) SetAssociatedEipId(v string) *HaVipForDescribeHaVipsOutput {
	s.AssociatedEipId = &v
	return s
}

// SetAssociatedInstanceIds sets the AssociatedInstanceIds field's value.
func (s *HaVipForDescribeHaVipsOutput) SetAssociatedInstanceIds(v []*string) *HaVipForDescribeHaVipsOutput {
	s.AssociatedInstanceIds = v
	return s
}

// SetAssociatedInstanceType sets the AssociatedInstanceType field's value.
func (s *HaVipForDescribeHaVipsOutput) SetAssociatedInstanceType(v string) *HaVipForDescribeHaVipsOutput {
	s.AssociatedInstanceType = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *HaVipForDescribeHaVipsOutput) SetCreatedAt(v string) *HaVipForDescribeHaVipsOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *HaVipForDescribeHaVipsOutput) SetDescription(v string) *HaVipForDescribeHaVipsOutput {
	s.Description = &v
	return s
}

// SetHaVipId sets the HaVipId field's value.
func (s *HaVipForDescribeHaVipsOutput) SetHaVipId(v string) *HaVipForDescribeHaVipsOutput {
	s.HaVipId = &v
	return s
}

// SetHaVipName sets the HaVipName field's value.
func (s *HaVipForDescribeHaVipsOutput) SetHaVipName(v string) *HaVipForDescribeHaVipsOutput {
	s.HaVipName = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *HaVipForDescribeHaVipsOutput) SetIpAddress(v string) *HaVipForDescribeHaVipsOutput {
	s.IpAddress = &v
	return s
}

// SetMasterInstanceId sets the MasterInstanceId field's value.
func (s *HaVipForDescribeHaVipsOutput) SetMasterInstanceId(v string) *HaVipForDescribeHaVipsOutput {
	s.MasterInstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *HaVipForDescribeHaVipsOutput) SetProjectName(v string) *HaVipForDescribeHaVipsOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *HaVipForDescribeHaVipsOutput) SetStatus(v string) *HaVipForDescribeHaVipsOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *HaVipForDescribeHaVipsOutput) SetSubnetId(v string) *HaVipForDescribeHaVipsOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *HaVipForDescribeHaVipsOutput) SetTags(v []*TagForDescribeHaVipsOutput) *HaVipForDescribeHaVipsOutput {
	s.Tags = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *HaVipForDescribeHaVipsOutput) SetUpdatedAt(v string) *HaVipForDescribeHaVipsOutput {
	s.UpdatedAt = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *HaVipForDescribeHaVipsOutput) SetVpcId(v string) *HaVipForDescribeHaVipsOutput {
	s.VpcId = &v
	return s
}

type TagFilterForDescribeHaVipsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeHaVipsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeHaVipsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeHaVipsInput) SetKey(v string) *TagFilterForDescribeHaVipsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeHaVipsInput) SetValues(v []*string) *TagFilterForDescribeHaVipsInput {
	s.Values = v
	return s
}

type TagForDescribeHaVipsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeHaVipsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeHaVipsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeHaVipsOutput) SetKey(v string) *TagForDescribeHaVipsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeHaVipsOutput) SetValue(v string) *TagForDescribeHaVipsOutput {
	s.Value = &v
	return s
}
