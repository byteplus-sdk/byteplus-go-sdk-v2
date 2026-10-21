package byteplusquery

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/custom"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/response"
)

// UnmarshalErrorHandler is a name request handler to unmarshal request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "byteplussdk.byteplusquery.UnmarshalError", Fn: UnmarshalError}

// UnmarshalError unmarshals an error response for an BYTEPLUS Query service.
func UnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	processUnmarshalError(unmarshalErrorInfo{
		Request: r,
	})
}

type unmarshalErrorInfo struct {
	Request  *request.Request
	Response *response.ByteplusResponse
	Body     []byte
	Err      error
}

func processUnmarshalError(info unmarshalErrorInfo) {
	r := info.Request
	if info.Response == nil {
		info.Response = &response.ByteplusResponse{}
		if info.Body == nil {
			body, err := ioutil.ReadAll(r.HTTPResponse.Body)
			info.Body = body
			if err != nil {
				r.Error = newResponseError(info, request.ErrCodeRead, "failed to read error response", err)
				return
			}
		}
		if err := json.Unmarshal(info.Body, info.Response); err != nil {
			r.Error = newResponseError(info, request.ErrCodeSerialization, "failed to decode error response", err)
			return
		}
	}

	if r.Config.CustomerUnmarshalError != nil {
		customerErr := r.Config.CustomerUnmarshalError(r.Context(), custom.RequestMetadata{
			ServiceName: r.ClientInfo.ServiceName,
			Version:     r.ClientInfo.APIVersion,
			Action:      r.Operation.Name,
			HttpMethod:  r.Operation.HTTPMethod,
			Region:      *r.Config.Region,
		}, *info.Response)
		if customerErr != nil {
			r.Error = customerErr
			return
		}
	}

	// Non-standard responses (non-JSON, or JSON without a TOP ResponseMetadata)
	// are surfaced as RequestFailure carrying the HTTP status, request id and raw
	// body. They return early and never touch r.Data.Metadata, preserving the
	// original behavior for these paths.
	if info.Response.ResponseMetadata == nil {
		simple := response.ByteplusSimpleError{}
		if err := json.Unmarshal(info.Body, &simple); err != nil {
			r.Error = newResponseError(info, request.ErrCodeSerialization, "failed to decode error response", err)
			return
		}
		code, message := simple.ErrorCode, simple.Message
		if code == "" {
			code, message = "UnknownError", "non-standard error response"
		}
		r.Error = newResponseError(info, code, message, nil)
		return
	}
	if info.Err == nil && (info.Response.ResponseMetadata.Error == nil || info.Response.ResponseMetadata.Error.Code == "") {
		r.Error = newResponseError(info, "UnknownError", "non-standard error response", nil)
		return
	}

	if info.Err != nil {
		r.Error = info.Err
	} else {
		r.Error = bytepluserr.NewRequestFailure(
			bytepluserr.New(info.Response.ResponseMetadata.Error.Code, info.Response.ResponseMetadata.Error.Message, nil),
			r.HTTPResponse.StatusCode,
			info.Response.ResponseMetadata.RequestId,
			r.Config.SimpleError,
		)
	}
	if r.DataFilled() && reflect.TypeOf(r.Data) != reflect.TypeOf(&map[string]interface{}{}) {

		if _, ok := reflect.TypeOf(r.Data).Elem().FieldByName("Metadata"); ok {
			if info.Response.ResponseMetadata != nil {
				info.Response.ResponseMetadata.HTTPCode = r.HTTPResponse.StatusCode
			}
			r.Metadata = *(info.Response.ResponseMetadata)
			reflect.ValueOf(r.Data).Elem().FieldByName("Metadata").Set(reflect.ValueOf(info.Response.ResponseMetadata))
		}
	}
}

// newResponseError surfaces HTTP diagnostics (status code, request id and raw
// response body) when the body cannot be decoded as a standard TOP error, so
// callers no longer lose the real failure. SimpleError only changes Error(),
// not the StatusCode()/RequestID()/Message() accessors.
func newResponseError(info unmarshalErrorInfo, code, message string, err error) bytepluserr.RequestFailure {
	return bytepluserr.NewRequestFailure(
		bytepluserr.New(code, fmt.Sprintf("%s; response body: %s", message, info.Body), err),
		info.Request.HTTPResponse.StatusCode,
		responseRequestID(info),
		info.Request.Config.SimpleError,
	)
}

// responseRequestID resolves the request id from the parsed metadata first, then
// the request context, then the TOP gateway headers, so non-standard responses
// still expose a usable id for troubleshooting.
func responseRequestID(info unmarshalErrorInfo) string {
	if info.Response != nil && info.Response.ResponseMetadata != nil && info.Response.ResponseMetadata.RequestId != "" {
		return info.Response.ResponseMetadata.RequestId
	}
	if info.Request.RequestID != "" {
		return info.Request.RequestID
	}
	for _, header := range []string{"X-Top-Requestid", "X-Top-Request-Id", "X-Tt-Logid"} {
		if id := info.Request.HTTPResponse.Header.Get(header); id != "" {
			return id
		}
	}
	return ""
}
