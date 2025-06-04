package arkruntime

import (
	"context"
	"net/http"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime/model"
)

const generateImagesPath = "/images/generations"

func (c *Client) GenerateImages(
	ctx context.Context,
	request model.GenerateImagesRequest,
	setters ...requestOption,
) (response model.ImagesResponse, err error) {
	if !c.isAPIKeyAuthentication() {
		return response, model.ErrAKSKNotSupported
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(generateImagesPath), resourceTypeEndpoint, request.Model, &response, requestOptions...)
	return
}
