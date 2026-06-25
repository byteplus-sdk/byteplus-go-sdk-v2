[← Endpoint 配置](2-Endpoint-zh.md) | Transport[(English)](3-Transport.md) | [代理配置 →](4-Proxy-zh.md)

---

## HTTP 连接池配置

> **默认**
>
> - 最大空闲连接数（`MaxIdleConns`）：100
> - 空闲连接存活时间（`IdleConnTimeout`）：90s
> - 每个路由最大连接数（`MaxIdleConnsPerHost`）：2

最大空闲连接数、空闲连接存活时间、每个路由最大连接数没有直接提供参数设置，可以通过自定义 HTTPClient 实现。

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
		MaxIdleConns:          100,              // 所有 host 加起来的最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接最大存活时间
		MaxIdleConnsPerHost:   10,               // 每个 host（路由）最大空闲连接数
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second, // 设置 ReadTimeout
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

## HTTPS 请求配置

### 指定 scheme

> **默认**
>
> - HTTPS

`disableSSL=true` 表示使用 `http`，`disableSSL=false` 表示使用 `https`；建议使用 HTTPS 以提升数据传输安全性。

```go
func main() {
	region := "ap-southeast-1"
	config := byteplus.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials())

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

### 忽略 SSL 验证

> **默认**
>
> - 不忽略（开启 SSL 认证）

没有直接提供参数设置，可以通过自定义 HTTPClient 实现。

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
		MaxIdleConnsPerHost:   10,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, // 跳过服务器证书校验
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

### 指定 TLS 协议版本

> **默认**
>
> - ≥ TLS 1.2

目前只支持自定义 HTTPClient 的方式实现；如果没有特殊要求，建议不要修改。

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
		MaxIdleConnsPerHost:   10,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
		},
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

---

[← Endpoint 配置](2-Endpoint-zh.md) | Transport[(English)](3-Transport.md) | [代理配置 →](4-Proxy-zh.md)
