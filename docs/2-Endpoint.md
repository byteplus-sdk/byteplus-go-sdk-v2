[← Credentials](1-Credentials.md) | Endpoint | [Transport →](3-Transport.md)

---

# Endpoint Configuration

## Custom Endpoint

> - **Default**: if `endpoint` is not specified, the SDK uses [Automatic Endpoint Resolution](#automatic-endpoint-resolution).

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

## Custom RegionId

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

## Automatic Endpoint Resolution

> - **Default**: automatic resolution is enabled; no manual endpoint configuration is required.

BytePlus provides a flexible endpoint resolution mechanism. The SDK automatically builds the endpoint based on service name and region, and supports DualStack.

### Default Endpoint Resolution

**Logic**

1. Whether the region is in the bootstrap list.
   - Only predefined regions or user-configured regions are auto-resolved; others fall back to the default endpoint behavior documented in the SDK integration guide.
2. DualStack support (IPv6)
   - Enable via `useDualStack=true` or env var `VOLC_ENABLE_DUALSTACK=true`.
   - When enabled, the suffix changes from `byteplusapi.com` to `byteplus-api.com`.
3. Construct endpoint:
   - Global services: `<service>.byteplusapi.com`.
   - Regional services: `<service>.<region>.byteplusapi.com`.

**Example**

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

### Standard Endpoint Resolution

| Global service | DualStack | Format |
|---|---|---|
| Yes | Yes | `{Service}.byteplus-api.com` |
| Yes | No  | `{Service}.byteplusapi.com` |
| No  | Yes | `{Service}.{region}.byteplus-api.com` |
| No  | No  | `{Service}.{region}.byteplusapi.com` |

Whether a service is global depends on the service itself and cannot be changed.

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

[← Credentials](1-Credentials.md) | Endpoint | [Transport →](3-Transport.md)
