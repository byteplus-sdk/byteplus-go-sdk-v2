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

`WithEndpoint` 配置业务 API 地址，不会被 STS 凭证请求继承。
AssumeRole 依次使用 Provider 显式 `Host`、独立配置的 `STSConfig.Host`、
自身默认 STS 地址。独立 STS Host 会在 session/客户端创建及 `Session.Copy`
后保留；OIDC 和 SAML 使用各自 Provider 的 `Endpoint`。

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

    - 服务未在映射中登记：`ResolveDefaultEndpoint` 返回 `byteplusutil.ErrServiceEndpointInfoMissing`。使用 `session.ClientConfigWithError(...)` 可以在构造阶段直接拿到错误；兼容签名 `session.ClientConfig(...)` 保留原有的无 error 返回，会通过在 Config 上安装 `Validate` handler，将同一错误在首次业务调用时抛出。可通过 `errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing)` 识别。参见 [错误处理](#错误处理)。

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

##### `WithBootstrapRegion` / `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF`（已废弃）

> **⚠️ Deprecated**：`Config.WithBootstrapRegion(map[string]struct{}{...})`、`Config.BootstrapRegion` 字段以及 `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF` 环境变量已被标记为**废弃**，**不再参与**默认寻址链路。相关方法和字段仅为 API 二进制/源码兼容而保留，运行时视为 no-op。请**勿在新代码中使用**，已有代码建议改用 `WithRegion` + `WithUseDualStack` 让 SDK 自动寻址，或用 `WithEndpoint` 显式覆盖。

##### 代码示例

> 注：示例中带删除线的调用 (`WithBootstrapRegion` / `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF`) 为**已废弃**用法，出于兼容性目的保留，请勿在新代码中使用。

<pre><code class="language-go">func main() {
	regionId := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：BYTEPLUS_ACCESS_KEY、BYTEPLUS_SECRET_KEY、BYTEPLUS_SESSION_TOKEN
		WithRegion(regionId).
		WithUseDualStack(true). // 定义是否启用双栈网络（IPv4 + IPv6）访问地址，默认false；也可以使用环境变量BYTEPLUS_ENABLE_DUALSTACK=true
<s>		WithBootstrapRegion(map[string]struct{}{</s>
<s>			"custom_example_region1": {},</s>
<s>			"custom_example_region2": {},</s>
<s>		}) // 已废弃：自定义自动寻址Region列表；也可以使用环境变量BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF</s>
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
</code></pre>

##### 错误处理

如果请求的服务名不在内置 `defaultEndpoint` 映射中，SDK 会返回 `byteplusutil.ErrServiceEndpointInfoMissing`。两条 API 路径的错误暴露时机不同：

- **`session.ClientConfigWithError(serviceName, ...)`**：在**构造阶段**直接返回该错误，适合希望立即感知失败的调用方；
- **`session.ClientConfig(serviceName, ...)`**（兼容签名）：保留无 `error` 返回值，同一错误会通过挂在返回 Config 上的 `Validate` handler，在**首次业务请求**发出时抛出，避免下游 `xxx.New(sess)` 处的 nil 指针 panic。

两个入口的错误消息文本完全一致，且都可通过 `errors.Is` 识别底层 sentinel：

```go
import (
	"errors"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
)

// 构造阶段感知
cfg, err := sess.ClientConfigWithError("unknown-service")
if errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing) {
	// SDK 版本可能不识别该服务，请升级依赖或显式指定 Endpoint。
}
_ = cfg

// 或使用兼容签名 + 请求期感知
svc := someservice.New(sess) // 内部走 sess.ClientConfig(...)
if _, err := svc.DoSomething(nil); errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing) {
	// 同样能识别到
}
```

遇到该错误时，建议先升级 SDK 版本；若确认 SDK 尚未内置该服务的寻址元数据，可通过 `WithEndpoint(...)` 或自定义 `EndpointResolver` 显式指定。

##### region / service 合法性校验

service 和 region 会被拼接进 Endpoint 的主机名。为防止通过构造 region / service 改写请求主机（SSRF），SDK 在寻址之前会先校验二者都是合法的单个主机名标签：

- 只能包含 ASCII 字母、数字、`-` 和 `_`；
- 必须以字母开头；
- 不能以 `-` 或 `_` 结尾。

例如 `ap-southeast-1`、`rds_mysql` 合法；`ap-southeast-1@127.0.0.1`、`ap-southeast-1:443`、`evil.com`、`ap southeast-1`、`1ap-southeast-1`、`ap-southeast-1-` 非法。先校验 service，再校验 region。空字符串不在此处校验，沿用原有逻辑处理（例如默认寻址中 Regional 服务缺少 region 仍返回 `byteplusutil.ErrMissingRegion`）。

校验时机与覆盖范围：

- 未显式设置 Endpoint（`Config.Endpoint` 为空）时，Session 在调用任何寻址逻辑之前执行校验，默认寻址和自定义 `EndpointResolver` 都在其后，非法值不会传给 Resolver；生成的服务客户端（如 `ecs.New(sess)`）和 `universal` 客户端都会经过这一步。错误暴露时机与上文 [错误处理](#错误处理) 相同：`session.ClientConfigWithError` 在构造阶段返回；`session.ClientConfig`（服务客户端和 `universal` 客户端使用）在请求的 Validate 阶段返回，早于签名和发送，不会发出网络请求。
- `byteplusutil.ResolveDefaultEndpoint` 在查找服务注册信息之前执行该校验，因此非法 service 返回 `InvalidService`，而不是 `ErrServiceEndpointInfoMissing`。
- 标准寻址 `endpoints.NewStandardEndpointResolver()` 的 `EndpointFor` 自身也会执行该校验，直接调用时同样返回错误。
- `feature/rds/auth` 的 `BuildAuthToken` 在生成 token 之前执行该校验，失败时返回错误，不生成 token。
- 显式设置 Endpoint（`WithEndpoint`）时，主机名不由 region / service 拼接，不执行该校验。

校验失败的底层错误是 `bytepluserr.Error`，错误码为 `endpoints.ErrCodeInvalidRegion`（`"InvalidRegion"`）或 `endpoints.ErrCodeInvalidService`（`"InvalidService"`），通过 `Code()` 获取。`ResolveDefaultEndpoint` 直接返回该错误，信息格式如下，原因部分为 `illegal character at index N`、`must start with a letter` 或 `must not end with '-' or '_'` 之一：

```
InvalidRegion: invalid region "ap-southeast-1@127.0.0.1": illegal character at index 14, region must start with a letter, end with a letter or digit, and contain only letters, digits, '-' and '_'
```

通过 Session 返回的错误（`ClientConfigWithError`、服务客户端、`universal` 客户端）带有上文相同的 `byteplus session: resolve client config for service "<service>": ` 前缀；`BuildAuthToken` 再包装一层 `unable to resolve client config: ...`。这些错误都不能直接做类型断言，请使用 `errors.As`：

```go
package main

import (
	"errors"
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/endpoints"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/ecs"
)

func main() {
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：BYTEPLUS_ACCESS_KEY、BYTEPLUS_SECRET_KEY、BYTEPLUS_SESSION_TOKEN
		WithRegion("ap-southeast-1@127.0.0.1")            // 非法 region
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	// 构造阶段感知
	_, err = sess.ClientConfigWithError("ecs")
	var bpErr bytepluserr.Error
	if errors.As(err, &bpErr) {
		switch bpErr.Code() {
		case endpoints.ErrCodeInvalidRegion, endpoints.ErrCodeInvalidService:
			fmt.Println(bpErr.Code(), bpErr.Message())
		}
	}

	// 或使用兼容签名 + 请求期感知
	svc := ecs.New(sess)
	if _, err := svc.DescribeInstances(&ecs.DescribeInstancesInput{}); errors.As(err, &bpErr) {
		fmt.Println(bpErr.Code())
	}
}
```

也可以使用导出的 `endpoints.ValidateRegion`、`endpoints.ValidateService`、`endpoints.ValidateServiceAndRegion` 提前校验，返回的错误与 `ResolveDefaultEndpoint` 相同。

**与旧版本不同：**

- region 首尾带空白（如 `" ap-southeast-1 "`）不会被 trim：直接调用 `byteplusutil.ResolveDefaultEndpoint` 和客户端请求都返回 `InvalidRegion`。旧版本的 `ResolveDefaultEndpoint` 会先去掉首尾空白再拼接主机名，而签名使用的始终是带空白的原值。region 的大小写处理不变：`AP-Southeast-1` 仍合法，主机名中按小写拼接。
- 兼容入口 `byteplusutil.GetDefaultEndpointByServiceInfo` 委托给 `ResolveDefaultEndpoint`，遇到非法 service / region 时不返回错误，而是记录 debug 日志并回退到默认 Endpoint `open.ap-southeast-1.byteplusapi.com`。需要感知错误的调用方请使用 `ResolveDefaultEndpoint`，或先调用 `endpoints.ValidateServiceAndRegion`。

#### Endpoint 标准寻址

##### 标准寻址规则

下表中 **CN mainland** 指服务在 `ServiceInfos` 中为 Regional 且传入 Region 为中国大陆 Region（不含 `cn-hongkong` 等港澳台白名单）。

| Global 服务 | 双栈 | CN mainland | 格式 |
|---|---|---|---|
| 是 | 是 | 任意 | `{Service}.byteplus-api.com` |
| 是 | 否 | 任意 | `{Service}.byteplusapi.com` |
| 否 | 是 | 是 | `{Service}.{region}.byteplus-api.com.cn` |
| 否 | 是 | 否 | `{Service}.{region}.byteplus-api.com` |
| 否 | 否 | 是 | `{Service}.{region}.byteplusapi.com.cn` |
| 否 | 否 | 否 | `{Service}.{region}.byteplusapi.com` |

Global 服务不追加 `.cn` 后缀；`.cn` 仅在 Regional 服务命中中国大陆 Region 时追加。本次默认寻址升级不改变 Standard Resolver 的语义。

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
