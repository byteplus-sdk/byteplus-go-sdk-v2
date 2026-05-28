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

Byteplus provides a flexible endpoint resolution mechanism. The SDK automatically builds the endpoint based on service name and region, and supports DualStack.

#### Default Endpoint Resolution

##### Resolution Logic

1. **Whether the region is in the bootstrap list**

    Built-in list implementation: [`./byteplus/byteplusutil/url.go#bootstrapRegion`](./byteplus/byteplusutil/url.go#L463).

    Only predefined regions (e.g., `ap-southeast-2`) or user-configured regions are auto-resolved; others fall back to `open.byteplusapi.com`.

    You can extend the list via env var `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF` or `customBootstrapRegion`.

2. **DualStack support (IPv6)**

    Enable via `useDualStack=true` or env var `BYTEPLUS_ENABLE_DUALSTACK=true`. Priority: `useDualStack` > `BYTEPLUS_ENABLE_DUALSTACK`.

    When enabled, the suffix changes from `byteplusapi.com` to `byteplus-api.com`.

3. **Construct endpoint based on service name and region**

    - **Global services (e.g., `CDN`, `IAM`)**: `<service>.byteplusapi.com` (or `byteplus-api.com` when DualStack is enabled). Example: `cdn.byteplusapi.com`.
    - **Regional services (e.g., `ECS`, `RDS`)**: `<service>.<region>.byteplusapi.com` is used as the default endpoint. Example: `ecs.ap-southeast-1.byteplusapi.com`.

##### Code Example

```go
func main() {
	regionId := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion(regionId).
		WithUseDualStack(true).
		WithBootstrapRegion(map[string]struct{}{
			"custom_example_region1": {},
			"custom_example_region2": {},
		})
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

#### Standard Endpoint Resolution

##### Resolution Rules

| Global service | DualStack | Format |
|---|---|---|
| Yes | Yes | `{Service}.byteplus-api.com` |
| Yes | No  | `{Service}.byteplusapi.com` |
| No  | Yes | `{Service}.{region}.byteplus-api.com` |
| No  | No  | `{Service}.{region}.byteplusapi.com` |

Whether a service is global depends on the service itself and cannot be changed. See: [`./byteplus/endpoints/standard_resolver.go#ServiceInfos`](./byteplus/endpoints/standard_resolver.go#L69).

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
