# Property Listing API

A RESTful API service built with Beego (Go) that aggregates property listings from **two** external services. The API fetches property IDs 
by location and retrieves detailed property information concurrently for optimal performance.


## Project Structure
```
property-listing-api/
├── conf/                        # Configuration files
│   ├── app.conf                 # Main config (gitignored)
│   └── app.conf.example         # Config template
├── controllers/                 # Request handlers
│   └── property_controller.go   # Property endpoint controller
    └── home_controller.go       # Root endpoint controller
├── models/                      # Data models
│   └── property_response.go     # Response structures
├── routers/                     # Route definitions
│   └── router.go                # API routes with versioning
├── services/                    # Business logic
│   ├── http_client.go           # HTTP client wrapper
│   ├── location_service.go      # Location API integration
│   └── property_service.go      # Property API integration (concurrent)
├── utils/                       # Utility functions
│   └── validator.go             # Input validation
├── .gitignore                   # Git ignore rules
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies lock
├── main.go                      # Application entry point
└── README.md                    # This file
```

## Features

- Location-based property search
- Concurrent property details fetching
- API key authentication
- Input validation and comprehensive error handling
- Clean REST API design with versioning
- Data transformation from external APIs

## Tech Stack

- **Framework**: Beego v2
- **Language**: Go 1.21+
- **Architecture**: RESTful API with concurrent processing


## Getting Started

### Prerequisites

- Go 1.21 or higher
- Access to external property APIs

### Installation

1. Clone the repository or download the zip file of the project
```bash
git clone 
cd property-listing-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. To set configuration:
- Rename **conf/app.conf.example** to **conf/app.conf**
- Replace **api_key** with actual-api-key.

**Note:** 
- You may change the httpport if needed; it will not affect the main functionality.
- location_api_base_url and property_api_base_url are two External API Configuration
- Please feel free to reach out to me for the actual API key

4. Run the application:
```bash
go run main.go
```
or 
```bash
bee run
```

The API will be available at `http://localhost:5000` and you will able to see the home page.



## API Documentation

### Endpoint
```
GET /v1/properties/:location?items=true
```

### Request

**Path Parameter:**
- `location` (required): Colon-separated location (e.g., `usa` or `usa:florida` or `usa:florida:destin`)

**Query Parameter:**
- `items` (required): Must be `true`

**Headers:**
- `x-api-key` (required): Your API key

### Example Request
```bash
curl -X GET "http://localhost:5000/v1/properties/usa:florida:destin?items=true" \
  -H "x-api-key: actual-api-key"
```
**Or you can use postman to test the endpoint** 

### Success Response (200 OK)
```json
{
  "Items": [
    {
      "ID": "BC-12810439",
      "Feed": 11,
      "Published": true,
      "GeoInfo": {
        "Categories": [
          {
            "Name": "Nepal",
            "Slug": "nepal",
            "Type": "country",
            "Display": ["nepal"]
          }
        ],
        "City": "Nārāyangarh",
        "Country": "Central Development Region",
        "CountryCode": "NP",
        "Display": "Bharatpur, Nepal",
        "LocationID": "571",
        "Lat": "27.627859",
        "Lng": "84.40818",
        "Slug": "dhakal-villa"
      },
      "Property": {
        "Amenities": {
          "1": "Air Conditioner",
          "2": "Balcony/Terrace"
        },
        "Counts": {
          "Bedroom": 3,
          "Bathroom": 3,
          "Occupancy": 3
        },
        "FeatureImage": "dhakal-villa-np-n%C4%81r%C4%81yangarh-bc-12810439-0.jpg",
        "IsPetFriendly": true,
        "MinStay": 1,
        "PropertyName": "Dhakal Villa",
        "PropertySlug": "dhakal-villa",
        "PropertyType": "House",
        "PropertyTypeCategoryId": "6",
        "RoomSize": 1341.9
      },
      "Partner": {
        "ID": "12810439",
        "OwnerID": "",
        "Archived": [],
        "PropertyType": "Homestays",
        "URL": "https://www.booking.com/hotel/np/dhakal-villa.html?aid=affiliate_id"
      }
    }
  ]
}
```

```

### Error Responses

#### 400 Bad Request - Missing Parameters
```json
{
  "error": "location parameter is required"
}
```
```json
{
  "error": "items parameter must be true"
}
```

#### 401 Unauthorized - Invalid API Key
```json
{
  "error": "x-api-key header is required"
}
```
```json
{
  "error": "invalid API key"
}
```

#### 502 Bad Gateway - External API Error
```json
{
  "error": "failed to fetch property IDs from location service"
}
```
```json
{
  "error": "failed to fetch property details from property service"
}
```


## Key Features

### Concurrent Processing
The API fetches property details concurrently using goroutines, maintaining the original order of property IDs while maximizing performance.

### Data Transformation
External API responses are automatically transformed to match the standardized response format defined in the assignment.

### Comprehensive Validation
All inputs are validated with appropriate error messages:
- API key authentication
- Required parameters check
- Parameter value validation

### Proper HTTP Status Codes
- `200` - Success
- `400` - Bad Request (validation errors)
- `401` - Unauthorized (invalid API key)
- `502` - Bad Gateway (external API failures)
- `500` - Internal Server Error

## Development

### Running Tests
```bash
go test ./...
```

### Building the Binary
```bash
go build -o property-listing-api
```

### Running the Binary
```bash
./property-listing-api
```

## Configuration

All configuration is managed through `conf/app.conf`:

- `httpport`: Server port (default: 5000)
- `runmode`: dev/prod mode
- `location_api_base_url`: Location service endpoint
- `property_api_base_url`: Property service endpoint
- `api_key`: API authentication key




