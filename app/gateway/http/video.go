package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/idl/video/videoPb"
	"ByteRhythm/util"
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FeedHandler(ctx *gin.Context) {
	var req videoPb.FeedRequest
	LatestTime, _ := strconv.Atoi(ctx.Query("latest_time"))
	req.LatestTime = int64(LatestTime)
	req.Token = ctx.Query("token")
	res, err := rpc.Feed(ctx, &req)
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

	res, err := rpc.Publish(ctx, &req)
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
	res, err := rpc.PublishList(ctx, &req)
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
