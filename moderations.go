package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ModerationRequest struct {
	Input string `json:"input"`
}

type ModerationResponse struct {
	Data struct {
		Object string  `json:"object"`
		ID     string  `json:"id"`
		Model  string  `json:"model"`
		Score  float32 `json:"score"`
	} `json:"data"`
}

func GetModeration(apiKey string, input string) (float32, error) {
	client := &http.Client{}

	requestData := &ModerationRequest{
		Input: input,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return -1.0, fmt.Errorf("failed to encode request data: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/moderations", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		return -1.0, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return -1.0, fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1.0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response ModerationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return -1.0, fmt.Errorf("failed to decode response body: %v", err)
	}

	return response.Data.Score, nil
}
