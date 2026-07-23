package credentials

import (
	"fmt"
	"strings"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/sts"
	"github.com/google/uuid"
)

const defaultSTSRegion = "ap-southeast-1"

// STSHostResolver, if non-nil, overrides the built-in inline STS host
// resolution used by defaultSTSHostFor. The byteplus/credentials/stswiring
// package installs a resolver at init time that delegates to
// byteplusutil.ResolveDefaultEndpoint("sts", ...) so every STS entrypoint
// shares the same defaultEndpoint decision table as the main request
// pipeline.
//
// The indirection exists because credentials cannot import byteplusutil
// directly (that would form a credentials -> byteplusutil -> byteplus ->
// credentials import cycle). Callers that do not transitively import
// stswiring (rare in practice — byteplus/session blank-imports it) fall back
// to the inline decision table below.
//
// The variable is written exactly once, from stswiring.init(), before any
// user code runs. Tests may swap it under t.Cleanup, but application code
// must treat it as read-only after init.
var STSHostResolver func(region string, useDualStack *bool) (string, error)

// cnNonMainlandRegionSet mirrors the Go China detection used by
// byteplusutil.ResolveDefaultEndpoint. Kept local as a fallback so that
// callers that do not transitively import stswiring (which installs the
// STSHostResolver hook) still get a correct .cn suffix decision.
var cnNonMainlandRegionSet = map[string]struct{}{
	"cn-hongkong": {},
}

// defaultSTSHostFor returns the STS host for the given region.
//
// If STSHostResolver is installed (byteplus/credentials/stswiring does this
// at init time, transitively pulled in by byteplus/session), it delegates so
// STS shares the exact same defaultEndpoint decision table as the main
// request pipeline. Otherwise it falls back to the inline decision table
// below, which follows the standard rules for a Regional service with
// GoChinaEnabled=true. Empty region falls back to defaultSTSRegion.
func defaultSTSHostFor(region string) string {
	normalized := strings.ToLower(strings.TrimSpace(region))
	if normalized == "" {
		normalized = defaultSTSRegion
	}
	if STSHostResolver != nil {
		if host, err := STSHostResolver(normalized, nil); err == nil && host != "" {
			return host
		}
	}
	suffix := "byteplusapi.com"
	if strings.HasPrefix(normalized, "cn-") {
		if _, nonMainland := cnNonMainlandRegionSet[normalized]; !nonMainland {
			suffix = "byteplusapi.com.cn"
		}
	}
	return "sts." + normalized + "." + suffix
}

type StsAssumeRoleProvider struct {
	AccessKey       string
	SecurityKey     string
	SessionToken    string
	RoleName        string
	AccountId       string
	Host            string
	Region          string
	Schema          string
	Timeout         time.Duration
	DurationSeconds int
	Policy          string
	// MaxRetries controls AssumeRole retry attempts.
	// Zero or negative values fall back to DefaultRetryerMaxNumRetries.
	MaxRetries    int
	RetryInterval time.Duration
}

type StsAssumeRoleTime struct {
	CurrentTime string
	ExpiredTime string
}

// StsAssumeRoleOptions contains optional configuration for STS AssumeRole.
type StsAssumeRoleOptions struct {
	SessionToken    string
	Host            string
	Region          string
	Schema          string
	Timeout         time.Duration
	DurationSeconds int
	Policy          string
	// MaxRetries controls AssumeRole retry attempts.
	// Zero or negative values fall back to DefaultRetryerMaxNumRetries.
	MaxRetries    int
	RetryInterval time.Duration
}

// StsAssumeRoleWithOptions creates a one-shot AssumeRole call with required
// parameters and optional functional options.
func StsAssumeRoleWithOptions(accessKey, securityKey, roleName, accountId string, optFns ...func(*StsAssumeRoleOptions)) (*Credentials, *StsAssumeRoleTime, error) {
	opts := StsAssumeRoleOptions{
		DurationSeconds: 3600,
	}
	for _, fn := range optFns {
		fn(&opts)
	}
	p := &StsAssumeRoleProvider{
		AccessKey:       accessKey,
		SecurityKey:     securityKey,
		SessionToken:    opts.SessionToken,
		RoleName:        roleName,
		AccountId:       accountId,
		Host:            opts.Host,
		Region:          opts.Region,
		Schema:          opts.Schema,
		Timeout:         opts.Timeout,
		DurationSeconds: opts.DurationSeconds,
		Policy:          opts.Policy,
		MaxRetries:      opts.MaxRetries,
		RetryInterval:   opts.RetryInterval,
	}
	return stsAssumeRoleInternal(p)
}

func StsAssumeRole(p *StsAssumeRoleProvider) (*Credentials, *StsAssumeRoleTime, error) {
	return stsAssumeRoleInternal(p)
}

// stsAssumeRoleInternal is the shared implementation for StsAssumeRole and
// StsAssumeRoleWithOptions.
func stsAssumeRoleInternal(p *StsAssumeRoleProvider) (*Credentials, *StsAssumeRoleTime, error) {
	ins := sts.NewInstance()
	region := p.Region
	if region == "" {
		region = defaultSTSRegion
	}
	ins.Client.ServiceInfo.Credentials.Region = region
	if p.Host != "" {
		ins.SetHost(p.Host)
	} else {
		ins.SetHost(defaultSTSHostFor(region))
	}
	if p.Schema != "" {
		ins.SetSchema(p.Schema)
	}
	if p.Timeout > 0 {
		ins.Client.SetTimeout(p.Timeout)
	}

	ins.Client.SetAccessKey(p.AccessKey)
	ins.Client.SetSecretKey(p.SecurityKey)
	if p.SessionToken != "" {
		ins.Client.SetSessionToken(p.SessionToken)
	}
	input := &sts.AssumeRoleRequest{
		DurationSeconds: p.DurationSeconds,
		RoleTrn:         fmt.Sprintf("trn:iam::%s:role/%s", p.AccountId, p.RoleName),
		RoleSessionName: uuid.New().String(),
		Policy:          p.Policy,
	}
	var maxRetriesPtr *int
	if p.MaxRetries > 0 {
		maxRetriesPtr = &p.MaxRetries
	}
	output, statusCode, err := assumeRoleWithRetry(ins, input, maxRetriesPtr, p.RetryInterval)
	var reqId string
	if output != nil {
		reqId = output.ResponseMetadata.RequestId
	}
	if err != nil {
		return nil, nil, fmt.Errorf("AssumeRole error,httpcode is %v and reqId is %s error is %s", statusCode, reqId, err.Error())
	}
	if output != nil && output.ResponseMetadata.Error != nil {
		return nil, nil, fmt.Errorf("AssumeRole service error, reqId: %s, code: %s, message: %s",
			reqId, output.ResponseMetadata.Error.Code, output.ResponseMetadata.Error.Message)
	}
	if statusCode >= 300 || statusCode < 200 {
		return nil, nil, fmt.Errorf("AssumeRole error,httpcode is %v and reqId is %s", statusCode, reqId)
	}
	if output == nil || output.Result == nil || output.Result.Credentials == nil {
		return nil, nil, fmt.Errorf("AssumeRole response error,httpcode is %v and reqId is %s", statusCode, reqId)
	}
	return NewCredentials(&StaticProvider{Value: Value{
			AccessKeyID:     output.Result.Credentials.AccessKeyId,
			SecretAccessKey: output.Result.Credentials.SecretAccessKey,
			SessionToken:    output.Result.Credentials.SessionToken,
		}}), &StsAssumeRoleTime{
			CurrentTime: output.Result.Credentials.CurrentTime,
			ExpiredTime: output.Result.Credentials.ExpiredTime,
		}, nil
}

// assumeRoleWithRetry invokes sts.AssumeRole with retries
func assumeRoleWithRetry(ins *sts.STS, input *sts.AssumeRoleRequest, maxRetries *int, retryInterval time.Duration) (*sts.AssumeRoleResp, int, error) {
	resolvedMaxRetries := resolveCredentialMaxRetries(maxRetries)
	if retryInterval <= 0 {
		retryInterval = DefaultRetryerRetryDelay
	}
	var (
		output     *sts.AssumeRoleResp
		statusCode int
		err        error
	)
	for attempt := 0; attempt <= resolvedMaxRetries; attempt++ {
		output, statusCode, err = ins.AssumeRole(input)
		failed := err != nil ||
			statusCode < 200 || statusCode >= 300 ||
			output == nil ||
			output.ResponseMetadata.Error != nil
		if !failed {
			return output, statusCode, err
		}
		// Determine STS error code for retryability check.
		var errCode string
		if output != nil && output.ResponseMetadata.Error != nil {
			errCode = output.ResponseMetadata.Error.Code
		}
		httpStatusCode := statusCode
		if err != nil && statusCode == 0 {
			httpStatusCode = 0
		}
		if !isRetryableSTSError(httpStatusCode, errCode) {
			return output, statusCode, err
		}
		if attempt < resolvedMaxRetries {
			time.Sleep(retryInterval)
		}
	}
	return output, statusCode, err
}
