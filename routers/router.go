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
		// Property listing endpoint
		bee.NSRouter("/properties/:location", &controllers.PropertyController{}),
	)

	// Register the namespace
	bee.AddNamespace(ns)
}