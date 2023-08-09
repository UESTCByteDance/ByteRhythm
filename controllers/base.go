package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/object"
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"mime/multipart"
	"strings"
	"time"
)

type baseController struct {
	web.Controller
	o orm.Ormer
}

func (c *baseController) Prepare() {
	c.o = orm.NewOrm()
}

func (c *baseController) Upload(file multipart.File, header *multipart.FileHeader, err error) string {
	var reader io.Reader = file
	var size = header.Size

	defer file.Close()

	if err != nil {
		return ""
	}

	// 取文件后缀
	suffix := ""
	if header.Filename != "" && strings.Contains(header.Filename, ".") {
		suffix = header.Filename[strings.LastIndex(header.Filename, ".")+1:]
		suffix = strings.ToLower(suffix)
	}

	secretKey, _ := web.AppConfig.String("SecretKey")
	accessKey, _ := web.AppConfig.String("AccessKey")
	bucket, _ := web.AppConfig.String("Bucket")
	domain, _ := web.AppConfig.String("Domain")
	key := fmt.Sprintf("%d.%s", time.Now().UnixNano(), suffix)
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, key),
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Region = &storage.ZoneHuanan
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": header.Filename,
		},
	}

	err = uploader.Put(context.Background(), &ret, upToken, key, reader, size, &putExtra)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s/%s", domain, ret.Key)
}

func (c *baseController) GetUserInfo(uid int) object.UserInfo {
	var (
		user           models.User
		videos         []*models.Video
		isFollow       bool
		totalFavorited int
	)
	//获取用户信息
	c.o.QueryTable(new(models.User)).Filter("id", uid).One(&user)
	followerCount, _ := c.o.QueryTable(new(models.Follow)).Filter("user_id", uid).Count()
	if followerCount > 0 {
		isFollow = true
	} else {
		isFollow = false
	}

	followCount, _ := c.o.QueryTable(new(models.Follow)).Filter("followed_user_id", uid).Count()

	favoriteCount, _ := c.o.QueryTable(new(models.Favorite)).Filter("user_id", uid).Count()

	workCount, _ := c.o.QueryTable(new(models.Video)).Filter("author_id", uid).All(&videos)

	//获取视频获赞数量
	for _, video := range videos {
		count, _ := c.o.QueryTable(new(models.Favorite)).Filter("video_id", video.Id).Count()
		totalFavorited += int(count)
	}

	userInfo := object.UserInfo{
		ID:              user.Id,
		Name:            user.Username,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		FollowCount:     int(followCount),
		FollowerCount:   int(followerCount),
		WorkCount:       int(workCount),
		FavoriteCount:   int(favoriteCount),
		TotalFavorited:  totalFavorited,
		IsFollow:        isFollow,
	}
	return userInfo
}
