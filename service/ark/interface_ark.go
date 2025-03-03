// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package arkiface provides an interface to enable mocking the ARK service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ark

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

// ARKAPI provides an interface to enable mocking the
// ark.ARK service client's API operation,
//
//	// byteplus sdk func uses an SDK service client to make a request to
//	// ARK.
//	func myFunc(svc ARKAPI) bool {
//	    // Make svc.CreateBatchInferenceJob request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := ark.New(sess)
//
//	    myFunc(svc)
//	}
type ARKAPI interface {
	CreateBatchInferenceJobCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateBatchInferenceJobCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateBatchInferenceJobCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateBatchInferenceJob(*CreateBatchInferenceJobInput) (*CreateBatchInferenceJobOutput, error)
	CreateBatchInferenceJobWithContext(byteplus.Context, *CreateBatchInferenceJobInput, ...request.Option) (*CreateBatchInferenceJobOutput, error)
	CreateBatchInferenceJobRequest(*CreateBatchInferenceJobInput) (*request.Request, *CreateBatchInferenceJobOutput)

	CreateEndpointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateEndpointCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateEndpointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateEndpoint(*CreateEndpointInput) (*CreateEndpointOutput, error)
	CreateEndpointWithContext(byteplus.Context, *CreateEndpointInput, ...request.Option) (*CreateEndpointOutput, error)
	CreateEndpointRequest(*CreateEndpointInput) (*request.Request, *CreateEndpointOutput)

	CreateEvaluationJobCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateEvaluationJobCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateEvaluationJobCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateEvaluationJob(*CreateEvaluationJobInput) (*CreateEvaluationJobOutput, error)
	CreateEvaluationJobWithContext(byteplus.Context, *CreateEvaluationJobInput, ...request.Option) (*CreateEvaluationJobOutput, error)
	CreateEvaluationJobRequest(*CreateEvaluationJobInput) (*request.Request, *CreateEvaluationJobOutput)

	CreateModelCustomizationJobCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateModelCustomizationJobCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateModelCustomizationJobCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateModelCustomizationJob(*CreateModelCustomizationJobInput) (*CreateModelCustomizationJobOutput, error)
	CreateModelCustomizationJobWithContext(byteplus.Context, *CreateModelCustomizationJobInput, ...request.Option) (*CreateModelCustomizationJobOutput, error)
	CreateModelCustomizationJobRequest(*CreateModelCustomizationJobInput) (*request.Request, *CreateModelCustomizationJobOutput)

	DeleteEndpointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteEndpointCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteEndpointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteEndpoint(*DeleteEndpointInput) (*DeleteEndpointOutput, error)
	DeleteEndpointWithContext(byteplus.Context, *DeleteEndpointInput, ...request.Option) (*DeleteEndpointOutput, error)
	DeleteEndpointRequest(*DeleteEndpointInput) (*request.Request, *DeleteEndpointOutput)

	GetApiKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetApiKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetApiKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetApiKey(*GetApiKeyInput) (*GetApiKeyOutput, error)
	GetApiKeyWithContext(byteplus.Context, *GetApiKeyInput, ...request.Option) (*GetApiKeyOutput, error)
	GetApiKeyRequest(*GetApiKeyInput) (*request.Request, *GetApiKeyOutput)

	GetEndpointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetEndpointCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetEndpointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetEndpoint(*GetEndpointInput) (*GetEndpointOutput, error)
	GetEndpointWithContext(byteplus.Context, *GetEndpointInput, ...request.Option) (*GetEndpointOutput, error)
	GetEndpointRequest(*GetEndpointInput) (*request.Request, *GetEndpointOutput)

	GetEndpointCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetEndpointCertificateCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetEndpointCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetEndpointCertificate(*GetEndpointCertificateInput) (*GetEndpointCertificateOutput, error)
	GetEndpointCertificateWithContext(byteplus.Context, *GetEndpointCertificateInput, ...request.Option) (*GetEndpointCertificateOutput, error)
	GetEndpointCertificateRequest(*GetEndpointCertificateInput) (*request.Request, *GetEndpointCertificateOutput)

	GetModelCustomizationJobCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetModelCustomizationJobCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetModelCustomizationJobCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetModelCustomizationJob(*GetModelCustomizationJobInput) (*GetModelCustomizationJobOutput, error)
	GetModelCustomizationJobWithContext(byteplus.Context, *GetModelCustomizationJobInput, ...request.Option) (*GetModelCustomizationJobOutput, error)
	GetModelCustomizationJobRequest(*GetModelCustomizationJobInput) (*request.Request, *GetModelCustomizationJobOutput)

	ListBatchInferenceJobsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBatchInferenceJobsCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBatchInferenceJobsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBatchInferenceJobs(*ListBatchInferenceJobsInput) (*ListBatchInferenceJobsOutput, error)
	ListBatchInferenceJobsWithContext(byteplus.Context, *ListBatchInferenceJobsInput, ...request.Option) (*ListBatchInferenceJobsOutput, error)
	ListBatchInferenceJobsRequest(*ListBatchInferenceJobsInput) (*request.Request, *ListBatchInferenceJobsOutput)

	ListModelCustomizationJobsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListModelCustomizationJobsCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListModelCustomizationJobsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListModelCustomizationJobs(*ListModelCustomizationJobsInput) (*ListModelCustomizationJobsOutput, error)
	ListModelCustomizationJobsWithContext(byteplus.Context, *ListModelCustomizationJobsInput, ...request.Option) (*ListModelCustomizationJobsOutput, error)
	ListModelCustomizationJobsRequest(*ListModelCustomizationJobsInput) (*request.Request, *ListModelCustomizationJobsOutput)
}

var _ ARKAPI = (*ARK)(nil)
