package oacg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListModelsResponse struct {
	Data []ModelInfo `json:"data"`
}

type ModelInfo struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	Created     int64  `json:"created"`
	ModelID     string `json:"model_id"`
	DisplayName string `json:"display_name"`
}

func ListModels(apiKey string) error {
	url := "https://api.openai.com/v1/models"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ListModels failed with status code %d", resp.StatusCode)
	}

	var response ListModelsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("error decoding response body: %v", err)
	}

	for _, model := range response.Data {
		fmt.Printf("%s (%s)\n", model.DisplayName, model.ModelID)
	}

	return nil
}
