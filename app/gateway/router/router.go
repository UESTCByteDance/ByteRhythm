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
	jaeger, closer, err := wrapper.NewJaegerTracer("HttpService", fmt.Sprintf("%s:%s", config.JaegerHost, config.JaegerPort))
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

	}
	return r

}
