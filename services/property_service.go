package services

import (
	"encoding/json"
	"fmt"
	"property-listing-api/models"
	"sync"

	"github.com/beego/beego/v2/server/web"
)

// PropertyService handles fetching property details
type PropertyService struct {
	httpClient *HTTPClient
	baseURL    string
}

// ExternalPropertyResponse represents the response from external API
type ExternalPropertyResponse struct {
	ID                     string                 `json:"id"`
	Feed                   int                    `json:"feed"`
	Published              bool                   `json:"published"`
	Categories             string                 `json:"categories"`
	City                   string                 `json:"city"`
	Country                string                 `json:"country"`
	CountryCode            string                 `json:"country_code"`
	Display                string                 `json:"display"`
	LocationID             string                 `json:"location_id"`
	Lonlat                 map[string]interface{} `json:"lonlat"`
	PropertySlug           string                 `json:"property_slug"`
	AmenityCategories      []string               `json:"amenity_categories"`
	BedroomCount           int                    `json:"bedroom_count"`
	BathroomCount          int                    `json:"bathroom_count"`
	Occupancy              int                    `json:"occupancy"`
	FeatureImage           string                 `json:"feature_image"`
	PropertyFlags          map[string]interface{} `json:"property_flags"`
	MinStay                int                    `json:"min_stay"`
	PropertyName           string                 `json:"property_name"`
	PropertyType           string                 `json:"property_type"`
	PropertyTypeCategory   string                 `json:"property_type_category"`
	RoomSizeSqft           float64                `json:"room_size_sqft"`
	FeedProviderID         string                 `json:"feed_provider_id"`
	OwnerID                string                 `json:"owner_id"`
	Archived               []string               `json:"archived"`
	FeedProviderURL        string                 `json:"feed_provider_url"`
	PropertyTypeCategoryID string                 `json:"property_type_category_id"`
}

// NewPropertyService creates a new property service instance
func NewPropertyService() *PropertyService {
	// Read base URL from configuration
	baseURL, _ := web.AppConfig.String("property_api_base_url")

	return &PropertyService{
		httpClient: NewHTTPClient(),
		baseURL:    baseURL,
	}
}

// GetPropertyDetails fetches details for a single property
func (s *PropertyService) GetPropertyDetails(propertyID string) (*models.PropertyItem, error) {
	// Build the full URL
	url := fmt.Sprintf("%s/%s", s.baseURL, propertyID)

	// Make the HTTP GET request
	body, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch property details: %w", err)
	}

	// Parse the external API response
	var externalResponse ExternalPropertyResponse
	if err := json.Unmarshal(body, &externalResponse); err != nil {
		return nil, fmt.Errorf("failed to parse property details response: %w", err)
	}

	// Transform to our model
	propertyItem := s.transformToPropertyItem(externalResponse)

	return &propertyItem, nil
}

// transformToPropertyItem transforms external API response to our model
func (s *PropertyService) transformToPropertyItem(external ExternalPropertyResponse) models.PropertyItem {
	// Parse categories JSON string
	var categories []models.Category
	if external.Categories != "" {
		json.Unmarshal([]byte(external.Categories), &categories)
	}

	// Extract lat/lng from lonlat
	lat := ""
	lng := ""
	if coords, ok := external.Lonlat["coordinates"].([]interface{}); ok && len(coords) == 2 {
		lng = fmt.Sprintf("%v", coords[0])
		lat = fmt.Sprintf("%v", coords[1])
	}

	// Build amenities map
	amenities := make(map[string]string)
	for i, amenity := range external.AmenityCategories {
		amenities[fmt.Sprintf("%d", i+1)] = amenity
	}

	// Get pet friendly flag
	isPetFriendly := false
	if external.PropertyFlags != nil {
		if val, ok := external.PropertyFlags["is_pet_friendly"].(bool); ok {
			isPetFriendly = val
		}
	}

	// Determine property type category ID (based on property type category)
	propertyTypeCategoryID := "6" // Default
	switch external.PropertyTypeCategory {
	case "House":
		propertyTypeCategoryID = "6"
	case "Apartment":
		propertyTypeCategoryID = "1"
	case "Villa":
		propertyTypeCategoryID = "2"
	}

	return models.PropertyItem{
		ID:        external.ID,
		Feed:      external.Feed,
		Published: external.Published,
		GeoInfo: models.GeoInfo{
			Categories:  categories,
			City:        external.City,
			Country:     external.Country,
			CountryCode: external.CountryCode,
			Display:     external.Display,
			LocationID:  external.LocationID,
			Lat:         lat,
			Lng:         lng,
			Slug:        external.PropertySlug,
		},
		Property: models.Property{
			Amenities:              amenities,
			Counts: models.Counts{
				Bedroom:   external.BedroomCount,
				Bathroom:  external.BathroomCount,
				Occupancy: external.Occupancy,
			},
			FeatureImage:           external.FeatureImage,
			IsPetFriendly:          isPetFriendly,
			MinStay:                external.MinStay,
			PropertyName:           external.PropertyName,
			PropertySlug:           external.PropertySlug,
			PropertyType:           external.PropertyTypeCategory,
			PropertyTypeCategoryId: propertyTypeCategoryID,
			RoomSize:               external.RoomSizeSqft,
		},
		Partner: models.Partner{
			ID:           external.FeedProviderID,
			OwnerID:      external.OwnerID,
			Archived:     external.Archived,
			PropertyType: external.PropertyType,
			URL:          external.FeedProviderURL,
		},
	}
}

// GetMultiplePropertyDetails fetches details for multiple properties concurrently
func (s *PropertyService) GetMultiplePropertyDetails(propertyIDs []string) ([]models.PropertyItem, error) {
	// Create a slice to store results in the same order as input
	results := make([]models.PropertyItem, len(propertyIDs))

	// WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Mutex to protect error collection
	var mu sync.Mutex
	var errors []error

	// Launch a goroutine for each property ID
	for i, propertyID := range propertyIDs {
		wg.Add(1)

		// Capture index and ID for the goroutine
		go func(index int, id string) {
			defer wg.Done()

			// Fetch property details
			propertyItem, err := s.GetPropertyDetails(id)
			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("failed to fetch property %s: %w", id, err))
				mu.Unlock()
				return
			}

			// Store result at the correct index to maintain order
			results[index] = *propertyItem
		}(i, propertyID)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Check if there were any errors
	if len(errors) > 0 {
		return nil, fmt.Errorf("encountered %d errors while fetching properties: %v", len(errors), errors[0])
	}

	return results, nil
}