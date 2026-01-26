package controllers

import (
	bee "github.com/beego/beego/v2/server/web"
)

// HomeController handles the root endpoint
type HomeController struct {
	bee.Controller
}

// Get handles GET requests to the root endpoint
func (c *HomeController) Get() {
	response := map[string]interface{}{
		"message": "Welcome to Property Listing API",
		"version": "v1",
		"endpoints": map[string]string{
			"properties": "/v1/properties/:location?items=true",
		},
		"documentation": "See README.md for full API documentation",
	}

	c.Data["json"] = response
	c.ServeJSON()
}