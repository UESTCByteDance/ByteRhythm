package http

import (
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/app/gateway/wrapper"
	"ByteRhythm/idl/message/messagePb"
	"ByteRhythm/util"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ActionMessageHandler(ctx *gin.Context) {
	var req messagePb.MessageActionRequest
	req.Token = ctx.Query("token")
	toUserId, _ := strconv.Atoi(ctx.Query("to_user_id"))
	req.ToUserId = int64(toUserId)
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	req.ActionType = int32(actionType)
	req.Content = ctx.Query("content")
	var res *messagePb.MessageActionResponse
	hystrix.ConfigureCommand("ActionMessage", wrapper.ActionMessageFuseConfig)
	err := hystrix.Do("ActionMessage", func() (err error) {
		res, err = rpc.ActionMessage(ctx, &req)
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

func ChatMessageHandler(ctx *gin.Context) {
	var req messagePb.MessageChatRequest
	req.Token = ctx.Query("token")
	toUserId, _ := strconv.Atoi(ctx.Query("to_user_id"))
	req.ToUserId = int64(toUserId)
	var res *messagePb.MessageChatResponse
	hystrix.ConfigureCommand("ChatMessage", wrapper.ChatMessageFuseConfig)
	err := hystrix.Do("ChatMessage", func() (err error) {
		res, err = rpc.ChatMessage(ctx, &req)
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
		"message_list": res.MessageList,
	})
}
