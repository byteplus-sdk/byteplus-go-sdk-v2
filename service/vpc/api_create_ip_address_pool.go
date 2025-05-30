// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreateIpAddressPoolCommon = "CreateIpAddressPool"

// CreateIpAddressPoolCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateIpAddressPoolCommon operation. The "output" return
// value will be populated with the CreateIpAddressPoolCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateIpAddressPoolCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateIpAddressPoolCommon Send returns without error.
//
// See CreateIpAddressPoolCommon for more information on using the CreateIpAddressPoolCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateIpAddressPoolCommonRequest method.
//    req, resp := client.CreateIpAddressPoolCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateIpAddressPoolCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateIpAddressPoolCommon,
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

// CreateIpAddressPoolCommon API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreateIpAddressPoolCommon for usage and error information.
func (c *VPC) CreateIpAddressPoolCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateIpAddressPoolCommonRequest(input)
	return out, req.Send()
}

// CreateIpAddressPoolCommonWithContext is the same as CreateIpAddressPoolCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateIpAddressPoolCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateIpAddressPoolCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateIpAddressPoolCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateIpAddressPool = "CreateIpAddressPool"

// CreateIpAddressPoolRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateIpAddressPool operation. The "output" return
// value will be populated with the CreateIpAddressPoolCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateIpAddressPoolCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateIpAddressPoolCommon Send returns without error.
//
// See CreateIpAddressPool for more information on using the CreateIpAddressPool
// API call, and error handling.
//
//    // Example sending a request using the CreateIpAddressPoolRequest method.
//    req, resp := client.CreateIpAddressPoolRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateIpAddressPoolRequest(input *CreateIpAddressPoolInput) (req *request.Request, output *CreateIpAddressPoolOutput) {
	op := &request.Operation{
		Name:       opCreateIpAddressPool,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateIpAddressPoolInput{}
	}

	output = &CreateIpAddressPoolOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateIpAddressPool API operation for VPC.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for VPC's
// API operation CreateIpAddressPool for usage and error information.
func (c *VPC) CreateIpAddressPool(input *CreateIpAddressPoolInput) (*CreateIpAddressPoolOutput, error) {
	req, out := c.CreateIpAddressPoolRequest(input)
	return out, req.Send()
}

// CreateIpAddressPoolWithContext is the same as CreateIpAddressPool with the addition of
// the ability to pass a context and additional request options.
//
// See CreateIpAddressPool for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateIpAddressPoolWithContext(ctx byteplus.Context, input *CreateIpAddressPoolInput, opts ...request.Option) (*CreateIpAddressPoolOutput, error) {
	req, out := c.CreateIpAddressPoolRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateIpAddressPoolInput struct {
	_ struct{} `type:"structure"`

	CidrBlock *string `type:"string"`

	CidrMask *int64 `type:"integer"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// ISP is a required field
	ISP *string `type:"string" required:"true"`

	Name *string `type:"string"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreateIpAddressPoolInput `type:"list"`
}

// String returns the string representation
func (s CreateIpAddressPoolInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateIpAddressPoolInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIpAddressPoolInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateIpAddressPoolInput"}
	if s.ISP == nil {
		invalidParams.Add(request.NewErrParamRequired("ISP"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *CreateIpAddressPoolInput) SetCidrBlock(v string) *CreateIpAddressPoolInput {
	s.CidrBlock = &v
	return s
}

// SetCidrMask sets the CidrMask field's value.
func (s *CreateIpAddressPoolInput) SetCidrMask(v int64) *CreateIpAddressPoolInput {
	s.CidrMask = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateIpAddressPoolInput) SetClientToken(v string) *CreateIpAddressPoolInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateIpAddressPoolInput) SetDescription(v string) *CreateIpAddressPoolInput {
	s.Description = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *CreateIpAddressPoolInput) SetISP(v string) *CreateIpAddressPoolInput {
	s.ISP = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateIpAddressPoolInput) SetName(v string) *CreateIpAddressPoolInput {
	s.Name = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateIpAddressPoolInput) SetProjectName(v string) *CreateIpAddressPoolInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateIpAddressPoolInput) SetTags(v []*TagForCreateIpAddressPoolInput) *CreateIpAddressPoolInput {
	s.Tags = v
	return s
}

type CreateIpAddressPoolOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	IpAddressPoolId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateIpAddressPoolOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateIpAddressPoolOutput) GoString() string {
	return s.String()
}

// SetIpAddressPoolId sets the IpAddressPoolId field's value.
func (s *CreateIpAddressPoolOutput) SetIpAddressPoolId(v string) *CreateIpAddressPoolOutput {
	s.IpAddressPoolId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateIpAddressPoolOutput) SetRequestId(v string) *CreateIpAddressPoolOutput {
	s.RequestId = &v
	return s
}

type TagForCreateIpAddressPoolInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateIpAddressPoolInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateIpAddressPoolInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateIpAddressPoolInput) SetKey(v string) *TagForCreateIpAddressPoolInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateIpAddressPoolInput) SetValue(v string) *TagForCreateIpAddressPoolInput {
	s.Value = &v
	return s
}
