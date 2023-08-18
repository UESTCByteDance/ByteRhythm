package router

import (
	"ByteRhythm/app/gateway/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/douyin")
	{
		v2 := v1.Group("/user")
		{
			v2.POST("/register/", http.RegisterHandler)
			v2.POST("/login/", http.LoginHandler)
			v2.GET("/", http.UserInfoHandler)
		}

		v2 = v1.Group("/feed")
		{
			v2.GET("/", http.FeedHandler)
		}

		v2 = v1.Group("/publish")
		{
			v2.POST("/action/", http.PublishHandler)
			v2.GET("/list/", http.PublishListHandler)
		}

	}
	return r

}
