package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CompletionRequest struct {
	Model            string                 `json:"model"`
	Prompt           string                 `json:"prompt"`
	MaxTokens        int                    `json:"max_tokens,omitempty"`
	Temperature      float32                `json:"temperature,omitempty"`
	TopP             float32                `json:"top_p,omitempty"`
	N                int                    `json:"n,omitempty"`
	Stream           bool                   `json:"stream,omitempty"`
	Logprobs         int                    `json:"logprobs,omitempty"`
	Echo             bool                   `json:"echo,omitempty"`
	Stop             []string               `json:"stop,omitempty"`
	PresencePenalty  float32                `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                `json:"frequency_penalty,omitempty"`
	BestOf           int                    `json:"best_of,omitempty"`
	LogitBias        map[string]interface{} `json:"logit_bias,omitempty"`
}

type CompletionResponse struct {
	Choices []struct {
		Text         string    `json:"text"`
		Index        int       `json:"index"`
		Logprobs     []Logprob `json:"logprobs"`
		FinishReason string    `json:"finish_reason"`
	} `json:"choices"`
}

type Logprob struct {
	Text  string  `json:"text"`
	Value float32 `json:"value"`
}

func GetCompletion(apiKey string, model string, prompt string, maxTokens int, temperature float32) (string, error) {
	client := &http.Client{}

	requestData := &CompletionRequest{
		Model:       model,
		Prompt:      prompt,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to encode request data: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestDataBytes))
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

	var response CompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no completions returned")
	}

	return response.Choices[0].Text, nil
}
