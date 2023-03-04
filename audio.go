package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type AudioTranscriptionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Text    string `json:"text"`
	Status  string `json:"status"`
	JobID   string `json:"job_id"`
	JobType string `json:"job_type"`
}

type AudioTranslationResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Text    string `json:"text"`
	Status  string `json:"status"`
	JobID   string `json:"job_id"`
	JobType string `json:"job_type"`
}

func TranscribeAudio(apiKey string, audioFilePath string, model string) (*AudioTranscriptionResponse, error) {
	url := "https://api.openai.com/v1/audio/transcriptions"
	responseType := &AudioTranscriptionResponse{}
	return audioRequest(apiKey, audioFilePath, model, url, responseType)
}

func TranslateAudio(apiKey string, audioFilePath string, model string) (*AudioTranslationResponse, error) {
	url := "https://api.openai.com/v1/audio/translations"
	responseType := &AudioTranslationResponse{}
	return audioRequest(apiKey, audioFilePath, model, url, responseType)
}

func audioRequest(apiKey string, audioFilePath string, model string, url string, responseType interface{}) (interface{}, error) {
	file, err := os.Open(audioFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", audioFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err)
	}
	io.Copy(part, file)
	writer.WriteField("model", model)
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", writer.FormDataContentType())

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
		Data interface{} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if err := json.Unmarshal(response.Data.(json.RawMessage), responseType); err != nil {
		return nil, fmt.Errorf("failed to decode response data: %v", err)
	}

	return responseType, nil
}
