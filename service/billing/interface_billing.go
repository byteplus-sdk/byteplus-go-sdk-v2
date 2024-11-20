// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package billingiface provides an interface to enable mocking the BILLING service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package billing

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

// BILLINGAPI provides an interface to enable mocking the
// billing.BILLING service client's API operation,
//
//    // byteplus sdk func uses an SDK service client to make a request to
//    // BILLING.
//    func myFunc(svc BILLINGAPI) bool {
//        // Make svc.CancelInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := billing.New(sess)
//
//        myFunc(svc)
//    }
//
type BILLINGAPI interface {
	CancelInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelInvitationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelInvitation(*CancelInvitationInput) (*CancelInvitationOutput, error)
	CancelInvitationWithContext(byteplus.Context, *CancelInvitationInput, ...request.Option) (*CancelInvitationOutput, error)
	CancelInvitationRequest(*CancelInvitationInput) (*request.Request, *CancelInvitationOutput)

	CreateFinancialRelationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateFinancialRelationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateFinancialRelationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateFinancialRelation(*CreateFinancialRelationInput) (*CreateFinancialRelationOutput, error)
	CreateFinancialRelationWithContext(byteplus.Context, *CreateFinancialRelationInput, ...request.Option) (*CreateFinancialRelationOutput, error)
	CreateFinancialRelationRequest(*CreateFinancialRelationInput) (*request.Request, *CreateFinancialRelationOutput)

	DeleteFinancialRelationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteFinancialRelationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteFinancialRelationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteFinancialRelation(*DeleteFinancialRelationInput) (*DeleteFinancialRelationOutput, error)
	DeleteFinancialRelationWithContext(byteplus.Context, *DeleteFinancialRelationInput, ...request.Option) (*DeleteFinancialRelationOutput, error)
	DeleteFinancialRelationRequest(*DeleteFinancialRelationInput) (*request.Request, *DeleteFinancialRelationOutput)

	HandleInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	HandleInvitationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	HandleInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	HandleInvitation(*HandleInvitationInput) (*HandleInvitationOutput, error)
	HandleInvitationWithContext(byteplus.Context, *HandleInvitationInput, ...request.Option) (*HandleInvitationOutput, error)
	HandleInvitationRequest(*HandleInvitationInput) (*request.Request, *HandleInvitationOutput)

	ListAvailableInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAvailableInstancesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAvailableInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAvailableInstances(*ListAvailableInstancesInput) (*ListAvailableInstancesOutput, error)
	ListAvailableInstancesWithContext(byteplus.Context, *ListAvailableInstancesInput, ...request.Option) (*ListAvailableInstancesOutput, error)
	ListAvailableInstancesRequest(*ListAvailableInstancesInput) (*request.Request, *ListAvailableInstancesOutput)

	ListBillDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBillDetailCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBillDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBillDetail(*ListBillDetailInput) (*ListBillDetailOutput, error)
	ListBillDetailWithContext(byteplus.Context, *ListBillDetailInput, ...request.Option) (*ListBillDetailOutput, error)
	ListBillDetailRequest(*ListBillDetailInput) (*request.Request, *ListBillDetailOutput)

	ListBillOverviewByCategoryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBillOverviewByCategoryCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBillOverviewByCategoryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBillOverviewByCategory(*ListBillOverviewByCategoryInput) (*ListBillOverviewByCategoryOutput, error)
	ListBillOverviewByCategoryWithContext(byteplus.Context, *ListBillOverviewByCategoryInput, ...request.Option) (*ListBillOverviewByCategoryOutput, error)
	ListBillOverviewByCategoryRequest(*ListBillOverviewByCategoryInput) (*request.Request, *ListBillOverviewByCategoryOutput)

	ListBillOverviewByProdCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBillOverviewByProdCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBillOverviewByProdCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBillOverviewByProd(*ListBillOverviewByProdInput) (*ListBillOverviewByProdOutput, error)
	ListBillOverviewByProdWithContext(byteplus.Context, *ListBillOverviewByProdInput, ...request.Option) (*ListBillOverviewByProdOutput, error)
	ListBillOverviewByProdRequest(*ListBillOverviewByProdInput) (*request.Request, *ListBillOverviewByProdOutput)

	ListFinancialRelationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListFinancialRelationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListFinancialRelationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListFinancialRelation(*ListFinancialRelationInput) (*ListFinancialRelationOutput, error)
	ListFinancialRelationWithContext(byteplus.Context, *ListFinancialRelationInput, ...request.Option) (*ListFinancialRelationOutput, error)
	ListFinancialRelationRequest(*ListFinancialRelationInput) (*request.Request, *ListFinancialRelationOutput)

	ListInvitationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListInvitationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListInvitationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListInvitation(*ListInvitationInput) (*ListInvitationOutput, error)
	ListInvitationWithContext(byteplus.Context, *ListInvitationInput, ...request.Option) (*ListInvitationOutput, error)
	ListInvitationRequest(*ListInvitationInput) (*request.Request, *ListInvitationOutput)

	ListSplitBillDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListSplitBillDetailCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListSplitBillDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListSplitBillDetail(*ListSplitBillDetailInput) (*ListSplitBillDetailOutput, error)
	ListSplitBillDetailWithContext(byteplus.Context, *ListSplitBillDetailInput, ...request.Option) (*ListSplitBillDetailOutput, error)
	ListSplitBillDetailRequest(*ListSplitBillDetailInput) (*request.Request, *ListSplitBillDetailOutput)

	UnsubscribeInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UnsubscribeInstanceCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UnsubscribeInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UnsubscribeInstance(*UnsubscribeInstanceInput) (*UnsubscribeInstanceOutput, error)
	UnsubscribeInstanceWithContext(byteplus.Context, *UnsubscribeInstanceInput, ...request.Option) (*UnsubscribeInstanceOutput, error)
	UnsubscribeInstanceRequest(*UnsubscribeInstanceInput) (*request.Request, *UnsubscribeInstanceOutput)
}

var _ BILLINGAPI = (*BILLING)(nil)
