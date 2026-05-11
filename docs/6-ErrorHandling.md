[← Retry](5-Retry.md) | Error Handling[(中文)](6-ErrorHandling-zh.md) | [Debugging →](7-Debugging.md)

---

## Error Handling

When calling APIs, different types of errors may be returned. You can adopt targeted handling strategies based on the specific error type and error code. For example, network errors can be retried, while business logic errors should be addressed by adjusting parameters or fixing business logic, thereby improving system robustness and user experience.

Error classification:


| Error Type | Description | Returned Error Type | Common Properties | Private Properties |
|---|---|---|---|---|
| 1. Client error | Request did not reach the server; parameter validation failed | bytepluserr.Error or native error | Code(): error code;  <br>Message(): error description;  <br>Error(): detailed error info;  <br>OrigErr(): original error | None |
| 2. Server error | Request successfully reached the server; business logic error returned | bytepluserr.RequestFailure | Same as above | RequestID() to obtain the request ID for server-side troubleshooting |
| 3. Network/timeout error | DNS resolution error or request timeout | bytepluserr.Error | Same as above | None |
| 4. Other errors | Other errors not covered by the above categories | bytepluserr.Error or native error | Same as above | None |

### Code Example

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/ecs"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
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
		KeyPairName: byteplus.String(("testKeyPairName")),
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
						var oPError *net.OpError
						if errors.Is(be.OrigErr(), context.DeadlineExceeded) {
							fmt.Println("3. Network/timeout error: http.Client Timeout (ReadTimeout)....", be.Code(), be.Error())
						} else if errors.As(be.OrigErr(), &oPError) && oPError.Op == "dial" {
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

[← Retry](5-Retry.md) | Error Handling[(中文)](6-ErrorHandling-zh.md) | [Debugging →](7-Debugging.md)
