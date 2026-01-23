package utils

import (
	"fmt"
)

// ValidateLocation checks if the location parameter is valid
func ValidateLocation(location string) error {
	if location == "" {
		return fmt.Errorf("location parameter is required")
	}
	return nil
}

// ValidateItems checks if the items parameter is true
func ValidateItems(items string) error {
	if items == "" {
		return fmt.Errorf("items parameter is required")
	}
	
	if items != "true" {
		return fmt.Errorf("items parameter must be true")
	}
	
	return nil
}

// ValidateAPIKey checks if the API key is valid
func ValidateAPIKey(providedKey, expectedKey string) error {
	if providedKey == "" {
		return fmt.Errorf("x-api-key header is required")
	}
	
	if providedKey != expectedKey {
		return fmt.Errorf("invalid API key")
	}
	
	return nil
}