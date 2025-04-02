// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyPrefixListCommon = "ModifyPrefixList"

// ModifyPrefixListCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyPrefixListCommon operation. The "output" return
// value will be populated with the ModifyPrefixListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyPrefixListCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyPrefixListCommon Send returns without error.
//
// See ModifyPrefixListCommon for more information on using the ModifyPrefixListCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyPrefixListCommonRequest method.
//    req, resp := client.ModifyPrefixListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyPrefixListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyPrefixListCommon,
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

// ModifyPrefixListCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyPrefixListCommon for usage and error information.
func (c *VPC) ModifyPrefixListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyPrefixListCommonRequest(input)
	return out, req.Send()
}

// ModifyPrefixListCommonWithContext is the same as ModifyPrefixListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyPrefixListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyPrefixListCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyPrefixListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyPrefixList = "ModifyPrefixList"

// ModifyPrefixListRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyPrefixList operation. The "output" return
// value will be populated with the ModifyPrefixListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyPrefixListCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyPrefixListCommon Send returns without error.
//
// See ModifyPrefixList for more information on using the ModifyPrefixList
// API call, and error handling.
//
//    // Example sending a request using the ModifyPrefixListRequest method.
//    req, resp := client.ModifyPrefixListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyPrefixListRequest(input *ModifyPrefixListInput) (req *request.Request, output *ModifyPrefixListOutput) {
	op := &request.Operation{
		Name:       opModifyPrefixList,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyPrefixListInput{}
	}

	output = &ModifyPrefixListOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyPrefixList API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyPrefixList for usage and error information.
func (c *VPC) ModifyPrefixList(input *ModifyPrefixListInput) (*ModifyPrefixListOutput, error) {
	req, out := c.ModifyPrefixListRequest(input)
	return out, req.Send()
}

// ModifyPrefixListWithContext is the same as ModifyPrefixList with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyPrefixList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyPrefixListWithContext(ctx byteplus.Context, input *ModifyPrefixListInput, opts ...request.Option) (*ModifyPrefixListOutput, error) {
	req, out := c.ModifyPrefixListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddPrefixListEntryForModifyPrefixListInput struct {
	_ struct{} `type:"structure"`

	Cidr *string `type:"string"`

	Description *string `type:"string"`
}

// String returns the string representation
func (s AddPrefixListEntryForModifyPrefixListInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AddPrefixListEntryForModifyPrefixListInput) GoString() string {
	return s.String()
}

// SetCidr sets the Cidr field's value.
func (s *AddPrefixListEntryForModifyPrefixListInput) SetCidr(v string) *AddPrefixListEntryForModifyPrefixListInput {
	s.Cidr = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AddPrefixListEntryForModifyPrefixListInput) SetDescription(v string) *AddPrefixListEntryForModifyPrefixListInput {
	s.Description = &v
	return s
}

type ModifyPrefixListInput struct {
	_ struct{} `type:"structure"`

	AddPrefixListEntries []*AddPrefixListEntryForModifyPrefixListInput `type:"list"`

	ClientToken *string `type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	DryRun *bool `type:"boolean"`

	MaxEntries *int64 `type:"integer"`

	// PrefixListId is a required field
	PrefixListId *string `type:"string" required:"true"`

	PrefixListName *string `min:"1" max:"128" type:"string"`

	RemovePrefixListEntries []*RemovePrefixListEntryForModifyPrefixListInput `type:"list"`
}

// String returns the string representation
func (s ModifyPrefixListInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyPrefixListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyPrefixListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyPrefixListInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.PrefixListId == nil {
		invalidParams.Add(request.NewErrParamRequired("PrefixListId"))
	}
	if s.PrefixListName != nil && len(*s.PrefixListName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("PrefixListName", 1))
	}
	if s.PrefixListName != nil && len(*s.PrefixListName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("PrefixListName", 128, *s.PrefixListName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAddPrefixListEntries sets the AddPrefixListEntries field's value.
func (s *ModifyPrefixListInput) SetAddPrefixListEntries(v []*AddPrefixListEntryForModifyPrefixListInput) *ModifyPrefixListInput {
	s.AddPrefixListEntries = v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyPrefixListInput) SetClientToken(v string) *ModifyPrefixListInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyPrefixListInput) SetDescription(v string) *ModifyPrefixListInput {
	s.Description = &v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *ModifyPrefixListInput) SetDryRun(v bool) *ModifyPrefixListInput {
	s.DryRun = &v
	return s
}

// SetMaxEntries sets the MaxEntries field's value.
func (s *ModifyPrefixListInput) SetMaxEntries(v int64) *ModifyPrefixListInput {
	s.MaxEntries = &v
	return s
}

// SetPrefixListId sets the PrefixListId field's value.
func (s *ModifyPrefixListInput) SetPrefixListId(v string) *ModifyPrefixListInput {
	s.PrefixListId = &v
	return s
}

// SetPrefixListName sets the PrefixListName field's value.
func (s *ModifyPrefixListInput) SetPrefixListName(v string) *ModifyPrefixListInput {
	s.PrefixListName = &v
	return s
}

// SetRemovePrefixListEntries sets the RemovePrefixListEntries field's value.
func (s *ModifyPrefixListInput) SetRemovePrefixListEntries(v []*RemovePrefixListEntryForModifyPrefixListInput) *ModifyPrefixListInput {
	s.RemovePrefixListEntries = v
	return s
}

type ModifyPrefixListOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyPrefixListOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyPrefixListOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyPrefixListOutput) SetRequestId(v string) *ModifyPrefixListOutput {
	s.RequestId = &v
	return s
}

type RemovePrefixListEntryForModifyPrefixListInput struct {
	_ struct{} `type:"structure"`

	Cidr *string `type:"string"`
}

// String returns the string representation
func (s RemovePrefixListEntryForModifyPrefixListInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s RemovePrefixListEntryForModifyPrefixListInput) GoString() string {
	return s.String()
}

// SetCidr sets the Cidr field's value.
func (s *RemovePrefixListEntryForModifyPrefixListInput) SetCidr(v string) *RemovePrefixListEntryForModifyPrefixListInput {
	s.Cidr = &v
	return s
}
