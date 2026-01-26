package routers

import (
	"property-listing-api/controllers"

	bee "github.com/beego/beego/v2/server/web"
)

func init() {
	// Root endpoint - greeting message
	bee.Router("/", &controllers.HomeController{})

	// API v1 namespace
	ns := bee.NewNamespace("/v1",
		// Property listing endpoint with location
		bee.NSRouter("/properties/:location", &controllers.PropertyController{}),
		// Property listing endpoint without location (handles empty location)
		bee.NSRouter("/properties", &controllers.PropertyController{}),
		bee.NSRouter("/properties/", &controllers.PropertyController{}),
	)

	// Register the namespace
	bee.AddNamespace(ns)
}