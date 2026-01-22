package services

import (
	"io"
	"net/http"
	"time"
)

// HTTPClient is a wrapper around http.Client for making HTTP requests
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTP client with timeout
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 10 * time.Second, // 10 second timeout for requests
		},
	}
}

// Get performs a GET request to the specified URL
func (c *HTTPClient) Get(url string) ([]byte, error) {
	// Create the GET request
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}