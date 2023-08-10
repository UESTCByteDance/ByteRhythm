package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/object"
	"ByteRhythm/utils"
	"fmt"
	"time"
)

type VideoController struct {
	baseController
}

// Feed 获取视频流
func (c *VideoController) Feed() {
	var (
		videoList []*object.VideoInfo
		videos    []*models.Video
	)
	latestTimeStamp, _ := c.GetInt("latest_time")
	if latestTimeStamp == 0 {
		latestTimeStamp = int(time.Now().Unix())
	}
	latestTime := time.Unix(int64(latestTimeStamp), 0)
	fmt.Println(latestTime)
	token := c.GetString("token")
	if token != "" {
		if err := utils.ValidateToken(token); err != nil {
			c.Data["json"] = map[string]interface{}{
				"status_code": 1,
				"status_msg":  "token验证失败",
				"video_list":  nil,
				"next_time":   nil,
			}
			c.ServeJSON()
			return
		} else {
			c.O.QueryTable(new(models.Video)).Filter("create_time__lte", latestTime).OrderBy("-create_time").Limit(30, 0).All(&videos)
		}
	}
	//经过测试，如果不登录，提交的latestTimeStamp为55574-03-08 08:53:51 +0800 CST，这个数据毫无意义还会干扰查询
	c.O.QueryTable(new(models.Video)).OrderBy("-create_time").Limit(30, 0).All(&videos)
	if len(videos) == 0 {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "获取视频流失败",
			"video_list":  nil,
			"next_time":   nil,
		}
		c.ServeJSON()
		return
	}

	for _, video := range videos {
		var isFavorite bool
		commentCount, _ := c.O.QueryTable(new(models.Comment)).Filter("video_id", video.Id).Count()
		favoriteCount, _ := c.O.QueryTable(new(models.Favorite)).Filter("video_id", video.Id).Count()
		if favoriteCount == 0 {
			isFavorite = false
		} else {
			isFavorite = true
		}
		userInfo := c.GetUserInfo(video.AuthorId.Id)
		videoList = append(videoList, &object.VideoInfo{
			ID:            video.Id,
			Title:         video.Title,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			CommentCount:  int(commentCount),
			FavoriteCount: int(favoriteCount),
			IsFavorite:    isFavorite,
			Author:        &userInfo,
		})
	}
	next := videos[len(videos)-1].CreateTime.Unix()
	c.Data["json"] = map[string]interface{}{
		"status_code": 0,
		"status_msg":  "获取视频流成功",
		"video_list":  videoList,
		"next_time":   next,
	}
	c.ServeJSON()
	return

}

// Publish 发布视频
func (c *VideoController) Publish() {
	token := c.GetString("token")
	title := c.GetString("title")
	if err := utils.ValidateToken(token); err != nil {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "token验证失败",
		}
		c.ServeJSON()
		return
	}
	var user models.User

	username, _ := utils.GetUsernameFromToken(token)
	c.O.QueryTable(new(models.User)).Filter("username", username).One(&user)
	if url := c.Upload(c.GetFile("data")); url == "" {
		c.Data["json"] = map[string]interface{}{
			"status_code": 1,
			"status_msg":  "发布失败",
		}
		c.ServeJSON()
		return
	} else {
		//需要修改cover_url，此处写死了
		video := models.Video{
			AuthorId: &user,
			PlayUrl:  url,
			Title:    title,
			CoverUrl: "http://rz2n87yck.hn-bkt.clouddn.com/cover.jpg",
		}
		if _, err := c.O.Insert(&video); err != nil {
			c.Data["json"] = map[string]interface{}{
				"status_code": 1,
				"status_msg":  "发布失败",
			}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{
				"status_code": 0,
				"status_msg":  "发布成功",
			}
			c.ServeJSON()
			return
		}

	}

}

// List 获取发布视频列表
func (c *VideoController) List() {
	uid, _ := c.GetInt("user_id")
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
	var (
		videos    []*models.Video
		videoList []*object.VideoInfo
		userInfo  = c.GetUserInfo(uid)
	)
	c.O.QueryTable(new(models.Video)).Filter("author_id", uid).All(&videos)
	for _, video := range videos {
		var isFavorite bool
		commentCount, _ := c.O.QueryTable(new(models.Comment)).Filter("video_id", video.Id).Count()
		favoriteCount, _ := c.O.QueryTable(new(models.Favorite)).Filter("video_id", video.Id).Count()
		if favoriteCount == 0 {
			isFavorite = false
		} else {
			isFavorite = true
		}

		videoList = append(videoList, &object.VideoInfo{
			ID:            video.Id,
			Title:         video.Title,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			CommentCount:  int(commentCount),
			FavoriteCount: int(favoriteCount),
			IsFavorite:    isFavorite,
			Author:        &userInfo,
		})
	}
	c.Data["json"] = map[string]interface{}{
		"status_code": 0,
		"status_msg":  "获取视频列表成功",
		"video_list":  videoList,
	}
	c.ServeJSON()
	return
}
