中文 | [English](README.MD)

### Byteplus SDK for Go
####  概述
1. Byteplus Go语言SDK
2. Go版本最低1.14+，如使用ModelArk服务(service/arkruntime)，需要使用go1.18+
3. service目录下的文件请不要修改，都是由代码生成器自动生成
4. 建议使用go mod方式进行依赖管理

SDK 详细使用文档（凭证配置、Endpoint、Transport、超时、重试、异常处理、Debug）请参考：[SDK 接入文档](./docs/0-Overview-zh.md)

#### Endpoint设置

如果您要自定义SDK的Endpoint，可以按照以下示例代码设置：
```go
config = byteplus.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithRegion(region).WithEndpoint("ecs.ap-southeast-1.byteplusapi.com")
```

Byteplus标准的Endpoint规则说明：

| Regional 服务                                                                                                                          | Global 服务                                                                        |
|--------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------|
| `{service}.{region}.byteplusapi.com` <br> 例如：云服务ecs在ap-southeast-1 Region域名为： `ecs.ap-southeast-1.byteplusapi.com` | `{service}.byteplusapi.com` <br> 例如：访问控制iam为Global服务，域名为：`iam.byteplusapi.com` |

注：

- Service中存在_符号时，Endpoint时需转为-符号。存在大写字母时需转成小写。


#### 代码示例
代码示例在example目录下

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

## Security and privacy
This project takes security seriously.
For vulnerability reporting and supported versions, see [SECURITY.md](SECURITY.md)
