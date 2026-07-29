[← Credentials](1-Credentials.md) | Endpoint[(中文)](2-Endpoint-zh.md) | [Transport →](3-Transport.md)

---

## Endpoint Configuration

> **Default**
>
> If `Endpoint` is not specified, the SDK uses [Automatic Endpoint Resolution](#automatic-endpoint-resolution).

### Custom Endpoint

```go
func main() {
	region := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion(region).
		WithEndpoint("<service>.<regionId>.byteplusapi.com")
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

An explicitly configured Endpoint has the highest priority and skips every subsequent resolution step (including any custom `EndpointResolver`).

### Custom RegionId

```go
func main() {
	regionId := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion(regionId)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### Automatic Endpoint Resolution

Byteplus provides a flexible endpoint resolution mechanism. The SDK automatically builds the endpoint based on the service name, region and the service's Go China flag, and supports DualStack.

#### Default Endpoint Resolution

##### Resolution Logic

1. **Service registration check**

    Every service in the built-in map carries an `IsGlobal` and a `GoChinaEnabled` bool. The SDK builds the endpoint following the rules below.

    - Service missing from the map: `ResolveDefaultEndpoint` returns `byteplusutil.ErrServiceEndpointInfoMissing`. Use `session.ClientConfigWithError(...)` to receive the error at construction time; the compatible `session.ClientConfig(...)` keeps its historic no-error signature and surfaces the same error on the first API call via a `Validate` handler on the returned Config. Detect it with `errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing)`. See [Error handling](#error-handling).

    Built-in service map: `defaultEndpoint` in [`../byteplus/byteplusutil/url.go`](../byteplus/byteplusutil/url.go).

2. **DualStack support (IPv6)**

    Enable via `useDualStack=true` or env var `BYTEPLUS_ENABLE_DUALSTACK=true`. Priority: `useDualStack` > `BYTEPLUS_ENABLE_DUALSTACK`.

    When enabled, the suffix changes from `byteplusapi.com` to `byteplus-api.com`.

3. **Go China suffix**

    When a service entry has `GoChinaEnabled=true` and the request region is in the Chinese mainland (a `cn-*` prefix but not one of the non-mainland regions such as `cn-hongkong`), the resolver appends the `.cn` suffix.

    Whether Go China applies is decided by the service itself and cannot be overridden.

4. **Endpoint construction**

    - **Global services (e.g., `IAM`, `Billing`)**: `<service>.byteplusapi.com` (or `byteplus-api.com` when DualStack is enabled; `.cn` is appended when Go China applies).
    - **Regional services (e.g., `ECS`, `RDS`)**: `<service>.<region>.byteplusapi.com` (DualStack / Go China rules identical to global services).

##### Decision Table

The table lists every effective combination. `RegionType` is derived from the service's `IsGlobal` flag; "Region is Go China" refers to the request region.

| RegionType | GoChinaEnabled | Region is Go China | Endpoint | Region embedded |
|---|---|---|---|---|
| Global | true | yes | `{service}.byteplusapi.com.cn` | no |
| Global | true | no | `{service}.byteplusapi.com` | no |
| Global | false | any | `{service}.byteplusapi.com` | no |
| Regional | true | yes | `{service}.{region}.byteplusapi.com.cn` | yes |
| Regional | true | no | `{service}.{region}.byteplusapi.com` | yes |
| Regional | false | any | `{service}.{region}.byteplusapi.com` | yes |

When DualStack is enabled, replace every occurrence of `byteplusapi.com` in the table with `byteplus-api.com`.

##### `WithBootstrapRegion` / `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF` (Deprecated)

> **⚠️ Deprecated**: `Config.WithBootstrapRegion(map[string]struct{}{...})`, the `Config.BootstrapRegion` field, and the `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF` environment variable are **deprecated** and **no longer participate** in the default addressing pipeline. The related methods and fields are retained only for API/binary compatibility and are treated as no-ops at runtime. **Do not use them in new code.** Existing callers should switch to `WithRegion` + `WithUseDualStack` and let the SDK auto-resolve the endpoint, or override it explicitly via `WithEndpoint`.

##### Code Example

> Note: the calls rendered with strikethrough (`WithBootstrapRegion` / `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF`) are **deprecated** and kept only for backward compatibility. Do not use them in new code.

<pre><code class="language-go">func main() {
	regionId := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion(regionId).
		WithUseDualStack(true). // enable dual stack; also honors env BYTEPLUS_ENABLE_DUALSTACK=true
<s>		WithBootstrapRegion(map[string]struct{}{</s>
<s>			"custom_example_region1": {},</s>
<s>			"custom_example_region2": {},</s>
<s>		}) // Deprecated: custom auto-resolve region list; also honors env BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF</s>
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
</code></pre>

##### Error handling

If the requested service is not registered in `defaultEndpoint`, the SDK returns `byteplusutil.ErrServiceEndpointInfoMissing` on the first client-config build. Detect it with:

```go
import (
	"errors"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
)

if errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing) {
	// The installed SDK likely does not know this service.
	// Upgrade the dependency or set the endpoint explicitly.
}
```

When you hit this error, first try upgrading the SDK. If the service is genuinely not carried by the SDK yet, set the endpoint explicitly via `WithEndpoint(...)` or supply a custom `EndpointResolver`.

#### Standard Endpoint Resolution

##### Resolution Rules

| Global service | DualStack | Format |
|---|---|---|
| Yes | Yes | `{Service}.byteplus-api.com` |
| Yes | No  | `{Service}.byteplusapi.com` |
| No  | Yes | `{Service}.{region}.byteplus-api.com` |
| No  | No  | `{Service}.{region}.byteplusapi.com` |

Standard resolution keeps its historical Go China behavior: when the service is registered as regional and the region is a Chinese mainland region, the `.cn` suffix is appended (`{Service}.{region}.byteplusapi.com.cn`). Non-mainland Chinese regions (for example `cn-hongkong`) stay on `.com`. This addressing upgrade does not change Standard Resolver semantics.

Whether a service is global depends on the service itself and cannot be changed. See `ServiceInfos` in [`../byteplus/endpoints/standard_resolver.go`](../byteplus/endpoints/standard_resolver.go).

##### Code Example

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
		WithEndpointResolver(endpoints.NewStandardEndpointResolver()).
		WithRegion(regionId).
		WithUseDualStack(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

---

[← Credentials](1-Credentials.md) | Endpoint[(中文)](2-Endpoint-zh.md) | [Transport →](3-Transport.md)
