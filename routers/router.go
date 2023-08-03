// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ByteRhythm/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/douyin",
		beego.NSNamespace("/user",
			beego.NSRouter("/", &controllers.UserController{}, "get:GetUser"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:LoginUser"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:RegisterUser"),
		),
		beego.NSNamespace("/feed",
			beego.NSRouter("/", &controllers.VideoController{}, "get:GetFeed"),
		),
		beego.NSNamespace("/publish",
			beego.NSRouter("/list", &controllers.VideoController{}, "get:ListPublish"),
			beego.NSRouter("/action", &controllers.VideoController{}, "post:ActionPublish"),
		),
		beego.NSNamespace("/favorite",
			beego.NSRouter("/list", &controllers.FavoriteController{}, "get:ListFavorite"),
			beego.NSRouter("/action", &controllers.FavoriteController{}, "post:ActionFavorite"),
		),
		beego.NSNamespace("/comment",
			beego.NSRouter("/list", &controllers.CommentController{}, "get:ListComment"),
			beego.NSRouter("/action", &controllers.CommentController{}, "post:ActionComment"),
		),
		beego.NSNamespace("/relation",
			beego.NSRouter("/action", &controllers.FollowController{}, "post:ActionRelation"),
			beego.NSRouter("/follow/list", &controllers.FollowController{}, "get:ListFollowRelation"),
			beego.NSRouter("/follower/list", &controllers.FollowController{}, "get:ListFollowerRelation"),
			beego.NSRouter("/friend/list", &controllers.FollowController{}, "get:ListFriendRelation"),
		),
		beego.NSNamespace("/message",
			beego.NSRouter("/chat", &controllers.MessageController{}, "get:ChatMessage"),
			beego.NSRouter("/action", &controllers.MessageController{}, "post:ActionMessage"),
		),
	)
	beego.AddNamespace(ns)
}
