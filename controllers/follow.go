package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/object"
	"ByteRhythm/utils"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

// FollowController operations for Follow
type FollowController struct {
	baseController
}

func (c *FollowController) ActionRelation() {
	token := c.GetString("token")
	toUserId, _ := c.GetInt("to_user_id")
	actionType, _ := c.GetInt("action_type")
	// todo 用token中解析得到id
	fromUserId, _ := utils.GetUserIdFromToken(token)

	if actionType == 1 {
		// 关注
		_, err := AddFollow(c, fromUserId, toUserId)
		if err == nil {
			c.Data["json"] = map[string]interface{}{
				"status_code": 0,
				"status_msg":  "关注成功",
			}
			c.ServeJSON()
			return
		}

		c.Data["json"] = map[string]interface{}{
			"status_code": 0,
			"status_msg":  err.Error(),
		}
		c.ServeJSON()
		return

	} else if actionType == 2 {
		// 取消关注
		err := CancelFollow(c, fromUserId, toUserId)
		if err == nil {
			c.Data["json"] = map[string]interface{}{
				"status_code": 0,
				"status_msg":  "取消关注成功",
			}
			c.ServeJSON()
			return
		}

		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  err.Error(),
		}
		c.ServeJSON()
		return

	}

	c.Data["json"] = map[string]interface{}{
		"status_code": 1,
		"status_msg":  "检查参数信息",
	}
	c.ServeJSON()
	return

}

func (c *FollowController) ListFollowRelation() {
	//token := c.GetString("token")
	userId, _ := c.GetInt("user_id")
	followList, err := GetAllFollowByUserId(c, userId)
	if err != nil {

		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  err.Error(),
		}
		c.ServeJSON()
		return

	}

	c.Data["json"] = map[string]interface{}{
		"status_code": 0,
		"status_msg":  "成功！",
		"user_list":   followList,
	}
	c.ServeJSON()

	return

}

func (c *FollowController) ListFollowerRelation() {
	userId, err := strconv.Atoi(c.GetString("user_id"))
	if err != nil {
		c.handleError(err)
		return
	}
	token := c.GetString("token")
	if err := utils.ValidateToken(token); err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "token验证失败",
			"video_list":  nil,
		}
		c.ServeJSON()
		return
	}

	// 查询当前用户的粉丝关系
	var follows []models.Follow
	_, err = c.o.QueryTable("follow").Filter("user_id", userId).All(&follows)
	if err != nil {
		c.handleError(err)
		return
	}

	// 获取所有粉丝的被关注用户ID
	var userInfos []object.UserInfo
	for _, follow := range follows {
		userInfo := c.GetUserInfo(follow.FollowedUserId.Id)
		userInfos = append(userInfos, userInfo)
	}

	// 构建响应结果
	response := map[string]interface{}{
		"status_code": "0",
	}

	// 批量查询被关注用户的详细信息
	if len(userInfos) != 0 {
		response["status_msg"] = "获取粉丝列表成功"
		response["user_list"] = userInfos
	} else {
		response["status_msg"] = "粉丝列表为空"
		response["user_list"] = nil
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *FollowController) ListFriendRelation() {
	userId, err := strconv.Atoi(c.GetString("user_id"))
	if err != nil {
		c.handleError(err)
		return
	}

	token := c.GetString("token")
	if err := utils.ValidateToken(token); err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "token验证失败",
			"video_list":  nil,
		}
		c.ServeJSON()
		return
	}

	// 定义一个切片来存储多个粉丝关系查询结果
	var follows []models.Follow

	// 查询当前用户的粉丝关系
	_, err = c.o.QueryTable(new(models.Follow)).Filter("user_id", userId).All(&follows)
	if err != nil {
		c.handleError(err)
		return
	}

	// 获取所有粉丝的被关注用户ID
	var userInfos []object.UserInfo
	for _, follow := range follows {
		count, _ := c.o.QueryTable(new(models.Follow)).Filter("user_id", follow.FollowedUserId.Id).Filter("followed_user_id", userId).Count()
		if count == 1 {
			userInfo := c.GetUserInfo(follow.FollowedUserId.Id)
			userInfos = append(userInfos, userInfo)
		} else {
			continue
		}
	}

	// 构建响应结果
	response := map[string]interface{}{
		"status_code": "0",
	}

	// 批量查询被关注用户的详细信息
	if len(userInfos) != 0 {
		response["status_msg"] = "获取好友列表成功"
		response["user_list"] = userInfos
	} else {
		response["status_msg"] = "好友列表为空"
		response["user_list"] = nil
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *FollowController) handleError(err error) {
	c.Data["json"] = map[string]interface{}{
		"status_code": "1",
		"status_msg":  err.Error(),
		"user_list":   nil,
	}
	c.ServeJSON()
}

// AddFollow 关注
func AddFollow(c *FollowController, fromUseId int, toUseId int) (id int, err error) {
	user := c.o.QueryTable(new(models.User)).Filter("id", toUseId)
	// 被关注的用户未查询到
	if !user.Exist() {
		return -1, errors.New("关注的用户不存在！")
	}
	follow := models.Follow{UserId: &models.User{Id: toUseId}, FollowedUserId: &models.User{Id: fromUseId}, CreateTime: time.Now()}

	// 如果该关注不存在则关注
	fmt.Println(follow.UserId)
	if created, id, err := c.o.ReadOrCreate(&follow, "UserId", "FollowedUserId"); err == nil {
		if created {
			return int(id), nil
		} else {
			return int(id), errors.New("已经关注过该用户了！")
		}
	}

	return -1, err

}

// CancelFollow 取消关注
func CancelFollow(c *FollowController, fromUseId int, toUseId int) (err error) {
	user := c.o.QueryTable(new(models.User)).Filter("id", toUseId)
	// 被关注的用户未查询到
	if !user.Exist() {
		return errors.New("关注的用户不存在！")
	}

	follow := models.Follow{UserId: &models.User{Id: toUseId}, FollowedUserId: &models.User{Id: fromUseId}}

	// 还未关注该用户
	err = c.o.Read(&follow, "UserId", "FollowedUserId")
	if err != nil {
		return errors.New("您还未关注该用户！")
	}

	// 删除改关注
	_, err = c.o.Delete(&follow, "UserId", "FollowedUserId")
	if err != nil {
		return nil
	}

	return err

}

// GetAllFollowByUserId  获取关注列表
func GetAllFollowByUserId(c *FollowController, userId int) (followList []object.UserInfo, err error) {
	// 查询出被关注者的id集合
	var list []orm.ParamsList

	c.o.Raw(`select user_id from follow where followed_user_id = ?`).SetArgs(userId).ValuesList(&list)
	// 查出每个被关注者的相关信息
	for i := range list {
		id := list[i][0]
		id_int, _ := strconv.Atoi(id.(string))

		user := c.GetUserInfo(id_int)
		if err == nil {
			followList = append(followList, user)

		}

	}

	return followList, nil
}
