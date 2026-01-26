package main

import (
	_ "property-listing-api/routers"
	bee "github.com/beego/beego/v2/server/web"
)

func main() {
	// Run the Beego application
	bee.Run()
}