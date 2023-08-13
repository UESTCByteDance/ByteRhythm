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
		web.NSNamespace("/favorite",
			web.NSRouter("/action/", &controllers.FavoriteController{}, "post:FavoriteAction"),
			web.NSRouter("/list/", &controllers.FavoriteController{}, "get:FavoriteList"),
		),
		web.NSNamespace("/comment",
			web.NSRouter("/action/", &controllers.CommentController{}, "post:CommentAction"),
			web.NSRouter("/list/", &controllers.CommentController{}, "get:CommentList"),
		),
		web.NSNamespace("/relation",
			web.NSRouter("/action", &controllers.FollowController{}, "post:ActionRelation"),
			web.NSRouter("/follow/list", &controllers.FollowController{}, "get:ListFollowRelation"),
			web.NSRouter("/follower/list", &controllers.FollowController{}, "get:ListFollowerRelation"),
			web.NSRouter("/friend/list", &controllers.FollowController{}, "get:ListFriendRelation"),
		),
		web.NSNamespace("/message",
			web.NSRouter("/chat", &controllers.MessageController{}, "get:ChatMessage"),
			web.NSRouter("/action", &controllers.MessageController{}, "post:ActionMessage"),
		),
	)
	web.AddNamespace(ns)
}
