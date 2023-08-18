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
	}
	return r

}
