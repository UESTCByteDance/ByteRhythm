package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/object"
	"ByteRhythm/utils"
	"strconv"
)

type FavoriteController struct {
	baseController
}

// 赞操作
func (c *FavoriteController) FavoriteAction() {
	// 获取必要参数
	tokenString := c.GetString("token")                       // 用户鉴权
	videoId, _ := strconv.Atoi(c.GetString("video_id"))       // 视频id
	actionType, _ := strconv.Atoi(c.GetString("action_type")) // 1-点赞，2-取消点赞
	// 鉴权
	if err := utils.ValidateToken(tokenString); err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "token鉴权失败",
		}
		c.ServeJSON()
		return
	}
	// 不能给自己点赞
	username, err := utils.GetUsernameFromToken(tokenString)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "根据token获取用户信息失败",
		}
		c.ServeJSON()
		return
	}
	var (
		user  models.User
		video models.Video
	)
	query := c.o.QueryTable(new(models.User)).Filter("username", username)
	if exist := query.Exist(); !exist {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "用户名不存在",
		}
		c.ServeJSON()
		return
	}
	query.One(&user)
	c.o.QueryTable(new(models.Video)).Filter("id", videoId).One(&video)
	if video.AuthorId.Id == user.Id {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "不能给自己点赞",
		}
		c.ServeJSON()
		return
	}
	// 点赞 or 取消点赞
	if actionType == 1 {
		// 不能重复点赞
		exist := c.o.QueryTable(new(models.Favorite)).Filter("user_id", user.Id).Filter("video_id", videoId).Exist()
		if exist {
			c.Data["json"] = map[string]interface{}{
				"status_code": 1,
				"status_msg":  "不能重复点赞",
			}
			c.ServeJSON()
			return
		}
		// 创建点赞记录
		favorite := models.Favorite{
			UserId:  &user,
			VideoId: &video,
		}
		c.o.Insert(&favorite)
		c.Data["json"] = map[string]interface{}{
			"status_code": 0,
			"status_msg":  "点赞成功",
		}
		c.ServeJSON()
	} else if actionType == 2 {
		// 删除点赞记录
		c.o.Delete(&models.Favorite{UserId: &user, VideoId: &video})
		c.Data["json"] = map[string]interface{}{
			"status_code": 0,
			"status_msg":  "取消点赞成功",
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "操作类型未知",
		}
		c.ServeJSON()
	}
}

// 喜欢列表
func (c *FavoriteController) FavoriteList() {
	// 获取必要参数
	userId, _ := strconv.Atoi(c.GetString("user_id")) // 用户id
	tokenString := c.GetString("token")               // 用户鉴权token
	// 鉴权
	if err := utils.ValidateToken(tokenString); err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "token鉴权失败",
			"video_list":  nil,
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
	// 获取用户喜欢的视频列表
	var (
		favs   []models.Favorite
		videos []object.VideoInfo
	)
	userInfo := c.GetUserInfo(userId)
	c.o.QueryTable(new(models.Favorite)).Filter("user_id", user.Id).All(&favs)
	for _, fav := range favs {
		var video models.Video
		c.o.QueryTable(new(models.Video)).Filter("id", fav.VideoId.Id).One(&video)
		var isFavorite bool
		commentCount, _ := c.o.QueryTable(new(models.Comment)).Filter("video_id", video.Id).Count()
		favoriteCount, _ := c.o.QueryTable(new(models.Favorite)).Filter("video_id", video.Id).Count()
		if favoriteCount == 0 {
			isFavorite = false
		} else {
			isFavorite = true
		}

		videos = append(videos, object.VideoInfo{
			Author:        &userInfo,
			CommentCount:  int(commentCount),
			CoverURL:      video.CoverUrl,
			FavoriteCount: int(favoriteCount),
			ID:            video.Id,
			IsFavorite:    isFavorite,
			PlayURL:       video.PlayUrl,
			Title:         video.Title,
		})
	}
	// 返回响应
	c.Data["json"] = map[string]interface{}{
		"status_code": "0",
		"status_msg":  "获取喜欢列表成功",
		"video_list":  videos,
	}
	c.ServeJSON()
}
