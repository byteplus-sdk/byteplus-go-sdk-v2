// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeInstanceTypesCommon = "DescribeInstanceTypes"

// DescribeInstanceTypesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInstanceTypesCommon operation. The "output" return
// value will be populated with the DescribeInstanceTypesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceTypesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceTypesCommon Send returns without error.
//
// See DescribeInstanceTypesCommon for more information on using the DescribeInstanceTypesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceTypesCommonRequest method.
//    req, resp := client.DescribeInstanceTypesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstanceTypesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstanceTypesCommon,
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

// DescribeInstanceTypesCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInstanceTypesCommon for usage and error information.
func (c *ECS) DescribeInstanceTypesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceTypesCommonRequest(input)
	return out, req.Send()
}

// DescribeInstanceTypesCommonWithContext is the same as DescribeInstanceTypesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceTypesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstanceTypesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceTypesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstanceTypes = "DescribeInstanceTypes"

// DescribeInstanceTypesRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeInstanceTypes operation. The "output" return
// value will be populated with the DescribeInstanceTypesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceTypesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceTypesCommon Send returns without error.
//
// See DescribeInstanceTypes for more information on using the DescribeInstanceTypes
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceTypesRequest method.
//    req, resp := client.DescribeInstanceTypesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstanceTypesRequest(input *DescribeInstanceTypesInput) (req *request.Request, output *DescribeInstanceTypesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstanceTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceTypesInput{}
	}

	output = &DescribeInstanceTypesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInstanceTypes API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeInstanceTypes for usage and error information.
func (c *ECS) DescribeInstanceTypes(input *DescribeInstanceTypesInput) (*DescribeInstanceTypesOutput, error) {
	req, out := c.DescribeInstanceTypesRequest(input)
	return out, req.Send()
}

// DescribeInstanceTypesWithContext is the same as DescribeInstanceTypes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceTypes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstanceTypesWithContext(ctx byteplus.Context, input *DescribeInstanceTypesInput, opts ...request.Option) (*DescribeInstanceTypesOutput, error) {
	req, out := c.DescribeInstanceTypesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInstanceTypesInput struct {
	_ struct{} `type:"structure"`

	ImageId *string `type:"string"`

	InstanceTypeIds []*string `type:"list"`

	InstanceTypes []*string `type:"list"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceTypesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceTypesInput) GoString() string {
	return s.String()
}

// SetImageId sets the ImageId field's value.
func (s *DescribeInstanceTypesInput) SetImageId(v string) *DescribeInstanceTypesInput {
	s.ImageId = &v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *DescribeInstanceTypesInput) SetInstanceTypeIds(v []*string) *DescribeInstanceTypesInput {
	s.InstanceTypeIds = v
	return s
}

// SetInstanceTypes sets the InstanceTypes field's value.
func (s *DescribeInstanceTypesInput) SetInstanceTypes(v []*string) *DescribeInstanceTypesInput {
	s.InstanceTypes = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeInstanceTypesInput) SetMaxResults(v int32) *DescribeInstanceTypesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstanceTypesInput) SetNextToken(v string) *DescribeInstanceTypesInput {
	s.NextToken = &v
	return s
}

type DescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceTypes []*InstanceTypeForDescribeInstanceTypesOutput `type:"list"`

	NextToken *string `type:"string"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetInstanceTypes sets the InstanceTypes field's value.
func (s *DescribeInstanceTypesOutput) SetInstanceTypes(v []*InstanceTypeForDescribeInstanceTypesOutput) *DescribeInstanceTypesOutput {
	s.InstanceTypes = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstanceTypesOutput) SetNextToken(v string) *DescribeInstanceTypesOutput {
	s.NextToken = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeInstanceTypesOutput) SetTotalCount(v int32) *DescribeInstanceTypesOutput {
	s.TotalCount = &v
	return s
}

type GpuDeviceForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	Count *int32 `type:"int32"`

	Memory *MemoryForDescribeInstanceTypesOutput `type:"structure"`

	ProductName *string `type:"string"`
}

// String returns the string representation
func (s GpuDeviceForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s GpuDeviceForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *GpuDeviceForDescribeInstanceTypesOutput) SetCount(v int32) *GpuDeviceForDescribeInstanceTypesOutput {
	s.Count = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *GpuDeviceForDescribeInstanceTypesOutput) SetMemory(v *MemoryForDescribeInstanceTypesOutput) *GpuDeviceForDescribeInstanceTypesOutput {
	s.Memory = v
	return s
}

// SetProductName sets the ProductName field's value.
func (s *GpuDeviceForDescribeInstanceTypesOutput) SetProductName(v string) *GpuDeviceForDescribeInstanceTypesOutput {
	s.ProductName = &v
	return s
}

type GpuForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	GpuDevices []*GpuDeviceForDescribeInstanceTypesOutput `type:"list"`
}

// String returns the string representation
func (s GpuForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s GpuForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetGpuDevices sets the GpuDevices field's value.
func (s *GpuForDescribeInstanceTypesOutput) SetGpuDevices(v []*GpuDeviceForDescribeInstanceTypesOutput) *GpuForDescribeInstanceTypesOutput {
	s.GpuDevices = v
	return s
}

type InstanceTypeForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	AffinityGroupSizes []*int32 `type:"list"`

	BaselineCredit *int64 `type:"int64"`

	Gpu *GpuForDescribeInstanceTypesOutput `type:"structure"`

	InitialCredit *int64 `type:"int64"`

	InstanceTypeFamily *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	IsSupportAffinityGroup *bool `type:"boolean"`

	LocalVolumes []*LocalVolumeForDescribeInstanceTypesOutput `type:"list"`

	Memory *MemoryForDescribeInstanceTypesOutput `type:"structure"`

	Network *NetworkForDescribeInstanceTypesOutput `type:"structure"`

	Processor *ProcessorForDescribeInstanceTypesOutput `type:"structure"`

	Rdma *RdmaForDescribeInstanceTypesOutput `type:"structure"`

	Volume *VolumeForDescribeInstanceTypesOutput `type:"structure"`
}

// String returns the string representation
func (s InstanceTypeForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceTypeForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetAffinityGroupSizes sets the AffinityGroupSizes field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetAffinityGroupSizes(v []*int32) *InstanceTypeForDescribeInstanceTypesOutput {
	s.AffinityGroupSizes = v
	return s
}

// SetBaselineCredit sets the BaselineCredit field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetBaselineCredit(v int64) *InstanceTypeForDescribeInstanceTypesOutput {
	s.BaselineCredit = &v
	return s
}

// SetGpu sets the Gpu field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetGpu(v *GpuForDescribeInstanceTypesOutput) *InstanceTypeForDescribeInstanceTypesOutput {
	s.Gpu = v
	return s
}

// SetInitialCredit sets the InitialCredit field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetInitialCredit(v int64) *InstanceTypeForDescribeInstanceTypesOutput {
	s.InitialCredit = &v
	return s
}

// SetInstanceTypeFamily sets the InstanceTypeFamily field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetInstanceTypeFamily(v string) *InstanceTypeForDescribeInstanceTypesOutput {
	s.InstanceTypeFamily = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetInstanceTypeId(v string) *InstanceTypeForDescribeInstanceTypesOutput {
	s.InstanceTypeId = &v
	return s
}

// SetIsSupportAffinityGroup sets the IsSupportAffinityGroup field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetIsSupportAffinityGroup(v bool) *InstanceTypeForDescribeInstanceTypesOutput {
	s.IsSupportAffinityGroup = &v
	return s
}

// SetLocalVolumes sets the LocalVolumes field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetLocalVolumes(v []*LocalVolumeForDescribeInstanceTypesOutput) *InstanceTypeForDescribeInstanceTypesOutput {
	s.LocalVolumes = v
	return s
}

// SetMemory sets the Memory field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetMemory(v *MemoryForDescribeInstanceTypesOutput) *InstanceTypeForDescribeInstanceTypesOutput {
	s.Memory = v
	return s
}

// SetNetwork sets the Network field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetNetwork(v *NetworkForDescribeInstanceTypesOutput) *InstanceTypeForDescribeInstanceTypesOutput {
	s.Network = v
	return s
}

// SetProcessor sets the Processor field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetProcessor(v *ProcessorForDescribeInstanceTypesOutput) *InstanceTypeForDescribeInstanceTypesOutput {
	s.Processor = v
	return s
}

// SetRdma sets the Rdma field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetRdma(v *RdmaForDescribeInstanceTypesOutput) *InstanceTypeForDescribeInstanceTypesOutput {
	s.Rdma = v
	return s
}

// SetVolume sets the Volume field's value.
func (s *InstanceTypeForDescribeInstanceTypesOutput) SetVolume(v *VolumeForDescribeInstanceTypesOutput) *InstanceTypeForDescribeInstanceTypesOutput {
	s.Volume = v
	return s
}

type LocalVolumeForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	Count *int32 `type:"int32"`

	Size *int32 `type:"int32"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s LocalVolumeForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s LocalVolumeForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *LocalVolumeForDescribeInstanceTypesOutput) SetCount(v int32) *LocalVolumeForDescribeInstanceTypesOutput {
	s.Count = &v
	return s
}

// SetSize sets the Size field's value.
func (s *LocalVolumeForDescribeInstanceTypesOutput) SetSize(v int32) *LocalVolumeForDescribeInstanceTypesOutput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *LocalVolumeForDescribeInstanceTypesOutput) SetVolumeType(v string) *LocalVolumeForDescribeInstanceTypesOutput {
	s.VolumeType = &v
	return s
}

type MemoryForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	EncryptedSize *int32 `type:"int32"`

	Size *int32 `type:"int32"`
}

// String returns the string representation
func (s MemoryForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s MemoryForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetEncryptedSize sets the EncryptedSize field's value.
func (s *MemoryForDescribeInstanceTypesOutput) SetEncryptedSize(v int32) *MemoryForDescribeInstanceTypesOutput {
	s.EncryptedSize = &v
	return s
}

// SetSize sets the Size field's value.
func (s *MemoryForDescribeInstanceTypesOutput) SetSize(v int32) *MemoryForDescribeInstanceTypesOutput {
	s.Size = &v
	return s
}

type NetworkForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	BaselineBandwidthMbps *int32 `type:"int32"`

	MaximumBandwidthMbps *int32 `type:"int32"`

	MaximumNetworkInterfaces *int32 `type:"int32"`

	MaximumPrivateIpv4AddressesPerNetworkInterface *int32 `type:"int32"`

	MaximumQueuesPerNetworkInterface *int32 `type:"int32"`

	MaximumThroughputKpps *int32 `type:"int32"`
}

// String returns the string representation
func (s NetworkForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetBaselineBandwidthMbps sets the BaselineBandwidthMbps field's value.
func (s *NetworkForDescribeInstanceTypesOutput) SetBaselineBandwidthMbps(v int32) *NetworkForDescribeInstanceTypesOutput {
	s.BaselineBandwidthMbps = &v
	return s
}

// SetMaximumBandwidthMbps sets the MaximumBandwidthMbps field's value.
func (s *NetworkForDescribeInstanceTypesOutput) SetMaximumBandwidthMbps(v int32) *NetworkForDescribeInstanceTypesOutput {
	s.MaximumBandwidthMbps = &v
	return s
}

// SetMaximumNetworkInterfaces sets the MaximumNetworkInterfaces field's value.
func (s *NetworkForDescribeInstanceTypesOutput) SetMaximumNetworkInterfaces(v int32) *NetworkForDescribeInstanceTypesOutput {
	s.MaximumNetworkInterfaces = &v
	return s
}

// SetMaximumPrivateIpv4AddressesPerNetworkInterface sets the MaximumPrivateIpv4AddressesPerNetworkInterface field's value.
func (s *NetworkForDescribeInstanceTypesOutput) SetMaximumPrivateIpv4AddressesPerNetworkInterface(v int32) *NetworkForDescribeInstanceTypesOutput {
	s.MaximumPrivateIpv4AddressesPerNetworkInterface = &v
	return s
}

// SetMaximumQueuesPerNetworkInterface sets the MaximumQueuesPerNetworkInterface field's value.
func (s *NetworkForDescribeInstanceTypesOutput) SetMaximumQueuesPerNetworkInterface(v int32) *NetworkForDescribeInstanceTypesOutput {
	s.MaximumQueuesPerNetworkInterface = &v
	return s
}

// SetMaximumThroughputKpps sets the MaximumThroughputKpps field's value.
func (s *NetworkForDescribeInstanceTypesOutput) SetMaximumThroughputKpps(v int32) *NetworkForDescribeInstanceTypesOutput {
	s.MaximumThroughputKpps = &v
	return s
}

type ProcessorForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	BaseFrequency *float64 `type:"float"`

	Cpus *int32 `type:"int32"`

	Model *string `type:"string"`

	TurboFrequency *float64 `type:"float"`
}

// String returns the string representation
func (s ProcessorForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ProcessorForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetBaseFrequency sets the BaseFrequency field's value.
func (s *ProcessorForDescribeInstanceTypesOutput) SetBaseFrequency(v float64) *ProcessorForDescribeInstanceTypesOutput {
	s.BaseFrequency = &v
	return s
}

// SetCpus sets the Cpus field's value.
func (s *ProcessorForDescribeInstanceTypesOutput) SetCpus(v int32) *ProcessorForDescribeInstanceTypesOutput {
	s.Cpus = &v
	return s
}

// SetModel sets the Model field's value.
func (s *ProcessorForDescribeInstanceTypesOutput) SetModel(v string) *ProcessorForDescribeInstanceTypesOutput {
	s.Model = &v
	return s
}

// SetTurboFrequency sets the TurboFrequency field's value.
func (s *ProcessorForDescribeInstanceTypesOutput) SetTurboFrequency(v float64) *ProcessorForDescribeInstanceTypesOutput {
	s.TurboFrequency = &v
	return s
}

type RdmaForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	RdmaNetworkInterfaces *int32 `type:"int32"`

	RdmaProductName *string `type:"string"`
}

// String returns the string representation
func (s RdmaForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s RdmaForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetRdmaNetworkInterfaces sets the RdmaNetworkInterfaces field's value.
func (s *RdmaForDescribeInstanceTypesOutput) SetRdmaNetworkInterfaces(v int32) *RdmaForDescribeInstanceTypesOutput {
	s.RdmaNetworkInterfaces = &v
	return s
}

// SetRdmaProductName sets the RdmaProductName field's value.
func (s *RdmaForDescribeInstanceTypesOutput) SetRdmaProductName(v string) *RdmaForDescribeInstanceTypesOutput {
	s.RdmaProductName = &v
	return s
}

type VolumeForDescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	MaximumCount *int32 `type:"int32"`

	SupportedVolumeTypes []*string `type:"list"`
}

// String returns the string representation
func (s VolumeForDescribeInstanceTypesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForDescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

// SetMaximumCount sets the MaximumCount field's value.
func (s *VolumeForDescribeInstanceTypesOutput) SetMaximumCount(v int32) *VolumeForDescribeInstanceTypesOutput {
	s.MaximumCount = &v
	return s
}

// SetSupportedVolumeTypes sets the SupportedVolumeTypes field's value.
func (s *VolumeForDescribeInstanceTypesOutput) SetSupportedVolumeTypes(v []*string) *VolumeForDescribeInstanceTypesOutput {
	s.SupportedVolumeTypes = v
	return s
}
