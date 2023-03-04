package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmbeddingsRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type EmbeddingsResponse struct {
	Data []float32 `json:"data"`
}

func GetEmbeddings(apiKey string, model string, input string) ([]float32, error) {
	client := &http.Client{}

	requestData := &EmbeddingsRequest{
		Model: model,
		Input: input,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request data: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response EmbeddingsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return response.Data, nil
}
