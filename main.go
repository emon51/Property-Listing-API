package main

import (
	_ "property-listing-api/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Run the Beego application
	web.Run()
}