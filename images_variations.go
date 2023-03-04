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
	"strconv"
)

type ImageVariationResponse struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	CreatedAt  int64  `json:"created_at"`
	Model      string `json:"model"`
	Iterations int    `json:"iterations"`
	Url        string `json:"url"`
}

func VariationsImages(apiKey string, imagePath string, n int, size string) (*ImageVariationResponse, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file: %v", err)
	}

	nString := strconv.Itoa(n)
	if err := writer.WriteField("n", nString); err != nil {
		return nil, fmt.Errorf("failed to write field: %v", err)
	}

	if size != "" {
		if err := writer.WriteField("size", size); err != nil {
			return nil, fmt.Errorf("failed to write field: %v", err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/variations", body)
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
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status code %d, response body: %s", resp.StatusCode, string(responseBody))
	}

	var response struct {
		Data ImageVariationResponse `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response.Data, nil
}
