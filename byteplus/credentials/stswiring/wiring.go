// Package stswiring installs credentials.STSHostResolver so every STS
// entrypoint in the credentials package (AssumeRole, StsProvider, OIDC, SAML)
// shares the same defaultEndpoint decision table used by the main SDK request
// pipeline.
//
// This package is intentionally standalone so that neither
// byteplus/credentials nor byteplus/session needs to know about the other side
// of the wiring:
//   - credentials cannot import byteplusutil (that would form a
//     credentials -> byteplusutil -> byteplus -> credentials import cycle).
//   - session should not carry STS-specific business logic in its own file.
//
// The wiring activates automatically whenever this package is imported. In
// practice byteplus/session imports it via a blank import, so ordinary SDK
// users get the hook for free without touching stswiring directly.
package stswiring

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
)

// The lambda closes over no state and compiles to a static function pointer,
// so this init amounts to a single indirect assignment; it introduces no
// runtime side effects until an STS entrypoint actually resolves a host.
//
// If ResolveDefaultEndpoint ever returns an error, defaultSTSHostFor's
// safety net falls back to the inline decision table. The current sts
// registration in defaultEndpoint makes this branch dead in practice, but
// it protects against future map schema changes.
func init() {
	credentials.STSHostResolver = func(region string, useDualStack *bool) (string, error) {
		return byteplusutil.ResolveDefaultEndpoint("sts", region, useDualStack, nil)
	}
}
