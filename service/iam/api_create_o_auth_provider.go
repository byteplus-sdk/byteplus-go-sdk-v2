// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opCreateOAuthProviderCommon = "CreateOAuthProvider"

// CreateOAuthProviderCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateOAuthProviderCommon operation. The "output" return
// value will be populated with the CreateOAuthProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateOAuthProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateOAuthProviderCommon Send returns without error.
//
// See CreateOAuthProviderCommon for more information on using the CreateOAuthProviderCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateOAuthProviderCommonRequest method.
//	req, resp := client.CreateOAuthProviderCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) CreateOAuthProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateOAuthProviderCommon,
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

// CreateOAuthProviderCommon API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation CreateOAuthProviderCommon for usage and error information.
func (c *IAM) CreateOAuthProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateOAuthProviderCommonRequest(input)
	return out, req.Send()
}

// CreateOAuthProviderCommonWithContext is the same as CreateOAuthProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateOAuthProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateOAuthProviderCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateOAuthProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateOAuthProvider = "CreateOAuthProvider"

// CreateOAuthProviderRequest generates a "byteplus/request.Request" representing the
// client's request for the CreateOAuthProvider operation. The "output" return
// value will be populated with the CreateOAuthProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateOAuthProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateOAuthProviderCommon Send returns without error.
//
// See CreateOAuthProvider for more information on using the CreateOAuthProvider
// API call, and error handling.
//
//	// Example sending a request using the CreateOAuthProviderRequest method.
//	req, resp := client.CreateOAuthProviderRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) CreateOAuthProviderRequest(input *CreateOAuthProviderInput) (req *request.Request, output *CreateOAuthProviderOutput) {
	op := &request.Operation{
		Name:       opCreateOAuthProvider,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateOAuthProviderInput{}
	}

	output = &CreateOAuthProviderOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateOAuthProvider API operation for IAM.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for IAM's
// API operation CreateOAuthProvider for usage and error information.
func (c *IAM) CreateOAuthProvider(input *CreateOAuthProviderInput) (*CreateOAuthProviderOutput, error) {
	req, out := c.CreateOAuthProviderRequest(input)
	return out, req.Send()
}

// CreateOAuthProviderWithContext is the same as CreateOAuthProvider with the addition of
// the ability to pass a context and additional request options.
//
// See CreateOAuthProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateOAuthProviderWithContext(ctx byteplus.Context, input *CreateOAuthProviderInput, opts ...request.Option) (*CreateOAuthProviderOutput, error) {
	req, out := c.CreateOAuthProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateOAuthProviderInput struct {
	_ struct{} `type:"structure"`

	// AuthorizeTemplate is a required field
	AuthorizeTemplate *string `type:"string" required:"true"`

	// AuthorizeURL is a required field
	AuthorizeURL *string `type:"string" required:"true"`

	// ClientId is a required field
	ClientId *string `type:"string" required:"true"`

	// ClientSecret is a required field
	ClientSecret *string `type:"string" required:"true"`

	Description *string `type:"string"`

	// IdentityMapType is a required field
	IdentityMapType *int32 `type:"int32" required:"true"`

	// IdpIdentityKey is a required field
	IdpIdentityKey *string `type:"string" required:"true"`

	// OAuthProviderName is a required field
	OAuthProviderName *string `type:"string" required:"true"`

	// SSOType is a required field
	SSOType *int32 `type:"int32" required:"true"`

	Scope *string `type:"string"`

	Status *int32 `type:"int32"`

	// TokenURL is a required field
	TokenURL *string `type:"string" required:"true"`

	// UserInfoURL is a required field
	UserInfoURL *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateOAuthProviderInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateOAuthProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOAuthProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateOAuthProviderInput"}
	if s.AuthorizeTemplate == nil {
		invalidParams.Add(request.NewErrParamRequired("AuthorizeTemplate"))
	}
	if s.AuthorizeURL == nil {
		invalidParams.Add(request.NewErrParamRequired("AuthorizeURL"))
	}
	if s.ClientId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClientId"))
	}
	if s.ClientSecret == nil {
		invalidParams.Add(request.NewErrParamRequired("ClientSecret"))
	}
	if s.IdentityMapType == nil {
		invalidParams.Add(request.NewErrParamRequired("IdentityMapType"))
	}
	if s.IdpIdentityKey == nil {
		invalidParams.Add(request.NewErrParamRequired("IdpIdentityKey"))
	}
	if s.OAuthProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("OAuthProviderName"))
	}
	if s.SSOType == nil {
		invalidParams.Add(request.NewErrParamRequired("SSOType"))
	}
	if s.TokenURL == nil {
		invalidParams.Add(request.NewErrParamRequired("TokenURL"))
	}
	if s.UserInfoURL == nil {
		invalidParams.Add(request.NewErrParamRequired("UserInfoURL"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAuthorizeTemplate sets the AuthorizeTemplate field's value.
func (s *CreateOAuthProviderInput) SetAuthorizeTemplate(v string) *CreateOAuthProviderInput {
	s.AuthorizeTemplate = &v
	return s
}

// SetAuthorizeURL sets the AuthorizeURL field's value.
func (s *CreateOAuthProviderInput) SetAuthorizeURL(v string) *CreateOAuthProviderInput {
	s.AuthorizeURL = &v
	return s
}

// SetClientId sets the ClientId field's value.
func (s *CreateOAuthProviderInput) SetClientId(v string) *CreateOAuthProviderInput {
	s.ClientId = &v
	return s
}

// SetClientSecret sets the ClientSecret field's value.
func (s *CreateOAuthProviderInput) SetClientSecret(v string) *CreateOAuthProviderInput {
	s.ClientSecret = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateOAuthProviderInput) SetDescription(v string) *CreateOAuthProviderInput {
	s.Description = &v
	return s
}

// SetIdentityMapType sets the IdentityMapType field's value.
func (s *CreateOAuthProviderInput) SetIdentityMapType(v int32) *CreateOAuthProviderInput {
	s.IdentityMapType = &v
	return s
}

// SetIdpIdentityKey sets the IdpIdentityKey field's value.
func (s *CreateOAuthProviderInput) SetIdpIdentityKey(v string) *CreateOAuthProviderInput {
	s.IdpIdentityKey = &v
	return s
}

// SetOAuthProviderName sets the OAuthProviderName field's value.
func (s *CreateOAuthProviderInput) SetOAuthProviderName(v string) *CreateOAuthProviderInput {
	s.OAuthProviderName = &v
	return s
}

// SetSSOType sets the SSOType field's value.
func (s *CreateOAuthProviderInput) SetSSOType(v int32) *CreateOAuthProviderInput {
	s.SSOType = &v
	return s
}

// SetScope sets the Scope field's value.
func (s *CreateOAuthProviderInput) SetScope(v string) *CreateOAuthProviderInput {
	s.Scope = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CreateOAuthProviderInput) SetStatus(v int32) *CreateOAuthProviderInput {
	s.Status = &v
	return s
}

// SetTokenURL sets the TokenURL field's value.
func (s *CreateOAuthProviderInput) SetTokenURL(v string) *CreateOAuthProviderInput {
	s.TokenURL = &v
	return s
}

// SetUserInfoURL sets the UserInfoURL field's value.
func (s *CreateOAuthProviderInput) SetUserInfoURL(v string) *CreateOAuthProviderInput {
	s.UserInfoURL = &v
	return s
}

type CreateOAuthProviderOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	OAuthProviderName *string `type:"string"`

	SSOType *int32 `type:"int32"`

	Status *int32 `type:"int32"`

	Trn *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s CreateOAuthProviderOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateOAuthProviderOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *CreateOAuthProviderOutput) SetCreateDate(v string) *CreateOAuthProviderOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateOAuthProviderOutput) SetDescription(v string) *CreateOAuthProviderOutput {
	s.Description = &v
	return s
}

// SetOAuthProviderName sets the OAuthProviderName field's value.
func (s *CreateOAuthProviderOutput) SetOAuthProviderName(v string) *CreateOAuthProviderOutput {
	s.OAuthProviderName = &v
	return s
}

// SetSSOType sets the SSOType field's value.
func (s *CreateOAuthProviderOutput) SetSSOType(v int32) *CreateOAuthProviderOutput {
	s.SSOType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CreateOAuthProviderOutput) SetStatus(v int32) *CreateOAuthProviderOutput {
	s.Status = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *CreateOAuthProviderOutput) SetTrn(v string) *CreateOAuthProviderOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *CreateOAuthProviderOutput) SetUpdateDate(v string) *CreateOAuthProviderOutput {
	s.UpdateDate = &v
	return s
}
