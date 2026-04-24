// Package defaults is a collection of helpers to retrieve the SDK's default
// configuration and handlers.
//
// Generally this package shouldn't be used directly, but session.Session
// instead. This package is useful when you need to reset the defaults
// of a session or service client to the SDK defaults before setting
// additional parameters.
package defaults

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/corehandlers"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials/clicreds"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials/endpointcreds"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

// A Defaults provides a collection of default values for SDK clients.
type Defaults struct {
	Config   *byteplus.Config
	Handlers request.Handlers
}

// Get returns the SDK's default values with Config and handlers pre-configured.
func Get() Defaults {
	cfg := Config()
	handlers := Handlers()
	cfg.Credentials = CredChain(cfg, handlers)

	return Defaults{
		Config:   cfg,
		Handlers: handlers,
	}
}

// Config returns the default configuration without credentials.
// To retrieve a config with credentials also included use
// `defaults.Get().Config` instead.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the configuration of an
// existing service client or session.
func Config() *byteplus.Config {
	return byteplus.NewConfig().
		WithCredentials(credentials.AnonymousCredentials).
		WithRegion(os.Getenv("BYTEPLUS_REGION")).
		WithHTTPClient(http.DefaultClient).
		WithMaxRetries(byteplus.UseServiceDefaultRetries).
		WithLogger(byteplus.NewDefaultLogger()).
		WithLogLevel(byteplus.LogOff)
}

// Handlers returns the default request handlers.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the request handlers of an
// existing service client or session.
func Handlers() request.Handlers {
	var handlers request.Handlers

	handlers.Validate.PushBackNamed(corehandlers.ValidateEndpointHandler)
	handlers.Validate.AfterEachFn = request.HandlerListStopOnError
	handlers.Build.PushBackNamed(corehandlers.CustomerRequestHandler)
	handlers.Build.AfterEachFn = request.HandlerListStopOnError
	handlers.Sign.PushBackNamed(corehandlers.BuildContentLengthHandler)
	handlers.Send.PushBackNamed(corehandlers.ValidateReqSigHandler)
	handlers.Send.PushBackNamed(corehandlers.SendHandler)
	handlers.AfterRetry.PushBackNamed(corehandlers.AfterRetryHandler)
	handlers.ValidateResponse.PushBackNamed(corehandlers.ValidateResponseHandler)

	return handlers
}

// CredChain returns the default credential chain.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the credentials of an
// existing service client or session's Config.
//
// Note: this now uses DefaultCredentialProvider instead of ChainProvider.
// Error messages always include details from every provider in the chain
// (the old CredentialsChainVerboseErrors config is no longer consulted).
func CredChain(cfg *byteplus.Config, handlers request.Handlers) *credentials.Credentials {
	return credentials.NewDefaultCredentialProviderFromProviders(
		CredProviders(cfg, handlers),
		true,
	)
}

// CredProviders returns the slice of providers used in
// the default credential chain.
//
// The 4-step default chain:
//  1. EnvProvider (AK/SK/STS from environment variables)
//  2. OIDCCredentialsProvider (from environment variables)
//  3. CliProvider (from ~/.byteplus/config.json)
//  4. EcsRoleProvider (from IMDS)
func CredProviders(cfg *byteplus.Config, handlers request.Handlers) []credentials.Provider {
	return []credentials.Provider{
		&credentials.EnvProvider{},
		credentials.NewOIDCCredentialsProviderFromEnv(),
		clicreds.NewCliProvider("", ""),
		credentials.NewEcsRoleProvider(""),
	}
}

// NewDefaultCredentialProvider creates a default credential chain with the
// given options.
func NewDefaultCredentialProvider(optFns ...func(*credentials.DefaultCredentialProviderOptions)) *credentials.Credentials {
	opts := credentials.DefaultCredentialProviderOptions{}
	for _, fn := range optFns {
		fn(&opts)
	}

	providers := []credentials.Provider{
		&credentials.EnvProvider{},
		credentials.NewOIDCCredentialsProviderFromEnv(),
		clicreds.NewCliProvider("", ""),
		credentials.NewEcsRoleProvider(opts.RoleName),
	}

	return credentials.NewDefaultCredentialProviderFromProviders(
		providers,
		opts.IsReuseEnabled(),
	)
}

const (
	httpProviderAuthorizationEnvVar = "BYTEPLUS_CONTAINER_AUTHORIZATION_TOKEN"
	httpProviderEnvVar              = "BYTEPLUS_CONTAINER_CREDENTIALS_FULL_URI"
)

var lookupHostFn = net.LookupHost

func isLoopbackHost(host string) (bool, error) {
	ip := net.ParseIP(host)
	if ip != nil {
		return ip.IsLoopback(), nil
	}

	// Host is not an ip, perform lookup
	addrs, err := lookupHostFn(host)
	if err != nil {
		return false, err
	}
	for _, addr := range addrs {
		if !net.ParseIP(addr).IsLoopback() {
			return false, nil
		}
	}

	return true, nil
}

func localHTTPCredProvider(cfg byteplus.Config, handlers request.Handlers, u string) credentials.Provider {
	var errMsg string

	parsed, err := url.Parse(u)
	if err != nil {
		errMsg = fmt.Sprintf("invalid URL, %v", err)
	} else {
		host := byteplus.URLHostname(parsed)
		if len(host) == 0 {
			errMsg = "unable to parse host from local HTTP cred provider URL"
		} else if isLoopback, loopbackErr := isLoopbackHost(host); loopbackErr != nil {
			errMsg = fmt.Sprintf("failed to resolve host %q, %v", host, loopbackErr)
		} else if !isLoopback {
			errMsg = fmt.Sprintf("invalid endpoint host, %q, only loopback hosts are allowed.", host)
		}
	}

	if len(errMsg) > 0 {
		if cfg.Logger != nil {
			cfg.Logger.Error("Ignoring, HTTP credential provider", errMsg, err)
		}
		return credentials.ErrorProvider{
			Err:          bytepluserr.New("CredentialsEndpointError", errMsg, err),
			ProviderName: endpointcreds.ProviderName,
		}
	}

	return httpCredProvider(cfg, handlers, u)
}

func httpCredProvider(cfg byteplus.Config, handlers request.Handlers, u string) credentials.Provider {
	return endpointcreds.NewProviderClient(cfg, handlers, u,
		func(p *endpointcreds.Provider) {
			p.ExpiryWindow = 5 * time.Minute
			p.AuthorizationToken = os.Getenv(httpProviderAuthorizationEnvVar)
		},
	)
}
