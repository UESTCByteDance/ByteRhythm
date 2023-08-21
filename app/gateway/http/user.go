package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/app/gateway/wrapper"
	"ByteRhythm/idl/user/userPb"
	"ByteRhythm/util"
	"net/http"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(ctx *gin.Context) {
	var req userPb.UserRequest
	req.Username = ctx.Query("username")
	req.Password = ctx.Query("password")
	var res *userPb.UserResponse
	hystrix.ConfigureCommand("Register", wrapper.RegisterFuseConfig)
	err := hystrix.Do("Register", func() (err error) {
		res, err = rpc.Register(ctx, &req)
		return err
	}, func(err error) error {
		return err
	})
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
	var res *userPb.UserResponse
	hystrix.ConfigureCommand("Login", wrapper.LoginFuseConfig)
	err := hystrix.Do("Login", func() (err error) {
		res, err = rpc.Login(ctx, &req)
		return err
	}, func(err error) error {
		return err
	})
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
	var res *userPb.UserInfoResponse
	hystrix.ConfigureCommand("UserInfo", wrapper.UserInfoFuseConfig)
	err := hystrix.Do("UserInfo", func() (err error) {
		res, err = rpc.UserInfo(ctx, &req)
		return err
	}, func(err error) error {
		return err
	})
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
