package controllers

import (
	"property-listing-api/models"
	"property-listing-api/services"
	"property-listing-api/utils"

	"github.com/beego/beego/v2/server/web"
)

// PropertyController handles property listing requests
type PropertyController struct {
	web.Controller
}

// Get handles GET requests to fetch properties
func (c *PropertyController) Get() {
	// Read API key from configuration
	expectedAPIKey, _ := web.AppConfig.String("api_key")

	// Get API key from header
	providedAPIKey := c.Ctx.Input.Header("x-api-key")

	// Validate API key
	if err := utils.ValidateAPIKey(providedAPIKey, expectedAPIKey); err != nil {
		c.sendErrorResponse(401, err.Error())
		return
	}

	// Get location from path parameter
	location := c.Ctx.Input.Param(":location")

	// Validate location (check if empty) - return 400
	if err := utils.ValidateLocation(location); err != nil {
		c.sendErrorResponse(400, err.Error())
		return
	}

	// Get items query parameter
	items := c.GetString("items")

	// Validate items parameter
	if err := utils.ValidateItems(items); err != nil {
		c.sendErrorResponse(400, err.Error())
		return
	}

	// Fetch property IDs from location service
	locationService := services.NewLocationService()
	propertyIDs, err := locationService.GetPropertyIDs(location)
	if err != nil {
		c.sendErrorResponse(502, "failed to fetch property IDs from location service")
		return
	}

	// Fetch property details concurrently
	propertyService := services.NewPropertyService()
	properties, err := propertyService.GetMultiplePropertyDetails(propertyIDs)
	if err != nil {
		c.sendErrorResponse(502, "failed to fetch property details from property service")
		return
	}

	// Build the response
	response := models.PropertyListingResponse{
		Items: properties,
	}

	// Send success response
	c.Data["json"] = response
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
}

// sendErrorResponse sends an error response with the specified status code
func (c *PropertyController) sendErrorResponse(statusCode int, message string) {
	errorResponse := models.ErrorResponse{
		Error: message,
	}
	c.Data["json"] = errorResponse
	c.Ctx.Output.SetStatus(statusCode)
	c.ServeJSON()
}