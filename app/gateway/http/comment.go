package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/app/gateway/wrapper"
	"ByteRhythm/idl/comment/commentPb"
	"ByteRhythm/util"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentActionHandler(ctx *gin.Context) {
	var req commentPb.CommentActionRequest
	req.Token = ctx.Query("token")
	vid, _ := strconv.Atoi(ctx.Query("video_id"))
	req.VideoId = int64(vid)
	ActionType, _ := strconv.Atoi(ctx.Query("action_type"))
	req.ActionType = int32(ActionType)
	CommentId, _ := strconv.Atoi(ctx.Query("comment_id"))
	req.CommentText = ctx.Query("comment_text")
	req.CommentId = int64(CommentId)

	var res *commentPb.CommentActionResponse

	hystrix.ConfigureCommand("CommentAction", wrapper.CommentActionFuseConfig)
	err := hystrix.Do("CommentAction", func() (err error) {
		res, err = rpc.CommentAction(ctx, &req)
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
		"comment":     res.Comment,
	})
}

func CommentListHandler(ctx *gin.Context) {
	var req commentPb.CommentListRequest
	req.Token = ctx.Query("token")
	vid, _ := strconv.Atoi(ctx.Query("video_id"))
	req.VideoId = int64(vid)

	var res *commentPb.CommentListResponse

	hystrix.ConfigureCommand("CommentList", wrapper.CommentListFuseConfig)
	err := hystrix.Do("CommentList", func() (err error) {
		res, err = rpc.CommentList(ctx, &req)
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
		"status_code":  res.StatusCode,
		"status_msg":   res.StatusMsg,
		"comment_list": res.CommentList,
	})
}
