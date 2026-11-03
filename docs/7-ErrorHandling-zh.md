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

当 HTTP 状态码为非 2xx（如 504 等）时，响应体可能不是标准的 TOP 网关 JSON 结构。默认错误处理按以下规则返回 `bytepluserr.RequestFailure`，便于排查问题：

- **响应体无法解码为预期响应结构**（如 HTML、纯文本、空响应、顶层数组，或 `ResponseMetadata` 等字段的类型不兼容）：返回 `RequestFailure`，`Code()` 为 `SerializationError`，`Message()` 中附带原始响应正文，并通过 `StatusCode()`、`RequestID()` 暴露 HTTP 状态码与请求 ID。
- **响应体可解码但缺少有效错误码**（如缺少 `ResponseMetadata`，且没有非空的简易错误码，或元数据中缺少错误码）：返回 `RequestFailure`，`Code()` 为 `UnknownError`，同样保留 HTTP 状态码、请求 ID 与响应正文。
- **响应体为简易错误格式**（缺少 `ResponseMetadata`，但符合 `{"errorcode":"错误码","message":"错误描述"}` 且 `errorcode` 非空）：返回 `RequestFailure`，`Code()` 保留真实的 `errorcode`，`Message()` 附带错误描述与原始正文，同时保留 HTTP 状态码与请求 ID。
- **响应体为标准 TOP JSON 错误**：行为保持不变，返回 `RequestFailure`，可解析出服务端的 `Code()`、`Message()`、`RequestID()`。

上述非标准响应的请求 ID 依次取自响应体中非空的 `ResponseMetadata.RequestId`、请求对象中已有的 `r.RequestID`、以及 `X-Top-Requestid` / `X-Top-Request-Id` / `X-Tt-Logid` 响应头；所有来源均无值时，`RequestID()` 返回空字符串。

此外，有输出数据的结构体和 `*map[string]interface{}` 解码路径收到 HTTP 200、正文为 JSON `null` 时，均返回 `RequestFailure`，`Code()` 为 `UnknownError`，`StatusCode()` 为 200，并保留请求 ID 与原始正文。普通解码和 `ForceJsonNumberDecode` 均适用。结构体输出缺少 `ResponseMetadata` 时也会返回 `UnknownError`，这是既有保护行为。HTTP 200 下其它异常正文（如 HTML）不属于上述非 2xx 错误诊断范围，仍可能返回原生 JSON 解码错误。

map 解码拒绝 JSON `null` 时会保留原输出指针。若自定义重试器重试 `UnknownError`，后续响应会继续解析到原输出对象，`After` 拦截器也能获取该对象；重试耗尽仍返回错误。默认重试器不会仅因 `UnknownError` 而重试 HTTP 200 响应。

兼容性说明：非标准错误响应中的 JSON 解码错误封装为 `RequestFailure`，此前依赖原生错误类型的调用方可通过 `OrigErr()` 获取底层错误。默认重试行为受错误分类影响：例如 `MaxRetries=3` 时，HTTP 400、正文为 HTML 的请求仅尝试 1 次；HTTP 500、502、503、504 仍按既有状态码规则重试。`SimpleError` 仅改变 `Error()` 的显示形式，不影响 `StatusCode()`、`RequestID()` 和 `Message()`。自定义错误处理或重试器以其实现为准。

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
