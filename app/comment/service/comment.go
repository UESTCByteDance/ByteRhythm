package service

import (
	"ByteRhythm/app/comment/dao"
	"ByteRhythm/idl/comment/commentPb"
	"ByteRhythm/idl/favorite/favoritePb"
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"time"
)

var CommentSrvIns *CommentSrv
var CommentSrvOnce sync.Once

// GetCommentSrv 懒汉式的单例模式 lazy-loading --> 懒汉式
func GetCommentSrv() *CommentSrv {
	CommentSrvOnce.Do(func() {
		CommentSrvIns = &CommentSrv{}
	})
	return CommentSrvIns
}

type CommentSrv struct {
}

func (c CommentSrv) CommentAction(ctx context.Context, req *commentPb.CommentActionRequest, res *commentPb.CommentActionResponse) error {
	token := req.Token
	videoId := req.VideoId
	actionType := req.ActionType
	commentText := req.CommentText
	commentId := req.CommentId

	// 解析token
	user, _ := util.GetUserFromToken(token)

	// 发布评论
	if actionType == 1 {

		//构建model
		comment := model.Comment{
			UserID:  int(user.ID), // uint to int
			VideoID: int(videoId), // int64 to int
			Content: commentText,
		}

		Comment := commentPb.Comment{
			Id:         int64(comment.ID),
			User:       BuildUserPbModel(ctx, user, token),
			Content:    comment.Content,
			CreateDate: time.Now().Format("01-02 15:04"),
		}

		// 数据库创建comment
		if err := dao.NewCommentDao(ctx).CreateComment(&comment); err != nil {
			return err
		}

		//修改 redis 1号数据库
		key := fmt.Sprintf("%d", videoId)
		redisResult, err := dao.RedisNo1Client.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			CommentActionResponseData(res, 1, "操作失败")
			return err
		}

		if err != redis.Nil { // 在redis中找到了视频信息
			var video favoritePb.Video
			err = json.Unmarshal([]byte(redisResult), &video)
			if err != nil {
				CommentActionResponseData(res, 1, "操作失败")
				return err
			}

			video.CommentCount += 1

			videoJson, _ := json.Marshal(&video)
			dao.RedisNo1Client.Set(ctx, key, videoJson, time.Hour)
		}

		// 构建 Redis 键
		key = strconv.Itoa(int(videoId))

		// 尝试 删除 Redis 2号数据库 的记录
		err = dao.RedisNo2Client.Del(ctx, key).Err()
		if err != nil && err != redis.Nil {
			CommentActionResponseData(res, 1, "操作失败")
			return err
		}

		CommentActionResponseData(res, 0, "评论成功", &Comment)

		return nil
	}

	// 删除评论

	//修改mysql数据库
	comment := model.Comment{
		ID:      uint(commentId),
		UserID:  int(user.ID), // uint to int
		VideoID: int(videoId), // int64 to int
		Content: commentText,
	}
	if err := dao.NewCommentDao(ctx).DeleteComment(&comment); err != nil {
		return err
	}

	//修改 redis 1号数据库
	key := fmt.Sprintf("%d", videoId)
	redisResult, err := dao.RedisNo1Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		CommentActionResponseData(res, 1, "操作失败")
		return err
	}

	if err != redis.Nil { // 在redis中找到了视频信息
		var video favoritePb.Video
		err = json.Unmarshal([]byte(redisResult), &video)
		if err != nil {
			CommentActionResponseData(res, 1, "操作失败")
			return err
		}

		video.CommentCount -= 1

		videoJson, _ := json.Marshal(&video)
		dao.RedisNo1Client.Set(ctx, key, videoJson, time.Hour)
	}

	// 构建 Redis 键
	key = strconv.Itoa(int(videoId))

	// 尝试 删除 Redis 2号数据库 的记录
	err = dao.RedisNo2Client.Del(ctx, key).Err()
	if err != nil && err != redis.Nil {
		CommentActionResponseData(res, 1, "操作失败")
		return err
	}

	CommentActionResponseData(res, 0, "删除评论成功")

	return nil
}

func (c CommentSrv) CommentList(ctx context.Context, req *commentPb.CommentListRequest, res *commentPb.CommentListResponse) error {
	videoId := req.VideoId

	// 构建redis键
	key := strconv.Itoa(int(videoId))

	// 尝试从 Redis 缓存中获取数据
	redisResult, err := dao.RedisNo2Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		CommentListResponseData(res, 1, "获取评论列表失败")
		return err
	}

	if redisResult != "" {
		// 如果缓存中存在数据，则解码并直接返回缓存数据
		err = json.Unmarshal([]byte(redisResult), &res.CommentList)
		if err != nil {
			CommentListResponseData(res, 1, "获取评论列表失败")
			return err
		}

		CommentListResponseData(res, 0, "获取评论列表成功")
		return nil
	}

	// 缓存中不存在数据，则从数据库获取聊天记录
	comments, err := dao.NewCommentDao(ctx).GetCommentListByVideoId(videoId)
	if err != nil {
		CommentListResponseData(res, 1, "获取评论列表失败")
		return err
	}

	for _, comment := range comments {
		res.CommentList = append(res.CommentList, BuildCommentPbModel(ctx, comment))
	}

	// 将结果存入 Redis 缓存
	jsonBytes, err := json.Marshal(&res.CommentList)
	if err != nil {
		CommentListResponseData(res, 1, "获取评论列表失败")
		return err
	}

	err = dao.RedisNo2Client.Set(ctx, key, string(jsonBytes), time.Hour).Err()
	if err != nil {
		CommentListResponseData(res, 1, "评论列表存入redis失败")
		return err
	}

	CommentListResponseData(res, 0, "获取评论列表成功")
	return nil
}

func CommentActionResponseData(res *commentPb.CommentActionResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.Comment = params[0].(*commentPb.Comment)
	}
}

func CommentListResponseData(res *commentPb.CommentListResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.CommentList = params[0].([]*commentPb.Comment)
	}
}

func BuildUserPbModel(ctx context.Context, user *model.User, token string) *commentPb.User {
	uid := int(user.ID)
	FollowCount, _ := dao.NewCommentDao(ctx).GetFollowCount(uid)
	FollowerCount, _ := dao.NewCommentDao(ctx).GetFollowerCount(uid)
	WorkCount, _ := dao.NewCommentDao(ctx).GetWorkCount(uid)
	FavoriteCount, _ := dao.NewCommentDao(ctx).GetFavoriteCount(uid)
	TotalFavorited, _ := dao.NewCommentDao(ctx).GetTotalFavorited(uid)
	IsFollow, _ := dao.NewCommentDao(ctx).GetIsFollowed(uid, token)
	return &commentPb.User{
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

func BuildCommentPbModel(ctx context.Context, comment *model.Comment) *commentPb.Comment {
	name, _ := dao.NewCommentDao(ctx).GetUsernameByUid(int64(comment.UserID))
	avatar, _ := dao.NewCommentDao(ctx).GetAvatarByUid(int64(comment.UserID))

	return &commentPb.Comment{
		Id: int64(comment.ID),
		User: &commentPb.User{
			Id:     int64(comment.UserID),
			Name:   name,
			Avatar: avatar,
		},
		Content:    comment.Content,
		CreateDate: comment.CreatedAt.Format("01-02 15:04"),
	}
}
