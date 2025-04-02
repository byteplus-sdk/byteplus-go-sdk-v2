// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opUpdateNetworkAclEntriesCommon = "UpdateNetworkAclEntries"

// UpdateNetworkAclEntriesCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the UpdateNetworkAclEntriesCommon operation. The "output" return
// value will be populated with the UpdateNetworkAclEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateNetworkAclEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateNetworkAclEntriesCommon Send returns without error.
//
// See UpdateNetworkAclEntriesCommon for more information on using the UpdateNetworkAclEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateNetworkAclEntriesCommonRequest method.
//    req, resp := client.UpdateNetworkAclEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) UpdateNetworkAclEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateNetworkAclEntriesCommon,
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

// UpdateNetworkAclEntriesCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation UpdateNetworkAclEntriesCommon for usage and error information.
func (c *VPC) UpdateNetworkAclEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateNetworkAclEntriesCommonRequest(input)
	return out, req.Send()
}

// UpdateNetworkAclEntriesCommonWithContext is the same as UpdateNetworkAclEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateNetworkAclEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) UpdateNetworkAclEntriesCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateNetworkAclEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateNetworkAclEntries = "UpdateNetworkAclEntries"

// UpdateNetworkAclEntriesRequest generates a "byteplus/request.Request" representing the
// client's request for the UpdateNetworkAclEntries operation. The "output" return
// value will be populated with the UpdateNetworkAclEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateNetworkAclEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateNetworkAclEntriesCommon Send returns without error.
//
// See UpdateNetworkAclEntries for more information on using the UpdateNetworkAclEntries
// API call, and error handling.
//
//    // Example sending a request using the UpdateNetworkAclEntriesRequest method.
//    req, resp := client.UpdateNetworkAclEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) UpdateNetworkAclEntriesRequest(input *UpdateNetworkAclEntriesInput) (req *request.Request, output *UpdateNetworkAclEntriesOutput) {
	op := &request.Operation{
		Name:       opUpdateNetworkAclEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNetworkAclEntriesInput{}
	}

	output = &UpdateNetworkAclEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateNetworkAclEntries API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation UpdateNetworkAclEntries for usage and error information.
func (c *VPC) UpdateNetworkAclEntries(input *UpdateNetworkAclEntriesInput) (*UpdateNetworkAclEntriesOutput, error) {
	req, out := c.UpdateNetworkAclEntriesRequest(input)
	return out, req.Send()
}

// UpdateNetworkAclEntriesWithContext is the same as UpdateNetworkAclEntries with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateNetworkAclEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) UpdateNetworkAclEntriesWithContext(ctx byteplus.Context, input *UpdateNetworkAclEntriesInput, opts ...request.Option) (*UpdateNetworkAclEntriesOutput, error) {
	req, out := c.UpdateNetworkAclEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EgressAclEntryForUpdateNetworkAclEntriesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DestinationCidrIp *string `type:"string"`

	NetworkAclEntryId *string `type:"string"`

	NetworkAclEntryName *string `type:"string"`

	Policy *string `type:"string"`

	Port *string `type:"string"`

	Protocol *string `type:"string"`
}

// String returns the string representation
func (s EgressAclEntryForUpdateNetworkAclEntriesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s EgressAclEntryForUpdateNetworkAclEntriesInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *EgressAclEntryForUpdateNetworkAclEntriesInput) SetDescription(v string) *EgressAclEntryForUpdateNetworkAclEntriesInput {
	s.Description = &v
	return s
}

// SetDestinationCidrIp sets the DestinationCidrIp field's value.
func (s *EgressAclEntryForUpdateNetworkAclEntriesInput) SetDestinationCidrIp(v string) *EgressAclEntryForUpdateNetworkAclEntriesInput {
	s.DestinationCidrIp = &v
	return s
}

// SetNetworkAclEntryId sets the NetworkAclEntryId field's value.
func (s *EgressAclEntryForUpdateNetworkAclEntriesInput) SetNetworkAclEntryId(v string) *EgressAclEntryForUpdateNetworkAclEntriesInput {
	s.NetworkAclEntryId = &v
	return s
}

// SetNetworkAclEntryName sets the NetworkAclEntryName field's value.
func (s *EgressAclEntryForUpdateNetworkAclEntriesInput) SetNetworkAclEntryName(v string) *EgressAclEntryForUpdateNetworkAclEntriesInput {
	s.NetworkAclEntryName = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *EgressAclEntryForUpdateNetworkAclEntriesInput) SetPolicy(v string) *EgressAclEntryForUpdateNetworkAclEntriesInput {
	s.Policy = &v
	return s
}

// SetPort sets the Port field's value.
func (s *EgressAclEntryForUpdateNetworkAclEntriesInput) SetPort(v string) *EgressAclEntryForUpdateNetworkAclEntriesInput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *EgressAclEntryForUpdateNetworkAclEntriesInput) SetProtocol(v string) *EgressAclEntryForUpdateNetworkAclEntriesInput {
	s.Protocol = &v
	return s
}

type IngressAclEntryForUpdateNetworkAclEntriesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	NetworkAclEntryId *string `type:"string"`

	NetworkAclEntryName *string `type:"string"`

	Policy *string `type:"string"`

	Port *string `type:"string"`

	Protocol *string `type:"string"`

	SourceCidrIp *string `type:"string"`
}

// String returns the string representation
func (s IngressAclEntryForUpdateNetworkAclEntriesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s IngressAclEntryForUpdateNetworkAclEntriesInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *IngressAclEntryForUpdateNetworkAclEntriesInput) SetDescription(v string) *IngressAclEntryForUpdateNetworkAclEntriesInput {
	s.Description = &v
	return s
}

// SetNetworkAclEntryId sets the NetworkAclEntryId field's value.
func (s *IngressAclEntryForUpdateNetworkAclEntriesInput) SetNetworkAclEntryId(v string) *IngressAclEntryForUpdateNetworkAclEntriesInput {
	s.NetworkAclEntryId = &v
	return s
}

// SetNetworkAclEntryName sets the NetworkAclEntryName field's value.
func (s *IngressAclEntryForUpdateNetworkAclEntriesInput) SetNetworkAclEntryName(v string) *IngressAclEntryForUpdateNetworkAclEntriesInput {
	s.NetworkAclEntryName = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *IngressAclEntryForUpdateNetworkAclEntriesInput) SetPolicy(v string) *IngressAclEntryForUpdateNetworkAclEntriesInput {
	s.Policy = &v
	return s
}

// SetPort sets the Port field's value.
func (s *IngressAclEntryForUpdateNetworkAclEntriesInput) SetPort(v string) *IngressAclEntryForUpdateNetworkAclEntriesInput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *IngressAclEntryForUpdateNetworkAclEntriesInput) SetProtocol(v string) *IngressAclEntryForUpdateNetworkAclEntriesInput {
	s.Protocol = &v
	return s
}

// SetSourceCidrIp sets the SourceCidrIp field's value.
func (s *IngressAclEntryForUpdateNetworkAclEntriesInput) SetSourceCidrIp(v string) *IngressAclEntryForUpdateNetworkAclEntriesInput {
	s.SourceCidrIp = &v
	return s
}

type UpdateNetworkAclEntriesInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	EgressAclEntries []*EgressAclEntryForUpdateNetworkAclEntriesInput `type:"list"`

	IngressAclEntries []*IngressAclEntryForUpdateNetworkAclEntriesInput `type:"list"`

	// NetworkAclId is a required field
	NetworkAclId *string `type:"string" required:"true"`

	UpdateEgressAclEntries *bool `type:"boolean"`

	UpdateIngressAclEntries *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateNetworkAclEntriesInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateNetworkAclEntriesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNetworkAclEntriesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateNetworkAclEntriesInput"}
	if s.NetworkAclId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkAclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *UpdateNetworkAclEntriesInput) SetClientToken(v string) *UpdateNetworkAclEntriesInput {
	s.ClientToken = &v
	return s
}

// SetEgressAclEntries sets the EgressAclEntries field's value.
func (s *UpdateNetworkAclEntriesInput) SetEgressAclEntries(v []*EgressAclEntryForUpdateNetworkAclEntriesInput) *UpdateNetworkAclEntriesInput {
	s.EgressAclEntries = v
	return s
}

// SetIngressAclEntries sets the IngressAclEntries field's value.
func (s *UpdateNetworkAclEntriesInput) SetIngressAclEntries(v []*IngressAclEntryForUpdateNetworkAclEntriesInput) *UpdateNetworkAclEntriesInput {
	s.IngressAclEntries = v
	return s
}

// SetNetworkAclId sets the NetworkAclId field's value.
func (s *UpdateNetworkAclEntriesInput) SetNetworkAclId(v string) *UpdateNetworkAclEntriesInput {
	s.NetworkAclId = &v
	return s
}

// SetUpdateEgressAclEntries sets the UpdateEgressAclEntries field's value.
func (s *UpdateNetworkAclEntriesInput) SetUpdateEgressAclEntries(v bool) *UpdateNetworkAclEntriesInput {
	s.UpdateEgressAclEntries = &v
	return s
}

// SetUpdateIngressAclEntries sets the UpdateIngressAclEntries field's value.
func (s *UpdateNetworkAclEntriesInput) SetUpdateIngressAclEntries(v bool) *UpdateNetworkAclEntriesInput {
	s.UpdateIngressAclEntries = &v
	return s
}

type UpdateNetworkAclEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s UpdateNetworkAclEntriesOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateNetworkAclEntriesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *UpdateNetworkAclEntriesOutput) SetRequestId(v string) *UpdateNetworkAclEntriesOutput {
	s.RequestId = &v
	return s
}
