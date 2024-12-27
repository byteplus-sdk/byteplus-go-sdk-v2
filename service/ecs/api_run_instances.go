// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opRunInstancesCommon = "RunInstances"

// RunInstancesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the RunInstancesCommon operation. The "output" return
// value will be populated with the RunInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RunInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after RunInstancesCommon Send returns without error.
//
// See RunInstancesCommon for more information on using the RunInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the RunInstancesCommonRequest method.
//    req, resp := client.RunInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) RunInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRunInstancesCommon,
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

// RunInstancesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation RunInstancesCommon for usage and error information.
func (c *ECS) RunInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RunInstancesCommonRequest(input)
	return out, req.Send()
}

// RunInstancesCommonWithContext is the same as RunInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RunInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RunInstancesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RunInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRunInstances = "RunInstances"

// RunInstancesRequest generates a "byteplus/request.Request" representing the
// client's request for the RunInstances operation. The "output" return
// value will be populated with the RunInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RunInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after RunInstancesCommon Send returns without error.
//
// See RunInstances for more information on using the RunInstances
// API call, and error handling.
//
//    // Example sending a request using the RunInstancesRequest method.
//    req, resp := client.RunInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) RunInstancesRequest(input *RunInstancesInput) (req *request.Request, output *RunInstancesOutput) {
	op := &request.Operation{
		Name:       opRunInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RunInstancesInput{}
	}

	output = &RunInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RunInstances API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation RunInstances for usage and error information.
func (c *ECS) RunInstances(input *RunInstancesInput) (*RunInstancesOutput, error) {
	req, out := c.RunInstancesRequest(input)
	return out, req.Send()
}

// RunInstancesWithContext is the same as RunInstances with the addition of
// the ability to pass a context and additional request options.
//
// See RunInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RunInstancesWithContext(ctx byteplus.Context, input *RunInstancesInput, opts ...request.Option) (*RunInstancesOutput, error) {
	req, out := c.RunInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConvertNetworkInterfaceForRunInstancesInput struct {
	_ struct{} `type:"structure"`

	PrimaryIpAddress *string `type:"string"`
}

// String returns the string representation
func (s ConvertNetworkInterfaceForRunInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertNetworkInterfaceForRunInstancesInput) GoString() string {
	return s.String()
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *ConvertNetworkInterfaceForRunInstancesInput) SetPrimaryIpAddress(v string) *ConvertNetworkInterfaceForRunInstancesInput {
	s.PrimaryIpAddress = &v
	return s
}

type EipAddressForRunInstancesInput struct {
	_ struct{} `type:"structure"`

	BandwidthMbps *int32 `type:"int32"`

	BandwidthPackageId *string `type:"string"`

	ChargeType *string `type:"string"`

	ISP *string `type:"string"`

	ReleaseWithInstance *bool `type:"boolean"`

	SecurityProtectionInstanceId *int32 `type:"int32"`

	SecurityProtectionTypes []*string `type:"list"`
}

// String returns the string representation
func (s EipAddressForRunInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s EipAddressForRunInstancesInput) GoString() string {
	return s.String()
}

// SetBandwidthMbps sets the BandwidthMbps field's value.
func (s *EipAddressForRunInstancesInput) SetBandwidthMbps(v int32) *EipAddressForRunInstancesInput {
	s.BandwidthMbps = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *EipAddressForRunInstancesInput) SetBandwidthPackageId(v string) *EipAddressForRunInstancesInput {
	s.BandwidthPackageId = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *EipAddressForRunInstancesInput) SetChargeType(v string) *EipAddressForRunInstancesInput {
	s.ChargeType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipAddressForRunInstancesInput) SetISP(v string) *EipAddressForRunInstancesInput {
	s.ISP = &v
	return s
}

// SetReleaseWithInstance sets the ReleaseWithInstance field's value.
func (s *EipAddressForRunInstancesInput) SetReleaseWithInstance(v bool) *EipAddressForRunInstancesInput {
	s.ReleaseWithInstance = &v
	return s
}

// SetSecurityProtectionInstanceId sets the SecurityProtectionInstanceId field's value.
func (s *EipAddressForRunInstancesInput) SetSecurityProtectionInstanceId(v int32) *EipAddressForRunInstancesInput {
	s.SecurityProtectionInstanceId = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *EipAddressForRunInstancesInput) SetSecurityProtectionTypes(v []*string) *EipAddressForRunInstancesInput {
	s.SecurityProtectionTypes = v
	return s
}

type NetworkInterfaceForRunInstancesInput struct {
	_ struct{} `type:"structure"`

	Ipv6AddressCount *int32 `type:"int32"`

	PrimaryIpAddress *string `type:"string"`

	PrivateIpAddresses []*string `type:"list"`

	// SecurityGroupIds is a required field
	SecurityGroupIds []*string `type:"list" required:"true"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s NetworkInterfaceForRunInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkInterfaceForRunInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NetworkInterfaceForRunInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "NetworkInterfaceForRunInstancesInput"}
	if s.SecurityGroupIds == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupIds"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIpv6AddressCount sets the Ipv6AddressCount field's value.
func (s *NetworkInterfaceForRunInstancesInput) SetIpv6AddressCount(v int32) *NetworkInterfaceForRunInstancesInput {
	s.Ipv6AddressCount = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *NetworkInterfaceForRunInstancesInput) SetPrimaryIpAddress(v string) *NetworkInterfaceForRunInstancesInput {
	s.PrimaryIpAddress = &v
	return s
}

// SetPrivateIpAddresses sets the PrivateIpAddresses field's value.
func (s *NetworkInterfaceForRunInstancesInput) SetPrivateIpAddresses(v []*string) *NetworkInterfaceForRunInstancesInput {
	s.PrivateIpAddresses = v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *NetworkInterfaceForRunInstancesInput) SetSecurityGroupIds(v []*string) *NetworkInterfaceForRunInstancesInput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NetworkInterfaceForRunInstancesInput) SetSubnetId(v string) *NetworkInterfaceForRunInstancesInput {
	s.SubnetId = &v
	return s
}

type PlacementForRunInstancesInput struct {
	_ struct{} `type:"structure"`

	Affinity *string `type:"string"`

	DedicatedHostClusterId *string `type:"string"`

	DedicatedHostId *string `type:"string"`

	Tenancy *string `type:"string"`
}

// String returns the string representation
func (s PlacementForRunInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s PlacementForRunInstancesInput) GoString() string {
	return s.String()
}

// SetAffinity sets the Affinity field's value.
func (s *PlacementForRunInstancesInput) SetAffinity(v string) *PlacementForRunInstancesInput {
	s.Affinity = &v
	return s
}

// SetDedicatedHostClusterId sets the DedicatedHostClusterId field's value.
func (s *PlacementForRunInstancesInput) SetDedicatedHostClusterId(v string) *PlacementForRunInstancesInput {
	s.DedicatedHostClusterId = &v
	return s
}

// SetDedicatedHostId sets the DedicatedHostId field's value.
func (s *PlacementForRunInstancesInput) SetDedicatedHostId(v string) *PlacementForRunInstancesInput {
	s.DedicatedHostId = &v
	return s
}

// SetTenancy sets the Tenancy field's value.
func (s *PlacementForRunInstancesInput) SetTenancy(v string) *PlacementForRunInstancesInput {
	s.Tenancy = &v
	return s
}

type RunInstancesInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	AutoRenewPeriod *int32 `type:"int32"`

	ClientToken *string `type:"string"`

	Count *int32 `type:"int32"`

	CreditSpecification *string `type:"string"`

	DeploymentSetGroupNumber *int32 `type:"int32"`

	DeploymentSetId *string `type:"string"`

	Description *string `type:"string"`

	DryRun *bool `type:"boolean"`

	EipAddress *EipAddressForRunInstancesInput `type:"structure"`

	HostName *string `type:"string"`

	Hostname *string `type:"string"`

	HpcClusterId *string `type:"string"`

	ImageId *string `type:"string"`

	ImageReleaseVersion *string `type:"string"`

	InstallRunCommandAgent *bool `type:"boolean"`

	InstanceChargeType *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceType *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	KeepImageCredential *bool `type:"boolean"`

	KeyPairName *string `type:"string"`

	MinCount *int32 `type:"int32"`

	NetworkInterface []*ConvertNetworkInterfaceForRunInstancesInput `type:"list"`

	NetworkInterfaces []*NetworkInterfaceForRunInstancesInput `type:"list"`

	Password *string `type:"string"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string"`

	Placement *PlacementForRunInstancesInput `type:"structure"`

	ProjectName *string `type:"string"`

	SecurityEnhancementStrategy *string `type:"string"`

	SpotPriceLimit *float64 `type:"double"`

	SpotStrategy *string `type:"string"`

	SuffixIndex *int32 `type:"int32"`

	Tags []*TagForRunInstancesInput `type:"list"`

	UniqueSuffix *bool `type:"boolean"`

	UserData *string `type:"string"`

	Volumes []*VolumeForRunInstancesInput `type:"list"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s RunInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s RunInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RunInstancesInput"}
	if s.NetworkInterfaces != nil {
		for i, v := range s.NetworkInterfaces {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "NetworkInterfaces", i), err.(request.ErrInvalidParams))
			}
		}
	}
	if s.Volumes != nil {
		for i, v := range s.Volumes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Volumes", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *RunInstancesInput) SetAutoRenew(v bool) *RunInstancesInput {
	s.AutoRenew = &v
	return s
}

// SetAutoRenewPeriod sets the AutoRenewPeriod field's value.
func (s *RunInstancesInput) SetAutoRenewPeriod(v int32) *RunInstancesInput {
	s.AutoRenewPeriod = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *RunInstancesInput) SetClientToken(v string) *RunInstancesInput {
	s.ClientToken = &v
	return s
}

// SetCount sets the Count field's value.
func (s *RunInstancesInput) SetCount(v int32) *RunInstancesInput {
	s.Count = &v
	return s
}

// SetCreditSpecification sets the CreditSpecification field's value.
func (s *RunInstancesInput) SetCreditSpecification(v string) *RunInstancesInput {
	s.CreditSpecification = &v
	return s
}

// SetDeploymentSetGroupNumber sets the DeploymentSetGroupNumber field's value.
func (s *RunInstancesInput) SetDeploymentSetGroupNumber(v int32) *RunInstancesInput {
	s.DeploymentSetGroupNumber = &v
	return s
}

// SetDeploymentSetId sets the DeploymentSetId field's value.
func (s *RunInstancesInput) SetDeploymentSetId(v string) *RunInstancesInput {
	s.DeploymentSetId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RunInstancesInput) SetDescription(v string) *RunInstancesInput {
	s.Description = &v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *RunInstancesInput) SetDryRun(v bool) *RunInstancesInput {
	s.DryRun = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *RunInstancesInput) SetEipAddress(v *EipAddressForRunInstancesInput) *RunInstancesInput {
	s.EipAddress = v
	return s
}

// SetHostName sets the HostName field's value.
func (s *RunInstancesInput) SetHostName(v string) *RunInstancesInput {
	s.HostName = &v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *RunInstancesInput) SetHostname(v string) *RunInstancesInput {
	s.Hostname = &v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *RunInstancesInput) SetHpcClusterId(v string) *RunInstancesInput {
	s.HpcClusterId = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *RunInstancesInput) SetImageId(v string) *RunInstancesInput {
	s.ImageId = &v
	return s
}

// SetImageReleaseVersion sets the ImageReleaseVersion field's value.
func (s *RunInstancesInput) SetImageReleaseVersion(v string) *RunInstancesInput {
	s.ImageReleaseVersion = &v
	return s
}

// SetInstallRunCommandAgent sets the InstallRunCommandAgent field's value.
func (s *RunInstancesInput) SetInstallRunCommandAgent(v bool) *RunInstancesInput {
	s.InstallRunCommandAgent = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *RunInstancesInput) SetInstanceChargeType(v string) *RunInstancesInput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *RunInstancesInput) SetInstanceName(v string) *RunInstancesInput {
	s.InstanceName = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *RunInstancesInput) SetInstanceType(v string) *RunInstancesInput {
	s.InstanceType = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *RunInstancesInput) SetInstanceTypeId(v string) *RunInstancesInput {
	s.InstanceTypeId = &v
	return s
}

// SetKeepImageCredential sets the KeepImageCredential field's value.
func (s *RunInstancesInput) SetKeepImageCredential(v bool) *RunInstancesInput {
	s.KeepImageCredential = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *RunInstancesInput) SetKeyPairName(v string) *RunInstancesInput {
	s.KeyPairName = &v
	return s
}

// SetMinCount sets the MinCount field's value.
func (s *RunInstancesInput) SetMinCount(v int32) *RunInstancesInput {
	s.MinCount = &v
	return s
}

// SetNetworkInterface sets the NetworkInterface field's value.
func (s *RunInstancesInput) SetNetworkInterface(v []*ConvertNetworkInterfaceForRunInstancesInput) *RunInstancesInput {
	s.NetworkInterface = v
	return s
}

// SetNetworkInterfaces sets the NetworkInterfaces field's value.
func (s *RunInstancesInput) SetNetworkInterfaces(v []*NetworkInterfaceForRunInstancesInput) *RunInstancesInput {
	s.NetworkInterfaces = v
	return s
}

// SetPassword sets the Password field's value.
func (s *RunInstancesInput) SetPassword(v string) *RunInstancesInput {
	s.Password = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *RunInstancesInput) SetPeriod(v int32) *RunInstancesInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *RunInstancesInput) SetPeriodUnit(v string) *RunInstancesInput {
	s.PeriodUnit = &v
	return s
}

// SetPlacement sets the Placement field's value.
func (s *RunInstancesInput) SetPlacement(v *PlacementForRunInstancesInput) *RunInstancesInput {
	s.Placement = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *RunInstancesInput) SetProjectName(v string) *RunInstancesInput {
	s.ProjectName = &v
	return s
}

// SetSecurityEnhancementStrategy sets the SecurityEnhancementStrategy field's value.
func (s *RunInstancesInput) SetSecurityEnhancementStrategy(v string) *RunInstancesInput {
	s.SecurityEnhancementStrategy = &v
	return s
}

// SetSpotPriceLimit sets the SpotPriceLimit field's value.
func (s *RunInstancesInput) SetSpotPriceLimit(v float64) *RunInstancesInput {
	s.SpotPriceLimit = &v
	return s
}

// SetSpotStrategy sets the SpotStrategy field's value.
func (s *RunInstancesInput) SetSpotStrategy(v string) *RunInstancesInput {
	s.SpotStrategy = &v
	return s
}

// SetSuffixIndex sets the SuffixIndex field's value.
func (s *RunInstancesInput) SetSuffixIndex(v int32) *RunInstancesInput {
	s.SuffixIndex = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *RunInstancesInput) SetTags(v []*TagForRunInstancesInput) *RunInstancesInput {
	s.Tags = v
	return s
}

// SetUniqueSuffix sets the UniqueSuffix field's value.
func (s *RunInstancesInput) SetUniqueSuffix(v bool) *RunInstancesInput {
	s.UniqueSuffix = &v
	return s
}

// SetUserData sets the UserData field's value.
func (s *RunInstancesInput) SetUserData(v string) *RunInstancesInput {
	s.UserData = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *RunInstancesInput) SetVolumes(v []*VolumeForRunInstancesInput) *RunInstancesInput {
	s.Volumes = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *RunInstancesInput) SetZoneId(v string) *RunInstancesInput {
	s.ZoneId = &v
	return s
}

type RunInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceIds []*string `type:"list"`
}

// String returns the string representation
func (s RunInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s RunInstancesOutput) GoString() string {
	return s.String()
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *RunInstancesOutput) SetInstanceIds(v []*string) *RunInstancesOutput {
	s.InstanceIds = v
	return s
}

type TagForRunInstancesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForRunInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForRunInstancesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForRunInstancesInput) SetKey(v string) *TagForRunInstancesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForRunInstancesInput) SetValue(v string) *TagForRunInstancesInput {
	s.Value = &v
	return s
}

type VolumeForRunInstancesInput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *string `type:"string"`

	ExtraPerformanceIOPS *int32 `type:"int32"`

	ExtraPerformanceThroughputMB *int32 `type:"int32"`

	ExtraPerformanceTypeId *string `type:"string"`

	// Size is a required field
	Size *int32 `type:"int32" required:"true"`

	SnapshotId *string `type:"string"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s VolumeForRunInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForRunInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VolumeForRunInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "VolumeForRunInstancesInput"}
	if s.Size == nil {
		invalidParams.Add(request.NewErrParamRequired("Size"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *VolumeForRunInstancesInput) SetDeleteWithInstance(v string) *VolumeForRunInstancesInput {
	s.DeleteWithInstance = &v
	return s
}

// SetExtraPerformanceIOPS sets the ExtraPerformanceIOPS field's value.
func (s *VolumeForRunInstancesInput) SetExtraPerformanceIOPS(v int32) *VolumeForRunInstancesInput {
	s.ExtraPerformanceIOPS = &v
	return s
}

// SetExtraPerformanceThroughputMB sets the ExtraPerformanceThroughputMB field's value.
func (s *VolumeForRunInstancesInput) SetExtraPerformanceThroughputMB(v int32) *VolumeForRunInstancesInput {
	s.ExtraPerformanceThroughputMB = &v
	return s
}

// SetExtraPerformanceTypeId sets the ExtraPerformanceTypeId field's value.
func (s *VolumeForRunInstancesInput) SetExtraPerformanceTypeId(v string) *VolumeForRunInstancesInput {
	s.ExtraPerformanceTypeId = &v
	return s
}

// SetSize sets the Size field's value.
func (s *VolumeForRunInstancesInput) SetSize(v int32) *VolumeForRunInstancesInput {
	s.Size = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *VolumeForRunInstancesInput) SetSnapshotId(v string) *VolumeForRunInstancesInput {
	s.SnapshotId = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *VolumeForRunInstancesInput) SetVolumeType(v string) *VolumeForRunInstancesInput {
	s.VolumeType = &v
	return s
}
