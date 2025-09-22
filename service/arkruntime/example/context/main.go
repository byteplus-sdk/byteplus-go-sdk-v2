package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime/model"
)

/**
 * Authentication
 * 1.If you authorize your endpoint using an API key, you can set your api key to environment variable "ARK_API_KEY"
 * client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
 * Note: If you use an API key, this API key will not be refreshed.
 * To prevent the API from expiring and failing after some time, choose an API key with no expiration date.
 *
 * 2.If you authorize your endpoint with Volcengine Identity and Access Management（IAM), set your api key to environment variable "BYTEPLUS_ACCESSKEY", "BYTEPLUS_SECRETKEY"
 * client := arkruntime.NewClientWithAkSk(os.Getenv("BYTEPLUS_ACCESSKEY"), os.Getenv("BYTEPLUS_SECRETKEY"))
 * To get your ak&sk, please refer to this document(https://www.byteplus.com/docs/6291/65568)
 * For more information，please check this document（https://www.byteplus.com/docs/82379/1263279）
 */

func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	goCtx := context.Background()

	fmt.Println("----- create context -----")
	createCtxReq := model.CreateContextRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Mode:  model.ContextModeSession,
		// initial messages
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleSystem,
				Content: &model.ChatCompletionMessageContent{
					StringValue: byteplus.String("You are Francis, a helpful AI assistant."),
				},
			},
		},
		TTL: byteplus.Int(3600),
	}

	createCtxRsp, err := client.CreateContext(goCtx, createCtxReq)
	if err != nil {
		fmt.Printf("create context error: %v\n", err)
		return
	}
	fmt.Printf("create context response: %v\n", createCtxRsp)

	fmt.Println("----- chat round 1 (non-stream) -----")
	req := model.ContextChatCompletionRequest{
		ContextID: createCtxRsp.ID,
		Model:     "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: byteplus.String("Hello, I am Jack."),
				},
			},
		},
	}

	resp, err := client.CreateContextChatCompletion(goCtx, req)
	if err != nil {
		fmt.Printf("non-stream chat error: %v\n", err)
		return
	}
	if len(resp.Choices) > 0 {
		fmt.Println(*resp.Choices[0].Message.Content.StringValue)
	}

	fmt.Println("----- chat round 2 (stream) -----")
	req = model.ContextChatCompletionRequest{
		ContextID: createCtxRsp.ID,
		Model:     "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: byteplus.String("Who am I?"),
				},
			},
		},
	}
	stream, err := client.CreateContextChatCompletionStream(goCtx, req)
	if err != nil {
		fmt.Printf("stream chat error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Stream chat error: %v\n", err)
			return
		}
		if len(recv.Choices) > 0 {
			fmt.Print(recv.Choices[0].Delta.Content)
		}
	}
}
