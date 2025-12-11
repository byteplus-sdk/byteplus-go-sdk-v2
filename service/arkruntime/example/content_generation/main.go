package main

import (
	"context"
	"fmt"
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
 */
func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()
	modelEp := "YOUR_ENDPOINT_ID"

	fmt.Println("----- create content generation task -----")
	createReq := model.CreateContentGenerationTaskRequest{
		Model: modelEp, // Replace with your endpoint ID
		Content: []*model.CreateContentGenerationContentItem{
			{
				Type: model.ContentGenerationContentItemTypeText,
				Text: byteplus.String("Bird soaring above vast grasslands --ratio 1:1"),
			},
			{
				Type: model.ContentGenerationContentItemTypeImage,
				ImageURL: &model.ImageURL{
					URL: "${YOUR URL HERE}", // Replace with URL
				},
			},
		},
		// CallbackUrl: byteplus.String("CALLBACK_URL"),
	}

	createResponse, err := client.CreateContentGenerationTask(ctx, createReq)
	if err != nil {
		fmt.Printf("create content generation error: %v\n", err)
		return
	}
	fmt.Printf("Task Created with ID: %s\n", createResponse.ID)

	fmt.Println("----- get content generation task -----")
	taskID := createResponse.ID

	getRequest := model.GetContentGenerationTaskRequest{ID: taskID}

	task, err := client.GetContentGenerationTask(ctx, getRequest)
	if err != nil {
		fmt.Printf("get content generation task error: %v\n", err)
		return
	}
	fmt.Printf("----- GetContentGenerationTask content generation task service_tier :%v\n -----", *task.ServiceTier)
	fmt.Printf("----- GetContentGenerationTask content generation task expire_time :%v\n -----", *task.ExecutionExpiresAfter)

	fmt.Println("----- list default content generation tasks -----")

	listDefaultReq := model.ListContentGenerationTasksRequest{
		PageNum:  byteplus.Int(1),
		PageSize: byteplus.Int(2),
		Filter: &model.ListContentGenerationTasksFilter{
			ServiceTier: byteplus.String("default"),
		},
	}

	listDefaultResp, err := client.ListContentGenerationTasks(ctx, listDefaultReq)
	if err != nil {
		fmt.Printf("failed to list default content generation tasks: %v\n", err)
	} else {
		fmt.Printf("List default returned %v results\n", listDefaultResp.Total)
		for _, item := range listDefaultResp.Items {
			fmt.Printf("Item ID: %s\n", item.ID)
			if item.ServiceTier != nil {
				fmt.Printf("  Service Tier: %s\n", *item.ServiceTier)
			}
			if item.ExecutionExpiresAfter != nil {
				fmt.Printf("  Execution Expires After: %d\n", *item.ExecutionExpiresAfter)
			}
		}
	}

	fmt.Println("----- delete default content generation task -----")

	deleteRequest := model.DeleteContentGenerationTaskRequest{ID: taskID}

	err = client.DeleteContentGenerationTask(ctx, deleteRequest)
	if err != nil {
		fmt.Printf("delete default content generation task error: %v\n", err)
	} else {
		fmt.Println("successfully deleted task id: ", taskID)
	}

	fmt.Println("----- create flex content generation task -----")
	createFlexReq := model.CreateContentGenerationTaskRequest{
		Model: modelEp, // Replace with your endpoint ID
		Content: []*model.CreateContentGenerationContentItem{
			{
				Type: model.ContentGenerationContentItemTypeText,
				Text: byteplus.String("A cat is sleeping --ratio 1:1"),
			},
			{
				Type: model.ContentGenerationContentItemTypeImage,
				ImageURL: &model.ImageURL{
					URL: "https://ark-project.tos-cn-beijing.volces.com/doc_image/see_i2v.jpeg", // Replace with URL
				},
			},
		},
		ServiceTier:           byteplus.String("flex"),
		ExecutionExpiresAfter: byteplus.Int64(3600), // 单位秒，示例为10分钟
	}

	createFlexResp, err := client.CreateContentGenerationTask(ctx, createFlexReq)
	if err != nil {
		fmt.Printf("create flex content generation error: %v\n", err)
		return
	}
	fmt.Printf("Flex Task Created with ID: %s\n", createFlexResp.ID)

	fmt.Println("----- get flex content generation task -----")
	flexTaskID := createFlexResp.ID

	getFlexReq := model.GetContentGenerationTaskRequest{ID: flexTaskID}

	getFlexResp, err := client.GetContentGenerationTask(ctx, getFlexReq)
	if err != nil {
		fmt.Printf("get flex content generation task error: %v\n", err)
		return
	}

	fmt.Printf("Flex Task ID: %s\n", getFlexResp.ID)
	if getFlexResp.ServiceTier != nil {
		fmt.Printf("Service Tier: %s\n", *getFlexResp.ServiceTier)
	}
	if getFlexResp.ExecutionExpiresAfter != nil {
		fmt.Printf("Execution Expires After: %d\n", *getFlexResp.ExecutionExpiresAfter)
	}
	if getFlexResp.Error != nil {
		fmt.Printf("Error Code: %s\n", getFlexResp.Error.Code)
		fmt.Printf("Error Message: %s\n", getFlexResp.Error.Message)
	}

	fmt.Println("----- list flex content generation tasks -----")

	listFlexReq := model.ListContentGenerationTasksRequest{
		PageNum:  byteplus.Int(1),
		PageSize: byteplus.Int(2),
		Filter: &model.ListContentGenerationTasksFilter{
			ServiceTier: byteplus.String("flex"),
		},
	}

	listFlexResp, err := client.ListContentGenerationTasks(ctx, listFlexReq)
	if err != nil {
		fmt.Printf("failed to list flex content generation tasks: %v\n", err)
	}

	fmt.Printf("List flex returned %v results\n", listFlexResp.Total)

	for _, item := range listFlexResp.Items {
		fmt.Printf("Item ID: %s\n", item.ID)
		if item.ServiceTier != nil {
			fmt.Printf("  Service Tier: %s\n", *item.ServiceTier)
		}
		if item.ExecutionExpiresAfter != nil {
			fmt.Printf("  Execution Expires After: %d\n", *item.ExecutionExpiresAfter)
		}
	}

	fmt.Println("----- delete flex content generation task -----")

	deleteFlexReq := model.DeleteContentGenerationTaskRequest{ID: flexTaskID}

	err = client.DeleteContentGenerationTask(ctx, deleteFlexReq)
	if err != nil {
		fmt.Printf("delete flex content generation task error: %v\n", err)
	} else {
		fmt.Println("successfully deleted flex task id: ", flexTaskID)
	}
}
