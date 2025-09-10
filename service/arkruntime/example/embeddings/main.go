package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime"
	"github.com/byteplus-sdk/byteplus-go-sdk-v2/service/arkruntime/model"
)

func main() {
	client := arkruntime.NewClientWithApiKey(
		os.Getenv("ARK_API_KEY"),
	)
	ctx := context.Background()

	fmt.Println("----- embeddings request -----")
	req := model.EmbeddingRequestStrings{
		Input: []string{
			"花椰菜又称菜花、花菜，是一种常见的蔬菜。",
		},
		Model: "${YOUR_ENDPOINT_ID}",
	}

	resp, err := client.CreateEmbeddings(ctx, req)
	if err != nil {
		fmt.Printf("embeddings error: %v\n", err)
		return
	}

	s, _ := json.Marshal(resp)
	fmt.Println(string(s))
}
