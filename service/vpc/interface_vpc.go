// Code generated by byteplus with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package vpciface provides an interface to enable mocking the VPC service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package vpc

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
)

// VPCAPI provides an interface to enable mocking the
// vpc.VPC service client's API operation,
//
//	// byteplus sdk func uses an SDK service client to make a request to
//	// VPC.
//	func myFunc(svc VPCAPI) bool {
//	    // Make svc.AddBandwidthPackageIp request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := vpc.New(sess)
//
//	    myFunc(svc)
//	}
type VPCAPI interface {
	DescribeVpcs(*DescribeVpcsInput) (*DescribeVpcsOutput, error)
	DescribeVpcsWithContext(byteplus.Context, *DescribeVpcsInput, ...request.Option) (*DescribeVpcsOutput, error)
	DescribeVpcsRequest(*DescribeVpcsInput) (*request.Request, *DescribeVpcsOutput)
}

var _ VPCAPI = (*VPC)(nil)
