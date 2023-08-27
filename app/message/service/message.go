package service

import (
	"ByteRhythm/app/message/dao"
	"ByteRhythm/idl/message/messagePb"
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type MessageSrv struct {
}

var MessageSrvIns *MessageSrv
var MessageSrvOnce sync.Once

// GetMessageSrv 懒汉式的单例模式 lazy-loading --> 懒汉式
func GetMessageSrv() *MessageSrv {
	MessageSrvOnce.Do(func() {
		MessageSrvIns = &MessageSrv{}
	})
	return MessageSrvIns
}

func (m MessageSrv) ChatMessage(ctx context.Context, req *messagePb.MessageChatRequest, res *messagePb.MessageChatResponse) error {
	token := req.Token
	toUserId := req.ToUserId

	fromUserId, err := util.GetUserIdFromToken(token)
	if err != nil {
		return nil
	}

	if int64(fromUserId) == toUserId {
		MessageChatResponseData(res, 1, "不能查看与自己的聊天记录！")
		return nil
	}

	// 构建 Redis 键
	var redisKey string
	if strconv.Itoa(fromUserId) < strconv.Itoa(int(toUserId)) {
		redisKey = "chat_messages:" + strconv.Itoa(fromUserId) + ":" + strconv.Itoa(int(toUserId))
	} else {
		redisKey = "chat_messages:" + strconv.Itoa(int(toUserId)) + ":" + strconv.Itoa(fromUserId)
	}

	// 尝试从 Redis 缓存中获取数据
	redisResult, err := dao.RedisClient.Get(ctx, redisKey).Result()
	if err != nil && err != redis.Nil {
		MessageChatResponseData(res, 1, "获取聊天记录失败！")
		return err
	}

	if redisResult != "" {
		// 如果缓存中存在数据，则解码并直接返回缓存数据
		err = json.Unmarshal([]byte(redisResult), &res.MessageList)
		if err != nil {
			MessageChatResponseData(res, 1, "获取聊天记录失败！")
			return err
		}

		MessageChatResponseData(res, 0, "获取聊天记录成功！")
		return nil
	}

	// 缓存中不存在数据，则从数据库获取聊天记录
	messages, err := dao.NewMessageDao(ctx).FindAllMessages(int64(fromUserId), toUserId)
	if err != nil {
		MessageChatResponseData(res, 1, "获取聊天记录失败！")
		return err
	}

	for _, message := range messages {
		res.MessageList = append(res.MessageList, BuildMessagePbModel(message))
	}

	// 将结果存入 Redis 缓存
	jsonBytes, err := json.Marshal(&res.MessageList)
	if err != nil {
		MessageChatResponseData(res, 1, "获取聊天记录失败！")
		return err
	}

	err = dao.RedisClient.Set(ctx, redisKey, string(jsonBytes), time.Hour).Err()
	if err != nil {
		MessageChatResponseData(res, 1, "获取聊天记录失败！")
		return err
	}

	MessageChatResponseData(res, 0, "获取聊天记录成功！")
	return nil
}

func (m MessageSrv) ActionMessage(ctx context.Context, req *messagePb.MessageActionRequest, res *messagePb.MessageActionResponse) error {
	token := req.Token
	fromUserID, err := util.GetUserIdFromToken(token)
	if err != nil {
		return err
	}

	actionType := req.ActionType
	if actionType == 1 {
		toUserID := req.ToUserId
		content := req.Content

		message := BuildMessageModel(fromUserID, int(toUserID), content)
		id, err := dao.NewMessageDao(ctx).CreateMessage(&message)
		if err != nil {
			MessageActionResponseData(res, 1, "发送消息失败！")
			return err
		}
		message.ID = uint(id)

		// 构建 Redis 键
		var redisKey string
		if strconv.Itoa(fromUserID) < strconv.Itoa(int(toUserID)) {
			redisKey = "chat_messages:" + strconv.Itoa(fromUserID) + ":" + strconv.Itoa(int(toUserID))
		} else {
			redisKey = "chat_messages:" + strconv.Itoa(int(toUserID)) + ":" + strconv.Itoa(fromUserID)
		}

		// 尝试从 Redis 缓存中获取数据
		redisResult, err := dao.RedisClient.Get(ctx, redisKey).Result()
		if err != nil && err != redis.Nil {
			MessageActionResponseData(res, 1, "操作失败！")
			return err
		}

		var messageList []*messagePb.Message
		// 如果缓存中存在数据，则解码并合并到 messageList 中
		if redisResult != "" {
			// 解码 Redis 结果
			err = json.Unmarshal([]byte(redisResult), &messageList)
			if err != nil {
				MessageActionResponseData(res, 1, "操作失败！")
				return err
			}
			messageList = append(messageList, BuildMessagePbModel(&message))
		} else {
			// 如果缓存中不存在数据，则创建新的 messageList 切片
			messageList = []*messagePb.Message{}
		}

		// 将结果存入 Redis 缓存
		jsonBytes, err := json.Marshal(&messageList)
		if err != nil {
			MessageActionResponseData(res, 1, "操作失败！")
			return err
		}

		err = dao.RedisClient.Set(ctx, redisKey, string(jsonBytes), time.Hour).Err()
		if err != nil {
			MessageActionResponseData(res, 1, "操作失败！")
			return err
		}

		MessageActionResponseData(res, 0, "发送消息成功！")
		return nil

	} else {
		MessageActionResponseData(res, 1, "非发送消息操作！")
		return nil
	}
}

func MessageChatResponseData(res *messagePb.MessageChatResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.MessageList = params[0].([]*messagePb.Message)
	}
}

func MessageActionResponseData(res *messagePb.MessageActionResponse, StatusCode int32, StatusMsg string) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
}

func BuildMessageModel(fromUserID int, toUserID int, content string) model.Message {
	return model.Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
		CreatedAt:  time.Now(),
	}
}

func BuildMessagePbModel(message *model.Message) *messagePb.Message {
	return &messagePb.Message{
		Id:         int64(message.ID),
		FromUserId: int64(message.FromUserID),
		ToUserId:   int64(message.ToUserID),
		Content:    message.Content,
		//CreateTime: message.CreatedAt.Format("2006-01-02 15:04:05"),
		//CreateTime: time.ParseInLocation("2006-01-02 15:04:05", message.CreatedAt, time.Local),
		CreateTime: "0",
	}
}
