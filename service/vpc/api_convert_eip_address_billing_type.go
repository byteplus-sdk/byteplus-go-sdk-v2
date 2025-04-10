// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opConvertEipAddressBillingTypeCommon = "ConvertEipAddressBillingType"

// ConvertEipAddressBillingTypeCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ConvertEipAddressBillingTypeCommon operation. The "output" return
// value will be populated with the ConvertEipAddressBillingTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ConvertEipAddressBillingTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ConvertEipAddressBillingTypeCommon Send returns without error.
//
// See ConvertEipAddressBillingTypeCommon for more information on using the ConvertEipAddressBillingTypeCommon
// API call, and error handling.
//
//    // Example sending a request using the ConvertEipAddressBillingTypeCommonRequest method.
//    req, resp := client.ConvertEipAddressBillingTypeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ConvertEipAddressBillingTypeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opConvertEipAddressBillingTypeCommon,
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

// ConvertEipAddressBillingTypeCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ConvertEipAddressBillingTypeCommon for usage and error information.
func (c *VPC) ConvertEipAddressBillingTypeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ConvertEipAddressBillingTypeCommonRequest(input)
	return out, req.Send()
}

// ConvertEipAddressBillingTypeCommonWithContext is the same as ConvertEipAddressBillingTypeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ConvertEipAddressBillingTypeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ConvertEipAddressBillingTypeCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ConvertEipAddressBillingTypeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opConvertEipAddressBillingType = "ConvertEipAddressBillingType"

// ConvertEipAddressBillingTypeRequest generates a "byteplus/request.Request" representing the
// client's request for the ConvertEipAddressBillingType operation. The "output" return
// value will be populated with the ConvertEipAddressBillingTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ConvertEipAddressBillingTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ConvertEipAddressBillingTypeCommon Send returns without error.
//
// See ConvertEipAddressBillingType for more information on using the ConvertEipAddressBillingType
// API call, and error handling.
//
//    // Example sending a request using the ConvertEipAddressBillingTypeRequest method.
//    req, resp := client.ConvertEipAddressBillingTypeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ConvertEipAddressBillingTypeRequest(input *ConvertEipAddressBillingTypeInput) (req *request.Request, output *ConvertEipAddressBillingTypeOutput) {
	op := &request.Operation{
		Name:       opConvertEipAddressBillingType,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConvertEipAddressBillingTypeInput{}
	}

	output = &ConvertEipAddressBillingTypeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ConvertEipAddressBillingType API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ConvertEipAddressBillingType for usage and error information.
func (c *VPC) ConvertEipAddressBillingType(input *ConvertEipAddressBillingTypeInput) (*ConvertEipAddressBillingTypeOutput, error) {
	req, out := c.ConvertEipAddressBillingTypeRequest(input)
	return out, req.Send()
}

// ConvertEipAddressBillingTypeWithContext is the same as ConvertEipAddressBillingType with the addition of
// the ability to pass a context and additional request options.
//
// See ConvertEipAddressBillingType for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ConvertEipAddressBillingTypeWithContext(ctx byteplus.Context, input *ConvertEipAddressBillingTypeInput, opts ...request.Option) (*ConvertEipAddressBillingTypeOutput, error) {
	req, out := c.ConvertEipAddressBillingTypeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConvertEipAddressBillingTypeInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`

	Bandwidth *int64 `type:"integer"`

	// BillingType is a required field
	BillingType *int64 `min:"1" max:"3" type:"integer" required:"true"`

	Period *int64 `min:"1" max:"60" type:"integer"`

	PeriodUnit *int64 `min:"1" max:"2" type:"integer"`
}

// String returns the string representation
func (s ConvertEipAddressBillingTypeInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertEipAddressBillingTypeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConvertEipAddressBillingTypeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ConvertEipAddressBillingTypeInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}
	if s.BillingType == nil {
		invalidParams.Add(request.NewErrParamRequired("BillingType"))
	}
	if s.BillingType != nil && *s.BillingType < 1 {
		invalidParams.Add(request.NewErrParamMinValue("BillingType", 1))
	}
	if s.BillingType != nil && *s.BillingType > 3 {
		invalidParams.Add(request.NewErrParamMaxValue("BillingType", 3))
	}
	if s.Period != nil && *s.Period < 1 {
		invalidParams.Add(request.NewErrParamMinValue("Period", 1))
	}
	if s.Period != nil && *s.Period > 60 {
		invalidParams.Add(request.NewErrParamMaxValue("Period", 60))
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

// SetAllocationId sets the AllocationId field's value.
func (s *ConvertEipAddressBillingTypeInput) SetAllocationId(v string) *ConvertEipAddressBillingTypeInput {
	s.AllocationId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ConvertEipAddressBillingTypeInput) SetBandwidth(v int64) *ConvertEipAddressBillingTypeInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *ConvertEipAddressBillingTypeInput) SetBillingType(v int64) *ConvertEipAddressBillingTypeInput {
	s.BillingType = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ConvertEipAddressBillingTypeInput) SetPeriod(v int64) *ConvertEipAddressBillingTypeInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ConvertEipAddressBillingTypeInput) SetPeriodUnit(v int64) *ConvertEipAddressBillingTypeInput {
	s.PeriodUnit = &v
	return s
}

type ConvertEipAddressBillingTypeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ConvertEipAddressBillingTypeOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertEipAddressBillingTypeOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ConvertEipAddressBillingTypeOutput) SetRequestId(v string) *ConvertEipAddressBillingTypeOutput {
	s.RequestId = &v
	return s
}
