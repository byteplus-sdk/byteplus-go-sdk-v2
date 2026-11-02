[← Retry](6-Retry.md) | Error Handling[(中文)](7-ErrorHandling-zh.md) | [Debugging →](8-Debugging.md)

---

## Error Handling

When calling APIs, different types of errors may be returned. You can adopt targeted handling strategies based on the specific error type and error code. For example, network errors can be retried, while business logic errors should be addressed by adjusting parameters or fixing business logic, thereby improving system robustness and user experience.

Error classification:

| Error Type | Description | Returned Error Type | Common Properties | Private Properties |
|---|---|---|---|---|
| 1. Client error | Request did not reach the server; parameter validation failed | bytepluserr.Error or native error | Code(): error code; <br>Message(): error description; <br>Error(): detailed error info; <br>OrigErr(): original error | None |
| 2. Server error | Request successfully reached the server; business logic error returned | bytepluserr.RequestFailure | Same as above | RequestID() to obtain the request ID for server-side troubleshooting |
| 3. Network/timeout error | DNS resolution error or request timeout | bytepluserr.Error | Same as above | None |
| 4. Other errors | Other errors not covered by the above categories | bytepluserr.Error or native error | Same as above | None |

### Error diagnostics for non-standard responses

When the HTTP status code is non-2xx (e.g. 504), the response body may not be a standard TOP gateway JSON structure. The default error handler returns `bytepluserr.RequestFailure` according to the following rules to make troubleshooting easier:

- **Body cannot be decoded into the expected response structure** (e.g. HTML, plain text, an empty response, a top-level array, or incompatible types for fields such as `ResponseMetadata`): returns `RequestFailure` with `Code()` set to `SerializationError`, the raw response body attached to `Message()`, and the HTTP status code and request ID exposed via `StatusCode()` and `RequestID()`.
- **Body can be decoded but has no valid error code** (e.g. missing `ResponseMetadata` with no non-empty simple error code, or metadata with no error code): returns `RequestFailure` with `Code()` set to `UnknownError`, also preserving the HTTP status code, request ID, and response body.
- **Body uses the simple error format** (missing `ResponseMetadata`, but matching `{"errorcode":"error code","message":"error description"}` with a non-empty `errorcode`): returns `RequestFailure` with `Code()` preserving the actual `errorcode`, `Message()` including the description and raw body, and the HTTP status code and request ID preserved.
- **Body is a standard TOP JSON error**: behavior is unchanged; returns `RequestFailure` from which the server-side `Code()`, `Message()`, and `RequestID()` can be parsed.

For the non-standard responses above, the request ID is resolved in order from a non-empty `ResponseMetadata.RequestId` in the body, an existing `r.RequestID` on the request object, then the `X-Top-Requestid` / `X-Top-Request-Id` / `X-Tt-Logid` response headers. `RequestID()` returns an empty string if none of these sources has a value.

In addition, when an HTTP 200 response has a JSON `null` body, decoding into a struct or `*map[string]interface{}` with output data returns `RequestFailure` with `Code()` set to `UnknownError` and `StatusCode()` set to 200, preserving the request ID and raw body. This applies to both normal decoding and `ForceJsonNumberDecode`. Struct decoding also returns `UnknownError` for a body without `ResponseMetadata`, preserving the existing protection. Other malformed HTTP 200 bodies (e.g. HTML) are outside the non-2xx diagnostics described above and may still return a native JSON decoding error.

Compatibility: JSON decoding errors for non-standard error responses are wrapped in `RequestFailure`; callers that previously relied on native error types can obtain the underlying error through `OrigErr()`. The error classification affects default retries: for example, with `MaxRetries=3`, an HTTP 400 response with an HTML body is attempted only once. HTTP 500, 502, 503, and 504 responses still retry under the existing status-code rules. `SimpleError` changes only the display of `Error()`, leaving `StatusCode()`, `RequestID()`, and `Message()` unchanged. Custom error handlers and retryers follow their own implementations.

### Code Example

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/ecs"
)

func main() {
	region := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewEnvCredentials())
	sess, err := session.NewSession(config)
	var be bytepluserr.Error
	if err != nil {
		if errors.As(err, &be) {
			fmt.Println("1. Client error (failed to create session)", be.Code(), be.Message(), be.Error())
		} else {
			fmt.Println("4. Other error", err.Error())
		}
		panic(err)
	}
	svc := ecs.New(sess)

	tags := make([]*ecs.TagForCreateKeyPairInput, 0, 2)
	tags = append(tags, &ecs.TagForCreateKeyPairInput{Key: byteplus.String("testTag")})
	createKeyPairInput := &ecs.CreateKeyPairInput{
		KeyPairName: byteplus.String("testKeyPairName"),
		Tags:        tags,
	}

	_, err = svc.CreateKeyPair(createKeyPairInput)
	if err != nil {
		var requestFailure bytepluserr.RequestFailure // Server-side error
		var errInvalidParam request.ErrInvalidParam   // Parameter validation error
		if errors.As(err, &errInvalidParam) {
			fmt.Println("1. Client error (parameter validation error):", errInvalidParam.Code(), errInvalidParam.Field(), errInvalidParam.Error())
		} else if errors.As(err, &requestFailure) {
			fmt.Println("2. Server error:", requestFailure.RequestID(), requestFailure.Code(), requestFailure.StatusCode(), requestFailure.Error())
		} else if errors.As(err, &be) {
			switch be.Code() {
			case "RequestCanceled":
				fmt.Println("3. Network/timeout error: context timeout passed to the API call")
			case "RequestError":
				if be.OrigErr() != nil {
					var netErr net.Error
					var dnsError *net.DNSError
					if errors.As(be.OrigErr(), &dnsError) {
						fmt.Println("3. Network/timeout error: DNS resolution error handling")
					} else if errors.As(be.OrigErr(), &netErr) && netErr.Timeout() {
						var opErr *net.OpError
						if errors.Is(be.OrigErr(), context.DeadlineExceeded) {
							fmt.Println("3. Network/timeout error: http.Client Timeout (ReadTimeout)....", be.Code(), be.Error())
						} else if errors.As(be.OrigErr(), &opErr) && opErr.Op == "dial" {
							fmt.Println("3. Network/timeout error: http.Client Transport.Dialer Timeout (ConnectTimeout)....", be.Code(), be.Error())
						} else {
							fmt.Println("3. Network/timeout error: other timeout handling", be.Code(), be.Message(), be.Error())
						}
					}
				}
			default:
				fmt.Println("4. Other error", be.Code(), be.Message(), be.Error())
			}
		} else {
			fmt.Println("4. Other error", err.Error())
		}
	}
}
```

---

[← Retry](6-Retry.md) | Error Handling[(中文)](7-ErrorHandling-zh.md) | [Debugging →](8-Debugging.md)
