package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/object"
	"ByteRhythm/utils"
	"os"
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
	//经过测试，提交的latestTimeStamp为55574-03-08 08:53:51 +0800 CST，这个数据毫无意义还会干扰查询
	//latestTimeStamp, _ := c.GetInt("latest_time")
	//if latestTimeStamp == 0 {
	//	latestTimeStamp = int(time.Now().Unix())
	//}
	latestTimeStamp := int(time.Now().Unix())
	latestTime := time.Unix(int64(latestTimeStamp), 0)
	//不需要验证token，不登录也能看视频流
	token := c.GetString("token")

	c.o.QueryTable(new(models.Video)).Filter("create_time__lte", latestTime).OrderBy("-create_time").Limit(30, 0).All(&videos)
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
		commentCount, _ := c.o.QueryTable(new(models.Comment)).Filter("video_id", video.Id).Count()
		favoriteCount, _ := c.o.QueryTable(new(models.Favorite)).Filter("video_id", video.Id).Count()
		if favoriteCount == 0 {
			isFavorite = false
		} else {
			isFavorite = true
		}
		userInfo := c.GetUserInfo(video.AuthorId.Id, token)
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

	user, _ := utils.GetUserFromToken(token)
	if url := c.UploadMP4(c.GetFile("data")); url == "" {
		c.PublishFail("发布失败")
		return
	} else {
		imgPath := utils.VideoGetNetImgCount(1, url)
		if imgPath == "" {
			c.PublishFail("发布失败")
			return

		}
		coverUrl := c.UploadJPG(imgPath, url)
		if coverUrl == "" {
			c.PublishFail("发布失败")
			return
		}
		os.Remove(imgPath)

		video := models.Video{
			AuthorId: user,
			PlayUrl:  url,
			Title:    title,
			CoverUrl: coverUrl,
		}
		if _, err := c.o.Insert(&video); err != nil {
			c.PublishFail("发布失败")
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

func (c *VideoController) PublishFail(msg string) {
	c.Data["json"] = map[string]interface{}{
		"status_code": 1,
		"status_msg":  msg,
	}
	c.ServeJSON()
}

// List 获取发布视频列表
func (c *VideoController) List() {
	uid, _ := c.GetInt("user_id")
	token := c.GetString("token")
	var (
		videos    []*models.Video
		videoList []*object.VideoInfo
		userInfo  = c.GetUserInfo(uid, token)
	)
	c.o.QueryTable(new(models.Video)).Filter("author_id", uid).All(&videos)
	for _, video := range videos {
		var isFavorite bool
		commentCount, _ := c.o.QueryTable(new(models.Comment)).Filter("video_id", video.Id).Count()
		favoriteCount, _ := c.o.QueryTable(new(models.Favorite)).Filter("video_id", video.Id).Count()
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
