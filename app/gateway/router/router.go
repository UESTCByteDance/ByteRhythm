package router

import (
	"ByteRhythm/app/gateway/http"
	"ByteRhythm/app/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//允许跨域请求
	r.Use(middleware.Cors())

	v1 := r.Group("/douyin")
	{
		v1.GET("/feed", http.FeedHandler)

		//jwt鉴权
		r.Use(middleware.JWT())

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

	}
	return r

}
