package main

import (
	"context"
	"fmt"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime/model/responses"
	"io"
)

/**
 * Authentication
 * If you authorize your endpoint using an API key, you can set your api key to environment variable "ARK_API_KEY"
 * client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
 * Note: If you use an API key, this API key will not be refreshed.
 * To prevent the API from expiring and failing after some time, choose an API key with no expiration date.
 */

func main() {
	client := arkruntime.NewClientWithApiKey("bad15305-3fed-4c35-986f-2676f1ad5828")
	ctx := context.Background()

	fmt.Println("----- round 1 message -----")
	// round 1 message
	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Image{
					Image: &responses.ContentItemImage{
						Type:     responses.ContentItemType_input_image,
						Detail:   responses.ContentItemImageDetail_auto.Enum(),
						ImageUrl: "https://ark-doc.tos-ap-southeast-1.bytepluses.com/see_i2v.jpeg",
					},
				},
			},
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: "What is the person doing in this picture?",
					},
				},
			},
		},
	}
	createResponsesReq := &responses.ResponsesRequest{
		Model: "seed-1-6-250615",
		Input: &responses.ResponsesInput{
			Union: &responses.ResponsesInput_ListValue{
				ListValue: &responses.InputItemList{ListValue: []*responses.InputItem{{
					Union: &responses.InputItem_InputMessage{
						InputMessage: inputMessage,
					},
				}}},
			},
		},
		Caching: &responses.ResponsesCaching{Type: responses.CacheType_enabled.Enum()},
	}

	resp, err := client.CreateResponsesStream(ctx, createResponsesReq)
	if err != nil {
		fmt.Printf("stream error: %v\n", err)
		return
	}
	var responseId string
	for {
		event, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream error: %v\n", err)
			return
		}
		handleEvent(event)
		if responseEvent := event.GetResponse(); responseEvent != nil {
			responseId = responseEvent.GetResponse().GetId()
		}
	}

	fmt.Println()
	// round 2 messages
	fmt.Println("-----round 2---------")
	createResponsesReq.Input.Union = &responses.ResponsesInput_StringValue{StringValue: "How many pictures are there in the former conversation? Describe them please"}
	createResponsesReq.PreviousResponseId = byteplus.String(responseId)
	resp, err = client.CreateResponsesStream(ctx, createResponsesReq)
	if err != nil {
		fmt.Printf("stream error: %v\n", err)
		return
	}
	for {
		event, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream error: %v\n", err)
			return
		}
		handleEvent(event)
		if responseEvent := event.GetResponse(); responseEvent != nil {
			responseId = responseEvent.GetResponse().GetId()
		}
	}
	fmt.Println(responseId)
}

func handleEvent(event *responses.Event) {
	switch event.GetEventType() {
	case responses.EventType_response_reasoning_summary_text_delta.String():
		print(event.GetReasoningText().GetDelta())
	case responses.EventType_response_reasoning_summary_text_done.String(): // aggregated reasoning text
		fmt.Printf("\naggregated reasoning text: %s\n", event.GetReasoningText().GetText())
	case responses.EventType_response_output_text_delta.String():
		print(event.GetText().GetDelta())
	case responses.EventType_response_output_text_done.String(): // aggregated output text
		fmt.Printf("\naggregated output text: %s\n", event.GetText().GetText())
	default:
		return
	}
}
