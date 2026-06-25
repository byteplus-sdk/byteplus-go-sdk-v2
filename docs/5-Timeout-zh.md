[← 代理配置](4-Proxy-zh.md) | 超时配置[(English)](5-Timeout.md) | [重试机制 →](6-Retry-zh.md)

---

## 超时配置

### 全局超时设置（Client 级别）

> **默认**
>
> - 默认 client：`http.DefaultClient`
> - `ConnectTimeout`：30s（即 `net.Dialer{Timeout: 30s}`）
> - `ReadTimeout` / 整体请求超时：不限制（`http.DefaultClient.Timeout == 0`）

暂不支持直接设置 `ConnectTimeOut` 和 `ReadTimeout` 配置，可以通过自定义 HttpClient 来实现。

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
	svc := ecs.New(sess)
}
```

### 单接口指定超时设置

单个接口超时配置需要调用以 **WithContext** 结尾的接口。

```go
func main() {
	region := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewEnvCredentials())
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := svc.AssociateInstancesIamRoleWithContext(ctx, associateInstancesIamRoleInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

---

[← 代理配置](4-Proxy-zh.md) | 超时配置[(English)](5-Timeout.md) | [重试机制 →](6-Retry-zh.md)
