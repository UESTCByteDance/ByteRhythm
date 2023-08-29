package middleware

import (
	"ByteRhythm/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Token string `json:"token" form:"token" binding:"required"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path != "/douyin/user/register" && c.Request.URL.Path != "/douyin/user/login" {
			var request Request
			c.ShouldBind(&request)
			token := request.Token
			if token != "" {
				if err := util.ValidateToken(token); err != nil {
					if c.Request.URL.Path != "/douyin/feed" {
						c.JSON(http.StatusForbidden, util.FailRequest("token验证失败，请重新登录"))
						c.Abort()
					}
				}
			}
		}
		c.Next()
	}
}
