// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeInstancesCommon = "DescribeInstances"

// DescribeInstancesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInstancesCommon operation. The "output" return
// value will be populated with the DescribeInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstancesCommon Send returns without error.
//
// See DescribeInstancesCommon for more information on using the DescribeInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstancesCommonRequest method.
//    req, resp := client.DescribeInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstancesCommon,
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

// DescribeInstancesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInstancesCommon for usage and error information.
func (c *ECS) DescribeInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstancesCommonRequest(input)
	return out, req.Send()
}

// DescribeInstancesCommonWithContext is the same as DescribeInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstancesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstances = "DescribeInstances"

// DescribeInstancesRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInstances operation. The "output" return
// value will be populated with the DescribeInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstancesCommon Send returns without error.
//
// See DescribeInstances for more information on using the DescribeInstances
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstancesRequest method.
//    req, resp := client.DescribeInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstancesRequest(input *DescribeInstancesInput) (req *request.Request, output *DescribeInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstancesInput{}
	}

	output = &DescribeInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInstances API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInstances for usage and error information.
func (c *ECS) DescribeInstances(input *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	return out, req.Send()
}

// DescribeInstancesWithContext is the same as DescribeInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstancesWithContext(ctx byteplus.Context, input *DescribeInstancesInput, opts ...request.Option) (*DescribeInstancesOutput, error) {
	req, out := c.DescribeInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CpuOptionsForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	CoreCount *int32 `type:"int32"`

	ThreadsPerCore *int32 `type:"int32"`
}

// String returns the string representation
func (s CpuOptionsForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CpuOptionsForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetCoreCount sets the CoreCount field's value.
func (s *CpuOptionsForDescribeInstancesOutput) SetCoreCount(v int32) *CpuOptionsForDescribeInstancesOutput {
	s.CoreCount = &v
	return s
}

// SetThreadsPerCore sets the ThreadsPerCore field's value.
func (s *CpuOptionsForDescribeInstancesOutput) SetThreadsPerCore(v int32) *CpuOptionsForDescribeInstancesOutput {
	s.ThreadsPerCore = &v
	return s
}

type DescribeInstancesInput struct {
	_ struct{} `type:"structure"`

	AffinityGroupIds []*string `type:"list"`

	DedicatedHostClusterId *string `type:"string"`

	DedicatedHostId *string `type:"string"`

	DeploymentSetGroupNumbers []*int32 `type:"list"`

	DeploymentSetIds []*string `type:"list"`

	EipAddresses []*string `type:"list"`

	HpcClusterId *string `type:"string"`

	InstanceChargeType *string `type:"string"`

	InstanceIds []*string `type:"list"`

	InstanceName *string `type:"string"`

	InstanceTypeFamilies []*string `type:"list"`

	InstanceTypeIds []*string `type:"list"`

	InstanceTypes []*string `type:"list"`

	Ipv6Addresses []*string `type:"list"`

	KeyPairName *string `type:"string"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	PrimaryIpAddress *string `type:"string"`

	ProjectName *string `type:"string"`

	ScheduledInstanceId *string `type:"string"`

	Status *string `type:"string"`

	TagFilters []*TagFilterForDescribeInstancesInput `type:"list"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancesInput) GoString() string {
	return s.String()
}

// SetAffinityGroupIds sets the AffinityGroupIds field's value.
func (s *DescribeInstancesInput) SetAffinityGroupIds(v []*string) *DescribeInstancesInput {
	s.AffinityGroupIds = v
	return s
}

// SetDedicatedHostClusterId sets the DedicatedHostClusterId field's value.
func (s *DescribeInstancesInput) SetDedicatedHostClusterId(v string) *DescribeInstancesInput {
	s.DedicatedHostClusterId = &v
	return s
}

// SetDedicatedHostId sets the DedicatedHostId field's value.
func (s *DescribeInstancesInput) SetDedicatedHostId(v string) *DescribeInstancesInput {
	s.DedicatedHostId = &v
	return s
}

// SetDeploymentSetGroupNumbers sets the DeploymentSetGroupNumbers field's value.
func (s *DescribeInstancesInput) SetDeploymentSetGroupNumbers(v []*int32) *DescribeInstancesInput {
	s.DeploymentSetGroupNumbers = v
	return s
}

// SetDeploymentSetIds sets the DeploymentSetIds field's value.
func (s *DescribeInstancesInput) SetDeploymentSetIds(v []*string) *DescribeInstancesInput {
	s.DeploymentSetIds = v
	return s
}

// SetEipAddresses sets the EipAddresses field's value.
func (s *DescribeInstancesInput) SetEipAddresses(v []*string) *DescribeInstancesInput {
	s.EipAddresses = v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *DescribeInstancesInput) SetHpcClusterId(v string) *DescribeInstancesInput {
	s.HpcClusterId = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *DescribeInstancesInput) SetInstanceChargeType(v string) *DescribeInstancesInput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DescribeInstancesInput) SetInstanceIds(v []*string) *DescribeInstancesInput {
	s.InstanceIds = v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DescribeInstancesInput) SetInstanceName(v string) *DescribeInstancesInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeFamilies sets the InstanceTypeFamilies field's value.
func (s *DescribeInstancesInput) SetInstanceTypeFamilies(v []*string) *DescribeInstancesInput {
	s.InstanceTypeFamilies = v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *DescribeInstancesInput) SetInstanceTypeIds(v []*string) *DescribeInstancesInput {
	s.InstanceTypeIds = v
	return s
}

// SetInstanceTypes sets the InstanceTypes field's value.
func (s *DescribeInstancesInput) SetInstanceTypes(v []*string) *DescribeInstancesInput {
	s.InstanceTypes = v
	return s
}

// SetIpv6Addresses sets the Ipv6Addresses field's value.
func (s *DescribeInstancesInput) SetIpv6Addresses(v []*string) *DescribeInstancesInput {
	s.Ipv6Addresses = v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *DescribeInstancesInput) SetKeyPairName(v string) *DescribeInstancesInput {
	s.KeyPairName = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeInstancesInput) SetMaxResults(v int32) *DescribeInstancesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstancesInput) SetNextToken(v string) *DescribeInstancesInput {
	s.NextToken = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *DescribeInstancesInput) SetPrimaryIpAddress(v string) *DescribeInstancesInput {
	s.PrimaryIpAddress = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeInstancesInput) SetProjectName(v string) *DescribeInstancesInput {
	s.ProjectName = &v
	return s
}

// SetScheduledInstanceId sets the ScheduledInstanceId field's value.
func (s *DescribeInstancesInput) SetScheduledInstanceId(v string) *DescribeInstancesInput {
	s.ScheduledInstanceId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeInstancesInput) SetStatus(v string) *DescribeInstancesInput {
	s.Status = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeInstancesInput) SetTagFilters(v []*TagFilterForDescribeInstancesInput) *DescribeInstancesInput {
	s.TagFilters = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeInstancesInput) SetVpcId(v string) *DescribeInstancesInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeInstancesInput) SetZoneId(v string) *DescribeInstancesInput {
	s.ZoneId = &v
	return s
}

type DescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Instances []*InstanceForDescribeInstancesOutput `type:"list"`

	NextToken *string `type:"string"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetInstances sets the Instances field's value.
func (s *DescribeInstancesOutput) SetInstances(v []*InstanceForDescribeInstancesOutput) *DescribeInstancesOutput {
	s.Instances = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstancesOutput) SetNextToken(v string) *DescribeInstancesOutput {
	s.NextToken = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeInstancesOutput) SetTotalCount(v int32) *DescribeInstancesOutput {
	s.TotalCount = &v
	return s
}

type EipAddressForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	IpAddress *string `type:"string"`
}

// String returns the string representation
func (s EipAddressForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s EipAddressForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *EipAddressForDescribeInstancesOutput) SetAllocationId(v string) *EipAddressForDescribeInstancesOutput {
	s.AllocationId = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *EipAddressForDescribeInstancesOutput) SetIpAddress(v string) *EipAddressForDescribeInstancesOutput {
	s.IpAddress = &v
	return s
}

type InstanceForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	AffinityGroupId *string `type:"string"`

	CpuOptions *CpuOptionsForDescribeInstancesOutput `type:"structure"`

	Cpus *int32 `type:"int32"`

	CreatedAt *string `type:"string"`

	DeletionProtection *bool `type:"boolean"`

	DeploymentSetGroupNumber *int32 `type:"int32"`

	DeploymentSetId *string `type:"string"`

	Description *string `type:"string"`

	EipAddress *EipAddressForDescribeInstancesOutput `type:"structure"`

	ElasticScheduledInstanceType *string `type:"string"`

	ExpiredAt *string `type:"string"`

	HostName *string `type:"string"`

	Hostname *string `type:"string"`

	HpcClusterId *string `type:"string"`

	ImageId *string `type:"string"`

	InstanceChargeType *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	KeyPairId *string `type:"string"`

	KeyPairName *string `type:"string"`

	LocalVolumes []*LocalVolumeForDescribeInstancesOutput `type:"list"`

	MemorySize *int32 `type:"int32"`

	MetadataOptions *MetadataOptionsForDescribeInstancesOutput `type:"structure"`

	NetworkInterfaces []*NetworkInterfaceForDescribeInstancesOutput `type:"list"`

	OsName *string `type:"string"`

	OsType *string `type:"string"`

	Placement *PlacementForDescribeInstancesOutput `type:"structure"`

	ProjectName *string `type:"string"`

	RdmaIpAddresses []*string `type:"list"`

	ScheduledInstanceId *string `type:"string"`

	SpotPriceLimit *float64 `type:"float"`

	SpotStrategy *string `type:"string"`

	Status *string `type:"string"`

	StoppedMode *string `type:"string"`

	Tags []*TagForDescribeInstancesOutput `type:"list"`

	UpdatedAt *string `type:"string"`

	Uuid *string `type:"string"`

	Volumes []*VolumeForDescribeInstancesOutput `type:"list"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s InstanceForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetAffinityGroupId sets the AffinityGroupId field's value.
func (s *InstanceForDescribeInstancesOutput) SetAffinityGroupId(v string) *InstanceForDescribeInstancesOutput {
	s.AffinityGroupId = &v
	return s
}

// SetCpuOptions sets the CpuOptions field's value.
func (s *InstanceForDescribeInstancesOutput) SetCpuOptions(v *CpuOptionsForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.CpuOptions = v
	return s
}

// SetCpus sets the Cpus field's value.
func (s *InstanceForDescribeInstancesOutput) SetCpus(v int32) *InstanceForDescribeInstancesOutput {
	s.Cpus = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *InstanceForDescribeInstancesOutput) SetCreatedAt(v string) *InstanceForDescribeInstancesOutput {
	s.CreatedAt = &v
	return s
}

// SetDeletionProtection sets the DeletionProtection field's value.
func (s *InstanceForDescribeInstancesOutput) SetDeletionProtection(v bool) *InstanceForDescribeInstancesOutput {
	s.DeletionProtection = &v
	return s
}

// SetDeploymentSetGroupNumber sets the DeploymentSetGroupNumber field's value.
func (s *InstanceForDescribeInstancesOutput) SetDeploymentSetGroupNumber(v int32) *InstanceForDescribeInstancesOutput {
	s.DeploymentSetGroupNumber = &v
	return s
}

// SetDeploymentSetId sets the DeploymentSetId field's value.
func (s *InstanceForDescribeInstancesOutput) SetDeploymentSetId(v string) *InstanceForDescribeInstancesOutput {
	s.DeploymentSetId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *InstanceForDescribeInstancesOutput) SetDescription(v string) *InstanceForDescribeInstancesOutput {
	s.Description = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *InstanceForDescribeInstancesOutput) SetEipAddress(v *EipAddressForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.EipAddress = v
	return s
}

// SetElasticScheduledInstanceType sets the ElasticScheduledInstanceType field's value.
func (s *InstanceForDescribeInstancesOutput) SetElasticScheduledInstanceType(v string) *InstanceForDescribeInstancesOutput {
	s.ElasticScheduledInstanceType = &v
	return s
}

// SetExpiredAt sets the ExpiredAt field's value.
func (s *InstanceForDescribeInstancesOutput) SetExpiredAt(v string) *InstanceForDescribeInstancesOutput {
	s.ExpiredAt = &v
	return s
}

// SetHostName sets the HostName field's value.
func (s *InstanceForDescribeInstancesOutput) SetHostName(v string) *InstanceForDescribeInstancesOutput {
	s.HostName = &v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *InstanceForDescribeInstancesOutput) SetHostname(v string) *InstanceForDescribeInstancesOutput {
	s.Hostname = &v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *InstanceForDescribeInstancesOutput) SetHpcClusterId(v string) *InstanceForDescribeInstancesOutput {
	s.HpcClusterId = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *InstanceForDescribeInstancesOutput) SetImageId(v string) *InstanceForDescribeInstancesOutput {
	s.ImageId = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceChargeType(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceId(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceName(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *InstanceForDescribeInstancesOutput) SetInstanceTypeId(v string) *InstanceForDescribeInstancesOutput {
	s.InstanceTypeId = &v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *InstanceForDescribeInstancesOutput) SetKeyPairId(v string) *InstanceForDescribeInstancesOutput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *InstanceForDescribeInstancesOutput) SetKeyPairName(v string) *InstanceForDescribeInstancesOutput {
	s.KeyPairName = &v
	return s
}

// SetLocalVolumes sets the LocalVolumes field's value.
func (s *InstanceForDescribeInstancesOutput) SetLocalVolumes(v []*LocalVolumeForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.LocalVolumes = v
	return s
}

// SetMemorySize sets the MemorySize field's value.
func (s *InstanceForDescribeInstancesOutput) SetMemorySize(v int32) *InstanceForDescribeInstancesOutput {
	s.MemorySize = &v
	return s
}

// SetMetadataOptions sets the MetadataOptions field's value.
func (s *InstanceForDescribeInstancesOutput) SetMetadataOptions(v *MetadataOptionsForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.MetadataOptions = v
	return s
}

// SetNetworkInterfaces sets the NetworkInterfaces field's value.
func (s *InstanceForDescribeInstancesOutput) SetNetworkInterfaces(v []*NetworkInterfaceForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.NetworkInterfaces = v
	return s
}

// SetOsName sets the OsName field's value.
func (s *InstanceForDescribeInstancesOutput) SetOsName(v string) *InstanceForDescribeInstancesOutput {
	s.OsName = &v
	return s
}

// SetOsType sets the OsType field's value.
func (s *InstanceForDescribeInstancesOutput) SetOsType(v string) *InstanceForDescribeInstancesOutput {
	s.OsType = &v
	return s
}

// SetPlacement sets the Placement field's value.
func (s *InstanceForDescribeInstancesOutput) SetPlacement(v *PlacementForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.Placement = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *InstanceForDescribeInstancesOutput) SetProjectName(v string) *InstanceForDescribeInstancesOutput {
	s.ProjectName = &v
	return s
}

// SetRdmaIpAddresses sets the RdmaIpAddresses field's value.
func (s *InstanceForDescribeInstancesOutput) SetRdmaIpAddresses(v []*string) *InstanceForDescribeInstancesOutput {
	s.RdmaIpAddresses = v
	return s
}

// SetScheduledInstanceId sets the ScheduledInstanceId field's value.
func (s *InstanceForDescribeInstancesOutput) SetScheduledInstanceId(v string) *InstanceForDescribeInstancesOutput {
	s.ScheduledInstanceId = &v
	return s
}

// SetSpotPriceLimit sets the SpotPriceLimit field's value.
func (s *InstanceForDescribeInstancesOutput) SetSpotPriceLimit(v float64) *InstanceForDescribeInstancesOutput {
	s.SpotPriceLimit = &v
	return s
}

// SetSpotStrategy sets the SpotStrategy field's value.
func (s *InstanceForDescribeInstancesOutput) SetSpotStrategy(v string) *InstanceForDescribeInstancesOutput {
	s.SpotStrategy = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *InstanceForDescribeInstancesOutput) SetStatus(v string) *InstanceForDescribeInstancesOutput {
	s.Status = &v
	return s
}

// SetStoppedMode sets the StoppedMode field's value.
func (s *InstanceForDescribeInstancesOutput) SetStoppedMode(v string) *InstanceForDescribeInstancesOutput {
	s.StoppedMode = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *InstanceForDescribeInstancesOutput) SetTags(v []*TagForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.Tags = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *InstanceForDescribeInstancesOutput) SetUpdatedAt(v string) *InstanceForDescribeInstancesOutput {
	s.UpdatedAt = &v
	return s
}

// SetUuid sets the Uuid field's value.
func (s *InstanceForDescribeInstancesOutput) SetUuid(v string) *InstanceForDescribeInstancesOutput {
	s.Uuid = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *InstanceForDescribeInstancesOutput) SetVolumes(v []*VolumeForDescribeInstancesOutput) *InstanceForDescribeInstancesOutput {
	s.Volumes = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *InstanceForDescribeInstancesOutput) SetVpcId(v string) *InstanceForDescribeInstancesOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *InstanceForDescribeInstancesOutput) SetZoneId(v string) *InstanceForDescribeInstancesOutput {
	s.ZoneId = &v
	return s
}

type LocalVolumeForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Count *int32 `type:"int32"`

	Size *int32 `type:"int32"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s LocalVolumeForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s LocalVolumeForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *LocalVolumeForDescribeInstancesOutput) SetCount(v int32) *LocalVolumeForDescribeInstancesOutput {
	s.Count = &v
	return s
}

// SetSize sets the Size field's value.
func (s *LocalVolumeForDescribeInstancesOutput) SetSize(v int32) *LocalVolumeForDescribeInstancesOutput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *LocalVolumeForDescribeInstancesOutput) SetVolumeType(v string) *LocalVolumeForDescribeInstancesOutput {
	s.VolumeType = &v
	return s
}

type MetadataOptionsForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	HttpTokens *string `type:"string"`
}

// String returns the string representation
func (s MetadataOptionsForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s MetadataOptionsForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetHttpTokens sets the HttpTokens field's value.
func (s *MetadataOptionsForDescribeInstancesOutput) SetHttpTokens(v string) *MetadataOptionsForDescribeInstancesOutput {
	s.HttpTokens = &v
	return s
}

type NetworkInterfaceForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Ipv6Addresses []*string `type:"list"`

	MacAddress *string `type:"string"`

	NetworkInterfaceId *string `type:"string"`

	PrimaryIpAddress *string `type:"string"`

	SecurityGroupIds []*string `type:"list"`

	SubnetId *string `type:"string"`

	Type *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s NetworkInterfaceForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkInterfaceForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetIpv6Addresses sets the Ipv6Addresses field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetIpv6Addresses(v []*string) *NetworkInterfaceForDescribeInstancesOutput {
	s.Ipv6Addresses = v
	return s
}

// SetMacAddress sets the MacAddress field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetMacAddress(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.MacAddress = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetNetworkInterfaceId(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetPrimaryIpAddress(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.PrimaryIpAddress = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetSecurityGroupIds(v []*string) *NetworkInterfaceForDescribeInstancesOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetSubnetId(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.SubnetId = &v
	return s
}

// SetType sets the Type field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetType(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.Type = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *NetworkInterfaceForDescribeInstancesOutput) SetVpcId(v string) *NetworkInterfaceForDescribeInstancesOutput {
	s.VpcId = &v
	return s
}

type PlacementForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Affinity *string `type:"string"`

	DedicatedHostClusterId *string `type:"string"`

	DedicatedHostId *string `type:"string"`

	Tenancy *string `type:"string"`
}

// String returns the string representation
func (s PlacementForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s PlacementForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetAffinity sets the Affinity field's value.
func (s *PlacementForDescribeInstancesOutput) SetAffinity(v string) *PlacementForDescribeInstancesOutput {
	s.Affinity = &v
	return s
}

// SetDedicatedHostClusterId sets the DedicatedHostClusterId field's value.
func (s *PlacementForDescribeInstancesOutput) SetDedicatedHostClusterId(v string) *PlacementForDescribeInstancesOutput {
	s.DedicatedHostClusterId = &v
	return s
}

// SetDedicatedHostId sets the DedicatedHostId field's value.
func (s *PlacementForDescribeInstancesOutput) SetDedicatedHostId(v string) *PlacementForDescribeInstancesOutput {
	s.DedicatedHostId = &v
	return s
}

// SetTenancy sets the Tenancy field's value.
func (s *PlacementForDescribeInstancesOutput) SetTenancy(v string) *PlacementForDescribeInstancesOutput {
	s.Tenancy = &v
	return s
}

type TagFilterForDescribeInstancesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeInstancesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeInstancesInput) SetKey(v string) *TagFilterForDescribeInstancesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeInstancesInput) SetValues(v []*string) *TagFilterForDescribeInstancesInput {
	s.Values = v
	return s
}

type TagForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeInstancesOutput) SetKey(v string) *TagForDescribeInstancesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeInstancesOutput) SetValue(v string) *TagForDescribeInstancesOutput {
	s.Value = &v
	return s
}

type VolumeForDescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	VolumeId *string `type:"string"`
}

// String returns the string representation
func (s VolumeForDescribeInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForDescribeInstancesOutput) GoString() string {
	return s.String()
}

// SetVolumeId sets the VolumeId field's value.
func (s *VolumeForDescribeInstancesOutput) SetVolumeId(v string) *VolumeForDescribeInstancesOutput {
	s.VolumeId = &v
	return s
}
