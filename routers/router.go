package routers

import (
	"property-listing-api/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// API v1 namespace
	ns := web.NewNamespace("/v1",
		// Property listing endpoint
		web.NSRouter("/properties/:location", &controllers.PropertyController{}),
	)

	// Register the namespace
	web.AddNamespace(ns)
}