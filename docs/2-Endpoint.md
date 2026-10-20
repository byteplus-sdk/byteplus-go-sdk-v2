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

`WithEndpoint` configures the business API endpoint; it is not inherited by
STS credential requests. AssumeRole uses its explicit provider `Host`, then an
independently configured `STSConfig.Host`, then its default STS endpoint.
An independent STS host is preserved across session/client creation and
`Session.Copy`. OIDC and SAML use their own provider `Endpoint`.

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

If the requested service is not registered in `defaultEndpoint`, the SDK returns `byteplusutil.ErrServiceEndpointInfoMissing`. The two API entry points surface this error at different stages:

- **`session.ClientConfigWithError(serviceName, ...)`** — returns the error **at construction time**, ideal for callers that want to fail fast.
- **`session.ClientConfig(serviceName, ...)`** (compatibility signature) — keeps its historic no-error signature and surfaces the same error on the **first API call** through a `Validate` handler installed on the returned Config. This prevents nil-pointer panics in downstream `xxx.New(sess)` callers.

Both paths produce the exact same error message and both can be matched with `errors.Is`:

```go
import (
	"errors"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/byteplusutil"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
)

// Fail fast at construction time
cfg, err := sess.ClientConfigWithError("unknown-service")
if errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing) {
	// The installed SDK likely does not know this service.
	// Upgrade the dependency or set the endpoint explicitly.
}
_ = cfg

// Or use the compatibility signature and detect on first call
svc := someservice.New(sess) // internally calls sess.ClientConfig(...)
if _, err := svc.DoSomething(nil); errors.Is(err, byteplusutil.ErrServiceEndpointInfoMissing) {
	// Same error, delivered at request time
}
```

When you hit this error, first try upgrading the SDK. If the service is genuinely not carried by the SDK yet, set the endpoint explicitly via `WithEndpoint(...)` or supply a custom `EndpointResolver`.

##### Region and service validation

The service and region are concatenated into the endpoint host name. To prevent a crafted region or service from rewriting the request host (SSRF), the SDK checks that both are a valid single host name label before resolving the endpoint:

- only ASCII letters, digits, `-` and `_`;
- starts with a letter;
- does not end with `-` or `_`.

For example `ap-southeast-1` and `rds_mysql` are valid; `ap-southeast-1@127.0.0.1`, `ap-southeast-1:443`, `evil.com`, `ap southeast-1`, `1ap-southeast-1` and `ap-southeast-1-` are not. The service is checked first, then the region. An empty value is not checked here and is handled by the existing logic (for example, a regional service without a region still returns `byteplusutil.ErrMissingRegion` in default resolution).

When the check runs and what it covers:

- When no endpoint is set explicitly (`Config.Endpoint` is empty), the session runs the check before any resolution; default resolution and any custom `EndpointResolver` come after it, so an invalid value is never passed to the resolver. Generated service clients (e.g. `ecs.New(sess)`) and the `universal` client are covered. The error surfaces at the same stage as described in [Error handling](#error-handling): `session.ClientConfigWithError` returns it at construction time; `session.ClientConfig` (used by service clients and the `universal` client) returns it in the request Validate stage, before signing and sending, so no network request is made.
- `byteplusutil.ResolveDefaultEndpoint` runs the check before looking up the service registration, so an invalid service returns `InvalidService` rather than `ErrServiceEndpointInfoMissing`.
- `EndpointFor` of the standard resolver `endpoints.NewStandardEndpointResolver()` runs the same check itself, so calling it directly also returns the error.
- `BuildAuthToken` in `feature/rds/auth` runs the check before building the token and returns the error without building a token.
- When an endpoint is set explicitly (`WithEndpoint`), the host is not built from the region or service and the check is skipped.

The underlying error of a failed check is a `bytepluserr.Error` with the code `endpoints.ErrCodeInvalidRegion` (`"InvalidRegion"`) or `endpoints.ErrCodeInvalidService` (`"InvalidService"`), available via `Code()`. `ResolveDefaultEndpoint` returns it as is, with the following message format, where the reason is one of `illegal character at index N`, `must start with a letter` or `must not end with '-' or '_'`:

```
InvalidRegion: invalid region "ap-southeast-1@127.0.0.1": illegal character at index 14, region must start with a letter, end with a letter or digit, and contain only letters, digits, '-' and '_'
```

Errors returned through the session (`ClientConfigWithError`, service clients, the `universal` client) carry the same `byteplus session: resolve client config for service "<service>": ` prefix described above; `BuildAuthToken` additionally wraps it as `unable to resolve client config: ...`. A plain type assertion does not work on these errors, use `errors.As`:

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
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion("ap-southeast-1@127.0.0.1") // invalid region
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	// Fail fast at construction time
	_, err = sess.ClientConfigWithError("ecs")
	var bpErr bytepluserr.Error
	if errors.As(err, &bpErr) {
		switch bpErr.Code() {
		case endpoints.ErrCodeInvalidRegion, endpoints.ErrCodeInvalidService:
			fmt.Println(bpErr.Code(), bpErr.Message())
		}
	}

	// Or use the compatibility signature and detect on first call
	svc := ecs.New(sess)
	if _, err := svc.DescribeInstances(&ecs.DescribeInstancesInput{}); errors.As(err, &bpErr) {
		fmt.Println(bpErr.Code())
	}
}
```

The exported functions `endpoints.ValidateRegion`, `endpoints.ValidateService` and `endpoints.ValidateServiceAndRegion` can be used to check values up front; they return the same errors as `ResolveDefaultEndpoint`.

**Behavior change:**

- A region with leading or trailing whitespace (e.g. `" ap-southeast-1 "`) is no longer trimmed: both a direct call to `byteplusutil.ResolveDefaultEndpoint` and a client request return `InvalidRegion`. Previously `ResolveDefaultEndpoint` trimmed the whitespace before building the host, while signing always used the raw value including the whitespace. Case handling is unchanged: `AP-Southeast-1` is still valid and is lowercased in the host.
- The compatibility entry point `byteplusutil.GetDefaultEndpointByServiceInfo` delegates to `ResolveDefaultEndpoint`. For an invalid service or region it returns no error; it writes a debug log and falls back to the default endpoint `open.ap-southeast-1.byteplusapi.com`. Callers that need to detect the error should use `ResolveDefaultEndpoint` or call `endpoints.ValidateServiceAndRegion` first.

#### Standard Endpoint Resolution

##### Resolution Rules

**CN mainland** below means the service is registered as regional in `ServiceInfos` and the caller-supplied region is a Chinese mainland region (i.e. NOT one of the non-mainland whitelist regions such as `cn-hongkong`).

| Global service | DualStack | CN mainland | Format |
|---|---|---|---|
| Yes | Yes | any | `{Service}.byteplus-api.com` |
| Yes | No  | any | `{Service}.byteplusapi.com` |
| No  | Yes | Yes | `{Service}.{region}.byteplus-api.com.cn` |
| No  | Yes | No  | `{Service}.{region}.byteplus-api.com` |
| No  | No  | Yes | `{Service}.{region}.byteplusapi.com.cn` |
| No  | No  | No  | `{Service}.{region}.byteplusapi.com` |

Global services never carry the `.cn` suffix; `.cn` is only appended for regional services when the caller-supplied region is a Chinese mainland region. This addressing upgrade does not change Standard Resolver semantics.

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
