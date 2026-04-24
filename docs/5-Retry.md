[← Timeout](4-Timeout.md) | Retry | [Error Handling →](6-ErrorHandling.md)

---

# Retries

The SDK includes retry logic for network errors and throttling. Business logic errors (invalid params, resource not found) are not retried.

## Enable/Disable Retries

> - **Default**: enabled

Set max retries to `0` to disable.

```go
func main() {
    region := "ap-southeast-1"
    config := byteplus.NewConfig().
            WithRegion(region).
            WithDisableSSL(true).
            WithCredentials(credentials.NewEnvCredentials()).
            WithMaxRetries(0)

    sess, err := session.NewSession(config)
    if err != nil {
            panic(err)
    }
    _ = sess
}
```

## Retry Count

> - **Default**: use service defaults unless overridden

```go
func main() {
    region := "ap-southeast-1"
    config := byteplus.NewConfig().
            WithRegion(region).
            WithDisableSSL(true).
            WithCredentials(credentials.NewEnvCredentials()).
            WithMaxRetries(4)

    sess, err := session.NewSession(config)
    if err != nil {
            panic(err)
    }
    _ = sess
}
```

## Custom Retry Error Codes

Configure retryable server error codes per request.

```go
resp, err := svc.DescribeAvailableResourceWithContext(ctx, input, func(request *request.Request) {
    request.RetryErrorCodes = []string{"InvalidAccessKey"}
})
_ = resp
_ = err
```

---

[← Timeout](4-Timeout.md) | Retry | [Error Handling →](6-ErrorHandling.md)
