[← 异常处理](7-ErrorHandling-zh.md) | Debug 机制[(English)](8-Debugging.md) | [概览 →](0-Overview-zh.md)

---

## Debug 机制

为便于客户在处理请求时进行问题排查和调试，SDK 支持日志功能，并提供多种日志级别设置。客户可根据实际需求配置日志级别，获取详细的请求与响应信息，以提升排障效率和系统可观测性。

### 开启 Debug 模式

debug 日志默认关闭，可通过 `WithDebug(true)` 开启。

```go
config := byteplus.NewConfig().
	WithRegion(region).
	WithDebug(true).
	WithCredentials(credentials.NewEnvCredentials())
```

### 指定日志输出位置

默认情况下，日志输出到 `os.Stdout`。如需输出到文件或其他 writer，可使用 `WithLogWriter`。

```go
file, _ := os.Create("sdk.log")
config := byteplus.NewConfig().
	WithRegion(region).
	WithDebug(true).
	WithLogWriter(file).
	WithCredentials(credentials.NewEnvCredentials())
```

---

[← 异常处理](7-ErrorHandling-zh.md) | Debug 机制[(English)](8-Debugging.md) | [概览 →](0-Overview-zh.md)
