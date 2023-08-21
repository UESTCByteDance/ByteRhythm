package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/app/gateway/wrapper"
	"ByteRhythm/idl/video/videoPb"
	"ByteRhythm/util"
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/gin-gonic/gin"
)

func FeedHandler(ctx *gin.Context) {
	var req videoPb.FeedRequest
	LatestTime, _ := strconv.Atoi(ctx.Query("latest_time"))
	req.LatestTime = int64(LatestTime)
	req.Token = ctx.Query("token")
	var res *videoPb.FeedResponse
	hystrix.ConfigureCommand("Feed", wrapper.FeedFuseConfig)
	err := hystrix.Do("Feed", func() (err error) {
		res, err = rpc.Feed(ctx, &req)
		return err
	}, func(err error) error {
		//降级处理
		wrapper.DefaultFeed(res)
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

func PublishHandler(ctx *gin.Context) {
	var req videoPb.PublishRequest
	req.Title = ctx.PostForm("title")
	req.Token = ctx.PostForm("token")
	//将获得的文件转为[]byte类型
	data, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FailRequest(err.Error()))
	}
	file, err := data.Open()
	defer file.Close()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FailRequest(err.Error()))
	}
	// 使用缓冲区逐块读取文件内容并写入 req.Data
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FailRequest(err.Error()))
		return
	}
	req.Data = buffer.Bytes()
	var res *videoPb.PublishResponse
	hystrix.ConfigureCommand("Publish", wrapper.PublishFuseConfig)
	err = hystrix.Do("Publish", func() (err error) {
		res, err = rpc.Publish(ctx, &req)
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

func PublishListHandler(ctx *gin.Context) {
	var req videoPb.PublishListRequest
	uid, _ := strconv.Atoi(ctx.Query("user_id"))
	req.UserId = int64(uid)
	req.Token = ctx.Query("token")
	var res *videoPb.PublishListResponse
	hystrix.ConfigureCommand("PublishList", wrapper.PublishListFuseConfig)
	err := hystrix.Do("PublishList", func() (err error) {
		res, err = rpc.PublishList(ctx, &req)
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
