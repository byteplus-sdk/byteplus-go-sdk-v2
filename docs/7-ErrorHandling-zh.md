[← 重试机制](6-Retry-zh.md) | 异常处理[(English)](7-ErrorHandling.md) | [Debug 机制 →](8-Debugging-zh.md)

---

## 异常处理

在调用接口时，可能会返回不同类型的错误。客户可根据具体的错误类型和错误码，采取有针对性的处理策略。例如，对于网络异常可选择重试，对于业务逻辑错误则应根据错误信息进行参数调整或业务逻辑修正，从而提升系统的健壮性和用户体验。

错误分类：

| 错误类型 | 错误描述 | 返回错误类型 | 公共属性 | 私有属性 |
|---|---|---|---|---|
| 1. 客户端错误 | 请求未到达服务端，对参数作验证 | bytepluserr.Error 或原生 error | Code()：错误码；<br>Message()：错误描述信息；<br>Error()：详细错误信息；<br>OrigErr()：原始错误 | 无 |
| 2. 服务端错误 | 请求成功到达服务器，返回业务逻辑错误 | bytepluserr.RequestFailure | 同上 | 可以通过 RequestID() 获取请求 id，方便服务端问题排查 |
| 3. 网络/超时错误 | DNS 解析错误或请求超时 | bytepluserr.Error | 同上 | 无 |
| 4. 其它错误 | 未包含在前 3 类错误中的其它错误 | bytepluserr.Error 或原生 error | 同上 | 无 |

### 非标准响应的错误诊断

当 HTTP 状态码异常（如 504 等）时，响应体可能不是标准的 TOP 网关 JSON 结构。此类响应统一返回 `bytepluserr.RequestFailure`，便于排查问题：

- **响应体非 JSON 格式**（如网关直接返回 HTML、纯文本或空响应）：返回 `RequestFailure`，`Code()` 为 `SerializationError`，`Message()` 中附带原始响应正文，并通过 `StatusCode()`、`RequestID()` 暴露 HTTP 状态码与请求 ID。
- **响应体为合法 JSON 但不满足 TOP 网关标准**（缺少 `ResponseMetadata` 或缺少错误码）：返回 `RequestFailure`，`Code()` 为 `UnknownError`，同样保留 HTTP 状态码、请求 ID 与响应正文。
- **响应体为标准 TOP JSON 错误**：行为保持不变，返回 `RequestFailure`，可解析出服务端的 `Code()`、`Message()`、`RequestID()`。

其中请求 ID 依次从响应体的 `ResponseMetadata.RequestId`、请求上下文、以及 `X-Top-Requestid` / `X-Top-Request-Id` / `X-Tt-Logid` 响应头中解析，确保非标准响应也能拿到可用的请求 ID。

### 代码示例

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
			fmt.Println("1. 客户端错误（创建 session 失败）", be.Code(), be.Message(), be.Error())
		} else {
			fmt.Println("4. 其它错误", err.Error())
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
		var requestFailure bytepluserr.RequestFailure // 服务端返回的错误
		var errInvalidParam request.ErrInvalidParam   // 参数验证错误
		if errors.As(err, &errInvalidParam) {
			fmt.Println("1. 客户端错误（参数验证错误）：", errInvalidParam.Code(), errInvalidParam.Field(), errInvalidParam.Error())
		} else if errors.As(err, &requestFailure) {
			fmt.Println("2. 服务端错误：", requestFailure.RequestID(), requestFailure.Code(), requestFailure.StatusCode(), requestFailure.Error())
		} else if errors.As(err, &be) {
			switch be.Code() {
			case "RequestCanceled":
				fmt.Println("3. 网络/超时错误：接口调用传入的 context 超时")
			case "RequestError":
				if be.OrigErr() != nil {
					var netErr net.Error
					var dnsError *net.DNSError
					if errors.As(be.OrigErr(), &dnsError) {
						fmt.Println("3. 网络/超时错误：DNS 解析错误处理")
					} else if errors.As(be.OrigErr(), &netErr) && netErr.Timeout() {
						var opErr *net.OpError
						if errors.Is(be.OrigErr(), context.DeadlineExceeded) {
							fmt.Println("3. 网络/超时错误：http.Client Timeout(ReadTimeout)....", be.Code(), be.Error())
						} else if errors.As(be.OrigErr(), &opErr) && opErr.Op == "dial" {
							fmt.Println("3. 网络/超时错误：http.Client Transport.Dialer Timeout(ConnectTimeout)....", be.Code(), be.Error())
						} else {
							fmt.Println("3. 网络/超时错误：其它超时处理", be.Code(), be.Message(), be.Error())
						}
					}
				}
			default:
				fmt.Println("4. 其它错误", be.Code(), be.Message(), be.Error())
			}
		} else {
			fmt.Println("4. 其它错误", err.Error())
		}
	}
}
```

---

[← 重试机制](6-Retry-zh.md) | 异常处理[(English)](7-ErrorHandling.md) | [Debug 机制 →](8-Debugging-zh.md)
