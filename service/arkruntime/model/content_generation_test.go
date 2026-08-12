package model

import (
	"encoding/json"
	"testing"
)

func TestCreateContentGenerationTaskRequestMarshalIncludesOptionalFields(t *testing.T) {
	outputFormat := "mp4"
	omniReferenceTaskType := "image"
	request := CreateContentGenerationTaskRequest{
		Model:                 "test-model",
		Content:               []*CreateContentGenerationContentItem{},
		OutputFormat:          &outputFormat,
		OmniReferenceTaskType: &omniReferenceTaskType,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	if got := body["output_format"]; got != outputFormat {
		t.Errorf("output_format = %v, want %q", got, outputFormat)
	}
	if got := body["omni_reference_task_type"]; got != omniReferenceTaskType {
		t.Errorf("omni_reference_task_type = %v, want %q", got, omniReferenceTaskType)
	}
}

func TestCreateContentGenerationTaskRequestMarshalOmitsUnsetOptionalFields(t *testing.T) {
	request := CreateContentGenerationTaskRequest{
		Model:   "test-model",
		Content: []*CreateContentGenerationContentItem{},
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	if _, ok := body["output_format"]; ok {
		t.Error("output_format should be omitted when unset")
	}
	if _, ok := body["omni_reference_task_type"]; ok {
		t.Error("omni_reference_task_type should be omitted when unset")
	}
}

func TestGetContentGenerationTaskResponseUnmarshalOutputFormat(t *testing.T) {
	var response GetContentGenerationTaskResponse
	if err := json.Unmarshal([]byte(`{"output_format":"mp4"}`), &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if response.OutputFormat == nil {
		t.Fatal("OutputFormat is nil")
	}
	if got, want := *response.OutputFormat, "mp4"; got != want {
		t.Errorf("OutputFormat = %q, want %q", got, want)
	}
}
