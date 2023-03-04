package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type EditRequest struct {
	Model       string `json:"model"`
	Input       string `json:"input"`
	Instruction string `json:"instruction"`
}

type EditResponse struct {
	Data struct {
		Object string `json:"object"`
		ID     string `json:"id"`
		Model  string `json:"model"`
		Result string `json:"result"`
	} `json:"data"`
}

func GetEdit(apiKey string, model string, input string, instruction string) (string, error) {
	client := &http.Client{}

	requestData := &EditRequest{
		Model:       model,
		Input:       input,
		Instruction: instruction,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to encode request data: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/edits", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response EditResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	return response.Data.Result, nil
}
