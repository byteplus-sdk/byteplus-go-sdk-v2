// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iamiface provides an interface to enable mocking the IAM service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iam

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

// IAMAPI provides an interface to enable mocking the
// iam.IAM service client's API operation,
//
//	// byteplus sdk func uses an SDK service client to make a request to
//	// IAM.
//	func myFunc(svc IAMAPI) bool {
//	    // Make svc.AddUserToGroup request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := iam.New(sess)
//
//	    myFunc(svc)
//	}
type IAMAPI interface {
	AddUserToGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddUserToGroupCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddUserToGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddUserToGroup(*AddUserToGroupInput) (*AddUserToGroupOutput, error)
	AddUserToGroupWithContext(byteplus.Context, *AddUserToGroupInput, ...request.Option) (*AddUserToGroupOutput, error)
	AddUserToGroupRequest(*AddUserToGroupInput) (*request.Request, *AddUserToGroupOutput)

	AttachRolePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachRolePolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachRolePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachRolePolicy(*AttachRolePolicyInput) (*AttachRolePolicyOutput, error)
	AttachRolePolicyWithContext(byteplus.Context, *AttachRolePolicyInput, ...request.Option) (*AttachRolePolicyOutput, error)
	AttachRolePolicyRequest(*AttachRolePolicyInput) (*request.Request, *AttachRolePolicyOutput)

	AttachUserGroupPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachUserGroupPolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachUserGroupPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachUserGroupPolicy(*AttachUserGroupPolicyInput) (*AttachUserGroupPolicyOutput, error)
	AttachUserGroupPolicyWithContext(byteplus.Context, *AttachUserGroupPolicyInput, ...request.Option) (*AttachUserGroupPolicyOutput, error)
	AttachUserGroupPolicyRequest(*AttachUserGroupPolicyInput) (*request.Request, *AttachUserGroupPolicyOutput)

	AttachUserPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachUserPolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachUserPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachUserPolicy(*AttachUserPolicyInput) (*AttachUserPolicyOutput, error)
	AttachUserPolicyWithContext(byteplus.Context, *AttachUserPolicyInput, ...request.Option) (*AttachUserPolicyOutput, error)
	AttachUserPolicyRequest(*AttachUserPolicyInput) (*request.Request, *AttachUserPolicyOutput)

	CreateAccessKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAccessKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAccessKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAccessKey(*CreateAccessKeyInput) (*CreateAccessKeyOutput, error)
	CreateAccessKeyWithContext(byteplus.Context, *CreateAccessKeyInput, ...request.Option) (*CreateAccessKeyOutput, error)
	CreateAccessKeyRequest(*CreateAccessKeyInput) (*request.Request, *CreateAccessKeyOutput)

	CreateGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateGroupCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateGroup(*CreateGroupInput) (*CreateGroupOutput, error)
	CreateGroupWithContext(byteplus.Context, *CreateGroupInput, ...request.Option) (*CreateGroupOutput, error)
	CreateGroupRequest(*CreateGroupInput) (*request.Request, *CreateGroupOutput)

	CreateLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLoginProfileCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLoginProfile(*CreateLoginProfileInput) (*CreateLoginProfileOutput, error)
	CreateLoginProfileWithContext(byteplus.Context, *CreateLoginProfileInput, ...request.Option) (*CreateLoginProfileOutput, error)
	CreateLoginProfileRequest(*CreateLoginProfileInput) (*request.Request, *CreateLoginProfileOutput)

	CreateOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateOAuthProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateOAuthProvider(*CreateOAuthProviderInput) (*CreateOAuthProviderOutput, error)
	CreateOAuthProviderWithContext(byteplus.Context, *CreateOAuthProviderInput, ...request.Option) (*CreateOAuthProviderOutput, error)
	CreateOAuthProviderRequest(*CreateOAuthProviderInput) (*request.Request, *CreateOAuthProviderOutput)

	CreatePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreatePolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreatePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreatePolicy(*CreatePolicyInput) (*CreatePolicyOutput, error)
	CreatePolicyWithContext(byteplus.Context, *CreatePolicyInput, ...request.Option) (*CreatePolicyOutput, error)
	CreatePolicyRequest(*CreatePolicyInput) (*request.Request, *CreatePolicyOutput)

	CreateRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRoleCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRole(*CreateRoleInput) (*CreateRoleOutput, error)
	CreateRoleWithContext(byteplus.Context, *CreateRoleInput, ...request.Option) (*CreateRoleOutput, error)
	CreateRoleRequest(*CreateRoleInput) (*request.Request, *CreateRoleOutput)

	CreateServiceLinkedRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateServiceLinkedRoleCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateServiceLinkedRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateServiceLinkedRole(*CreateServiceLinkedRoleInput) (*CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleWithContext(byteplus.Context, *CreateServiceLinkedRoleInput, ...request.Option) (*CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleRequest(*CreateServiceLinkedRoleInput) (*request.Request, *CreateServiceLinkedRoleOutput)

	CreateUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateUserCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateUser(*CreateUserInput) (*CreateUserOutput, error)
	CreateUserWithContext(byteplus.Context, *CreateUserInput, ...request.Option) (*CreateUserOutput, error)
	CreateUserRequest(*CreateUserInput) (*request.Request, *CreateUserOutput)

	DeleteAccessKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAccessKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAccessKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAccessKey(*DeleteAccessKeyInput) (*DeleteAccessKeyOutput, error)
	DeleteAccessKeyWithContext(byteplus.Context, *DeleteAccessKeyInput, ...request.Option) (*DeleteAccessKeyOutput, error)
	DeleteAccessKeyRequest(*DeleteAccessKeyInput) (*request.Request, *DeleteAccessKeyOutput)

	DeleteGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteGroupCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteGroup(*DeleteGroupInput) (*DeleteGroupOutput, error)
	DeleteGroupWithContext(byteplus.Context, *DeleteGroupInput, ...request.Option) (*DeleteGroupOutput, error)
	DeleteGroupRequest(*DeleteGroupInput) (*request.Request, *DeleteGroupOutput)

	DeleteLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLoginProfileCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLoginProfile(*DeleteLoginProfileInput) (*DeleteLoginProfileOutput, error)
	DeleteLoginProfileWithContext(byteplus.Context, *DeleteLoginProfileInput, ...request.Option) (*DeleteLoginProfileOutput, error)
	DeleteLoginProfileRequest(*DeleteLoginProfileInput) (*request.Request, *DeleteLoginProfileOutput)

	DeleteOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteOAuthProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteOAuthProvider(*DeleteOAuthProviderInput) (*DeleteOAuthProviderOutput, error)
	DeleteOAuthProviderWithContext(byteplus.Context, *DeleteOAuthProviderInput, ...request.Option) (*DeleteOAuthProviderOutput, error)
	DeleteOAuthProviderRequest(*DeleteOAuthProviderInput) (*request.Request, *DeleteOAuthProviderOutput)

	DeletePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeletePolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeletePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeletePolicy(*DeletePolicyInput) (*DeletePolicyOutput, error)
	DeletePolicyWithContext(byteplus.Context, *DeletePolicyInput, ...request.Option) (*DeletePolicyOutput, error)
	DeletePolicyRequest(*DeletePolicyInput) (*request.Request, *DeletePolicyOutput)

	DeleteRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRoleCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRole(*DeleteRoleInput) (*DeleteRoleOutput, error)
	DeleteRoleWithContext(byteplus.Context, *DeleteRoleInput, ...request.Option) (*DeleteRoleOutput, error)
	DeleteRoleRequest(*DeleteRoleInput) (*request.Request, *DeleteRoleOutput)

	DeleteSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSAMLProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSAMLProvider(*DeleteSAMLProviderInput) (*DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderWithContext(byteplus.Context, *DeleteSAMLProviderInput, ...request.Option) (*DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderRequest(*DeleteSAMLProviderInput) (*request.Request, *DeleteSAMLProviderOutput)

	DeleteServiceLinkedRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteServiceLinkedRoleCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteServiceLinkedRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteServiceLinkedRole(*DeleteServiceLinkedRoleInput) (*DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleWithContext(byteplus.Context, *DeleteServiceLinkedRoleInput, ...request.Option) (*DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleRequest(*DeleteServiceLinkedRoleInput) (*request.Request, *DeleteServiceLinkedRoleOutput)

	DeleteUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteUserCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteUser(*DeleteUserInput) (*DeleteUserOutput, error)
	DeleteUserWithContext(byteplus.Context, *DeleteUserInput, ...request.Option) (*DeleteUserOutput, error)
	DeleteUserRequest(*DeleteUserInput) (*request.Request, *DeleteUserOutput)

	DetachRolePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachRolePolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachRolePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachRolePolicy(*DetachRolePolicyInput) (*DetachRolePolicyOutput, error)
	DetachRolePolicyWithContext(byteplus.Context, *DetachRolePolicyInput, ...request.Option) (*DetachRolePolicyOutput, error)
	DetachRolePolicyRequest(*DetachRolePolicyInput) (*request.Request, *DetachRolePolicyOutput)

	DetachUserGroupPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachUserGroupPolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachUserGroupPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachUserGroupPolicy(*DetachUserGroupPolicyInput) (*DetachUserGroupPolicyOutput, error)
	DetachUserGroupPolicyWithContext(byteplus.Context, *DetachUserGroupPolicyInput, ...request.Option) (*DetachUserGroupPolicyOutput, error)
	DetachUserGroupPolicyRequest(*DetachUserGroupPolicyInput) (*request.Request, *DetachUserGroupPolicyOutput)

	DetachUserPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachUserPolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachUserPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachUserPolicy(*DetachUserPolicyInput) (*DetachUserPolicyOutput, error)
	DetachUserPolicyWithContext(byteplus.Context, *DetachUserPolicyInput, ...request.Option) (*DetachUserPolicyOutput, error)
	DetachUserPolicyRequest(*DetachUserPolicyInput) (*request.Request, *DetachUserPolicyOutput)

	GetAccessKeyLastUsedCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetAccessKeyLastUsedCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetAccessKeyLastUsedCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetAccessKeyLastUsed(*GetAccessKeyLastUsedInput) (*GetAccessKeyLastUsedOutput, error)
	GetAccessKeyLastUsedWithContext(byteplus.Context, *GetAccessKeyLastUsedInput, ...request.Option) (*GetAccessKeyLastUsedOutput, error)
	GetAccessKeyLastUsedRequest(*GetAccessKeyLastUsedInput) (*request.Request, *GetAccessKeyLastUsedOutput)

	GetGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetGroupCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetGroup(*GetGroupInput) (*GetGroupOutput, error)
	GetGroupWithContext(byteplus.Context, *GetGroupInput, ...request.Option) (*GetGroupOutput, error)
	GetGroupRequest(*GetGroupInput) (*request.Request, *GetGroupOutput)

	GetLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetLoginProfileCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetLoginProfile(*GetLoginProfileInput) (*GetLoginProfileOutput, error)
	GetLoginProfileWithContext(byteplus.Context, *GetLoginProfileInput, ...request.Option) (*GetLoginProfileOutput, error)
	GetLoginProfileRequest(*GetLoginProfileInput) (*request.Request, *GetLoginProfileOutput)

	GetOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetOAuthProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetOAuthProvider(*GetOAuthProviderInput) (*GetOAuthProviderOutput, error)
	GetOAuthProviderWithContext(byteplus.Context, *GetOAuthProviderInput, ...request.Option) (*GetOAuthProviderOutput, error)
	GetOAuthProviderRequest(*GetOAuthProviderInput) (*request.Request, *GetOAuthProviderOutput)

	GetPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetPolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetPolicy(*GetPolicyInput) (*GetPolicyOutput, error)
	GetPolicyWithContext(byteplus.Context, *GetPolicyInput, ...request.Option) (*GetPolicyOutput, error)
	GetPolicyRequest(*GetPolicyInput) (*request.Request, *GetPolicyOutput)

	GetRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetRoleCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetRole(*GetRoleInput) (*GetRoleOutput, error)
	GetRoleWithContext(byteplus.Context, *GetRoleInput, ...request.Option) (*GetRoleOutput, error)
	GetRoleRequest(*GetRoleInput) (*request.Request, *GetRoleOutput)

	GetSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetSAMLProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetSAMLProvider(*GetSAMLProviderInput) (*GetSAMLProviderOutput, error)
	GetSAMLProviderWithContext(byteplus.Context, *GetSAMLProviderInput, ...request.Option) (*GetSAMLProviderOutput, error)
	GetSAMLProviderRequest(*GetSAMLProviderInput) (*request.Request, *GetSAMLProviderOutput)

	GetSecurityConfigCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetSecurityConfigCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetSecurityConfigCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetSecurityConfig(*GetSecurityConfigInput) (*GetSecurityConfigOutput, error)
	GetSecurityConfigWithContext(byteplus.Context, *GetSecurityConfigInput, ...request.Option) (*GetSecurityConfigOutput, error)
	GetSecurityConfigRequest(*GetSecurityConfigInput) (*request.Request, *GetSecurityConfigOutput)

	GetUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetUserCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetUser(*GetUserInput) (*GetUserOutput, error)
	GetUserWithContext(byteplus.Context, *GetUserInput, ...request.Option) (*GetUserOutput, error)
	GetUserRequest(*GetUserInput) (*request.Request, *GetUserOutput)

	ListAccessKeysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAccessKeysCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAccessKeysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAccessKeys(*ListAccessKeysInput) (*ListAccessKeysOutput, error)
	ListAccessKeysWithContext(byteplus.Context, *ListAccessKeysInput, ...request.Option) (*ListAccessKeysOutput, error)
	ListAccessKeysRequest(*ListAccessKeysInput) (*request.Request, *ListAccessKeysOutput)

	ListAttachedRolePoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAttachedRolePoliciesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAttachedRolePoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAttachedRolePolicies(*ListAttachedRolePoliciesInput) (*ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesWithContext(byteplus.Context, *ListAttachedRolePoliciesInput, ...request.Option) (*ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesRequest(*ListAttachedRolePoliciesInput) (*request.Request, *ListAttachedRolePoliciesOutput)

	ListAttachedUserGroupPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAttachedUserGroupPoliciesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAttachedUserGroupPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAttachedUserGroupPolicies(*ListAttachedUserGroupPoliciesInput) (*ListAttachedUserGroupPoliciesOutput, error)
	ListAttachedUserGroupPoliciesWithContext(byteplus.Context, *ListAttachedUserGroupPoliciesInput, ...request.Option) (*ListAttachedUserGroupPoliciesOutput, error)
	ListAttachedUserGroupPoliciesRequest(*ListAttachedUserGroupPoliciesInput) (*request.Request, *ListAttachedUserGroupPoliciesOutput)

	ListAttachedUserPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAttachedUserPoliciesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAttachedUserPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAttachedUserPolicies(*ListAttachedUserPoliciesInput) (*ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesWithContext(byteplus.Context, *ListAttachedUserPoliciesInput, ...request.Option) (*ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesRequest(*ListAttachedUserPoliciesInput) (*request.Request, *ListAttachedUserPoliciesOutput)

	ListEntitiesForPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListEntitiesForPolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListEntitiesForPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListEntitiesForPolicy(*ListEntitiesForPolicyInput) (*ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyWithContext(byteplus.Context, *ListEntitiesForPolicyInput, ...request.Option) (*ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyRequest(*ListEntitiesForPolicyInput) (*request.Request, *ListEntitiesForPolicyOutput)

	ListGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListGroupsCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListGroups(*ListGroupsInput) (*ListGroupsOutput, error)
	ListGroupsWithContext(byteplus.Context, *ListGroupsInput, ...request.Option) (*ListGroupsOutput, error)
	ListGroupsRequest(*ListGroupsInput) (*request.Request, *ListGroupsOutput)

	ListGroupsForUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListGroupsForUserCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListGroupsForUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListGroupsForUser(*ListGroupsForUserInput) (*ListGroupsForUserOutput, error)
	ListGroupsForUserWithContext(byteplus.Context, *ListGroupsForUserInput, ...request.Option) (*ListGroupsForUserOutput, error)
	ListGroupsForUserRequest(*ListGroupsForUserInput) (*request.Request, *ListGroupsForUserOutput)

	ListIdentityProvidersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListIdentityProvidersCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListIdentityProvidersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListIdentityProviders(*ListIdentityProvidersInput) (*ListIdentityProvidersOutput, error)
	ListIdentityProvidersWithContext(byteplus.Context, *ListIdentityProvidersInput, ...request.Option) (*ListIdentityProvidersOutput, error)
	ListIdentityProvidersRequest(*ListIdentityProvidersInput) (*request.Request, *ListIdentityProvidersOutput)

	ListPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListPoliciesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListPolicies(*ListPoliciesInput) (*ListPoliciesOutput, error)
	ListPoliciesWithContext(byteplus.Context, *ListPoliciesInput, ...request.Option) (*ListPoliciesOutput, error)
	ListPoliciesRequest(*ListPoliciesInput) (*request.Request, *ListPoliciesOutput)

	ListRolesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListRolesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListRolesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListRoles(*ListRolesInput) (*ListRolesOutput, error)
	ListRolesWithContext(byteplus.Context, *ListRolesInput, ...request.Option) (*ListRolesOutput, error)
	ListRolesRequest(*ListRolesInput) (*request.Request, *ListRolesOutput)

	ListSAMLProvidersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListSAMLProvidersCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListSAMLProvidersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListSAMLProviders(*ListSAMLProvidersInput) (*ListSAMLProvidersOutput, error)
	ListSAMLProvidersWithContext(byteplus.Context, *ListSAMLProvidersInput, ...request.Option) (*ListSAMLProvidersOutput, error)
	ListSAMLProvidersRequest(*ListSAMLProvidersInput) (*request.Request, *ListSAMLProvidersOutput)

	ListTagsForResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListTagsForResourcesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListTagsForResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListTagsForResources(*ListTagsForResourcesInput) (*ListTagsForResourcesOutput, error)
	ListTagsForResourcesWithContext(byteplus.Context, *ListTagsForResourcesInput, ...request.Option) (*ListTagsForResourcesOutput, error)
	ListTagsForResourcesRequest(*ListTagsForResourcesInput) (*request.Request, *ListTagsForResourcesOutput)

	ListUsersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListUsersCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListUsersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListUsers(*ListUsersInput) (*ListUsersOutput, error)
	ListUsersWithContext(byteplus.Context, *ListUsersInput, ...request.Option) (*ListUsersOutput, error)
	ListUsersRequest(*ListUsersInput) (*request.Request, *ListUsersOutput)

	ListUsersForGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListUsersForGroupCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListUsersForGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListUsersForGroup(*ListUsersForGroupInput) (*ListUsersForGroupOutput, error)
	ListUsersForGroupWithContext(byteplus.Context, *ListUsersForGroupInput, ...request.Option) (*ListUsersForGroupOutput, error)
	ListUsersForGroupRequest(*ListUsersForGroupInput) (*request.Request, *ListUsersForGroupOutput)

	RemoveUserFromGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveUserFromGroupCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveUserFromGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveUserFromGroup(*RemoveUserFromGroupInput) (*RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupWithContext(byteplus.Context, *RemoveUserFromGroupInput, ...request.Option) (*RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupRequest(*RemoveUserFromGroupInput) (*request.Request, *RemoveUserFromGroupOutput)

	SetSecurityConfigCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetSecurityConfigCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetSecurityConfigCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetSecurityConfig(*SetSecurityConfigInput) (*SetSecurityConfigOutput, error)
	SetSecurityConfigWithContext(byteplus.Context, *SetSecurityConfigInput, ...request.Option) (*SetSecurityConfigOutput, error)
	SetSecurityConfigRequest(*SetSecurityConfigInput) (*request.Request, *SetSecurityConfigOutput)

	TagResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	TagResourcesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	TagResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	TagResources(*TagResourcesInput) (*TagResourcesOutput, error)
	TagResourcesWithContext(byteplus.Context, *TagResourcesInput, ...request.Option) (*TagResourcesOutput, error)
	TagResourcesRequest(*TagResourcesInput) (*request.Request, *TagResourcesOutput)

	UntagResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UntagResourcesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UntagResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UntagResources(*UntagResourcesInput) (*UntagResourcesOutput, error)
	UntagResourcesWithContext(byteplus.Context, *UntagResourcesInput, ...request.Option) (*UntagResourcesOutput, error)
	UntagResourcesRequest(*UntagResourcesInput) (*request.Request, *UntagResourcesOutput)

	UpdateAccessKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateAccessKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateAccessKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateAccessKey(*UpdateAccessKeyInput) (*UpdateAccessKeyOutput, error)
	UpdateAccessKeyWithContext(byteplus.Context, *UpdateAccessKeyInput, ...request.Option) (*UpdateAccessKeyOutput, error)
	UpdateAccessKeyRequest(*UpdateAccessKeyInput) (*request.Request, *UpdateAccessKeyOutput)

	UpdateGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateGroupCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateGroup(*UpdateGroupInput) (*UpdateGroupOutput, error)
	UpdateGroupWithContext(byteplus.Context, *UpdateGroupInput, ...request.Option) (*UpdateGroupOutput, error)
	UpdateGroupRequest(*UpdateGroupInput) (*request.Request, *UpdateGroupOutput)

	UpdateLoginProfileCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateLoginProfileCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateLoginProfileCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateLoginProfile(*UpdateLoginProfileInput) (*UpdateLoginProfileOutput, error)
	UpdateLoginProfileWithContext(byteplus.Context, *UpdateLoginProfileInput, ...request.Option) (*UpdateLoginProfileOutput, error)
	UpdateLoginProfileRequest(*UpdateLoginProfileInput) (*request.Request, *UpdateLoginProfileOutput)

	UpdatePolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdatePolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdatePolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdatePolicy(*UpdatePolicyInput) (*UpdatePolicyOutput, error)
	UpdatePolicyWithContext(byteplus.Context, *UpdatePolicyInput, ...request.Option) (*UpdatePolicyOutput, error)
	UpdatePolicyRequest(*UpdatePolicyInput) (*request.Request, *UpdatePolicyOutput)

	UpdateRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateRoleCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateRole(*UpdateRoleInput) (*UpdateRoleOutput, error)
	UpdateRoleWithContext(byteplus.Context, *UpdateRoleInput, ...request.Option) (*UpdateRoleOutput, error)
	UpdateRoleRequest(*UpdateRoleInput) (*request.Request, *UpdateRoleOutput)

	UpdateUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateUserCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateUser(*UpdateUserInput) (*UpdateUserOutput, error)
	UpdateUserWithContext(byteplus.Context, *UpdateUserInput, ...request.Option) (*UpdateUserOutput, error)
	UpdateUserRequest(*UpdateUserInput) (*request.Request, *UpdateUserOutput)

	AddSAMLProviderCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddSAMLProviderCertificateCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddSAMLProviderCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddSAMLProviderCertificate(*AddSAMLProviderCertificateInput) (*AddSAMLProviderCertificateOutput, error)
	AddSAMLProviderCertificateWithContext(byteplus.Context, *AddSAMLProviderCertificateInput, ...request.Option) (*AddSAMLProviderCertificateOutput, error)
	AddSAMLProviderCertificateRequest(*AddSAMLProviderCertificateInput) (*request.Request, *AddSAMLProviderCertificateOutput)

	CreateSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSAMLProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSAMLProvider(*CreateSAMLProviderInput) (*CreateSAMLProviderOutput, error)
	CreateSAMLProviderWithContext(byteplus.Context, *CreateSAMLProviderInput, ...request.Option) (*CreateSAMLProviderOutput, error)
	CreateSAMLProviderRequest(*CreateSAMLProviderInput) (*request.Request, *CreateSAMLProviderOutput)

	UpdateOAuthProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateOAuthProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateOAuthProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateOAuthProvider(*UpdateOAuthProviderInput) (*UpdateOAuthProviderOutput, error)
	UpdateOAuthProviderWithContext(byteplus.Context, *UpdateOAuthProviderInput, ...request.Option) (*UpdateOAuthProviderOutput, error)
	UpdateOAuthProviderRequest(*UpdateOAuthProviderInput) (*request.Request, *UpdateOAuthProviderOutput)

	UpdateSAMLProviderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSAMLProviderCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSAMLProviderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSAMLProvider(*UpdateSAMLProviderInput) (*UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderWithContext(byteplus.Context, *UpdateSAMLProviderInput, ...request.Option) (*UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderRequest(*UpdateSAMLProviderInput) (*request.Request, *UpdateSAMLProviderOutput)
}

var _ IAMAPI = (*IAM)(nil)
