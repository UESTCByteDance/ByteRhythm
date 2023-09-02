package service

import (
	"ByteRhythm/app/favorite/dao"
	"ByteRhythm/consts"
	"ByteRhythm/idl/favorite/favoritePb"
	"ByteRhythm/model"
	"ByteRhythm/mq"
	"ByteRhythm/util"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type FavoriteSrv struct {
}

var FavoriteSrvIns *FavoriteSrv
var FavoriteSrvOnce sync.Once

// GetFavoriteSrv 懒汉式的单例模式 lazy-loading --> 懒汉式
func GetFavoriteSrv() *FavoriteSrv {
	FavoriteSrvOnce.Do(func() {
		FavoriteSrvIns = &FavoriteSrv{}
	})
	return FavoriteSrvIns
}

func (c FavoriteSrv) FavoriteAction(ctx context.Context, req *favoritePb.FavoriteActionRequest, res *favoritePb.FavoriteActionResponse) error {
	actionType := req.ActionType
	vid := req.VideoId
	uid, _ := util.GetUserIdFromToken(req.Token)

	body, _ := json.Marshal(&req)

	// 点赞
	if actionType == 1 {
		// 不能重复点赞
		isFavorite, _ := dao.NewFavoriteDao(ctx).GetIsFavoriteByUserIdAndVid(int64(uid), vid)

		if isFavorite {
			FavoriteActionResponseData(res, 1, "重复点赞")
			return nil
		}

		//修改redis
		key := fmt.Sprintf("%d", vid)
		redisResult, err := dao.RedisClient.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			FavoriteActionResponseData(res, 1, "点赞失败")
			return err
		}

		if err != redis.Nil { // 在redis中找到了视频信息
			var video favoritePb.Video
			err = json.Unmarshal([]byte(redisResult), &video)
			if err != nil {
				FavoriteActionResponseData(res, 1, "点赞失败")
				return err
			}

			video.FavoriteCount += 1

			videoJson, _ := json.Marshal(&video)
			dao.RedisClient.Set(ctx, key, videoJson, time.Hour)
		}

		// 加入消息队列
		err = mq.SendMessage2MQ(body, consts.CreateFavorite2MQ)
		if err != nil {
			FavoriteActionResponseData(res, 1, "点赞失败")
			return err
		}
	}

	// 取消点赞
	if actionType == 2 {
		//修改redis
		key := fmt.Sprintf("%d", vid)
		redisResult, err := dao.RedisClient.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			FavoriteActionResponseData(res, 1, "取消点赞失败")
			return err
		}

		if err != redis.Nil { // 在redis中找到了视频信息
			var video favoritePb.Video
			err = json.Unmarshal([]byte(redisResult), &video)
			if err != nil {
				FavoriteActionResponseData(res, 1, "取消点赞失败")
				return err
			}

			video.FavoriteCount -= 1

			videoJson, _ := json.Marshal(&video)
			dao.RedisClient.Set(ctx, key, videoJson, time.Hour)
		}

		// 加入消息队列
		err = mq.SendMessage2MQ(body, consts.DeleteFavorite2MQ)
		if err != nil {
			FavoriteActionResponseData(res, 1, "取消点赞失败")
			return err
		}
	}

	return nil
}

func (c FavoriteSrv) FavoriteList(ctx context.Context, req *favoritePb.FavoriteListRequest, res *favoritePb.FavoriteListResponse) error {
	uid := req.UserId

	favorites, err := dao.NewFavoriteDao(ctx).GetFavoriteListByUserId(uid)

	if err != nil {
		return err
	}

	res.StatusCode = 0
	res.StatusMsg = "获取喜欢列表成功"

	for _, favorite := range favorites {
		vid := favorite.VideoID
		res.VideoList = append(res.VideoList, BuildVideoPbModelByVid(ctx, vid))
	}
	return nil
}

func BuildVideoPbModelByVid(ctx context.Context, vid int) *favoritePb.Video {
	FavoriteCount, _ := dao.NewFavoriteDao(ctx).GetFavoriteCount(vid)
	PlayUrl, _ := dao.NewFavoriteDao(ctx).GetPlayUrlByVid(vid)
	CoverUrl, _ := dao.NewFavoriteDao(ctx).GetCoverUrlByVid(vid)

	return &favoritePb.Video{
		Id:            int64(vid),
		PlayUrl:       PlayUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: FavoriteCount,
	}
}

func BuildUserPbModel(ctx context.Context, user *model.User, token string) *favoritePb.User {
	uid := int(user.ID)
	FollowCount, _ := dao.NewFavoriteDao(ctx).GetFollowCount(uid)
	FollowerCount, _ := dao.NewFavoriteDao(ctx).GetFollowerCount(uid)
	WorkCount, _ := dao.NewFavoriteDao(ctx).GetWorkCount(uid)
	FavoriteCount, _ := dao.NewFavoriteDao(ctx).GetFavoriteCount(uid)
	TotalFavorited, _ := dao.NewFavoriteDao(ctx).GetTotalFavorited(uid)
	IsFollow, _ := dao.NewFavoriteDao(ctx).GetIsFollowed(uid, token)
	return &favoritePb.User{
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

func FavoriteActionResponseData(res *favoritePb.FavoriteActionResponse, StatusCode int32, StatusMsg string) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
}

func FavoriteMQ2DB(ctx context.Context, req *favoritePb.FavoriteActionRequest) error {
	token := req.Token
	videoId := req.VideoId
	actionType := req.ActionType

	// 解析token
	user, _ := util.GetUserFromToken(token)

	// 点赞
	if actionType == 1 {
		//加入redis
		//加入校队列

		favorite := model.Favorite{
			UserID:  int(user.ID), // uint to int
			VideoID: int(videoId), // int64 to int
		}
		if err := dao.NewFavoriteDao(ctx).CreateFavorite(&favorite); err != nil {
			return err
		}

		return nil
	}

	// 取消点赞
	favorite := model.Favorite{
		UserID:  int(user.ID), // uint to int
		VideoID: int(videoId), // int64 to int
	}

	if err := dao.NewFavoriteDao(ctx).DeleteFavorite(&favorite); err != nil {
		return err
	}

	return nil
}
