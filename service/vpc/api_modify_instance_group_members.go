// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opModifyInstanceGroupMembersCommon = "ModifyInstanceGroupMembers"

// ModifyInstanceGroupMembersCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyInstanceGroupMembersCommon operation. The "output" return
// value will be populated with the ModifyInstanceGroupMembersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceGroupMembersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceGroupMembersCommon Send returns without error.
//
// See ModifyInstanceGroupMembersCommon for more information on using the ModifyInstanceGroupMembersCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceGroupMembersCommonRequest method.
//    req, resp := client.ModifyInstanceGroupMembersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyInstanceGroupMembersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyInstanceGroupMembersCommon,
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

// ModifyInstanceGroupMembersCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyInstanceGroupMembersCommon for usage and error information.
func (c *VPC) ModifyInstanceGroupMembersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceGroupMembersCommonRequest(input)
	return out, req.Send()
}

// ModifyInstanceGroupMembersCommonWithContext is the same as ModifyInstanceGroupMembersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceGroupMembersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyInstanceGroupMembersCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceGroupMembersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyInstanceGroupMembers = "ModifyInstanceGroupMembers"

// ModifyInstanceGroupMembersRequest generates a "byteplus/request.Request" representing the
// client's request for the ModifyInstanceGroupMembers operation. The "output" return
// value will be populated with the ModifyInstanceGroupMembersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceGroupMembersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceGroupMembersCommon Send returns without error.
//
// See ModifyInstanceGroupMembers for more information on using the ModifyInstanceGroupMembers
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceGroupMembersRequest method.
//    req, resp := client.ModifyInstanceGroupMembersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyInstanceGroupMembersRequest(input *ModifyInstanceGroupMembersInput) (req *request.Request, output *ModifyInstanceGroupMembersOutput) {
	op := &request.Operation{
		Name:       opModifyInstanceGroupMembers,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceGroupMembersInput{}
	}

	output = &ModifyInstanceGroupMembersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyInstanceGroupMembers API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation ModifyInstanceGroupMembers for usage and error information.
func (c *VPC) ModifyInstanceGroupMembers(input *ModifyInstanceGroupMembersInput) (*ModifyInstanceGroupMembersOutput, error) {
	req, out := c.ModifyInstanceGroupMembersRequest(input)
	return out, req.Send()
}

// ModifyInstanceGroupMembersWithContext is the same as ModifyInstanceGroupMembers with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceGroupMembers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyInstanceGroupMembersWithContext(ctx byteplus.Context, input *ModifyInstanceGroupMembersInput, opts ...request.Option) (*ModifyInstanceGroupMembersOutput, error) {
	req, out := c.ModifyInstanceGroupMembersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AppendMemberForModifyInstanceGroupMembersInput struct {
	_ struct{} `type:"structure"`

	Id *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s AppendMemberForModifyInstanceGroupMembersInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s AppendMemberForModifyInstanceGroupMembersInput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *AppendMemberForModifyInstanceGroupMembersInput) SetId(v string) *AppendMemberForModifyInstanceGroupMembersInput {
	s.Id = &v
	return s
}

// SetType sets the Type field's value.
func (s *AppendMemberForModifyInstanceGroupMembersInput) SetType(v string) *AppendMemberForModifyInstanceGroupMembersInput {
	s.Type = &v
	return s
}

type ModifyInstanceGroupMembersInput struct {
	_ struct{} `type:"structure"`

	AppendMembers []*AppendMemberForModifyInstanceGroupMembersInput `type:"list"`

	// InstanceGroupId is a required field
	InstanceGroupId *string `type:"string" required:"true"`

	RemoveMembers []*RemoveMemberForModifyInstanceGroupMembersInput `type:"list"`
}

// String returns the string representation
func (s ModifyInstanceGroupMembersInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceGroupMembersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstanceGroupMembersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyInstanceGroupMembersInput"}
	if s.InstanceGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAppendMembers sets the AppendMembers field's value.
func (s *ModifyInstanceGroupMembersInput) SetAppendMembers(v []*AppendMemberForModifyInstanceGroupMembersInput) *ModifyInstanceGroupMembersInput {
	s.AppendMembers = v
	return s
}

// SetInstanceGroupId sets the InstanceGroupId field's value.
func (s *ModifyInstanceGroupMembersInput) SetInstanceGroupId(v string) *ModifyInstanceGroupMembersInput {
	s.InstanceGroupId = &v
	return s
}

// SetRemoveMembers sets the RemoveMembers field's value.
func (s *ModifyInstanceGroupMembersInput) SetRemoveMembers(v []*RemoveMemberForModifyInstanceGroupMembersInput) *ModifyInstanceGroupMembersInput {
	s.RemoveMembers = v
	return s
}

type ModifyInstanceGroupMembersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyInstanceGroupMembersOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceGroupMembersOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyInstanceGroupMembersOutput) SetRequestId(v string) *ModifyInstanceGroupMembersOutput {
	s.RequestId = &v
	return s
}

type RemoveMemberForModifyInstanceGroupMembersInput struct {
	_ struct{} `type:"structure"`

	Id *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s RemoveMemberForModifyInstanceGroupMembersInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveMemberForModifyInstanceGroupMembersInput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *RemoveMemberForModifyInstanceGroupMembersInput) SetId(v string) *RemoveMemberForModifyInstanceGroupMembersInput {
	s.Id = &v
	return s
}

// SetType sets the Type field's value.
func (s *RemoveMemberForModifyInstanceGroupMembersInput) SetType(v string) *RemoveMemberForModifyInstanceGroupMembersInput {
	s.Type = &v
	return s
}
