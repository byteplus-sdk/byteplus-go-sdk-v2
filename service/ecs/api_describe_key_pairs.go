// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

const opDescribeKeyPairsCommon = "DescribeKeyPairs"

// DescribeKeyPairsCommonRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeKeyPairsCommon operation. The "output" return
// value will be populated with the DescribeKeyPairsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyPairsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyPairsCommon Send returns without error.
//
// See DescribeKeyPairsCommon for more information on using the DescribeKeyPairsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyPairsCommonRequest method.
//    req, resp := client.DescribeKeyPairsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeKeyPairsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeKeyPairsCommon,
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

// DescribeKeyPairsCommon API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeKeyPairsCommon for usage and error information.
func (c *ECS) DescribeKeyPairsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyPairsCommonRequest(input)
	return out, req.Send()
}

// DescribeKeyPairsCommonWithContext is the same as DescribeKeyPairsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKeyPairsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeKeyPairsCommonWithContext(ctx byteplus.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeKeyPairsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeKeyPairs = "DescribeKeyPairs"

// DescribeKeyPairsRequest generates a "byteplus/request.Request" representing the
// client's request for the DescribeKeyPairs operation. The "output" return
// value will be populated with the DescribeKeyPairsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeKeyPairsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeKeyPairsCommon Send returns without error.
//
// See DescribeKeyPairs for more information on using the DescribeKeyPairs
// API call, and error handling.
//
//    // Example sending a request using the DescribeKeyPairsRequest method.
//    req, resp := client.DescribeKeyPairsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeKeyPairsRequest(input *DescribeKeyPairsInput) (req *request.Request, output *DescribeKeyPairsOutput) {
	op := &request.Operation{
		Name:       opDescribeKeyPairs,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeKeyPairsInput{}
	}

	output = &DescribeKeyPairsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeKeyPairs API operation for ECS.
//
// Returns bytepluserr.Error for service API and SDK errors. Use runtime type assertions
// with bytepluserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the BYTEPLUS API reference guide for ECS's
// API operation DescribeKeyPairs for usage and error information.
func (c *ECS) DescribeKeyPairs(input *DescribeKeyPairsInput) (*DescribeKeyPairsOutput, error) {
	req, out := c.DescribeKeyPairsRequest(input)
	return out, req.Send()
}

// DescribeKeyPairsWithContext is the same as DescribeKeyPairs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeKeyPairs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeKeyPairsWithContext(ctx byteplus.Context, input *DescribeKeyPairsInput, opts ...request.Option) (*DescribeKeyPairsOutput, error) {
	req, out := c.DescribeKeyPairsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeKeyPairsInput struct {
	_ struct{} `type:"structure"`

	FingerPrint *string `type:"string"`

	KeyPairIds []*string `type:"list"`

	KeyPairName *string `type:"string"`

	KeyPairNames []*string `type:"list"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	ProjectName *string `type:"string"`

	TagFilters []*TagFilterForDescribeKeyPairsInput `type:"list"`
}

// String returns the string representation
func (s DescribeKeyPairsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyPairsInput) GoString() string {
	return s.String()
}

// SetFingerPrint sets the FingerPrint field's value.
func (s *DescribeKeyPairsInput) SetFingerPrint(v string) *DescribeKeyPairsInput {
	s.FingerPrint = &v
	return s
}

// SetKeyPairIds sets the KeyPairIds field's value.
func (s *DescribeKeyPairsInput) SetKeyPairIds(v []*string) *DescribeKeyPairsInput {
	s.KeyPairIds = v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *DescribeKeyPairsInput) SetKeyPairName(v string) *DescribeKeyPairsInput {
	s.KeyPairName = &v
	return s
}

// SetKeyPairNames sets the KeyPairNames field's value.
func (s *DescribeKeyPairsInput) SetKeyPairNames(v []*string) *DescribeKeyPairsInput {
	s.KeyPairNames = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeKeyPairsInput) SetMaxResults(v int32) *DescribeKeyPairsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeKeyPairsInput) SetNextToken(v string) *DescribeKeyPairsInput {
	s.NextToken = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeKeyPairsInput) SetProjectName(v string) *DescribeKeyPairsInput {
	s.ProjectName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeKeyPairsInput) SetTagFilters(v []*TagFilterForDescribeKeyPairsInput) *DescribeKeyPairsInput {
	s.TagFilters = v
	return s
}

type DescribeKeyPairsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	KeyPairs []*KeyPairForDescribeKeyPairsOutput `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeKeyPairsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeKeyPairsOutput) GoString() string {
	return s.String()
}

// SetKeyPairs sets the KeyPairs field's value.
func (s *DescribeKeyPairsOutput) SetKeyPairs(v []*KeyPairForDescribeKeyPairsOutput) *DescribeKeyPairsOutput {
	s.KeyPairs = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeKeyPairsOutput) SetNextToken(v string) *DescribeKeyPairsOutput {
	s.NextToken = &v
	return s
}

type KeyPairForDescribeKeyPairsOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	FingerPrint *string `type:"string"`

	KeyPairId *string `type:"string"`

	KeyPairName *string `type:"string"`

	ProjectName *string `type:"string"`

	Tags []*TagForDescribeKeyPairsOutput `type:"list"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s KeyPairForDescribeKeyPairsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s KeyPairForDescribeKeyPairsOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetCreatedAt(v string) *KeyPairForDescribeKeyPairsOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetDescription(v string) *KeyPairForDescribeKeyPairsOutput {
	s.Description = &v
	return s
}

// SetFingerPrint sets the FingerPrint field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetFingerPrint(v string) *KeyPairForDescribeKeyPairsOutput {
	s.FingerPrint = &v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetKeyPairId(v string) *KeyPairForDescribeKeyPairsOutput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetKeyPairName(v string) *KeyPairForDescribeKeyPairsOutput {
	s.KeyPairName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetProjectName(v string) *KeyPairForDescribeKeyPairsOutput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetTags(v []*TagForDescribeKeyPairsOutput) *KeyPairForDescribeKeyPairsOutput {
	s.Tags = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *KeyPairForDescribeKeyPairsOutput) SetUpdatedAt(v string) *KeyPairForDescribeKeyPairsOutput {
	s.UpdatedAt = &v
	return s
}

type TagFilterForDescribeKeyPairsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeKeyPairsInput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeKeyPairsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeKeyPairsInput) SetKey(v string) *TagFilterForDescribeKeyPairsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeKeyPairsInput) SetValues(v []*string) *TagFilterForDescribeKeyPairsInput {
	s.Values = v
	return s
}

type TagForDescribeKeyPairsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeKeyPairsOutput) String() string {
	return byteplusutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeKeyPairsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeKeyPairsOutput) SetKey(v string) *TagForDescribeKeyPairsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeKeyPairsOutput) SetValue(v string) *TagForDescribeKeyPairsOutput {
	s.Value = &v
	return s
}
