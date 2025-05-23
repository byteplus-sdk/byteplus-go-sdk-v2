package corehandlers

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import "github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/request"

// ValidateParametersHandler is a request handler to validate the input parameters.
// Validating parameters only has meaning if done prior to the request being sent.
var ValidateParametersHandler = request.NamedHandler{Name: "core.ValidateParametersHandler", Fn: func(r *request.Request) {
	if !r.ParamsFilled() {
		return
	}

	if v, ok := r.Params.(request.Validator); ok {
		if err := v.Validate(); err != nil {
			r.Error = err
		}
	}
}}
