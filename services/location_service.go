package services

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

// LocationService handles fetching property IDs by location
type LocationService struct {
	httpClient *HTTPClient
	baseURL    string
}

// NewLocationService creates a new location service instance
func NewLocationService() *LocationService {
	// Read base URL from configuration
	baseURL, _ := web.AppConfig.String("location_api_base_url")

	return &LocationService{
		httpClient: NewHTTPClient(),
		baseURL:    baseURL,
	}
}

// GetPropertyIDs fetches property IDs for a given location
func (s *LocationService) GetPropertyIDs(location string) ([]string, error) {
	// Build the full URL
	url := fmt.Sprintf("%s/%s", s.baseURL, location)

	// Make the HTTP GET request
	body, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch property IDs: %w", err)
	}

	// Parse the JSON response
	var propertyIDs []string
	if err := json.Unmarshal(body, &propertyIDs); err != nil {
		return nil, fmt.Errorf("failed to parse property IDs response: %w", err)
	}

	return propertyIDs, nil
}


