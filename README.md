# Property Listing API

A RESTful API service built with Beego (Go) that aggregates property listings from **two** external services. The API fetches property IDs 
by location and retrieves detailed property information concurrently for optimal performance.

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

3. Copy the configuration template:
```bash
cp conf/app.conf.example conf/app.conf
```

4. Update `conf/app.conf` with your actual configuration:
```ini
appname = property-listing-api
httpport = 5000
runmode = dev

location_api_base_url = http://192.168.0.35:8099/api
property_api_base_url = http://192.168.0.35:8099/api

api_key = actual-api-key
```
**Note:** 
- You may change the httpport if needed; it will not affect the main functionality.
- location_api_base_url and property_api_base_url are two External API Configuration
- Please feel free to reach out to me for the actual API key

5. Run the application:
```bash
go run main.go
```
or 
```bash
bee run
```


The API will be available at `http://localhost:5000`

