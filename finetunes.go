package oacg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type FineTune struct {
	ID          string `json:"id"`
	ModelID     string `json:"model"`
	Status      string `json:"status"`
	TrainingLog string `json:"training_log"`
}

type FineTuneList struct {
	Data []FineTune `json:"data"`
}

type FineTuneResponse struct {
	Data FineTune `json:"data"`
}

type FineTuneEvents struct {
	Data []struct {
		Event      string `json:"event"`
		CreatedAt  int64  `json:"created_at"`
		Status     string `json:"status"`
		Percentage int    `json:"percentage"`
	} `json:"data"`
}

func CreateFineTune(apiKey string, trainingFileID string) (string, error) {
	url := "https://api.openai.com/v1/fine-tunes"

	data := fmt.Sprintf(`{
		"training_file": "%s"
	}`, trainingFileID)

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var response FineTuneResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return response.Data.ID, nil
}

func ListFineTunes(apiKey string) ([]FineTune, error) {
	url := "https://api.openai.com/v1/fine-tunes"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var response FineTuneList
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return response.Data, nil
}

func GetFineTune(apiKey string, fineTuneID string) (*FineTune, error) {
	url := fmt.Sprintf("https://api.openai.com/v1/fine-tunes/%s", fineTuneID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	var response FineTuneResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response.Data, nil
}

func CancelFineTune(apiKey string, fineTuneID string) error {
	url := fmt.Sprintf("https://api.openai.com/v1/fine-tunes/%s/cancel", fineTuneID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	return nil
}

func GetFineTuneEvents(apiKey string, fineTuneID string) (*FineTuneEvents, error) {
	url := fmt.Sprintf("https://api.openai.com/v1/fine-tunes/%s/events", fineTuneID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	var response FineTuneEvents
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}

func DeleteFineTune(apiKey string, fineTuneID string) error {
	url := fmt.Sprintf("https://api.openai.com/v1/models/curie:ft-%s", fineTuneID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	return nil
}
