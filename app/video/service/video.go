package service

import (
	"ByteRhythm/app/video/dao"
	"ByteRhythm/app/video/mq"
	"ByteRhythm/idl/video/videoPb"
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"
)

type VideoSrv struct {
}

var VideoSrvIns *VideoSrv
var VideoSrvOnce sync.Once

func GetVideoSrv() *VideoSrv {
	VideoSrvOnce.Do(func() {
		VideoSrvIns = &VideoSrv{}
	})
	return VideoSrvIns
}

func (v *VideoSrv) Feed(ctx context.Context, req *videoPb.FeedRequest, res *videoPb.FeedResponse) error {

	latestTimeStamp := time.Now().Unix()
	latestTime := time.Unix(latestTimeStamp, 0)
	token := req.Token

	videos, err := dao.NewVideoDao(ctx).GetVideoListByLatestTime(latestTime)
	if err != nil {
		return err
	}
	for _, video := range videos {
		res.VideoList = append(res.VideoList, BuildVideoPbModel(ctx, video, token))
	}
	return nil
}

func (v *VideoSrv) Publish(ctx context.Context, req *videoPb.PublishRequest, res *videoPb.PublishResponse) error {
	//加入消息队列
	body, _ := json.Marshal(&req)
	err := mq.SendMessage2MQ(body)
	if err != nil {
		PublishResponseData(res, 1, "发布失败")
		return err
	}
	PublishResponseData(res, 0, "发布成功")
	return nil
}

func (v *VideoSrv) PublishList(ctx context.Context, req *videoPb.PublishListRequest, res *videoPb.PublishListResponse) error {
	token := req.Token
	uid := int(req.UserId)

	videos, err := dao.NewVideoDao(ctx).GetVideoListByUserId(uid)
	if err != nil {
		return err
	}
	for _, video := range videos {
		res.VideoList = append(res.VideoList, BuildVideoPbModel(ctx, video, token))
	}
	return nil
}

func PublishResponseData(res *videoPb.PublishResponse, StatusCode int32, StatusMsg string) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
}

func BuildVideoPbModel(ctx context.Context, video *model.Video, token string) *videoPb.Video {
	Author, _ := dao.NewVideoDao(ctx).FindUser(video)
	vid := int(video.ID)
	FavoriteCount, _ := dao.NewVideoDao(ctx).GetFavoriteCount(vid)
	CommentCount, _ := dao.NewVideoDao(ctx).GetCommentCount(vid)
	IsFavorite, _ := dao.NewVideoDao(ctx).GetIsFavorite(vid, token)
	return &videoPb.Video{
		Id:            int64(vid),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		Title:         video.Title,
		FavoriteCount: FavoriteCount,
		CommentCount:  CommentCount,
		IsFavorite:    IsFavorite,
		Author:        BuildUserPbModel(ctx, Author, token),
	}
}

func BuildUserPbModel(ctx context.Context, user *model.User, token string) *videoPb.User {
	uid := int(user.ID)
	FollowCount, _ := dao.NewVideoDao(ctx).GetFollowCount(uid)
	FollowerCount, _ := dao.NewVideoDao(ctx).GetFollowerCount(uid)
	WorkCount, _ := dao.NewVideoDao(ctx).GetWorkCount(uid)
	FavoriteCount, _ := dao.NewVideoDao(ctx).GetFavoriteCount(uid)
	TotalFavorited, _ := dao.NewVideoDao(ctx).GetTotalFavorited(uid)
	IsFollow, _ := dao.NewVideoDao(ctx).GetIsFollowed(uid, token)
	return &videoPb.User{
		Id:              int64(uid),
		Name:            user.Username,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		FollowCount:     FollowCount,
		FollowerCount:   FollowerCount,
		WorkCount:       WorkCount,
		FavoriteCount:   FavoriteCount,
		TotalFavorited:  TotalFavorited,
		IsFollow:        IsFollow,
	}
}

func BuildVideoModel(uid int, VideoUrl string, coverUrl string, title string) model.Video {
	return model.Video{
		AuthorID: uid,
		PlayUrl:  VideoUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
}

func VideoMQ2MySQL(ctx context.Context, req *videoPb.PublishRequest) error {
	token := req.Token
	data := req.Data
	title := req.Title
	uid, _ := util.GetUserIdFromToken(token)
	VideoUrl, _ := util.UploadVideo(data)
	imgPath := util.VideoGetNetImgCount(1, VideoUrl)
	coverUrl := util.UploadJPG(imgPath, VideoUrl)
	os.Remove(imgPath)
	video := BuildVideoModel(uid, VideoUrl, coverUrl, title)
	if err := dao.NewVideoDao(ctx).CreateVideo(&video); err != nil {
		return err
	}
	return nil
}
