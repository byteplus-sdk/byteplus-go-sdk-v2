English | [简体中文](./SDK_Integration_zh.md)

# Table of Contents

* [Integrate the SDK](#integrate-the-sdk)
* [Environment Requirements](#environment-requirements)
* [Securely Configure Credentials](#securely-configure-credentials)

  * [Environment Variables](#environment-variables)

    * [Linux](#linux)
    * [Windows](#windows)

      * [Graphical Interface](#graphical-interface)
      * [Command Line](#command-line)
* [Credentials](#credentials)

  * [AK / SK](#aksk)
  * [STS Token](#sts-token)
  * [AssumeRole](#assumerole)
* [Endpoint Configuration](#endpoint-configuration)

  * [Custom Endpoint](#custom-endpoint)
  * [Custom RegionId](#custom-regionid)
  * [Automatic Endpoint Resolution](#automatic-endpoint-resolution)

    * [Default Resolution Logic](#default-endpoint-resolution)
* [HTTP Connection-Pool Settings](#http-connection-pool-settings)
* [HTTPS Request Settings](#https-request-settings)

  * [Specify the Scheme](#specify-the-scheme)
  * [Skip SSL Verification](#skip-ssl-verification)
  * [Specify the TLS Version](#specify-the-tls-version)
* [Timeouts](#timeouts)

  * [Global Timeouts (Client Level)](#global-timeouts-client-level)
  * [Per-API Timeout](#per-api-timeout)
* [Retry Mechanism](#retry-mechanism)

  * [Enable Retries](#enable-retries)
  * [Retry Attempts](#retry-attempts)
  * [Custom Retry Error Codes](#custom-retry-error-codes)
* [Error Handling](#error-handling)
* [Debugging](#debugging)
* [Custom Logger](#custom-logger)

  * [Implementing Your Own Logger](#implementing-your-own-logger)

---

# Integrate the SDK

When calling BytePlus APIs, we recommend integrating the official SDK into your project.
Using the SDK simplifies development, speeds up integration, and significantly reduces long-term maintenance costs.
SDK integration generally involves three steps:

1. **Add the SDK** to your project.
2. **Configure access credentials**.
3. **Write the API-calling code**.

This guide explains each step in detail, with tips for common scenarios.

---

# Environment Requirements

1. **Go 1.14+ is required.**
   If you use the **ArkRuntime** service (`service/arkruntime`), Go 1.18+ is required.
2. **Go Modules** (`go mod`) is the recommended dependency-management method.

---

# Securely Configure Credentials

To avoid credential leakage, **never hard-code your keys** in source code.
BytePlus offers several secure ways to load credentials, such as environment variables.

## Environment Variables

### Linux

> ⚠️ **Notes**
> Environment variables set with `export` are **session-scoped**.
> For persistence, add the `export` commands to your shell’s start-up file.


| Key                                                       | Command                                                 |
| --------------------------------------------------------- | ------------------------------------------------------- |
| `BYTEPLUS_ACCESS_KEY_ID`*(or `BYTEPLUS_ACCESS_KEY`)*      | `export BYTEPLUS_ACCESS_KEY_ID=yourAccessKeyID`         |
| `BYTEPLUS_SECRET_ACCESS_KEY` *(or `BYTEPLUS_SECRET_KEY`)* | `export BYTEPLUS_SECRET_ACCESS_KEY=yourSecretAccessKey` |
| `BYTEPLUS_SESSION_TOKEN`                                  | `export BYTEPLUS_SESSION_TOKEN=yourSessionToken`        |

**Verify:** run `echo $BYTEPLUS_ACCESS_KEY_ID`; the correct key indicates success.

### Windows

Two options are available: **Graphical Interface** or **Command Line**.
To verify, open *Command Prompt* and run:

```bat
echo %BYTEPLUS_ACCESS_KEY_ID%
echo %BYTEPLUS_SECRET_ACCESS_KEY%
echo %BYTEPLUS_SESSION_TOKEN%
```

If the expected values are returned, the variables are set.

#### Graphical Interface

*Windows 10 example*

1. Right-click **This PC → Properties → Advanced system settings → Environment Variables**.
2. Click **New** under either *User variables* or *System variables*.
3. Add the variables:


| Variable             | Example                                         |
| -------------------- | ----------------------------------------------- |
| **AccessKey Id**     | Name`BYTEPLUS_ACCESS_KEY_ID`  Value `*****`     |
| **AccessKey Secret** | Name`BYTEPLUS_SECRET_ACCESS_KEY`  Value `*****` |
| **Session Token**    | Name`BYTEPLUS_SESSION_TOKEN`  Value `*****`     |

#### Command Line

Run *as Administrator*:

```bat
setx BYTEPLUS_ACCESS_KEY_ID yourAccessKeyID /M
setx BYTEPLUS_SECRET_ACCESS_KEY yourAccessKeySecret /M
setx BYTEPLUS_SESSION_TOKEN yourSessionToken /M
```

> `/M` stores the variable at the **system** level.
> Omit `/M` for **user-level** variables.

---

# Credentials

BytePlus SDK supports three mainstream authentication methods:


| Method         | Typical Scenario                                    |
| -------------- | --------------------------------------------------- |
| **AK / SK**    | Long-term server workloads.                         |
| **STS Token**  | Short-lived access, e.g. mobile or browser clients. |
| **AssumeRole** | Cross-account or fine-grained delegated access.     |

Environment-variable loading is described in [Environment Variables](#environment-variables).

## AK/SK

AK (AccessKey) and SK (SecretKey) are permanent credentials created in the BytePlus console.

> ⚠️ **Best Practices**
>
> 1. **Never** embed AK/SK in client apps.
> 2. Store keys in a **configuration service** or **environment variables**.
> 3. Apply the **principle of least privilege**.

```go
func main() {
    ak, sk, region := "Your AK", "Your SK", "ap-southeast-1"
    cfg := byteplus.NewConfig().
       WithRegion(region).
       // 1. Static credentials—*NOT* recommended for production
       WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
       // 2. Preferred: load from env vars
       // WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(cfg)
    if err != nil { panic(err) }
}
```

## STS Token

Security Token Service (STS) issues **temporary** credentials (AK, SK, Token).

> ⚠️ **Best Practices**
>
> 1. Grant **least privilege**; avoid `*` wild-cards.
> 2. Set an appropriate **expiry**—preferably ≤ 1 hour.

```go
func main() {
    ak, sk, token, region := "Your AK", "Your SK", "Your token", "ap-southeast-1"
    cfg := byteplus.NewConfig().
       WithRegion(region).
       WithCredentials(credentials.NewStaticCredentials(ak, sk, token))
       // Or load from env vars:
       // WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(cfg)
    if err != nil { panic(err) }
}
```

## AssumeRole

Obtains **auto-refreshing temporary credentials** before your STS token expires (refreshed 60 s early).

> ⚠️ **Best Practices**
>
> 1. Grant **least privilege**; no `*` wild-cards.
> 2. Expiry ≤ 12 hours; shorter is safer.
> 3. Use **fine-grained roles** with minimal permissions.

```go
func main() {
    ak, sk, region := "Your AK", "Your SK", "ap-southeast-1"
    cfg := byteplus.NewConfig().
        WithRegion(region).
        WithCredentials(credentials.NewStsCredentials(credentials.StsValue{
            AccessKey:  ak,                   // or os.Getenv("BYTEPLUS_ACCESS_KEY_ID")
            SecurityKey: sk,                  // or os.Getenv("BYTEPLUS_SECRET_ACCESS_KEY")
            RoleName:   "RoleName",
            Host:       "Host",
            Region:     "Region",
            AccountId:  "123456",
            Schema:     "Schema",
            Timeout:    5 * time.Second,
            DurationSeconds: 900,             // 15 min
        }))

    sess, err := session.NewSession(cfg)
    if err != nil { panic(err) }
}
```

---

# Endpoint Configuration

## Custom Endpoint

> **Default:**
> if no endpoint is specified, the SDK falls back to [Automatic Endpoint Resolution](#automatic-endpoint-resolution).

```go
func main() {
    region := "ap-southeast-1"
    cfg := byteplus.NewConfig().
       WithCredentials(credentials.NewEnvCredentials()).
       WithRegion(region).
       WithEndpoint("<service>.<regionid>.byteplusapi.com")

    sess, err := session.NewSession(cfg)
    if err != nil { 
	panic(err) 
    }
}
```

## Custom RegionId

```go
func main() {
    regionId := "ap-southeast-1"
    cfg := byteplus.NewConfig().
       WithCredentials(credentials.NewEnvCredentials()).
       WithRegion(regionId)
    sess, err := session.NewSession(cfg)
    if err != nil {
	panic(err) 
    }
}
```

## Automatic Endpoint Resolution

> **Default:**
> enabled; no manual configuration required.

BytePlus constructs the endpoint from service name and region, and supports IPv6 “DualStack”.

### Default Endpoint Resolution

1. **Auto-discoverable Regions**  
   Built-in automatic addressing Region list code:[./byteplus/byteplusutil/url.go#bootstrapRegion](./byteplus/byteplusutil/url.go#L179)   
   Automatic resolution applies only to predefined regions (e.g. `ap-southeast-2`).
   Extend the list via the `BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF` env var or `WithBootstrapRegion()`.
2. **DualStack (IPv6)**
   Enable with `WithUseDualStack(true)` or `BYTEPLUS_ENABLE_DUALSTACK=true`.
   The domain suffix changes from `byteplusapi.com` to `byteplus-api.com`.
3. **Address Format**

   * **Global services** (e.g. `BILLING`, `IAM`): `service.byteplusapi.com`
   * **Regional services** (e.g. `ECS`, `VPC`): `service.region.byteplusapi.com`

```go
func main() {
    region := "ap-southeast-1"
    cfg := byteplus.NewConfig().
        WithCredentials(credentials.NewEnvCredentials()).
        WithRegion(region).
        WithUseDualStack(true). // or env var
        WithBootstrapRegion(map[string]struct{}{
           "custom_example_region1": {},
           "custom_example_region2": {},
        }) // Custom auto-discoverable Regions can also be configured using the environment variable BYTEPLUS_BOOTSTRAP_REGION_LIST_CONF

    sess, err := session.NewSession(cfg)
    if err != nil { 
       panic(err) 
    }
}
```

---

# HTTP Connection-Pool Settings

> **Default values**
>
> * `MaxIdleConns`: 100
> * `IdleConnTimeout`: 90 s
> * `MaxIdleConnsPerHost`: 2

Configure by supplying a custom `http.Client`:

```go
func main() {
    region := "ap-southeast-1"
    tr := &http.Transport{
        Proxy: http.ProxyFromEnvironment,
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
            DualStack: true,
        }).DialContext,
        MaxIdleConns:        100,
        IdleConnTimeout:     90 * time.Second,
        MaxIdleConnsPerHost: 10,
        TLSHandshakeTimeout: 10 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
    }
    client := &http.Client{
        Transport: tr,
        Timeout:   60 * time.Second, // ReadTimeout
    }
    cfg := byteplus.NewConfig().
        WithRegion(region).
        WithHTTPClient(client).
        WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(cfg)
    if err != nil { panic(err) }
    svc := ecs.New(sess)
}
```

---

# HTTPS Request Settings

## Specify the Scheme

> **Default:**
> `https`

`DisableSSL=true` switches to `http`; `false` (default) keeps `https`.

```go
func main() {
    region := "ap-southeast-1"
    cfg := byteplus.NewConfig().
       WithRegion(region).
       WithDisableSSL(true). //true means http, false means https
       WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(cfg)
    if err != nil { panic(err) }
}
```

## Skip SSL Verification

Provide a custom `http.Client`:

```go
tr := &http.Transport{
    /* ... other settings ... */
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ⚠️ Not recommended for production
}
```

## Specify the TLS Version

```go
tr := &http.Transport{
    /* ... */
    TLSClientConfig: &tls.Config{
        MinVersion: tls.VersionTLS12,
        MaxVersion: tls.VersionTLS13,
    },
}
```

---

# Timeouts

## Global Timeouts (Client Level)

> **Defaults**
>
> * `ConnectTimeout`: 30 s
> * `ReadTimeout`: unlimited (uses `http.DefaultClient`)

Set by injecting a custom `http.Client`.

```Go
func main() {
    region := "ap-southeast-1"
    transport := &http.Transport{
       Proxy: http.ProxyFromEnvironment,
       DialContext: (&net.Dialer{
          Timeout:   30 * time.Second, // ConnectTimeOut
          KeepAlive: 30 * time.Second,
          DualStack: true,
       }).DialContext,
       MaxIdleConns:          100,
       IdleConnTimeout:       90 * time.Second,
       TLSHandshakeTimeout:   10 * time.Second,
       ExpectContinueTimeout: 1 * time.Second,
    }

    client := &http.Client{
       Transport: transport,
       Timeout:   60 * time.Second, // ReadTimeout
    }
    config := byteplus.NewConfig().
       WithRegion(region).
       WithHTTPClient(client). 
       WithCredentials(credentials.NewEnvCredentials()) // env set：BYTEPLUS_ACCESS_KEY_ID、BYTEPLUS_SECRET_ACCESS_KEY、BYTEPLUS_SESSION_TOKEN

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    svc := ecs.New(sess)
}
```

## Per-API Timeout

Call the `WithContext` variant and pass a `context.Context` with timeout.

```go
// Create a context with a 5-second timeout; note: if a context already exists, replace context.Background() here with the existing context.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

resp, err := svc.AssociateInstancesIamRoleWithContext(ctx, input)
```

---

# Retry Mechanism

Network errors and throttling are automatically retried; business-logic errors are not.

> **Back-off formula**
> `delay = min(MaxDelay, 2ⁿ × minDelay × (1 + Rand[0,1)) + Retry-After`


| Parameter       | Description                                    |
|-----------------| ---------------------------------------------- |
| `MaxDelay`      | Upper bound (default 500 ms)                   |
| `2ⁿ`            | Pure exponential growth                        |
| `1+Rand[0,1)`   | Uniform jitter to mitigate “thundering herd” |
| `min(...)`      |     To prevent unbounded growth beyond MaxDelay.                                         |
| `Retry-After`   | Server-specified sleep time, if provided       |

**Example** (`MaxDelay=500 ms`, `minDelay=30 ms`):


| Retry# (n) | Back-off + Jitter (ms) | Retry-After (ms) | Final Delay (ms) |
| ---------- | ---------------------- | ---------------- | ---------------- |
| 0          | 30 – 60               | 10               | 40 – 70         |
| 1          | 60 – 120              | 20               | 80 – 140        |
| 2          | 120 – 240             | 30               | 150 – 270       |
| …         | ≤ 500                 | 10               | ≤ 510           |

## Enable Retries

> **Default:**
> enabled (3 attempts)

Disable by setting `WithMaxRetries(0)`.

```go
func main() {  
        region := "ap-southeast-1"
        // Configure retry settings
        config := byteplus.NewConfig().
                WithRegion(region).
                WithDisableSSL(true).
                WithCredentials(credentials.NewEnvCredentials()). //env set：BYTEPLUS_ACCESS_KEY_ID、BYTEPLUS_SECRET_ACCESS_KEY、BYTEPLUS_SESSION_TOKEN
                // close retry
                WithMaxRetries(0)

        sess, err := session.NewSession(config)
        if err != nil {
            panic(err)
        }
        svc := ecs.New(sess)
}
```

## Retry Attempts

Set with `WithMaxRetries(n)`.

```go
func main() {
        region := "ap-southeast-1"
        // Configure retry settings
        config := byteplus.NewConfig().
                WithRegion(region).
                WithDisableSSL(true).
                WithCredentials(credentials.NewEnvCredentials()). //env set：BYTEPLUS_ACCESS_KEY_ID、BYTEPLUS_SECRET_ACCESS_KEY、BYTEPLUS_SESSION_TOKEN
                // set retry count (default is DefaultRetryerMaxNumRetries)
                WithMaxRetries(4)

        sess, err := session.NewSession(config)
        if err != nil {
            panic(err)
        }
        svc := ecs.New(sess)
}
```

## Custom Retry Error Codes

Pass a callback to override `request.RetryErrorCodes`.

```go
func main() {
    region := "ap-southeast-1"
    // Configure retry settings
    config := byteplus.NewConfig().
        WithRegion(region).
        WithDisableSSL(true).
        WithCredentials(credentials.NewEnvCredentials()). //env set：BYTEPLUS_ACCESS_KEY_ID、BYTEPLUS_SECRET_ACCESS_KEY、BYTEPLUS_SESSION_TOKEN
        // set retry count (default is DefaultRetryerMaxNumRetries)
        WithMaxRetries(4)
    
    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
    svc := ecs.New(sess)
    describeAvailableResourceInput := &ecs.DescribeAvailableResourceInput{
        DestinationResource: byteplus.String("InstanceType"),
        InstanceTypeId:      byteplus.String("ecs.g2i.large"),
        ZoneId:              byteplus.String("cn-*****"),
    }
    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    resp, err := svc.DescribeAvailableResourceWithContext(ctx, describeAvailableResourceInput, func(request *request.Request) {
        // custom retry error code
        request.RetryErrorCodes = []string{"InvalidAccessKey"}
    })
    if err != nil {
        panic(err)
    }
}
```

---

# Error Handling


| Type                     | Scenario               | Returns                        | Common Fields                                 | Extra Fields                  |
|--------------------------| ---------------------- | ------------------------------ | --------------------------------------------- | ----------------------------- |
| **1. Client Error**      | The request did not reach the server, verify the parameters  | `bytepluserr.Error` or `error` | `Code()`, `Message()`, `Error()`, `OrigErr()` | —                            |
| **2. Server Error**      | Request reached server | `bytepluserr.RequestFailure`   | same                                          | `RequestID()`, `StatusCode()` |
| **3. Network / Timeout** | DNS failure, deadline  | `bytepluserr.Error`            | same                                          | —                            |
| **4. Other Error**       | Uncategorized          | same as above                  | same                                          | —                            |

(See the full sample code for detailed handling.)

```go
package main

import (
   "context"
   "errors"
   "fmt"
   "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
   "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
   "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
   "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
   "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
   "github.com/byteplus-sdk/byteplus-go-sdk-v2/service/ecs"
   "net"
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
         fmt.Println("1. Client Error(Session Creation)", be.Code(), be.Message(), be.Error())
      } else {
         fmt.Println("4. Other Error", err.Error())
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
      var requestFailure bytepluserr.RequestFailure // Server Error
      var errInvalidParam request.ErrInvalidParam   // Parameter Validation
      if errors.As(err, &errInvalidParam) {
         fmt.Println("1. Client Error(Parameter Validation)：", errInvalidParam.Code(), errInvalidParam.Field(), errInvalidParam.Error())
      } else if errors.As(err, &requestFailure) {
         fmt.Println("2. Server Error：", requestFailure.RequestID(), requestFailure.Code(), requestFailure.StatusCode(), requestFailure.Error())
      } else if errors.As(err, &be) {
         switch be.Code() {
         case "RequestCanceled":
            fmt.Println("3. Network / Timeout：context timeout")
         case "RequestError":
            if be.OrigErr() != nil {
               var netErr net.Error
               var dnsError *net.DNSError
               if errors.As(be.OrigErr(), &dnsError) {
                  fmt.Println("3. Network / Timeout：DNS parse error")
               } else if errors.As(be.OrigErr(), &netErr) && netErr.Timeout() {
                  var oPError *net.OpError
                  if errors.Is(be.OrigErr(), context.DeadlineExceeded) {
                     fmt.Println("3. Network / Timeout：http.Client Timeout(ReadTimeout)....", be.Code(), be.Error())
                  } else if errors.As(be.OrigErr(), &oPError) && oPError.Op == "dial" {
                     fmt.Println("3. Network / Timeout：http.Client Transport.Dialer Timeout(ConnectTimeout)....", be.Code(), be.Error())
                  } else {
                     fmt.Println("3. Network / Timeout", be.Code(), be.Message(), be.Error())
                  }
               }
            }
         default:
            fmt.Println("4. Other Error", be.Code(), be.Message(), be.Error())
         }
      } else {
         fmt.Println("4. Other Error", err.Error())
      }

   }

}
```

---

# Debugging

Log levels:


| Level                        | Description             |
| ---------------------------- | ----------------------- |
| `LogOff`                     | Disable logs (default)  |
| `LogDebug`                   | Enable debug logs       |
| `LogDebugWithSigning`        | Include signing details |
| `LogDebugWithHTTPBody`       | Include HTTP bodies     |
| `LogDebugWithRequestRetries` | Include retry info      |
| `LogDebugWithInputAndOutput` | Include struct I/O      |

```go
cfg := byteplus.NewConfig().
    WithLogLevel(byteplus.LogDebugWithInputAndOutput)
```

---

# Custom Logger

## Default Logger

If no logger is supplied, `byteplus/logger.go` provides a minimal console logger.

## Implementing Your Own Logger

```go
type myLogger struct {
    logger *log.Logger
}
func (l *myLogger) Log(args ...interface{}) {
    for i, arg := range args {
        if s, ok := arg.(string); ok {
            args[i] = strings.ReplaceAll(s, "KeyWord", "***") // redact secrets
        }
    }
    l.logger.Println(args...)
}

func main() {
    file, _ := os.Create("ecs_test.log")
    multi := io.MultiWriter(os.Stdout, file)
    custom := &myLogger{
        logger: log.New(multi, "[MyApp] ", log.LstdFlags|log.Lshortfile),
    }
    cfg := byteplus.NewConfig().
        WithLogLevel(byteplus.LogDebugWithInputAndOutput).
        WithLogger(custom).
        WithCredentials(credentials.NewEnvCredentials()).
        WithRegion("ap-southeast-1")

    sess, err := session.NewSession(cfg)
    if err != nil { panic(err) }
}
```
