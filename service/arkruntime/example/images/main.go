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

	fmt.Println("----- [Seedream] generate images (response format: url) -----")
	generateReq := model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your Seedream endpoint ID
		Prompt:         "Bird soaring above vast grasslands",
		ResponseFormat: byteplus.String(model.GenerateImagesResponseFormatURL),
		Seed:           byteplus.Int64(1234567890),
		Watermark:      byteplus.Bool(true),
		Size:           byteplus.String("512x512"),
		GuidanceScale:  byteplus.Float64(2.5),
	}

	imagesResponse, err := client.GenerateImages(ctx, generateReq)
	if err != nil {
		fmt.Printf("generate images error: %v\n", err)
		return
	}

	if imagesResponse.Error != nil {
		fmt.Printf("Error Code: %s\n", imagesResponse.Error.Code)
		fmt.Printf("Error Message: %s\n", imagesResponse.Error.Message)
	}

	fmt.Printf("Model: %s\n", imagesResponse.Model)
	fmt.Printf("Image URL: %s\n", *imagesResponse.Data[0].Url)
	fmt.Printf("Generated Images: %d\n", imagesResponse.Usage.GeneratedImages)
	fmt.Printf("Created: %d\n", imagesResponse.Created)

	fmt.Println("----- [Seedream] generate images (response format: base64) -----")
	generateReq = model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your Seedream endpoint ID
		Prompt:         "Bird soaring above vast grasslands",
		ResponseFormat: byteplus.String(model.GenerateImagesResponseFormatBase64),
		Seed:           byteplus.Int64(1234567890),
		Watermark:      byteplus.Bool(true),
		Size:           byteplus.String("512x512"),
		GuidanceScale:  byteplus.Float64(2.5),
	}

	imagesResponse, err = client.GenerateImages(ctx, generateReq)
	if err != nil {
		fmt.Printf("generate images error: %v\n", err)
		return
	}

	if imagesResponse.Error != nil {
		fmt.Printf("Error Code: %s\n", imagesResponse.Error.Code)
		fmt.Printf("Error Message: %s\n", imagesResponse.Error.Message)
	}

	fmt.Printf("Model: %s\n", imagesResponse.Model)
	fmt.Printf("Image URL: %s\n", *imagesResponse.Data[0].B64Json)
	fmt.Printf("Generated Images: %d\n", imagesResponse.Usage.GeneratedImages)
	fmt.Printf("Created: %d\n", imagesResponse.Created)

	generateReq = model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		Image:          byteplus.String("YOUR_IMAGE_URL_HERE"), // Replace with your input image URL
		ResponseFormat: byteplus.String(model.GenerateImagesResponseFormatURL),
		Seed:           byteplus.Int64(1234567890),
		Watermark:      byteplus.Bool(true),
		Size:           byteplus.String(model.GenerateImagesSizeAdaptive),
		GuidanceScale:  byteplus.Float64(2.5),
	}
	imagesResponse, err = client.GenerateImages(ctx, generateReq)
	if err != nil {
		fmt.Printf("generate images error: %v\n", err)
		return
	}
	if imagesResponse.Error != nil {
		fmt.Printf("Error Code: %s\n", imagesResponse.Error.Code)
		fmt.Printf("Error Message: %s\n", imagesResponse.Error.Message)
	}
	fmt.Printf("Model: %s\n", imagesResponse.Model)
	fmt.Printf("Image URL: %s\n", *imagesResponse.Data[0].Url)
	fmt.Printf("Generated Images: %d\n", imagesResponse.Usage.GeneratedImages)
	fmt.Printf("Created: %d\n", imagesResponse.Created)
	fmt.Println("----- [Seededit] generate images (input format: base64) -----")
	generateReq = model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your Seededit endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		Image:          byteplus.String("YOUR_IMAGE_BASE64_HERE"), // Replace with your input image base64 data url
		ResponseFormat: byteplus.String(model.GenerateImagesResponseFormatURL),
		Seed:           byteplus.Int64(1234567890),
		Watermark:      byteplus.Bool(true),
		Size:           byteplus.String(model.GenerateImagesSizeAdaptive),
		GuidanceScale:  byteplus.Float64(2.5),
	}
	imagesResponse, err = client.GenerateImages(ctx, generateReq)
	if err != nil {
		fmt.Printf("generate images error: %v\n", err)
		return
	}
	if imagesResponse.Error != nil {
		fmt.Printf("Error Code: %s\n", imagesResponse.Error.Code)
		fmt.Printf("Error Message: %s\n", imagesResponse.Error.Message)
	}
	fmt.Printf("Model: %s\n", imagesResponse.Model)
	fmt.Printf("Image URL: %s\n", *imagesResponse.Data[0].Url)
	fmt.Printf("Generated Images: %d\n", imagesResponse.Usage.GeneratedImages)
	fmt.Printf("Created: %d\n", imagesResponse.Created)
}
