package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ImageGenerationRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n,omitempty"`
	Size   string `json:"size,omitempty"`
	Model  string `json:"model,omitempty"`
}

type ImageGenerationResponse struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	CreatedAt  int64  `json:"created_at"`
	Model      string `json:"model"`
	Iterations int    `json:"iterations"`
	Url        string `json:"url"`
}

func GenerateImages(apiKey string, prompt string, n int, size string) (*ImageGenerationResponse, error) {
	requestBody, err := json.Marshal(&ImageGenerationRequest{
		Prompt: prompt,
		N:      n,
		Size:   size,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	var response struct {
		Data ImageGenerationResponse `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response.Data, nil
}
