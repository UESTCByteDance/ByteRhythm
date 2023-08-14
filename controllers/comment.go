package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/object"
	"ByteRhythm/utils"
	"strconv"
)

type CommentController struct {
	baseController
}

// 评论操作
func (c *CommentController) CommentAction() {
	// 获取必要参数
	tokenString := c.GetString("token")                       // 用户鉴权token
	videoId, _ := strconv.Atoi(c.GetString("video_id"))       // 视频id
	actionType, _ := strconv.Atoi(c.GetString("action_type")) // 1-发布评论，2-删除评论
	commentContext := c.GetString("comment_text")             // 评论内容
	commentId, _ := strconv.Atoi(c.GetString("comment_id"))   // 评论id

	// 鉴权
	if err := utils.ValidateToken(tokenString); err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "token鉴权失败",
			"comment":     nil,
		}
		c.ServeJSON()
		return
	}

	// 解析token
	user, err := utils.GetUserFromToken(tokenString)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "获取用户信息失败",
			"comment":     nil,
		}
		c.ServeJSON()
		return
	}

	// 发布评论
	if actionType == 1 {
		var video models.Video
		c.o.QueryTable(new(models.Video)).Filter("id", videoId).One(&video)
		comment := models.Comment{
			UserId:  user,
			VideoId: &video,
			Content: commentContext,
		}
		if _, err := c.o.Insert(&comment); err != nil {
			c.Data["json"] = map[string]interface{}{
				"status_code": 1,
				"status_msg":  "发布评论失败",
				"comment":     nil,
			}
			c.ServeJSON()
			return
		}
		c.Data["json"] = map[string]interface{}{
			"status_code": 0,
			"status_msg":  "发布评论成功",
			"comment":     comment,
		}
		c.ServeJSON()
		return
	}
	// 删除评论
	//var comment models.Comment
	//c.o.QueryTable(new(models.Comment)).Filter("id", commentId).One(&comment)
	//c.o.Delete(comment)
	c.o.Delete(&models.Comment{Id: commentId}, "Id")
	c.Data["json"] = map[string]interface{}{
		"status_code": 0,
		"status_msg":  "删除评论成功",
		"comment":     nil,
	}
	c.ServeJSON()
}

// CommentList 视频评论列表
func (c *CommentController) CommentList() {
	// 获取必要参数
	tokenString := c.GetString("token")                 // 用户鉴权token
	videoId, _ := strconv.Atoi(c.GetString("video_id")) // 视频id
	// 鉴权
	if err := utils.ValidateToken(tokenString); err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "token鉴权失败",
		}
		c.ServeJSON()
		return
	}
	// 解析token
	user, err := utils.GetUserFromToken(tokenString)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "获取用户信息失败",
			"video_list":  nil,
		}
		c.ServeJSON()
		return
	}
	// 查询所有评论
	var comments []models.Comment
	var commentInfos []object.CommentInfo
	c.o.QueryTable(new(models.Comment)).Filter("video_id", videoId).OrderBy("-create_time").All(&comments)
	for _, comm := range comments {
		userInfo := c.GetUserInfo(user.Id)
		commentInfo := object.CommentInfo{
			Content:    comm.Content,
			CreateDate: comm.CreateTime.Format("2006-01-02 15:04"),
			ID:         int64(comm.Id),
			User:       userInfo,
		}
		commentInfos = append(commentInfos, commentInfo)
	}
	c.Data["json"] = map[string]interface{}{
		"status_code":  "0",
		"status_msg":   "获取评论列表成功",
		"comment_list": commentInfos,
	}
	c.ServeJSON()
}
