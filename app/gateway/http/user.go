package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/idl/user/userPb"
	"ByteRhythm/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {
	var req userPb.UserRequest
	req.Username = ctx.Query("username")
	req.Password = ctx.Query("password")
	res, err := rpc.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FailRequest(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
		"user_id":     res.UserId,
		"token":       res.Token,
	})
}

func LoginHandler(ctx *gin.Context) {
	var req userPb.UserRequest
	req.Username = ctx.Query("username")
	req.Password = ctx.Query("password")
	res, err := rpc.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FailRequest(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
		"user_id":     res.UserId,
		"token":       res.Token,
	})
}

func UserInfoHandler(ctx *gin.Context) {
	var req userPb.UserInfoRequest
	uid, _ := strconv.Atoi(ctx.Query("user_id"))
	req.UserId = int64(uid)
	req.Token = ctx.Query("token")
	res, err := rpc.UserInfo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FailRequest(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
		"user":        res.User,
	})
}
