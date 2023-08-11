package routers

import (
	"ByteRhythm/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/douyin",
		web.NSNamespace("/user",
			web.NSRouter("/", &controllers.UserController{}, "get:Info"),
			web.NSRouter("/login/", &controllers.UserController{}, "post:Login"),
			web.NSRouter("/register/", &controllers.UserController{}, "post:Register"),
		),
		web.NSNamespace("/feed",
			web.NSRouter("/", &controllers.VideoController{}, "get:Feed"),
		),
		web.NSNamespace("/publish",
			web.NSRouter("/list/", &controllers.VideoController{}, "get:List"),
			web.NSRouter("/action/", &controllers.VideoController{}, "post:Publish"),
		),
	)
	web.AddNamespace(ns)
}
