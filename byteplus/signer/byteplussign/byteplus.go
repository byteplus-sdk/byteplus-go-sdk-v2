package byteplussign

import (
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

var SignRequestHandler = request.NamedHandler{
	Name: "byteplussign.SignRequestHandler", Fn: SignSDKRequest,
}

func SignSDKRequest(req *request.Request) {

	region := req.ClientInfo.SigningRegion

	var (
		dynamicCredentials *credentials.Credentials
		dynamicRegion      *string
		c1                 base.Credentials
		err                error
	)

	if req.Config.DynamicCredentialsIncludeError != nil {
		dynamicCredentials, dynamicRegion, err = req.Config.DynamicCredentialsIncludeError(req.Context())
		if err != nil {
			req.Error = err
			return
		}
	} else if req.Config.DynamicCredentials != nil {
		dynamicCredentials, dynamicRegion = req.Config.DynamicCredentials(req.Context())
	}

	if req.Config.DynamicCredentials != nil || req.Config.DynamicCredentialsIncludeError != nil {
		if byteplus.StringValue(dynamicRegion) == "" {
			req.Error = byteplus.ErrMissingRegion
			return
		}
		region = byteplus.StringValue(dynamicRegion)
	} else if region == "" {
		region = byteplus.StringValue(req.Config.Region)
	}

	name := req.ClientInfo.SigningName
	if name == "" {
		name = req.ClientInfo.ServiceID
	}

	if dynamicCredentials == nil {
		c1, err = req.Config.Credentials.GetBase(region, name)
	} else {
		c1, err = dynamicCredentials.GetBase(region, name)
	}

	if err != nil {
		req.Error = err
		return
	}

	r := c1.Sign(req.HTTPRequest)
	req.HTTPRequest.Header = r.Header
}
