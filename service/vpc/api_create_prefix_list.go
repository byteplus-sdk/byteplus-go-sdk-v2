// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreatePrefixListCommon = "CreatePrefixList"

// CreatePrefixListCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreatePrefixListCommon operation. The "output" return
// value will be populated with the CreatePrefixListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreatePrefixListCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreatePrefixListCommon Send returns without error.
//
// See CreatePrefixListCommon for more information on using the CreatePrefixListCommon
// API call, and error handling.
//
//    // Example sending a request using the CreatePrefixListCommonRequest method.
//    req, resp := client.CreatePrefixListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreatePrefixListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreatePrefixListCommon,
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

// CreatePrefixListCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreatePrefixListCommon for usage and error information.
func (c *VPC) CreatePrefixListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreatePrefixListCommonRequest(input)
	return out, req.Send()
}

// CreatePrefixListCommonWithContext is the same as CreatePrefixListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreatePrefixListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreatePrefixListCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreatePrefixListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreatePrefixList = "CreatePrefixList"

// CreatePrefixListRequest generates a "byteplus/request.Request" representing the
// client's request for the CreatePrefixList operation. The "output" return
// value will be populated with the CreatePrefixListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreatePrefixListCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreatePrefixListCommon Send returns without error.
//
// See CreatePrefixList for more information on using the CreatePrefixList
// API call, and error handling.
//
//    // Example sending a request using the CreatePrefixListRequest method.
//    req, resp := client.CreatePrefixListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreatePrefixListRequest(input *CreatePrefixListInput) (req *request.Request, output *CreatePrefixListOutput) {
	op := &request.Operation{
		Name:       opCreatePrefixList,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePrefixListInput{}
	}

	output = &CreatePrefixListOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreatePrefixList API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreatePrefixList for usage and error information.
func (c *VPC) CreatePrefixList(input *CreatePrefixListInput) (*CreatePrefixListOutput, error) {
	req, out := c.CreatePrefixListRequest(input)
	return out, req.Send()
}

// CreatePrefixListWithContext is the same as CreatePrefixList with the addition of
// the ability to pass a context and additional request options.
//
// See CreatePrefixList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreatePrefixListWithContext(ctx byteplus.Context, input *CreatePrefixListInput, opts ...request.Option) (*CreatePrefixListOutput, error) {
	req, out := c.CreatePrefixListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreatePrefixListInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	DryRun *bool `type:"boolean"`

	IpVersion *string `type:"string" enum:"IpVersionForCreatePrefixListInput"`

	// MaxEntries is a required field
	MaxEntries *int64 `type:"integer" required:"true"`

	PrefixListEntries []*PrefixListEntryForCreatePrefixListInput `type:"list"`

	PrefixListName *string `min:"1" max:"128" type:"string"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreatePrefixListInput `type:"list"`
}

// String returns the string representation
func (s CreatePrefixListInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatePrefixListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePrefixListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreatePrefixListInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.MaxEntries == nil {
		invalidParams.Add(request.NewErrParamRequired("MaxEntries"))
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

// SetClientToken sets the ClientToken field's value.
func (s *CreatePrefixListInput) SetClientToken(v string) *CreatePrefixListInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreatePrefixListInput) SetDescription(v string) *CreatePrefixListInput {
	s.Description = &v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *CreatePrefixListInput) SetDryRun(v bool) *CreatePrefixListInput {
	s.DryRun = &v
	return s
}

// SetIpVersion sets the IpVersion field's value.
func (s *CreatePrefixListInput) SetIpVersion(v string) *CreatePrefixListInput {
	s.IpVersion = &v
	return s
}

// SetMaxEntries sets the MaxEntries field's value.
func (s *CreatePrefixListInput) SetMaxEntries(v int64) *CreatePrefixListInput {
	s.MaxEntries = &v
	return s
}

// SetPrefixListEntries sets the PrefixListEntries field's value.
func (s *CreatePrefixListInput) SetPrefixListEntries(v []*PrefixListEntryForCreatePrefixListInput) *CreatePrefixListInput {
	s.PrefixListEntries = v
	return s
}

// SetPrefixListName sets the PrefixListName field's value.
func (s *CreatePrefixListInput) SetPrefixListName(v string) *CreatePrefixListInput {
	s.PrefixListName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreatePrefixListInput) SetProjectName(v string) *CreatePrefixListInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreatePrefixListInput) SetTags(v []*TagForCreatePrefixListInput) *CreatePrefixListInput {
	s.Tags = v
	return s
}

type CreatePrefixListOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PrefixListId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreatePrefixListOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatePrefixListOutput) GoString() string {
	return s.String()
}

// SetPrefixListId sets the PrefixListId field's value.
func (s *CreatePrefixListOutput) SetPrefixListId(v string) *CreatePrefixListOutput {
	s.PrefixListId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreatePrefixListOutput) SetRequestId(v string) *CreatePrefixListOutput {
	s.RequestId = &v
	return s
}

type PrefixListEntryForCreatePrefixListInput struct {
	_ struct{} `type:"structure"`

	Cidr *string `type:"string"`

	Description *string `type:"string"`
}

// String returns the string representation
func (s PrefixListEntryForCreatePrefixListInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s PrefixListEntryForCreatePrefixListInput) GoString() string {
	return s.String()
}

// SetCidr sets the Cidr field's value.
func (s *PrefixListEntryForCreatePrefixListInput) SetCidr(v string) *PrefixListEntryForCreatePrefixListInput {
	s.Cidr = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *PrefixListEntryForCreatePrefixListInput) SetDescription(v string) *PrefixListEntryForCreatePrefixListInput {
	s.Description = &v
	return s
}

type TagForCreatePrefixListInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreatePrefixListInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreatePrefixListInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreatePrefixListInput) SetKey(v string) *TagForCreatePrefixListInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreatePrefixListInput) SetValue(v string) *TagForCreatePrefixListInput {
	s.Value = &v
	return s
}

const (
	// IpVersionForCreatePrefixListInputIpv4 is a IpVersionForCreatePrefixListInput enum value
	IpVersionForCreatePrefixListInputIpv4 = "IPv4"

	// IpVersionForCreatePrefixListInputIpv6 is a IpVersionForCreatePrefixListInput enum value
	IpVersionForCreatePrefixListInputIpv6 = "IPv6"
)
