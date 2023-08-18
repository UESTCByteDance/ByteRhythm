package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/idl/video/videoPb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FeedHandler(ctx *gin.Context) {
	var req videoPb.FeedRequest
	res, err := rpc.Feed(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
		"video":       res.VideoList,
	})
}

func PublishHandler(ctx *gin.Context) {
	var req videoPb.PublishRequest
	res, err := rpc.Publish(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
	})
}

func PublishListHandler(ctx *gin.Context) {
	var req videoPb.PublishListRequest
	res, err := rpc.PublishList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
		"video":       res.VideoList,
	})
}
