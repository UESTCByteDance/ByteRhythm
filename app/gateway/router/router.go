package router

import (
	"ByteRhythm/app/gateway/http"
	"ByteRhythm/app/gateway/middleware"
	"ByteRhythm/app/gateway/wrapper"
	"ByteRhythm/config"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
)

func NewRouter() *gin.Engine {
	config.Init()
	r := gin.Default()
	jaeger, closer, err := wrapper.InitJaeger("HttpService", fmt.Sprintf("%s:%s", config.JaegerHost, config.JaegerPort))
	defer closer.Close()
	if err != nil {
		logger.Info("HttpService init jaeger failed, err:", err)
	}
	r.Use(
		middleware.JWT(),
		cors.Default(),
		middleware.Jaeger(jaeger),
	)
	v1 := r.Group("/douyin")
	{
		v1.GET("/feed", http.FeedHandler)

		v2 := v1.Group("/user")
		{
			v2.POST("/register/", http.RegisterHandler)
			v2.POST("/login/", http.LoginHandler)
			v2.GET("/", http.UserInfoHandler)
		}

		v2 = v1.Group("/publish")
		{
			v2.POST("/action/", http.PublishHandler)
			v2.GET("/list/", http.PublishListHandler)
		}

		v2 = v1.Group("/relation")
		{
			v2.POST("/action/", http.ActionRelationHandler)
			v2.GET("/follow/list/", http.ListFollowRelationHandler)
			v2.GET("/follower/list/", http.ListFollowerRelationHandler)
			v2.GET("/friend/list/", http.ListFriendRelationHandler)
		}

		v2 = v1.Group("/message")
		{
			v2.POST("/action/", http.ActionMessageHandler)
			v2.GET("/chat/", http.ChatMessageHandler)
		}

		v2 = v1.Group("/comment")
		{
			v2.POST("/action/", http.CommentActionHandler)
			v2.GET("/list/", http.CommentListHandler)
		}

		v2 = v1.Group("/favorite")
		{
			v2.POST("/action/", http.FavoriteActionHandler)
			v2.GET("/list/", http.FavoriteListHandler)
		}

	}
	return r

}
