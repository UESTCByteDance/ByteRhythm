package rpc

import (
	"ByteRhythm/idl/message/messagePb"
	"context"
)

func ActionMessage(ctx context.Context, req *messagePb.MessageActionRequest) (res *messagePb.MessageActionResponse, err error) {
	res, err = MessageService.ActionMessage(ctx, req)
	if err != nil {
		return
	}
	return
}

func ChatMessage(ctx context.Context, req *messagePb.MessageChatRequest) (res *messagePb.MessageChatResponse, err error) {
	res, err = MessageService.ChatMessage(ctx, req)
	if err != nil {
		return
	}
	return
}
