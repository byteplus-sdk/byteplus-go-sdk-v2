English | [简体中文](./README_zh.MD)
### Byteplus SDK for Go

- Go =>1.14+; Go =>1.18+ if using the ModelArk service (service/arkruntime)
- please don't modify any source code under path /service, which is fully machine generated
- recommend using Go modules - built-in feature dependency versioning and dependency management

##### Sample:
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