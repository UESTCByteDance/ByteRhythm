package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"my-tiktok/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
