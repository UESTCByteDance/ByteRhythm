package service

import (
	"ByteRhythm/app/video/dao"
	"ByteRhythm/consts"
	"ByteRhythm/idl/video/videoPb"
	"ByteRhythm/model"
	"ByteRhythm/mq"
	"ByteRhythm/util"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
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

	// 使用Keys命令获取所有键
	keys, err := dao.RedisClient.Keys(ctx, "*").Result()
	if err != nil {
		FeedResponseData(res, 1, "获取视频流失败")
		return err
	}
	keys = util.SortKeys(keys)
	var videoList []*videoPb.Video

	//从缓存取对应的视频
	for _, key := range keys {
		// 尝试从 Redis 缓存中获取数据
		redisResult, err := dao.RedisClient.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			FeedResponseData(res, 1, "获取视频流失败")
			return err
		}
		if err != redis.Nil {
			var video videoPb.Video
			err = json.Unmarshal([]byte(redisResult), &video)
			if err != nil {
				FeedResponseData(res, 1, "获取视频流失败")
				return err
			}
			if token == "" {
				video.IsFavorite = false
				video.Author.IsFollow = false
			} else {
				video.IsFavorite, _ = dao.NewVideoDao(ctx).GetIsFavorite(int(video.Id), token)
				video.Author.IsFollow, _ = dao.NewVideoDao(ctx).GetIsFollowed(int(video.Author.Id), token)
			}
			videoList = append(videoList, &video)
		}
	}
	if len(keys) == 30 {
		FeedResponseData(res, 0, "获取视频流成功", videoList, latestTimeStamp)
		return nil
	}

	//从数据库取对应的视频
	videos, err := dao.NewVideoDao(ctx).GetVideoListByLatestTime(latestTime, util.StringArray2IntArray(keys), 30-len(keys))
	if err != nil {
		FeedResponseData(res, 1, "获取失败")
		return err
	}
	var nextTime int64
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	for _, video := range videos {
		videoPbModel := BuildVideoPbModel(ctx, video, token)
		videoList = append(videoList, videoPbModel)
		//将视频存入缓存，加入消息队列
		body, _ := json.Marshal(&videoPbModel)
		err := mq.SendMessage2MQ(body, consts.Video2RedisQueue)
		if err != nil {
			return err
		}
	}
	FeedResponseData(res, 0, "获取视频流成功", videoList, nextTime)

	return nil
}

func (v *VideoSrv) Publish(ctx context.Context, req *videoPb.PublishRequest, res *videoPb.PublishResponse) error {
	//加入消息队列
	body, _ := json.Marshal(&req)
	err := mq.SendMessage2MQ(body, consts.CreateVideoQueue)
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
		PublishListResponseData(res, 1, "获取失败")
		return err
	}
	var videoList []*videoPb.Video
	for _, video := range videos {
		videoList = append(videoList, BuildVideoPbModel(ctx, video, token))
	}
	PublishListResponseData(res, 0, "获取成功", videoList)
	return nil
}

func FeedResponseData(res *videoPb.FeedResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.VideoList = params[0].([]*videoPb.Video)
		res.NextTime = params[1].(int64)
	}
}
func PublishResponseData(res *videoPb.PublishResponse, StatusCode int32, StatusMsg string) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
}
func PublishListResponseData(res *videoPb.PublishListResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.VideoList = params[0].([]*videoPb.Video)
	}
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

func VideoMQ2DB(ctx context.Context, req *videoPb.PublishRequest) error {
	token := req.Token
	data := req.Data
	title := req.Title
	uid, _ := util.GetUserIdFromToken(token)
	VideoUrl, _ := util.UploadVideo(data)
	imgPath := util.VideoGetNetImgCount(1, VideoUrl)
	coverUrl := util.UploadJPG(imgPath, VideoUrl)
	os.Remove(imgPath)
	video := BuildVideoModel(uid, VideoUrl, coverUrl, title)
	//将视频存入数据库
	if err := dao.NewVideoDao(ctx).CreateVideo(&video); err != nil {
		return err
	}
	//将视频存入缓存
	var videoCache *videoPb.Video
	videoCache = BuildVideoPbModel(ctx, &video, token)
	videoJson, _ := json.Marshal(&videoCache)
	dao.RedisClient.Set(ctx, fmt.Sprintf("%d", video.ID), videoJson, time.Hour)
	return nil
}

func VideoMQ2Redis(ctx context.Context, req *videoPb.Video) error {
	videoJson, _ := json.Marshal(&req)
	dao.RedisClient.Set(ctx, fmt.Sprintf("%d", req.Id), videoJson, time.Hour)
	return nil
}
