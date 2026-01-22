package routers

import (
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &web.Controller{})
}