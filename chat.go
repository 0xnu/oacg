package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Text          string    `json:"text"`
		Index         int       `json:"index"`
		Logprobs      []Logprob `json:"logprobs"`
		FinishReason  string    `json:"finish_reason"`
		SelectedToken int       `json:"selected_token"`
	} `json:"choices"`
}

func GetChatCompletion(apiKey string, model string, messages []ChatMessage) (string, error) {
	client := &http.Client{}

	requestData := &ChatRequest{
		Model:    model,
		Messages: messages,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to encode request data: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestDataBytes))
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

	var response ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no completions returned")
	}

	return response.Choices[0].Text, nil
}
