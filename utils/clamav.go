package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type ClamAVScanResponse struct {
	Success bool                   `json:"success"`
	Data    ClamAVScanResponseData `json:"data"`
}

type ClamAVScanResponseData struct {
	Result []ClamAVScanResponseDataResult `json:"result"`
}

type ClamAVScanResponseDataResult struct {
	Name       string   `json:"name"`
	IsInfected bool     `json:"is_infected"`
	Viruses    []string `json:"viruses"`
}

func ClamAVScan(url string, files []*multipart.FileHeader) (*ClamAVScanResponse, error) {
	// Create a buffer to hold the multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add files to the multipart form
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		part, err := writer.CreateFormFile("FILES", file.Filename)
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(part, f); err != nil {
			return nil, err
		}
	}

	// Close the writer to finalize the multipart form
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	httpReq, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	// Check the response status
	if httpRes.StatusCode != http.StatusOK {
		body, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error: received status %s, body: %s", httpRes.Status, body)
	}

	// Decode the response into ClamAVScanResponse
	var resp ClamAVScanResponse
	if err = json.NewDecoder(httpRes.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &resp, nil
}
