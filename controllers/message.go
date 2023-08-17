package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/object"
	"ByteRhythm/utils"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

// MessageController operations for Message
type MessageController struct {
	baseController
}

func (c *MessageController) ChatMessage() {
	token := c.GetString("token")
	toUserId, _ := c.GetInt("to_user_id")

	fromUserId, _ := utils.GetUserIdFromToken(token)

	if fromUserId == toUserId {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "不能查看与自己的聊天记录",
		}
		c.ServeJSON()
		return
	}
	messageList, err := GetALLMessage(c, fromUserId, toUserId)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  err.Error(),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"status_code":  "0",
		"status_msg":   "获取聊天记录成功！",
		"message_list": messageList,
	}
	c.ServeJSON()

	return

}

func GetALLMessage(c *MessageController, fromUseId int, toUseId int) (messageList []object.MessageDto, err error) {
	// 构建 Redis 键
	redisKey := fmt.Sprintf("messages:%d:%d", fromUseId, toUseId)

	// 尝试从 Redis 缓存中获取数据
	redisResult, err := c.redisClient.Get(redisKey).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	// 如果缓存中存在数据，则直接返回缓存结果
	if redisResult != "" {
		// 解码 Redis 结果
		err = json.Unmarshal([]byte(redisResult), &messageList)
		if err != nil {
			return nil, err
		}
		return messageList, nil
	}

	// 缓存中不存在数据，则从数据库中查询
	var maps []orm.Params
	_, err = c.o.Raw(`select * from message where (from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?) `).
		SetArgs(fromUseId, toUseId, toUseId, fromUseId).Values(&maps)
	if err != nil {
		return nil, err
	}

	for i := range maps {

		message := maps[i]

		id, _ := strconv.Atoi(message["id"].(string))
		from, _ := strconv.Atoi(message["from_user_id"].(string))
		to, _ := strconv.Atoi(message["to_user_id"].(string))

		parseTime, _ := time.ParseInLocation("2006-01-02 15:04:05", message["create_time"].(string), time.Local)

		// 转化为我们需要的格式
		ms := object.MessageDto{Id: id, FromUserId: from, ToUserId: to, Content: message["content"].(string), CreateTime: int(parseTime.Unix())}

		messageList = append(messageList, ms)
	}

	// 将结果存入 Redis 缓存
	jsonBytes, err := json.Marshal(messageList)
	if err != nil {
		return nil, err
	}
	err = c.redisClient.Set(redisKey, string(jsonBytes), time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return messageList, nil
}

func (c *MessageController) ActionMessage() {
	token := c.GetString("token")
	fromUserId, err := utils.GetUserIdFromToken(token)
	if err != nil {
		c.handleError(err)
		return
	}

	actionType := c.GetString("action_type")
	if actionType == "1" {
		toUserId, err := strconv.Atoi(c.GetString("to_user_id"))
		if err != nil {
			c.handleError(err)
			return
		}

		content := c.GetString("content")

		user := &models.User{Id: fromUserId}
		if err = c.o.Read(user); err != nil {
			c.handleError(err)
			return
		}
		toUser := &models.User{Id: toUserId}
		if err = c.o.Read(toUser); err != nil {
			c.handleError(err)
			return
		}

		var message = models.Message{
			FromUserId: user,
			ToUserId:   toUser,
			Content:    content,
		}

		// 构建 Redis 键
		redisKey := fmt.Sprintf("messages:%d:%d", fromUserId, toUserId)

		// 尝试从 Redis 缓存中获取数据
		redisResult, err := c.redisClient.Get(redisKey).Result()
		if err != nil && err != redis.Nil {
			c.handleError(err)
			return
		}

		var messageList []object.MessageDto
		// 如果缓存中存在数据，则解码并合并到 messageList 中
		if redisResult != "" {
			// 解码 Redis 结果
			err = json.Unmarshal([]byte(redisResult), &messageList)
			if err != nil {
				c.handleError(err)
				return
			}

			// 将新的消息合并到 messageList 中
			newMessage := object.MessageDto{
				FromUserId: message.FromUserId.Id,
				ToUserId:   message.ToUserId.Id,
				Content:    message.Content,
			}
			messageList = append(messageList, newMessage)
		} else {
			// 如果缓存中不存在数据，则创建新的 messageList 切片
			messageList = []object.MessageDto{}
		}

		// 将结果存入 Redis 缓存
		jsonBytes, err := json.Marshal(messageList)
		if err != nil {
			c.handleError(err)
			return
		}

		err = c.redisClient.Set(redisKey, string(jsonBytes), time.Hour).Err()
		if err != nil {
			c.handleError(err)
			return
		}

		_, err = c.o.Insert(&message)
		if err != nil {
			c.handleError(err)
			return
		}
		c.Data["json"] = map[string]interface{}{
			"status_code": 0,
			"status_msg":  "发送消息成功！",
		}
		c.ServeJSON()
	} else {
		err := errors.New("非发送消息操作！")
		c.handleError(err)
		return
	}
}

func (c *MessageController) handleError(err error) {
	c.Data["json"] = map[string]interface{}{
		"status_code": 1,
		"status_msg":  err.Error(),
	}
	c.ServeJSON()
}
