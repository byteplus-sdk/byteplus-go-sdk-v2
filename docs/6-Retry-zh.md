[← 超时配置](5-Timeout-zh.md) | 重试机制[(English)](6-Retry.md) | [异常处理 →](7-ErrorHandling-zh.md)

---

## 重试机制

请求的处理逻辑内置了网络异常重试逻辑，即当遇到网络异常问题或限流错误时，系统会自动尝试重新发起请求，以确保服务的稳定性和可靠性。若请求因业务逻辑错误而报错，例如参数错误、资源不存在等情况，SDK 将不会执行重试操作。

### 开启/关闭重试机制

> **默认**
>
> - 开启（3 次）

如果想关闭，可以把最大尝试次数改为 0。

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
	svc := ecs.New(sess)
}
```

### 重试次数

> **默认**
>
> - 3 次

设置最大重试次数：

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
	svc := ecs.New(sess)
}
```

### 自定义重试错误码

在调用接口的时候可以根据业务需求，自定义重试的错误码（服务端返回的错误码）。

```go
resp, err := svc.DescribeAvailableResourceWithContext(ctx, describeAvailableResourceInput, func(request *request.Request) {
	request.RetryErrorCodes = []string{"InvalidAccessKey"}
})
```

---

[← 超时配置](5-Timeout-zh.md) | 重试机制[(English)](6-Retry.md) | [异常处理 →](7-ErrorHandling-zh.md)
