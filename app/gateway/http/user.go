package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/idl/pb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	req.Username = ctx.Query("username")
	req.Password = ctx.Query("password")
	Res, err := rpc.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": Res.StatusCode,
		"status_msg":  Res.StatusMsg,
		"user_id":     Res.UserId,
		"token":       Res.Token,
	})
}

func LoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	req.Username = ctx.Query("username")
	req.Password = ctx.Query("password")
	Res, err := rpc.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": Res.StatusCode,
		"status_msg":  Res.StatusMsg,
		"user_id":     Res.UserId,
		"token":       Res.Token,
	})
}

func UserInfoHandler(ctx *gin.Context) {
	var req pb.UserInfoRequest
	uid, _ := strconv.Atoi(ctx.Query("user_id"))
	req.UserId = int64(uid)
	req.Token = ctx.GetString("token")
	Res, err := rpc.UserInfo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": Res.StatusCode,
		"status_msg":  Res.StatusMsg,
		"user":        Res.User,
	})
}
