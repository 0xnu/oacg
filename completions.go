package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

type CompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text    string  `json:"text"`
	Index   int     `json:"index"`
	LogProb float32 `json:"logprobs"`
}

func GetCompletion(apiKey string, prompt string, modelName string, maxTokens int, temperature float32) ([]string, error) {
	url := "https://api.openai.com/v1/completions"

	// Create completion request payload
	requestBody := CompletionRequest{
		Model:       modelName,
		Prompt:      prompt,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}
	jsonPayload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error encoding request payload: %v", err)
	}

	// Send completion request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get completion failed with status code %d", resp.StatusCode)
	}

	// Handle the completion response
	var response CompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	var completions []string
	for _, choice := range response.Choices {
		completions = append(completions, choice.Text)
	}

	return completions, nil
}
