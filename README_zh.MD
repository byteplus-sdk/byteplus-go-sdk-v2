[English](./README.MD) | 简体中文

### Byteplus SDK for Go
####  概述
1. Go语言SDK
2. Go版本最低1.14+，如使用方舟服务(service/arkruntime)，需要使用go1.18+
3. service目录下的文件请不要修改，都是由代码生成器自动生成
4. 建议使用go mod方式进行依赖管理

##### 代码示例
```go 
package main

import (
	"fmt"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/vpc"
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
更多代码示例请参考：[SDK接入文档](./SDK_Integration_zh.md)