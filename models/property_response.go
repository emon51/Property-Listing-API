package models

// PropertyListingResponse represents the complete API response
type PropertyListingResponse struct {
	Items []PropertyItem `json:"Items"`
}

// PropertyItem represents a single property in the response
type PropertyItem struct {
	ID        string      `json:"ID"`
	Feed      int         `json:"Feed"`
	Published bool        `json:"Published"`
	GeoInfo   GeoInfo     `json:"GeoInfo"`
	Property  Property    `json:"Property"`
	Partner   Partner     `json:"Partner"`
}

// GeoInfo contains geographical information
type GeoInfo struct {
	Categories  []Category `json:"Categories"`
	City        string     `json:"City"`
	Country     string     `json:"Country"`
	CountryCode string     `json:"CountryCode"`
	Display     string     `json:"Display"`
	LocationID  string     `json:"LocationID"`
	Lat         string     `json:"Lat"`
	Lng         string     `json:"Lng"`
	Slug        string     `json:"Slug"`
}

// Category represents a geographical category
type Category struct {
	Name    string   `json:"Name"`
	Slug    string   `json:"Slug"`
	Type    string   `json:"Type"`
	Display []string `json:"Display"`
}

// Property contains property details
type Property struct {
	Amenities              map[string]string `json:"Amenities"`
	Counts                 Counts            `json:"Counts"`
	FeatureImage           string            `json:"FeatureImage"`
	IsPetFriendly          bool              `json:"IsPetFriendly"`
	MinStay                int               `json:"MinStay"`
	PropertyName           string            `json:"PropertyName"`
	PropertySlug           string            `json:"PropertySlug"`
	PropertyType           string            `json:"PropertyType"`
	PropertyTypeCategoryId string            `json:"PropertyTypeCategoryId"`
	RoomSize               float64           `json:"RoomSize"`
}

// Counts represents bedroom, bathroom, and occupancy counts
type Counts struct {
	Bedroom   int `json:"Bedroom"`
	Bathroom  int `json:"Bathroom"`
	Occupancy int `json:"Occupancy"`
}

// Partner contains partner/booking information
type Partner struct {
	ID           string   `json:"ID"`
	OwnerID      string   `json:"OwnerID"`
	Archived     []string `json:"Archived"`
	PropertyType string   `json:"PropertyType"`
	URL          string   `json:"URL"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}