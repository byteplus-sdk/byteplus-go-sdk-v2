English | [中文](README_zh.MD)

### Byteplus SDK for Go

#### Overview

1. Byteplus Cloud Go SDK.
2. Minimum Go version: **1.14+**. If using ModelArk service (`service/arkruntime`), **1.18+** is required.
3. Files in the `service` directory are auto-generated; do not modify them.
4. It is recommended to use `go mod` for dependency management.

For detailed SDK documentation (credentials, endpoint, transport, timeout, retry, error handling, debugging), see: [SDK Integration Guide](./docs/0-Overview.md)

#### Endpoint Configuration

If you want to customize the SDK endpoint, use the following code:

```go
config = byteplus.NewConfig().
        WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
        WithRegion(region).WithEndpoint("ecs.ap-southeast-1.byteplusapi.com")
```

Standard endpoint rules:

| Regional Service | Global Service |
|---|---|
| `{service}.{region}.byteplusapi.com` <br> Example: `ecs.ap-southeast-1.byteplusapi.com` | `{service}.byteplusapi.com` <br> Example: `iam.byteplusapi.com` |

Note:

- If the service name contains `_`, it should be converted to `-` in the endpoint. Use lowercase for all characters.

#### Code Example

Examples are available in the `example` directory.

```go
package main

import (
    "fmt"

    "github.com/byteplus-sdk/byteplus-go-sdk-v2/service/vpc"
    "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
    "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
    "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
)

func main() {
    var (
        ak     string
        sk     string
        region string
        config *byteplus.Config
        sess   *session.Session
        client *vpc.VPC
        resp   *vpc.DescribeVpcsOutput
        err    error
    )
    ak = "your ak"
    sk = "your sk"
    region = "ap-southeast-1"
    config = byteplus.NewConfig().
        WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
        WithRegion(region)

    sess, err = session.NewSession(config)
    if err != nil {
        panic(err)
    }

    client = vpc.New(sess)

    resp, err = client.DescribeVpcs(&vpc.DescribeVpcsInput{
        PageNumber: byteplus.Int64(1),
        PageSize:   byteplus.Int64(10),
    })

    if err != nil {
        panic(err)
    }

    fmt.Println(&resp)
}
```

