package oacg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Model struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	CreatedAt   int64  `json:"created"`
	ModelType   string `json:"model_type"`
	Description string `json:"description"`
}

type ModelListResponse struct {
	Data []Model `json:"data"`
}

func ListModels(apiKey string) ([]Model, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.openai.com/v1/models", nil)
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

	var response ModelListResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return response.Data, nil
}

func GetModel(apiKey string, modelID string) (*Model, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openai.com/v1/models/%s", modelID), nil)
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

	var model Model
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &model, nil
}
