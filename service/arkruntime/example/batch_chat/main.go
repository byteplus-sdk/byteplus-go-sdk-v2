package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime/model"
)

func main() {
	// Create a client with your api key
	client := arkruntime.NewClientWithApiKey(
		"YOUR_API_KEY",
		arkruntime.WithBatchMaxParallel(3000), // set batch max parallel to 3000
	)

	// Mock 50000 requests for testing. You can load real requests from file, message queue or database, etc.
	requests := MockRequests("YOUR_ENDPOINT_ID", 50000)

	wg, ctx := sync.WaitGroup{}, context.Background()

	// set global timeout for all requests. If timeout, all requests will be canceled.
	ctx, cancel := context.WithTimeout(ctx, 24*time.Hour)
	defer cancel()

	// do batch inference
	for request := range requests {
		wg.Add(1)
		// create a goroutine for each request
		go func(request model.CreateChatCompletionRequest) {
			defer wg.Done()

			// set timeout for each request. If timeout, the request will be canceled.
			ctx, cancel := context.WithTimeout(ctx, 24*time.Hour)
			defer cancel()

			// do batch inference
			result, err := client.CreateBatchChatCompletion(ctx, request)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(MustMarshalJson(result))
			}
		}(request)
	}
	// wait for all requests to finish
	wg.Wait()
}

// MockRequests mock requests for testing.
// You can load real requests from file, message queue or database, etc.
func MockRequests(endpoint string, count int) <-chan model.CreateChatCompletionRequest {
	requests := make(chan model.CreateChatCompletionRequest)

	go func() {
		defer close(requests)
		for i := 0; i < count; i++ {
			requests <- model.CreateChatCompletionRequest{
				Model: endpoint,
				Messages: []*model.ChatCompletionMessage{
					{
						Role: model.ChatMessageRoleSystem,
						Content: &model.ChatCompletionMessageContent{
							StringValue: byteplus.String("You are a helpful AI assistant."),
						},
					},
					{
						Role: model.ChatMessageRoleUser,
						Content: &model.ChatCompletionMessageContent{
							StringValue: byteplus.String("Hello, how are you?"),
						},
					},
				},
			}
		}
	}()

	return requests
}

func MustMarshalJson(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}
