package util

import "github.com/gin-gonic/gin"

func FailRequest(StatusMsg string) gin.H {
	return gin.H{
		"status_code": 1,
		"status_msg":  StatusMsg,
	}
}
