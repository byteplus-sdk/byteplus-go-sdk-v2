package defaults

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/internal/shareddefaults"
)

// SharedCredentialsFilename returns the SDK's default file path
// for the shared credentials file.
//
// Builds the shared config file path based on the OS's platform.
//
//   - Linux/Unix: $HOME/.byteplus/credentials
//   - Windows: %USERPROFILE%\.byteplus\credentials
func SharedCredentialsFilename() string {
	return shareddefaults.SharedCredentialsFilename()
}

// SharedConfigFilename returns the SDK's default file path for
// the shared config file.
//
// Builds the shared config file path based on the OS's platform.
//
//   - Linux/Unix: $HOME/.byteplus/config
//   - Windows: %USERPROFILE%\.byteplus\config
func SharedConfigFilename() string {
	return shareddefaults.SharedConfigFilename()
}

// SharedEndpointConfigFilename returns the SDK's default file path for
// the endpoint config file.
//
// Builds the shared config file path based on the OS's platform.
//
//   - Linux/Unix: $HOME/.byteplus/endpoint
//   - Windows: %USERPROFILE%\.byteplus\endpoint
func SharedEndpointConfigFilename() string {
	return shareddefaults.SharedEndpointConfigFilename()
}
