[← Transport](3-Transport.md) | Timeout | [Retry →](5-Retry.md)

---

# Timeouts

## Global Timeouts (Client Level)

> - **Default**
>   - `ConnectTimeout`: 30s
>   - `ReadTimeout`: unlimited
>   - Default client: `http.DefaultClient`

Configure timeouts via a custom `http.Client`.

```go
func main() {
    region := "ap-southeast-1"
    transport := &http.Transport{
       Proxy: http.ProxyFromEnvironment,
       DialContext: (&net.Dialer{
          Timeout:   30 * time.Second,
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
       Timeout:   60 * time.Second,
    }
    config := byteplus.NewConfig().
       WithRegion(region).
       WithHTTPClient(client).
       WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    _ = sess
}
```

## Per-API Timeout

Use API methods with suffix `WithContext`.

```go
func main() {
    region :=  "ap-southeast-1"
    config := byteplus.NewConfig().
       WithRegion(region).
       WithCredentials(credentials.NewEnvCredentials())
    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _ = ctx
}
```

---

[← Transport](3-Transport.md) | Timeout | [Retry →](5-Retry.md)
