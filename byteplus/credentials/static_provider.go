package credentials

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
)

// StaticProviderName provides a name of Static provider
const StaticProviderName = "StaticProvider"

var (
	// ErrStaticCredentialsEmpty is emitted when static credentials are empty.
	ErrStaticCredentialsEmpty = bytepluserr.New("EmptyStaticCreds", "static credentials are empty", nil)
)

// A StaticProvider is a set of credentials which are set programmatically,
// and will never expire.
type StaticProvider struct {
	Value
}

// NewStaticCredentials returns a pointer to a new Credentials object
// wrapping a static credentials value provider.
func NewStaticCredentials(id, secret, token string) *Credentials {
	return NewCredentials(&StaticProvider{Value: Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		SessionToken:    token,
	}})
}

// NewStaticCredentialsFromCreds returns a pointer to a new Credentials object
// wrapping the static credentials value provide. Same as NewStaticCredentials
// but takes the creds Value instead of individual fields
func NewStaticCredentialsFromCreds(creds Value) *Credentials {
	return NewCredentials(&StaticProvider{Value: creds})
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s *StaticProvider) Retrieve() (Value, error) {
	if s.AccessKeyID == "" || s.SecretAccessKey == "" {
		return Value{ProviderName: StaticProviderName}, ErrStaticCredentialsEmpty
	}

	if len(s.Value.ProviderName) == 0 {
		s.Value.ProviderName = StaticProviderName
	}
	return s.Value, nil
}

// IsExpired returns if the credentials are expired.
//
// For StaticProvider, the credentials never expired.
func (s *StaticProvider) IsExpired() bool {
	return false
}
