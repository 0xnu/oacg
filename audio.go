package oacg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func audioRequest(apiKey string, audioFilePath string, model string, url string, responseType interface{}) error {
	file, err := os.Open(audioFilePath)
	if err != nil {
		return fmt.Errorf("failed to open audio file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("audio", file.Name())
	if err != nil {
		return fmt.Errorf("failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	if model != "" {
		if err := writer.WriteField("model", model); err != nil {
			return fmt.Errorf("failed to write field: %v", err)
		}
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status code %d, response body: %s", resp.StatusCode, string(responseBody))
	}

	if err := json.NewDecoder(resp.Body).Decode(responseType); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	return nil
}

func TranscribeAudio(apiKey string, audioFilePath string, model string) (*AudioTranscriptionResponse, error) {
	var response AudioTranscriptionResponse
	url := "https://api.openai.com/v1/audio/transcriptions"

	err := audioRequest(apiKey, audioFilePath, model, url, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func TranslateAudio(apiKey string, audioFilePath string, model string) (*AudioTranslationResponse, error) {
	var response AudioTranslationResponse
	url := "https://api.openai.com/v1/audio/translations"

	err := audioRequest(apiKey, audioFilePath, model, url, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
