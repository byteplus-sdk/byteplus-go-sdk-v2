package arkruntime

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/credentials"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/session"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/ark"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime/model"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime/utils"
)

type Client struct {
	config clientConfig

	requestBuilder utils.RequestBuilder

	arkClient               *ark.ARK
	resourceStsTokens       sync.Map
	rwLock                  *sync.RWMutex
	advisoryRefreshTimeout  int
	mandatoryRefreshTimeout int

	batchHTTPClient      *http.Client
	modelBreakerProvider *utils.ModelBreakerProvider
}

func NewClientWithApiKey(apiKey string, setters ...ConfigOption) *Client {
	config := NewClientConfig(apiKey, "", "", setters...)
	return newClientWithConfig(config)
}

func NewClientWithAkSk(ak, sk string, setters ...ConfigOption) *Client {
	config := NewClientConfig("", ak, sk, setters...)
	return newClientWithConfig(config)
}

// NewClientWithConfig creates new API client for specified config.
func newClientWithConfig(config clientConfig) *Client {
	var arkClient *ark.ARK
	arkConfig := byteplus.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(config.ak, config.sk, "")).
		WithRegion(config.region)

	sess, _ := session.NewSession(arkConfig)
	arkClient = ark.New(sess)

	return &Client{
		config:                  config,
		requestBuilder:          utils.NewRequestBuilder(),
		arkClient:               arkClient,
		resourceStsTokens:       sync.Map{},
		rwLock:                  &sync.RWMutex{},
		advisoryRefreshTimeout:  model.DefaultAdvisoryRefreshTimeout,
		mandatoryRefreshTimeout: model.DefaultMandatoryRefreshTimeout,
		batchHTTPClient:         newBatchHTTPClient(config.batchMaxParallel),
		modelBreakerProvider:    utils.NewModelBreakerProvider(),
	}
}

func (c *Client) GetEndpointStsToken(ctx context.Context, endpointId string) (string, error) {
	return c.GetResourceStsToken(ctx, resourceTypeEndpoint, endpointId)
}

func (c *Client) GetResourceStsToken(ctx context.Context, resourceType string, resourceId string) (string, error) {
	err := c.refresh(ctx, resourceType, resourceId)
	if err != nil {
		return "", err
	}

	token, ok := c.resourceStsTokens.Load(fmt.Sprintf(stsTokenKeyPattern, resourceType, resourceId))
	if ok {
		return token.(tokenInfo).token, nil
	}
	return "", nil
}

func (c *Client) refresh(ctx context.Context, resourceType string, resourceId string) error {
	if !c.needRefresh(resourceType, resourceId, c.advisoryRefreshTimeout) {
		return nil
	}

	if c.rwLock.TryLock() {
		defer c.rwLock.Unlock()
		if !c.needRefresh(resourceType, resourceId, c.advisoryRefreshTimeout) {
			return nil
		}

		isMandatoryRefresh := c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout)
		return c.protectedRefresh(ctx, resourceType, resourceId, isMandatoryRefresh)
	} else if c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout) {
		c.rwLock.Lock()
		defer c.rwLock.Unlock()
		if !c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout) {
			return nil
		}
		return c.protectedRefresh(ctx, resourceType, resourceId, true)
	}
	return nil
}

func (c *Client) needRefresh(resourceType string, resourceId string, refreshIn int) bool {
	delta := c.advisoryRefreshTimeout
	if refreshIn > 0 {
		delta = refreshIn
	}

	token, ok := c.resourceStsTokens.Load(fmt.Sprintf(stsTokenKeyPattern, resourceType, resourceId))
	if ok {
		return token.(tokenInfo).expiredTime-time.Now().Unix() < int64(delta)
	}
	return true
}

func (c *Client) protectedRefresh(ctx context.Context, resourceType string, resourceId string, isMandatory bool) error {
	resp, err := c.arkClient.GetApiKeyWithContext(ctx, &ark.GetApiKeyInput{
		DurationSeconds: byteplus.Int32(model.DefaultStsTimeout),
		ResourceIds:     []*string{byteplus.String(resourceId)},
		ResourceType:    byteplus.String(resourceType),
	})
	if err != nil {
		if isMandatory {
			return err
		} else {
			return nil
		}
	}
	c.resourceStsTokens.Store(fmt.Sprintf(stsTokenKeyPattern, resourceType, resourceId), tokenInfo{*resp.ApiKey, int64(*resp.ExpiredTime)})
	return nil
}

type requestOptions struct {
	body   interface{}
	header http.Header
}

type requestOption func(*requestOptions)

func withBody(body interface{}) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func withContentType(contentType string) requestOption {
	return func(args *requestOptions) {
		args.header.Set("Content-Type", contentType)
	}
}

func WithCustomHeader(key, value string) requestOption {
	return func(args *requestOptions) {
		args.header.Set(key, value)
	}
}

func WithCustomHeaders(m map[string]string) requestOption {
	return func(args *requestOptions) {
		for k, v := range m {
			args.header.Set(k, v)
		}
	}
}

func (c *Client) newRequest(ctx context.Context, method, url, resourceType, resourceId string, setters ...requestOption) (*http.Request, *model.RequestError) {
	// Default Options
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}

	requestID := utils.GenRequestId()
	args.header.Set(model.ClientRequestHeader, requestID)
	errH := c.setCommonHeaders(ctx, args, resourceType, resourceId)
	if errH != nil {
		return nil, errH
	}

	for _, setter := range setters {
		setter(args)
	}
	req, err := c.requestBuilder.Build(ctx, method, url, args.body, args.header)
	if err != nil {
		return nil, model.NewRequestError(http.StatusBadRequest, err, requestID)
	}
	return req, nil
}

func (c *Client) sendRequest(client *http.Client, req *http.Request, v model.Response) error {
	requestID := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Accept", "application/json")

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return model.NewRequestError(http.StatusInternalServerError, err, requestID)
	}

	defer res.Body.Close()

	if v != nil {
		v.SetHeader(res.Header)
	}

	if isFailureStatusCode(res) {
		return c.handleErrorResp(res)
	}

	err = decodeResponse(res.Body, v)
	if err != nil {
		err = &model.RequestError{
			HTTPStatusCode: res.StatusCode,
			Err:            err,
			RequestId:      requestID,
		}
	}
	return err
}

func (c *Client) Do(ctx context.Context, method, url, resourceType, resourceId string, v model.Response, setters ...requestOption) (err error) {
	err = utils.Retry(
		ctx,
		utils.RetryPolicy{
			MaxAttempts:    c.config.RetryTimes,
			InitialBackoff: model.ErrorRetryBaseDelay,
			MaxBackoff:     model.ErrorRetryMaxDelay,
		},
		func() bool { return true },
		func() error {
			req, innerErr := c.newRequest(ctx, method, url, resourceType, resourceId, setters...)
			if innerErr != nil {
				return innerErr
			}

			return c.sendRequest(c.config.HTTPClient, req, v)
		},
		nil,
		needRetryError,
	)
	return
}

func (c *Client) DoBatch(ctx context.Context, method, url, resourceType, resourceId string, v model.Response, setters ...requestOption) error {
	breaker := c.modelBreakerProvider.GetOrCreateBreaker(resourceId)

	for {
		breaker.Wait()

		select {
		case <-ctx.Done(): // whole context finish
			return ctx.Err()
		default:
		}

		err := func() error {
			req, er := c.newRequest(ctx, method, url, resourceType, resourceId, setters...)
			if er != nil {
				return er
			}

			return c.sendRequest(c.batchHTTPClient, req, v)
		}()

		// no error: just return on this try
		if err == nil {
			return nil
		}

		// no need to retry error
		if !needRetryError(err) {
			return err
		}

		retryAfter := c.getRetryAfter(v)
		if retryAfter > 0 {
			breaker.Reset(time.Duration(retryAfter) * time.Second)
		}
	}
}

func (c *Client) sendRequestRaw(req *http.Request) (response model.RawResponse, err error) {
	requestID := req.Header.Get(model.ClientRequestHeader)
	resp, err := c.config.HTTPClient.Do(req) //nolint:bodyclose // body should be closed by outer function
	if err != nil {
		err = model.NewRequestError(http.StatusInternalServerError, err, requestID)
		return
	}

	if isFailureStatusCode(resp) {
		err = c.handleErrorResp(resp)
		return
	}

	response.SetHeader(resp.Header)
	response.ReadCloser = resp.Body
	return
}

func sendChatCompletionRequestStream(client *Client, req *http.Request) (*utils.ChatCompletionStreamReader, error) {
	requestID := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.config.HTTPClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return &utils.ChatCompletionStreamReader{}, model.NewRequestError(http.StatusInternalServerError, err, requestID)
	}
	if isFailureStatusCode(resp) {
		return &utils.ChatCompletionStreamReader{}, client.handleErrorResp(resp)
	}
	return &utils.ChatCompletionStreamReader{
		EmptyMessagesLimit: client.config.EmptyMessagesLimit,
		Reader:             bufio.NewReader(resp.Body),
		Response:           resp,
		ErrAccumulator:     utils.NewErrorAccumulator(),
		Unmarshaler:        &utils.JSONUnmarshaler{},
		HttpHeader:         model.HttpHeader(resp.Header),
	}, nil
}

func (c *Client) ChatCompletionRequestStreamDo(ctx context.Context, method, url, resourceId string, setters ...requestOption) (streamReader *utils.ChatCompletionStreamReader, err error) {
	err = utils.Retry(
		ctx,
		utils.RetryPolicy{
			MaxAttempts:    c.config.RetryTimes,
			InitialBackoff: model.ErrorRetryBaseDelay,
			MaxBackoff:     model.ErrorRetryMaxDelay,
		},
		func() bool { return true },
		func() error {
			req, innerErr := c.newRequest(ctx, method, url, resourceTypeEndpoint, resourceId, setters...)
			if innerErr != nil {
				return innerErr
			}

			streamReader, err = sendChatCompletionRequestStream(c, req)
			return err
		},
		nil,
		needRetryError,
	)

	return
}

func (c *Client) setCommonHeaders(ctx context.Context, args *requestOptions, resourceType string, resourceId string) *model.RequestError {
	requestID := args.header.Get(model.ClientRequestHeader)
	if len(c.config.apiKey) > 0 {
		args.header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.apiKey))
	} else {
		if resourceTypeEndpoint == resourceType && !strings.HasPrefix(resourceId, "ep-") {
			return model.NewRequestError(http.StatusBadRequest, model.ErrBodyWithoutEndpoint, requestID)
		} else if resourceTypeBot == resourceType && !strings.HasPrefix(resourceId, "bot-") {
			return model.NewRequestError(http.StatusBadRequest, model.ErrBodyWithoutBot, requestID)
		}
		token, err := c.GetResourceStsToken(ctx, resourceType, resourceId)
		if err != nil {
			if volcErr, ok := err.(bytepluserr.RequestFailure); ok {
				return model.NewRequestError(volcErr.StatusCode(), fmt.Errorf("failed to get resource sts token. err=%w", volcErr), volcErr.RequestID())
			}
			return model.NewRequestError(http.StatusInternalServerError, fmt.Errorf("failed to get resource sts token. err=%w", err), requestID)
		}
		args.header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	return nil
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func needRetryError(err error) bool {
	apiErr := &model.APIError{}
	reqErr := &model.RequestError{}
	if errors.As(err, &apiErr) {
		return apiErr.HTTPStatusCode >= http.StatusInternalServerError || apiErr.HTTPStatusCode == http.StatusTooManyRequests
	} else if errors.Is(err, io.EOF) {
		return true
	} else if errors.As(err, &reqErr) {
		return reqErr.HTTPStatusCode >= http.StatusInternalServerError
	}
	return false
}

func decodeResponse(body io.Reader, v interface{}) error {
	if v == nil {
		return nil
	}

	switch o := v.(type) {
	case *string:
		return decodeString(body, o)
	default:
		return json.NewDecoder(body).Decode(v)
	}
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	requestID := resp.Header.Get(model.ClientRequestHeader)
	var errRes model.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := model.NewRequestError(resp.StatusCode, err, requestID)
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}

	errRes.Error.HTTPStatusCode = resp.StatusCode
	errRes.Error.RequestId = requestID
	return errRes.Error
}

func (c *Client) getRetryAfter(v model.Response) int64 {
	header := v.GetHeader()
	retryAfter := header[model.RetryAfterHeader]
	if len(retryAfter) == 0 || retryAfter[0] == "" {
		return 0
	}
	retryAfterInterval, err := strconv.ParseInt(retryAfter[0], 10, 64)
	if err != nil {
		return 0
	}
	return retryAfterInterval
}

func (c *Client) isAPIKeyAuthentication() bool {
	return c.config.apiKey != ""
}
