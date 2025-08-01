// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kmsiface provides an interface to enable mocking the KMS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kms

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

// KMSAPI provides an interface to enable mocking the
// kms.KMS service client's API operation,
//
//    // byteplus sdk func uses an SDK service client to make a request to
//    // KMS.
//    func myFunc(svc KMSAPI) bool {
//        // Make svc.ArchiveKey request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kms.New(sess)
//
//        myFunc(svc)
//    }
//
type KMSAPI interface {
	ArchiveKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ArchiveKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ArchiveKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ArchiveKey(*ArchiveKeyInput) (*ArchiveKeyOutput, error)
	ArchiveKeyWithContext(byteplus.Context, *ArchiveKeyInput, ...request.Option) (*ArchiveKeyOutput, error)
	ArchiveKeyRequest(*ArchiveKeyInput) (*request.Request, *ArchiveKeyOutput)

	AsymmetricDecryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricDecryptCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricDecryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricDecrypt(*AsymmetricDecryptInput) (*AsymmetricDecryptOutput, error)
	AsymmetricDecryptWithContext(byteplus.Context, *AsymmetricDecryptInput, ...request.Option) (*AsymmetricDecryptOutput, error)
	AsymmetricDecryptRequest(*AsymmetricDecryptInput) (*request.Request, *AsymmetricDecryptOutput)

	AsymmetricEncryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricEncryptCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricEncryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricEncrypt(*AsymmetricEncryptInput) (*AsymmetricEncryptOutput, error)
	AsymmetricEncryptWithContext(byteplus.Context, *AsymmetricEncryptInput, ...request.Option) (*AsymmetricEncryptOutput, error)
	AsymmetricEncryptRequest(*AsymmetricEncryptInput) (*request.Request, *AsymmetricEncryptOutput)

	AsymmetricSignCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricSignCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricSignCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricSign(*AsymmetricSignInput) (*AsymmetricSignOutput, error)
	AsymmetricSignWithContext(byteplus.Context, *AsymmetricSignInput, ...request.Option) (*AsymmetricSignOutput, error)
	AsymmetricSignRequest(*AsymmetricSignInput) (*request.Request, *AsymmetricSignOutput)

	AsymmetricVerifyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AsymmetricVerifyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AsymmetricVerifyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AsymmetricVerify(*AsymmetricVerifyInput) (*AsymmetricVerifyOutput, error)
	AsymmetricVerifyWithContext(byteplus.Context, *AsymmetricVerifyInput, ...request.Option) (*AsymmetricVerifyOutput, error)
	AsymmetricVerifyRequest(*AsymmetricVerifyInput) (*request.Request, *AsymmetricVerifyOutput)

	BackupSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	BackupSecretCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BackupSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BackupSecret(*BackupSecretInput) (*BackupSecretOutput, error)
	BackupSecretWithContext(byteplus.Context, *BackupSecretInput, ...request.Option) (*BackupSecretOutput, error)
	BackupSecretRequest(*BackupSecretInput) (*request.Request, *BackupSecretOutput)

	BatchGetSecretValueCommon(*map[string]interface{}) (*map[string]interface{}, error)
	BatchGetSecretValueCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BatchGetSecretValueCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BatchGetSecretValue(*BatchGetSecretValueInput) (*BatchGetSecretValueOutput, error)
	BatchGetSecretValueWithContext(byteplus.Context, *BatchGetSecretValueInput, ...request.Option) (*BatchGetSecretValueOutput, error)
	BatchGetSecretValueRequest(*BatchGetSecretValueInput) (*request.Request, *BatchGetSecretValueOutput)

	CancelArchiveKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelArchiveKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelArchiveKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelArchiveKey(*CancelArchiveKeyInput) (*CancelArchiveKeyOutput, error)
	CancelArchiveKeyWithContext(byteplus.Context, *CancelArchiveKeyInput, ...request.Option) (*CancelArchiveKeyOutput, error)
	CancelArchiveKeyRequest(*CancelArchiveKeyInput) (*request.Request, *CancelArchiveKeyOutput)

	CancelKeyDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelKeyDeletionCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelKeyDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelKeyDeletion(*CancelKeyDeletionInput) (*CancelKeyDeletionOutput, error)
	CancelKeyDeletionWithContext(byteplus.Context, *CancelKeyDeletionInput, ...request.Option) (*CancelKeyDeletionOutput, error)
	CancelKeyDeletionRequest(*CancelKeyDeletionInput) (*request.Request, *CancelKeyDeletionOutput)

	CancelSecretDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelSecretDeletionCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelSecretDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelSecretDeletion(*CancelSecretDeletionInput) (*CancelSecretDeletionOutput, error)
	CancelSecretDeletionWithContext(byteplus.Context, *CancelSecretDeletionInput, ...request.Option) (*CancelSecretDeletionOutput, error)
	CancelSecretDeletionRequest(*CancelSecretDeletionInput) (*request.Request, *CancelSecretDeletionOutput)

	CreateKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateKey(*CreateKeyInput) (*CreateKeyOutput, error)
	CreateKeyWithContext(byteplus.Context, *CreateKeyInput, ...request.Option) (*CreateKeyOutput, error)
	CreateKeyRequest(*CreateKeyInput) (*request.Request, *CreateKeyOutput)

	CreateKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKeyringCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateKeyring(*CreateKeyringInput) (*CreateKeyringOutput, error)
	CreateKeyringWithContext(byteplus.Context, *CreateKeyringInput, ...request.Option) (*CreateKeyringOutput, error)
	CreateKeyringRequest(*CreateKeyringInput) (*request.Request, *CreateKeyringOutput)

	CreateSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSecretCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSecret(*CreateSecretInput) (*CreateSecretOutput, error)
	CreateSecretWithContext(byteplus.Context, *CreateSecretInput, ...request.Option) (*CreateSecretOutput, error)
	CreateSecretRequest(*CreateSecretInput) (*request.Request, *CreateSecretOutput)

	DecryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DecryptCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DecryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Decrypt(*DecryptInput) (*DecryptOutput, error)
	DecryptWithContext(byteplus.Context, *DecryptInput, ...request.Option) (*DecryptOutput, error)
	DecryptRequest(*DecryptInput) (*request.Request, *DecryptOutput)

	DeleteKeyMaterialCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteKeyMaterialCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteKeyMaterialCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteKeyMaterial(*DeleteKeyMaterialInput) (*DeleteKeyMaterialOutput, error)
	DeleteKeyMaterialWithContext(byteplus.Context, *DeleteKeyMaterialInput, ...request.Option) (*DeleteKeyMaterialOutput, error)
	DeleteKeyMaterialRequest(*DeleteKeyMaterialInput) (*request.Request, *DeleteKeyMaterialOutput)

	DeleteKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteKeyringCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteKeyring(*DeleteKeyringInput) (*DeleteKeyringOutput, error)
	DeleteKeyringWithContext(byteplus.Context, *DeleteKeyringInput, ...request.Option) (*DeleteKeyringOutput, error)
	DeleteKeyringRequest(*DeleteKeyringInput) (*request.Request, *DeleteKeyringOutput)

	DescribeKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKey(*DescribeKeyInput) (*DescribeKeyOutput, error)
	DescribeKeyWithContext(byteplus.Context, *DescribeKeyInput, ...request.Option) (*DescribeKeyOutput, error)
	DescribeKeyRequest(*DescribeKeyInput) (*request.Request, *DescribeKeyOutput)

	DescribeKeyringsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeyringsCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeyringsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKeyrings(*DescribeKeyringsInput) (*DescribeKeyringsOutput, error)
	DescribeKeyringsWithContext(byteplus.Context, *DescribeKeyringsInput, ...request.Option) (*DescribeKeyringsOutput, error)
	DescribeKeyringsRequest(*DescribeKeyringsInput) (*request.Request, *DescribeKeyringsOutput)

	DescribeKeysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeysCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKeys(*DescribeKeysInput) (*DescribeKeysOutput, error)
	DescribeKeysWithContext(byteplus.Context, *DescribeKeysInput, ...request.Option) (*DescribeKeysOutput, error)
	DescribeKeysRequest(*DescribeKeysInput) (*request.Request, *DescribeKeysOutput)

	DescribeRegionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*DescribeRegionsInput) (*DescribeRegionsOutput, error)
	DescribeRegionsWithContext(byteplus.Context, *DescribeRegionsInput, ...request.Option) (*DescribeRegionsOutput, error)
	DescribeRegionsRequest(*DescribeRegionsInput) (*request.Request, *DescribeRegionsOutput)

	DescribeSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecretCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecret(*DescribeSecretInput) (*DescribeSecretOutput, error)
	DescribeSecretWithContext(byteplus.Context, *DescribeSecretInput, ...request.Option) (*DescribeSecretOutput, error)
	DescribeSecretRequest(*DescribeSecretInput) (*request.Request, *DescribeSecretOutput)

	DescribeSecretVersionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecretVersionsCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecretVersionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecretVersions(*DescribeSecretVersionsInput) (*DescribeSecretVersionsOutput, error)
	DescribeSecretVersionsWithContext(byteplus.Context, *DescribeSecretVersionsInput, ...request.Option) (*DescribeSecretVersionsOutput, error)
	DescribeSecretVersionsRequest(*DescribeSecretVersionsInput) (*request.Request, *DescribeSecretVersionsOutput)

	DescribeSecretsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecretsCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecretsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecrets(*DescribeSecretsInput) (*DescribeSecretsOutput, error)
	DescribeSecretsWithContext(byteplus.Context, *DescribeSecretsInput, ...request.Option) (*DescribeSecretsOutput, error)
	DescribeSecretsRequest(*DescribeSecretsInput) (*request.Request, *DescribeSecretsOutput)

	DisableKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableKey(*DisableKeyInput) (*DisableKeyOutput, error)
	DisableKeyWithContext(byteplus.Context, *DisableKeyInput, ...request.Option) (*DisableKeyOutput, error)
	DisableKeyRequest(*DisableKeyInput) (*request.Request, *DisableKeyOutput)

	DisableKeyRotationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableKeyRotationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableKeyRotationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableKeyRotation(*DisableKeyRotationInput) (*DisableKeyRotationOutput, error)
	DisableKeyRotationWithContext(byteplus.Context, *DisableKeyRotationInput, ...request.Option) (*DisableKeyRotationOutput, error)
	DisableKeyRotationRequest(*DisableKeyRotationInput) (*request.Request, *DisableKeyRotationOutput)

	EnableKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableKey(*EnableKeyInput) (*EnableKeyOutput, error)
	EnableKeyWithContext(byteplus.Context, *EnableKeyInput, ...request.Option) (*EnableKeyOutput, error)
	EnableKeyRequest(*EnableKeyInput) (*request.Request, *EnableKeyOutput)

	EnableKeyRotationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableKeyRotationCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableKeyRotationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableKeyRotation(*EnableKeyRotationInput) (*EnableKeyRotationOutput, error)
	EnableKeyRotationWithContext(byteplus.Context, *EnableKeyRotationInput, ...request.Option) (*EnableKeyRotationOutput, error)
	EnableKeyRotationRequest(*EnableKeyRotationInput) (*request.Request, *EnableKeyRotationOutput)

	EncryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EncryptCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EncryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Encrypt(*EncryptInput) (*EncryptOutput, error)
	EncryptWithContext(byteplus.Context, *EncryptInput, ...request.Option) (*EncryptOutput, error)
	EncryptRequest(*EncryptInput) (*request.Request, *EncryptOutput)

	GenerateDataKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GenerateDataKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GenerateDataKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GenerateDataKey(*GenerateDataKeyInput) (*GenerateDataKeyOutput, error)
	GenerateDataKeyWithContext(byteplus.Context, *GenerateDataKeyInput, ...request.Option) (*GenerateDataKeyOutput, error)
	GenerateDataKeyRequest(*GenerateDataKeyInput) (*request.Request, *GenerateDataKeyOutput)

	GenerateMacCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GenerateMacCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GenerateMacCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GenerateMac(*GenerateMacInput) (*GenerateMacOutput, error)
	GenerateMacWithContext(byteplus.Context, *GenerateMacInput, ...request.Option) (*GenerateMacOutput, error)
	GenerateMacRequest(*GenerateMacInput) (*request.Request, *GenerateMacOutput)

	GetParametersForImportCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetParametersForImportCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetParametersForImportCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetParametersForImport(*GetParametersForImportInput) (*GetParametersForImportOutput, error)
	GetParametersForImportWithContext(byteplus.Context, *GetParametersForImportInput, ...request.Option) (*GetParametersForImportOutput, error)
	GetParametersForImportRequest(*GetParametersForImportInput) (*request.Request, *GetParametersForImportOutput)

	GetPublicKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetPublicKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetPublicKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetPublicKey(*GetPublicKeyInput) (*GetPublicKeyOutput, error)
	GetPublicKeyWithContext(byteplus.Context, *GetPublicKeyInput, ...request.Option) (*GetPublicKeyOutput, error)
	GetPublicKeyRequest(*GetPublicKeyInput) (*request.Request, *GetPublicKeyOutput)

	GetSecretValueCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetSecretValueCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetSecretValueCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetSecretValue(*GetSecretValueInput) (*GetSecretValueOutput, error)
	GetSecretValueWithContext(byteplus.Context, *GetSecretValueInput, ...request.Option) (*GetSecretValueOutput, error)
	GetSecretValueRequest(*GetSecretValueInput) (*request.Request, *GetSecretValueOutput)

	ImportKeyMaterialCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ImportKeyMaterialCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ImportKeyMaterialCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ImportKeyMaterial(*ImportKeyMaterialInput) (*ImportKeyMaterialOutput, error)
	ImportKeyMaterialWithContext(byteplus.Context, *ImportKeyMaterialInput, ...request.Option) (*ImportKeyMaterialOutput, error)
	ImportKeyMaterialRequest(*ImportKeyMaterialInput) (*request.Request, *ImportKeyMaterialOutput)

	ListTagsForResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListTagsForResourcesCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListTagsForResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListTagsForResources(*ListTagsForResourcesInput) (*ListTagsForResourcesOutput, error)
	ListTagsForResourcesWithContext(byteplus.Context, *ListTagsForResourcesInput, ...request.Option) (*ListTagsForResourcesOutput, error)
	ListTagsForResourcesRequest(*ListTagsForResourcesInput) (*request.Request, *ListTagsForResourcesOutput)

	QueryKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	QueryKeyringCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	QueryKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	QueryKeyring(*QueryKeyringInput) (*QueryKeyringOutput, error)
	QueryKeyringWithContext(byteplus.Context, *QueryKeyringInput, ...request.Option) (*QueryKeyringOutput, error)
	QueryKeyringRequest(*QueryKeyringInput) (*request.Request, *QueryKeyringOutput)

	ReEncryptCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReEncryptCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReEncryptCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReEncrypt(*ReEncryptInput) (*ReEncryptOutput, error)
	ReEncryptWithContext(byteplus.Context, *ReEncryptInput, ...request.Option) (*ReEncryptOutput, error)
	ReEncryptRequest(*ReEncryptInput) (*request.Request, *ReEncryptOutput)

	ReplicateKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReplicateKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReplicateKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReplicateKey(*ReplicateKeyInput) (*ReplicateKeyOutput, error)
	ReplicateKeyWithContext(byteplus.Context, *ReplicateKeyInput, ...request.Option) (*ReplicateKeyOutput, error)
	ReplicateKeyRequest(*ReplicateKeyInput) (*request.Request, *ReplicateKeyOutput)

	RestoreSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestoreSecretCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestoreSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestoreSecret(*RestoreSecretInput) (*RestoreSecretOutput, error)
	RestoreSecretWithContext(byteplus.Context, *RestoreSecretInput, ...request.Option) (*RestoreSecretOutput, error)
	RestoreSecretRequest(*RestoreSecretInput) (*request.Request, *RestoreSecretOutput)

	RotateSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RotateSecretCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RotateSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RotateSecret(*RotateSecretInput) (*RotateSecretOutput, error)
	RotateSecretWithContext(byteplus.Context, *RotateSecretInput, ...request.Option) (*RotateSecretOutput, error)
	RotateSecretRequest(*RotateSecretInput) (*request.Request, *RotateSecretOutput)

	ScheduleKeyDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ScheduleKeyDeletionCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ScheduleKeyDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ScheduleKeyDeletion(*ScheduleKeyDeletionInput) (*ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionWithContext(byteplus.Context, *ScheduleKeyDeletionInput, ...request.Option) (*ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionRequest(*ScheduleKeyDeletionInput) (*request.Request, *ScheduleKeyDeletionOutput)

	ScheduleSecretDeletionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ScheduleSecretDeletionCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ScheduleSecretDeletionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ScheduleSecretDeletion(*ScheduleSecretDeletionInput) (*ScheduleSecretDeletionOutput, error)
	ScheduleSecretDeletionWithContext(byteplus.Context, *ScheduleSecretDeletionInput, ...request.Option) (*ScheduleSecretDeletionOutput, error)
	ScheduleSecretDeletionRequest(*ScheduleSecretDeletionInput) (*request.Request, *ScheduleSecretDeletionOutput)

	SetSecretValueCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetSecretValueCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetSecretValueCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetSecretValue(*SetSecretValueInput) (*SetSecretValueOutput, error)
	SetSecretValueWithContext(byteplus.Context, *SetSecretValueInput, ...request.Option) (*SetSecretValueOutput, error)
	SetSecretValueRequest(*SetSecretValueInput) (*request.Request, *SetSecretValueOutput)

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

	UpdateKeyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateKeyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateKeyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateKey(*UpdateKeyInput) (*UpdateKeyOutput, error)
	UpdateKeyWithContext(byteplus.Context, *UpdateKeyInput, ...request.Option) (*UpdateKeyOutput, error)
	UpdateKeyRequest(*UpdateKeyInput) (*request.Request, *UpdateKeyOutput)

	UpdateKeyringCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateKeyringCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateKeyringCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateKeyring(*UpdateKeyringInput) (*UpdateKeyringOutput, error)
	UpdateKeyringWithContext(byteplus.Context, *UpdateKeyringInput, ...request.Option) (*UpdateKeyringOutput, error)
	UpdateKeyringRequest(*UpdateKeyringInput) (*request.Request, *UpdateKeyringOutput)

	UpdatePrimaryRegionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdatePrimaryRegionCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdatePrimaryRegionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdatePrimaryRegion(*UpdatePrimaryRegionInput) (*UpdatePrimaryRegionOutput, error)
	UpdatePrimaryRegionWithContext(byteplus.Context, *UpdatePrimaryRegionInput, ...request.Option) (*UpdatePrimaryRegionOutput, error)
	UpdatePrimaryRegionRequest(*UpdatePrimaryRegionInput) (*request.Request, *UpdatePrimaryRegionOutput)

	UpdateSecretCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSecretCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSecretCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSecret(*UpdateSecretInput) (*UpdateSecretOutput, error)
	UpdateSecretWithContext(byteplus.Context, *UpdateSecretInput, ...request.Option) (*UpdateSecretOutput, error)
	UpdateSecretRequest(*UpdateSecretInput) (*request.Request, *UpdateSecretOutput)

	UpdateSecretRotationPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSecretRotationPolicyCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSecretRotationPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSecretRotationPolicy(*UpdateSecretRotationPolicyInput) (*UpdateSecretRotationPolicyOutput, error)
	UpdateSecretRotationPolicyWithContext(byteplus.Context, *UpdateSecretRotationPolicyInput, ...request.Option) (*UpdateSecretRotationPolicyOutput, error)
	UpdateSecretRotationPolicyRequest(*UpdateSecretRotationPolicyInput) (*request.Request, *UpdateSecretRotationPolicyOutput)

	VerifyMacCommon(*map[string]interface{}) (*map[string]interface{}, error)
	VerifyMacCommonWithContext(byteplus.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	VerifyMacCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	VerifyMac(*VerifyMacInput) (*VerifyMacOutput, error)
	VerifyMacWithContext(byteplus.Context, *VerifyMacInput, ...request.Option) (*VerifyMacOutput, error)
	VerifyMacRequest(*VerifyMacInput) (*request.Request, *VerifyMacOutput)
}

var _ KMSAPI = (*KMS)(nil)
