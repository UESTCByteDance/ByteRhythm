package middleware

import (
	"ByteRhythm/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token != "" {
			if err := util.ValidateToken(token); err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"status_code": 1,
					"status_msg":  "token验证失败",
				})
				c.Abort()
			}
		}
		c.Next()
	}
}
