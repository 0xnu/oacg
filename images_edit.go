package oacg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type ImageEditRequest struct {
	Image  *os.File
	Mask   *os.File
	Prompt string
	N      int
	Size   string
}

type ImageEditResponse struct {
	Data []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func EditImages(apiKey string, request *ImageEditRequest) (*ImageEditResponse, error) {
	// Create a new multipart writer
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Add the image file
	imageFile, err := request.Image.Stat()
	if err != nil {
		return nil, err
	}
	image, err := writer.CreateFormFile("image", imageFile.Name())
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(image, request.Image)
	if err != nil {
		return nil, err
	}

	// Add the mask file
	maskFile, err := request.Mask.Stat()
	if err != nil {
		return nil, err
	}
	mask, err := writer.CreateFormFile("mask", maskFile.Name())
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(mask, request.Mask)
	if err != nil {
		return nil, err
	}

	// Add the prompt
	err = writer.WriteField("prompt", request.Prompt)
	if err != nil {
		return nil, err
	}

	// Add the n value
	err = writer.WriteField("n", string(request.N))
	if err != nil {
		return nil, err
	}

	// Add the size value
	err = writer.WriteField("size", request.Size)
	if err != nil {
		return nil, err
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request
	url := "https://api.openai.com/v1/images/edits"
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Make the API request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check for errors
	if resp.StatusCode != 200 {
		return nil, errors.New(string(responseBody))
	}

	// Parse the response
	var response ImageEditResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
