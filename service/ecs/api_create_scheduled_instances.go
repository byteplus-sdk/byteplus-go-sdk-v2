// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreateScheduledInstancesCommon = "CreateScheduledInstances"

// CreateScheduledInstancesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateScheduledInstancesCommon operation. The "output" return
// value will be populated with the CreateScheduledInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateScheduledInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateScheduledInstancesCommon Send returns without error.
//
// See CreateScheduledInstancesCommon for more information on using the CreateScheduledInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateScheduledInstancesCommonRequest method.
//    req, resp := client.CreateScheduledInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateScheduledInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateScheduledInstancesCommon,
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

// CreateScheduledInstancesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation CreateScheduledInstancesCommon for usage and error information.
func (c *ECS) CreateScheduledInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateScheduledInstancesCommonRequest(input)
	return out, req.Send()
}

// CreateScheduledInstancesCommonWithContext is the same as CreateScheduledInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateScheduledInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateScheduledInstancesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateScheduledInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateScheduledInstances = "CreateScheduledInstances"

// CreateScheduledInstancesRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateScheduledInstances operation. The "output" return
// value will be populated with the CreateScheduledInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateScheduledInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateScheduledInstancesCommon Send returns without error.
//
// See CreateScheduledInstances for more information on using the CreateScheduledInstances
// API call, and error handling.
//
//    // Example sending a request using the CreateScheduledInstancesRequest method.
//    req, resp := client.CreateScheduledInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateScheduledInstancesRequest(input *CreateScheduledInstancesInput) (req *request.Request, output *CreateScheduledInstancesOutput) {
	op := &request.Operation{
		Name:       opCreateScheduledInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateScheduledInstancesInput{}
	}

	output = &CreateScheduledInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateScheduledInstances API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation CreateScheduledInstances for usage and error information.
func (c *ECS) CreateScheduledInstances(input *CreateScheduledInstancesInput) (*CreateScheduledInstancesOutput, error) {
	req, out := c.CreateScheduledInstancesRequest(input)
	return out, req.Send()
}

// CreateScheduledInstancesWithContext is the same as CreateScheduledInstances with the addition of
// the ability to pass a context and additional request options.
//
// See CreateScheduledInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateScheduledInstancesWithContext(ctx byteplus.Context, input *CreateScheduledInstancesInput, opts ...request.Option) (*CreateScheduledInstancesOutput, error) {
	req, out := c.CreateScheduledInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	AutoReleaseAt *string `type:"string"`

	ClientToken *string `type:"string"`

	Count *int32 `type:"int32"`

	DeliveryType *string `type:"string"`

	Description *string `type:"string"`

	DryRun *bool `type:"boolean"`

	EipAddress *EipAddressForCreateScheduledInstancesInput `type:"structure"`

	ElasticScheduledInstanceType *string `type:"string"`

	EndDeliveryAt *string `type:"string"`

	Hostname *string `type:"string"`

	HpcClusterId *string `type:"string"`

	ImageId *string `type:"string"`

	InstallRunCommandAgent *bool `type:"boolean"`

	InstanceName *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	KeepImageCredential *bool `type:"boolean"`

	KeyPairName *string `type:"string"`

	NetworkInterfaces []*NetworkInterfaceForCreateScheduledInstancesInput `type:"list"`

	Password *string `type:"string"`

	ProjectName *string `type:"string"`

	ScheduledInstanceDescription *string `type:"string"`

	ScheduledInstanceName *string `type:"string"`

	SecurityEnhancementStrategy *string `type:"string"`

	StartDeliveryAt *string `type:"string"`

	SuffixIndex *int32 `type:"int32"`

	Tags []*TagForCreateScheduledInstancesInput `type:"list"`

	UniqueSuffix *bool `type:"boolean"`

	UserData *string `type:"string"`

	Volumes []*VolumeForCreateScheduledInstancesInput `type:"list"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s CreateScheduledInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateScheduledInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateScheduledInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateScheduledInstancesInput"}
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

// SetAutoReleaseAt sets the AutoReleaseAt field's value.
func (s *CreateScheduledInstancesInput) SetAutoReleaseAt(v string) *CreateScheduledInstancesInput {
	s.AutoReleaseAt = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateScheduledInstancesInput) SetClientToken(v string) *CreateScheduledInstancesInput {
	s.ClientToken = &v
	return s
}

// SetCount sets the Count field's value.
func (s *CreateScheduledInstancesInput) SetCount(v int32) *CreateScheduledInstancesInput {
	s.Count = &v
	return s
}

// SetDeliveryType sets the DeliveryType field's value.
func (s *CreateScheduledInstancesInput) SetDeliveryType(v string) *CreateScheduledInstancesInput {
	s.DeliveryType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateScheduledInstancesInput) SetDescription(v string) *CreateScheduledInstancesInput {
	s.Description = &v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *CreateScheduledInstancesInput) SetDryRun(v bool) *CreateScheduledInstancesInput {
	s.DryRun = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *CreateScheduledInstancesInput) SetEipAddress(v *EipAddressForCreateScheduledInstancesInput) *CreateScheduledInstancesInput {
	s.EipAddress = v
	return s
}

// SetElasticScheduledInstanceType sets the ElasticScheduledInstanceType field's value.
func (s *CreateScheduledInstancesInput) SetElasticScheduledInstanceType(v string) *CreateScheduledInstancesInput {
	s.ElasticScheduledInstanceType = &v
	return s
}

// SetEndDeliveryAt sets the EndDeliveryAt field's value.
func (s *CreateScheduledInstancesInput) SetEndDeliveryAt(v string) *CreateScheduledInstancesInput {
	s.EndDeliveryAt = &v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *CreateScheduledInstancesInput) SetHostname(v string) *CreateScheduledInstancesInput {
	s.Hostname = &v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *CreateScheduledInstancesInput) SetHpcClusterId(v string) *CreateScheduledInstancesInput {
	s.HpcClusterId = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *CreateScheduledInstancesInput) SetImageId(v string) *CreateScheduledInstancesInput {
	s.ImageId = &v
	return s
}

// SetInstallRunCommandAgent sets the InstallRunCommandAgent field's value.
func (s *CreateScheduledInstancesInput) SetInstallRunCommandAgent(v bool) *CreateScheduledInstancesInput {
	s.InstallRunCommandAgent = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *CreateScheduledInstancesInput) SetInstanceName(v string) *CreateScheduledInstancesInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *CreateScheduledInstancesInput) SetInstanceTypeId(v string) *CreateScheduledInstancesInput {
	s.InstanceTypeId = &v
	return s
}

// SetKeepImageCredential sets the KeepImageCredential field's value.
func (s *CreateScheduledInstancesInput) SetKeepImageCredential(v bool) *CreateScheduledInstancesInput {
	s.KeepImageCredential = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *CreateScheduledInstancesInput) SetKeyPairName(v string) *CreateScheduledInstancesInput {
	s.KeyPairName = &v
	return s
}

// SetNetworkInterfaces sets the NetworkInterfaces field's value.
func (s *CreateScheduledInstancesInput) SetNetworkInterfaces(v []*NetworkInterfaceForCreateScheduledInstancesInput) *CreateScheduledInstancesInput {
	s.NetworkInterfaces = v
	return s
}

// SetPassword sets the Password field's value.
func (s *CreateScheduledInstancesInput) SetPassword(v string) *CreateScheduledInstancesInput {
	s.Password = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateScheduledInstancesInput) SetProjectName(v string) *CreateScheduledInstancesInput {
	s.ProjectName = &v
	return s
}

// SetScheduledInstanceDescription sets the ScheduledInstanceDescription field's value.
func (s *CreateScheduledInstancesInput) SetScheduledInstanceDescription(v string) *CreateScheduledInstancesInput {
	s.ScheduledInstanceDescription = &v
	return s
}

// SetScheduledInstanceName sets the ScheduledInstanceName field's value.
func (s *CreateScheduledInstancesInput) SetScheduledInstanceName(v string) *CreateScheduledInstancesInput {
	s.ScheduledInstanceName = &v
	return s
}

// SetSecurityEnhancementStrategy sets the SecurityEnhancementStrategy field's value.
func (s *CreateScheduledInstancesInput) SetSecurityEnhancementStrategy(v string) *CreateScheduledInstancesInput {
	s.SecurityEnhancementStrategy = &v
	return s
}

// SetStartDeliveryAt sets the StartDeliveryAt field's value.
func (s *CreateScheduledInstancesInput) SetStartDeliveryAt(v string) *CreateScheduledInstancesInput {
	s.StartDeliveryAt = &v
	return s
}

// SetSuffixIndex sets the SuffixIndex field's value.
func (s *CreateScheduledInstancesInput) SetSuffixIndex(v int32) *CreateScheduledInstancesInput {
	s.SuffixIndex = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateScheduledInstancesInput) SetTags(v []*TagForCreateScheduledInstancesInput) *CreateScheduledInstancesInput {
	s.Tags = v
	return s
}

// SetUniqueSuffix sets the UniqueSuffix field's value.
func (s *CreateScheduledInstancesInput) SetUniqueSuffix(v bool) *CreateScheduledInstancesInput {
	s.UniqueSuffix = &v
	return s
}

// SetUserData sets the UserData field's value.
func (s *CreateScheduledInstancesInput) SetUserData(v string) *CreateScheduledInstancesInput {
	s.UserData = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *CreateScheduledInstancesInput) SetVolumes(v []*VolumeForCreateScheduledInstancesInput) *CreateScheduledInstancesInput {
	s.Volumes = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateScheduledInstancesInput) SetZoneId(v string) *CreateScheduledInstancesInput {
	s.ZoneId = &v
	return s
}

type CreateScheduledInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScheduledInstanceId *string `type:"string"`
}

// String returns the string representation
func (s CreateScheduledInstancesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateScheduledInstancesOutput) GoString() string {
	return s.String()
}

// SetScheduledInstanceId sets the ScheduledInstanceId field's value.
func (s *CreateScheduledInstancesOutput) SetScheduledInstanceId(v string) *CreateScheduledInstancesOutput {
	s.ScheduledInstanceId = &v
	return s
}

type EipAddressForCreateScheduledInstancesInput struct {
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
func (s EipAddressForCreateScheduledInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s EipAddressForCreateScheduledInstancesInput) GoString() string {
	return s.String()
}

// SetBandwidthMbps sets the BandwidthMbps field's value.
func (s *EipAddressForCreateScheduledInstancesInput) SetBandwidthMbps(v int32) *EipAddressForCreateScheduledInstancesInput {
	s.BandwidthMbps = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *EipAddressForCreateScheduledInstancesInput) SetBandwidthPackageId(v string) *EipAddressForCreateScheduledInstancesInput {
	s.BandwidthPackageId = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *EipAddressForCreateScheduledInstancesInput) SetChargeType(v string) *EipAddressForCreateScheduledInstancesInput {
	s.ChargeType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipAddressForCreateScheduledInstancesInput) SetISP(v string) *EipAddressForCreateScheduledInstancesInput {
	s.ISP = &v
	return s
}

// SetReleaseWithInstance sets the ReleaseWithInstance field's value.
func (s *EipAddressForCreateScheduledInstancesInput) SetReleaseWithInstance(v bool) *EipAddressForCreateScheduledInstancesInput {
	s.ReleaseWithInstance = &v
	return s
}

// SetSecurityProtectionInstanceId sets the SecurityProtectionInstanceId field's value.
func (s *EipAddressForCreateScheduledInstancesInput) SetSecurityProtectionInstanceId(v int32) *EipAddressForCreateScheduledInstancesInput {
	s.SecurityProtectionInstanceId = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *EipAddressForCreateScheduledInstancesInput) SetSecurityProtectionTypes(v []*string) *EipAddressForCreateScheduledInstancesInput {
	s.SecurityProtectionTypes = v
	return s
}

type NetworkInterfaceForCreateScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	PrimaryIpAddress *string `type:"string"`

	PrivateIpAddresses []*string `type:"list"`

	// SecurityGroupIds is a required field
	SecurityGroupIds []*string `type:"list" required:"true"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s NetworkInterfaceForCreateScheduledInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkInterfaceForCreateScheduledInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NetworkInterfaceForCreateScheduledInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "NetworkInterfaceForCreateScheduledInstancesInput"}
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

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *NetworkInterfaceForCreateScheduledInstancesInput) SetPrimaryIpAddress(v string) *NetworkInterfaceForCreateScheduledInstancesInput {
	s.PrimaryIpAddress = &v
	return s
}

// SetPrivateIpAddresses sets the PrivateIpAddresses field's value.
func (s *NetworkInterfaceForCreateScheduledInstancesInput) SetPrivateIpAddresses(v []*string) *NetworkInterfaceForCreateScheduledInstancesInput {
	s.PrivateIpAddresses = v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *NetworkInterfaceForCreateScheduledInstancesInput) SetSecurityGroupIds(v []*string) *NetworkInterfaceForCreateScheduledInstancesInput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NetworkInterfaceForCreateScheduledInstancesInput) SetSubnetId(v string) *NetworkInterfaceForCreateScheduledInstancesInput {
	s.SubnetId = &v
	return s
}

type TagForCreateScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateScheduledInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateScheduledInstancesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateScheduledInstancesInput) SetKey(v string) *TagForCreateScheduledInstancesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateScheduledInstancesInput) SetValue(v string) *TagForCreateScheduledInstancesInput {
	s.Value = &v
	return s
}

type VolumeForCreateScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *bool `type:"boolean"`

	ExtraPerformanceIOPS *int32 `type:"int32"`

	ExtraPerformanceThroughputMB *int32 `type:"int32"`

	ExtraPerformanceTypeId *string `type:"string"`

	// Size is a required field
	Size *int32 `type:"int32" required:"true"`

	SnapshotId *string `type:"string"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s VolumeForCreateScheduledInstancesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForCreateScheduledInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VolumeForCreateScheduledInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "VolumeForCreateScheduledInstancesInput"}
	if s.Size == nil {
		invalidParams.Add(request.NewErrParamRequired("Size"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *VolumeForCreateScheduledInstancesInput) SetDeleteWithInstance(v bool) *VolumeForCreateScheduledInstancesInput {
	s.DeleteWithInstance = &v
	return s
}

// SetExtraPerformanceIOPS sets the ExtraPerformanceIOPS field's value.
func (s *VolumeForCreateScheduledInstancesInput) SetExtraPerformanceIOPS(v int32) *VolumeForCreateScheduledInstancesInput {
	s.ExtraPerformanceIOPS = &v
	return s
}

// SetExtraPerformanceThroughputMB sets the ExtraPerformanceThroughputMB field's value.
func (s *VolumeForCreateScheduledInstancesInput) SetExtraPerformanceThroughputMB(v int32) *VolumeForCreateScheduledInstancesInput {
	s.ExtraPerformanceThroughputMB = &v
	return s
}

// SetExtraPerformanceTypeId sets the ExtraPerformanceTypeId field's value.
func (s *VolumeForCreateScheduledInstancesInput) SetExtraPerformanceTypeId(v string) *VolumeForCreateScheduledInstancesInput {
	s.ExtraPerformanceTypeId = &v
	return s
}

// SetSize sets the Size field's value.
func (s *VolumeForCreateScheduledInstancesInput) SetSize(v int32) *VolumeForCreateScheduledInstancesInput {
	s.Size = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *VolumeForCreateScheduledInstancesInput) SetSnapshotId(v string) *VolumeForCreateScheduledInstancesInput {
	s.SnapshotId = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *VolumeForCreateScheduledInstancesInput) SetVolumeType(v string) *VolumeForCreateScheduledInstancesInput {
	s.VolumeType = &v
	return s
}