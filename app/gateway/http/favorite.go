package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/app/gateway/wrapper"
	"ByteRhythm/idl/favorite/favoritePb"
	"ByteRhythm/util"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FavoriteActionHandler(ctx *gin.Context) {
	var req favoritePb.FavoriteActionRequest
	req.Token = ctx.Query("token")
	vid, _ := strconv.Atoi(ctx.Query("video_id"))
	req.VideoId = int64(vid)
	ActionType, _ := strconv.Atoi(ctx.Query("action_type"))
	req.ActionType = int32(ActionType)

	var res *favoritePb.FavoriteActionResponse

	hystrix.ConfigureCommand("FavoriteAction", wrapper.FavoriteActionFuseConfig)
	err := hystrix.Do("FavoriteAction", func() (err error) {
		res, err = rpc.FavoriteAction(ctx, &req)
		if err != nil {
			return err
		}
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
	})
}

func FavoriteListHandler(ctx *gin.Context) {
	var req favoritePb.FavoriteListRequest
	uid, _ := strconv.Atoi(ctx.Query("user_id"))
	req.UserId = int64(uid)
	req.Token = ctx.Query("token")

	var res *favoritePb.FavoriteListResponse

	hystrix.ConfigureCommand("FavoriteList", wrapper.FavoriteListFuseConfig)
	err := hystrix.Do("FavoriteList", func() (err error) {
		res, err = rpc.FavoriteList(ctx, &req)
		if err != nil {
			return err
		}
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
		"video_list":  res.VideoList,
	})
}
