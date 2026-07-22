[← 访问凭据](1-Credentials-zh.md) | Endpoint 配置[(English)](2-Endpoint.md) | [Transport →](3-Transport-zh.md)

---

## EndPoint 配置

> **默认**
>
> 不指定 endpoint 时，走 [自动化 Endpoint 寻址](#自动化-endpoint-寻址)。

### 自定义 Endpoint

用户可以通过在初始化客户端时指定 Endpoint：

```go
func main() {
	region := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：BYTEPLUS_ACCESS_KEY、BYTEPLUS_SECRET_KEY、BYTEPLUS_SESSION_TOKEN
		WithRegion(region).
		// 自定义Endpoint
		WithEndpoint("<example>.<regionId>.byteplusapi.com")
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

显式设置的 Endpoint 优先级最高，会跳过后续所有寻址逻辑（包括自定义 `EndpointResolver`）。

### 自定义 RegionId

```go
func main() {
	regionId := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：BYTEPLUS_ACCESS_KEY、BYTEPLUS_SECRET_KEY、BYTEPLUS_SESSION_TOKEN
		WithRegion(regionId)                              // 自定义regionId
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### 自动化 Endpoint 寻址

> **默认**
>
> 默认支持自动寻址，无需手动指定 Endpoint。

为了简化用户配置，Byteplus 提供了灵活的 Endpoint 自动寻址机制。用户无需手动指定服务地址，SDK 会根据服务名称、区域（Region）、服务是否标记为 Go China 等信息自动拼接出合理的访问地址，并支持用户自定义 DualStack（双栈）。

#### Endpoint 默认寻址

##### 寻址逻辑

1. **服务注册判定**

    每个服务在内置映射中都会登记 `IsGlobal` 和 `GoChinaEnabled` 两个 bool 字段，SDK 按下方"标准寻址规则"构造 Endpoint。

    - 服务未在映射中登记：`ResolveDefaultEndpoint` 返回 `byteplusutil.ErrServiceEndpointInfoMissing`；调用 `session.ClientConfig` 系列方法时会立即抛出错误，可通过 `errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing)` 识别。参见 [错误处理](#错误处理)。

    内置服务映射：[`../byteplus/byteplusutil/url.go`](../byteplus/byteplusutil/url.go) 中的 `defaultEndpoint`。

2. **DualStack 支持（IPv6）**

    SDK 支持双栈网络（IPv4 + IPv6）访问地址，自动启用条件如下：显式传入参数 `useDualStack = true`，或设置环境变量 `BYTEPLUS_ENABLE_DUALSTACK=true`。优先级：`useDualStack` > `BYTEPLUS_ENABLE_DUALSTACK`。

    启用后，域名后缀将从 `byteplusapi.com` 切换为 `byteplus-api.com`。

3. **Go China 后缀**

    当服务在内置映射中标记 `GoChinaEnabled=true`，并且请求 Region 属于中国大陆（`cn-*` 前缀且不属于 `cn-hongkong` 等非大陆港澳台 Region）时，在域名后追加 `.cn` 后缀。

    是否 GoChina 由服务侧决定，不可修改。

4. **根据服务名和区域自动构造 Endpoint 地址**

    - **Global 服务（如 `IAM`、`Billing`）**：使用 `<服务名>.byteplusapi.com`（DualStack 时使用 `byteplus-api.com`；命中 Go China 时追加 `.cn`）。
    - **Regional 服务（如 `ECS`、`RDS`）**：使用 `<服务名>.<区域名>.byteplusapi.com` 作为默认 Endpoint（DualStack / Go China 规则同上）。

##### 寻址决策表

下表列出所有生效组合。左侧列的 "RegionType" 由服务的 `IsGlobal` 决定；"Region 是否 GoChina" 指请求 Region 是否属于中国大陆。

| RegionType | GoChinaEnabled | 请求 Region 是否 Go China | Endpoint | 是否包含 Region |
|---|---|---|---|---|
| Global | true | 是 | `{service}.byteplusapi.com.cn` | 否 |
| Global | true | 否 | `{service}.byteplusapi.com` | 否 |
| Global | false | 任意 | `{service}.byteplusapi.com` | 否 |
| Regional | true | 是 | `{service}.{region}.byteplusapi.com.cn` | 是 |
| Regional | true | 否 | `{service}.{region}.byteplusapi.com` | 是 |
| Regional | false | 任意 | `{service}.{region}.byteplusapi.com` | 是 |

启用 DualStack 时，将上表中的 `byteplusapi.com` 整体替换为 `byteplus-api.com`。

##### 关于 `WithBootstrapRegion` / `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF`

自本次自动寻址重构起，`Config.BootstrapRegion` 字段和 `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF` 环境变量**不再参与**默认寻址链路。它们的类型和方法保留，仅为 API 兼容，新代码不建议使用。

##### 代码示例

```go
func main() {
	regionId := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：BYTEPLUS_ACCESS_KEY、BYTEPLUS_SECRET_KEY、BYTEPLUS_SESSION_TOKEN
		WithRegion(regionId).
		WithUseDualStack(true) // 定义是否启用双栈网络（IPv4 + IPv6）访问地址，默认false；也可以使用环境变量BYTEPLUS_ENABLE_DUALSTACK=true
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

##### 错误处理

如果请求的服务名不在内置 `defaultEndpoint` 映射中，SDK 在第一次构造 client 配置时就会返回 `byteplusutil.ErrServiceEndpointInfoMissing`。可用如下方式识别：

```go
import (
	"errors"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
)

if errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing) {
	// SDK 版本可能不识别该服务，请升级依赖或显式指定 Endpoint。
}
```

遇到该错误时，建议先升级 SDK 版本；若确认 SDK 尚未内置该服务的寻址元数据，可通过 `WithEndpoint(...)` 或自定义 `EndpointResolver` 显式指定。

#### Endpoint 标准寻址

##### 标准寻址规则

| Global 服务 | 双栈 | 格式 |
|---|---|---|
| 是 | 是 | `{Service}.byteplus-api.com` |
| 是 | 否 | `{Service}.byteplusapi.com` |
| 否 | 是 | `{Service}.{region}.byteplus-api.com` |
| 否 | 否 | `{Service}.{region}.byteplusapi.com` |

标准寻址不会追加 Go China 的 `.cn` 后缀，如需 `.cn` 请使用默认寻址（或显式指定 Endpoint）。

服务是否为 Global 由具体服务决定，不可修改。可以参考列表：[`../byteplus/endpoints/standard_resolver.go`](../byteplus/endpoints/standard_resolver.go) 中的 `ServiceInfos`。

##### 代码示例

```go
package main

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
)

func main() {
	regionId := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithEndpointResolver(endpoints.NewStandardEndpointResolver()). // 配置标准寻址
		WithRegion(regionId).                                          // 配置regionId
		WithUseDualStack(true)                                         // 配置是否双栈
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

---

[← 访问凭据](1-Credentials-zh.md) | Endpoint 配置[(English)](2-Endpoint.md) | [Transport →](3-Transport-zh.md)
