// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreateBandwidthPackageCommon = "CreateBandwidthPackage"

// CreateBandwidthPackageCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateBandwidthPackageCommon operation. The "output" return
// value will be populated with the CreateBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateBandwidthPackageCommon Send returns without error.
//
// See CreateBandwidthPackageCommon for more information on using the CreateBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateBandwidthPackageCommonRequest method.
//    req, resp := client.CreateBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateBandwidthPackageCommon,
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

// CreateBandwidthPackageCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreateBandwidthPackageCommon for usage and error information.
func (c *VPC) CreateBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// CreateBandwidthPackageCommonWithContext is the same as CreateBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateBandwidthPackageCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateBandwidthPackage = "CreateBandwidthPackage"

// CreateBandwidthPackageRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateBandwidthPackage operation. The "output" return
// value will be populated with the CreateBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateBandwidthPackageCommon Send returns without error.
//
// See CreateBandwidthPackage for more information on using the CreateBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the CreateBandwidthPackageRequest method.
//    req, resp := client.CreateBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateBandwidthPackageRequest(input *CreateBandwidthPackageInput) (req *request.Request, output *CreateBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opCreateBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateBandwidthPackageInput{}
	}

	output = &CreateBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateBandwidthPackage API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreateBandwidthPackage for usage and error information.
func (c *VPC) CreateBandwidthPackage(input *CreateBandwidthPackageInput) (*CreateBandwidthPackageOutput, error) {
	req, out := c.CreateBandwidthPackageRequest(input)
	return out, req.Send()
}

// CreateBandwidthPackageWithContext is the same as CreateBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See CreateBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateBandwidthPackageWithContext(ctx byteplus.Context, input *CreateBandwidthPackageInput, opts ...request.Option) (*CreateBandwidthPackageOutput, error) {
	req, out := c.CreateBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	// Bandwidth is a required field
	Bandwidth *int64 `min:"2" type:"integer" required:"true"`

	BandwidthPackageName *string `min:"1" max:"128" type:"string"`

	BillingType *int64 `min:"1" max:"4" type:"integer"`

	Description *string `min:"1" max:"255" type:"string"`

	ISP *string `type:"string" enum:"ISPForCreateBandwidthPackageInput"`

	Period *int64 `type:"integer"`

	PeriodUnit *int64 `min:"1" max:"2" type:"integer"`

	ProjectName *string `type:"string"`

	Protocol *string `type:"string" enum:"ProtocolForCreateBandwidthPackageInput"`

	SecurityProtectionTypes []*string `type:"list"`

	Tags []*TagForCreateBandwidthPackageInput `type:"list"`
}

// String returns the string representation
func (s CreateBandwidthPackageInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateBandwidthPackageInput"}
	if s.Bandwidth == nil {
		invalidParams.Add(request.NewErrParamRequired("Bandwidth"))
	}
	if s.Bandwidth != nil && *s.Bandwidth < 2 {
		invalidParams.Add(request.NewErrParamMinValue("Bandwidth", 2))
	}
	if s.BandwidthPackageName != nil && len(*s.BandwidthPackageName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("BandwidthPackageName", 1))
	}
	if s.BandwidthPackageName != nil && len(*s.BandwidthPackageName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("BandwidthPackageName", 128, *s.BandwidthPackageName))
	}
	if s.BillingType != nil && *s.BillingType < 1 {
		invalidParams.Add(request.NewErrParamMinValue("BillingType", 1))
	}
	if s.BillingType != nil && *s.BillingType > 4 {
		invalidParams.Add(request.NewErrParamMaxValue("BillingType", 4))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.PeriodUnit != nil && *s.PeriodUnit < 1 {
		invalidParams.Add(request.NewErrParamMinValue("PeriodUnit", 1))
	}
	if s.PeriodUnit != nil && *s.PeriodUnit > 2 {
		invalidParams.Add(request.NewErrParamMaxValue("PeriodUnit", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *CreateBandwidthPackageInput) SetBandwidth(v int64) *CreateBandwidthPackageInput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageName sets the BandwidthPackageName field's value.
func (s *CreateBandwidthPackageInput) SetBandwidthPackageName(v string) *CreateBandwidthPackageInput {
	s.BandwidthPackageName = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *CreateBandwidthPackageInput) SetBillingType(v int64) *CreateBandwidthPackageInput {
	s.BillingType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateBandwidthPackageInput) SetDescription(v string) *CreateBandwidthPackageInput {
	s.Description = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *CreateBandwidthPackageInput) SetISP(v string) *CreateBandwidthPackageInput {
	s.ISP = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *CreateBandwidthPackageInput) SetPeriod(v int64) *CreateBandwidthPackageInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *CreateBandwidthPackageInput) SetPeriodUnit(v int64) *CreateBandwidthPackageInput {
	s.PeriodUnit = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateBandwidthPackageInput) SetProjectName(v string) *CreateBandwidthPackageInput {
	s.ProjectName = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *CreateBandwidthPackageInput) SetProtocol(v string) *CreateBandwidthPackageInput {
	s.Protocol = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *CreateBandwidthPackageInput) SetSecurityProtectionTypes(v []*string) *CreateBandwidthPackageInput {
	s.SecurityProtectionTypes = v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateBandwidthPackageInput) SetTags(v []*TagForCreateBandwidthPackageInput) *CreateBandwidthPackageInput {
	s.Tags = v
	return s
}

type CreateBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BandwidthPackageId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateBandwidthPackageOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateBandwidthPackageOutput) GoString() string {
	return s.String()
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *CreateBandwidthPackageOutput) SetBandwidthPackageId(v string) *CreateBandwidthPackageOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateBandwidthPackageOutput) SetRequestId(v string) *CreateBandwidthPackageOutput {
	s.RequestId = &v
	return s
}

type TagForCreateBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateBandwidthPackageInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateBandwidthPackageInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateBandwidthPackageInput) SetKey(v string) *TagForCreateBandwidthPackageInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateBandwidthPackageInput) SetValue(v string) *TagForCreateBandwidthPackageInput {
	s.Value = &v
	return s
}

const (
	// ISPForCreateBandwidthPackageInputBgp is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputBgp = "BGP"

	// ISPForCreateBandwidthPackageInputSingleLineBgp is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputSingleLineBgp = "SingleLine_BGP"

	// ISPForCreateBandwidthPackageInputStaticBgp is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputStaticBgp = "Static_BGP"

	// ISPForCreateBandwidthPackageInputFusionBgp is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputFusionBgp = "Fusion_BGP"

	// ISPForCreateBandwidthPackageInputChinaMobile is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaMobile = "ChinaMobile"

	// ISPForCreateBandwidthPackageInputChinaUnicom is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaUnicom = "ChinaUnicom"

	// ISPForCreateBandwidthPackageInputChinaTelecom is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaTelecom = "ChinaTelecom"

	// ISPForCreateBandwidthPackageInputChinaMobileValue is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaMobileValue = "ChinaMobile_Value"

	// ISPForCreateBandwidthPackageInputChinaUnicomValue is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaUnicomValue = "ChinaUnicom_Value"

	// ISPForCreateBandwidthPackageInputChinaTelecomValue is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaTelecomValue = "ChinaTelecom_Value"
)

const (
	// ProtocolForCreateBandwidthPackageInputIpv4 is a ProtocolForCreateBandwidthPackageInput enum value
	ProtocolForCreateBandwidthPackageInputIpv4 = "IPv4"

	// ProtocolForCreateBandwidthPackageInputDualStack is a ProtocolForCreateBandwidthPackageInput enum value
	ProtocolForCreateBandwidthPackageInputDualStack = "Dual-stack"

	// ProtocolForCreateBandwidthPackageInputIpv6 is a ProtocolForCreateBandwidthPackageInput enum value
	ProtocolForCreateBandwidthPackageInputIpv6 = "IPv6"
)
