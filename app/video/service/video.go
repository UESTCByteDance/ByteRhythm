package service

import (
	"ByteRhythm/app/video/dao"
	"ByteRhythm/idl/video/videoPb"
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"
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
		AuthorID := video.AuthorID
		Author, _ := dao.NewVideoDao(ctx).FindUser(video)
		vid := int(video.ID)
		FavoriteCount, _ := dao.NewVideoDao(ctx).GetFavoriteCount(vid)
		CommentCount, _ := dao.NewVideoDao(ctx).GetCommentCount(vid)
		IsFavorite, _ := dao.NewVideoDao(ctx).GetIsFavorite(vid, token)
		FollowCount, _ := dao.NewVideoDao(ctx).GetFollowCount(AuthorID)
		FollowerCount, _ := dao.NewVideoDao(ctx).GetFollowerCount(AuthorID)
		WorkCount, _ := dao.NewVideoDao(ctx).GetWorkCount(AuthorID)
		UserFavoriteCount, _ := dao.NewVideoDao(ctx).GetUserFavoriteCount(AuthorID)
		TotalFavorited, _ := dao.NewVideoDao(ctx).GetTotalFavorited(AuthorID)
		IsFollow, _ := dao.NewVideoDao(ctx).GetIsFollowed(AuthorID, token)
		res.VideoList = append(res.VideoList, &videoPb.Video{
			Id:            int64(vid),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			Title:         video.Title,
			FavoriteCount: FavoriteCount,
			CommentCount:  CommentCount,
			IsFavorite:    IsFavorite,
			Author: &videoPb.User{
				Id:              int64(AuthorID),
				Name:            Author.Username,
				Avatar:          Author.Avatar,
				BackgroundImage: Author.BackgroundImage,
				Signature:       Author.Signature,
				FollowCount:     FollowCount,
				FollowerCount:   FollowerCount,
				WorkCount:       WorkCount,
				FavoriteCount:   UserFavoriteCount,
				TotalFavorited:  TotalFavorited,
				IsFollow:        IsFollow,
			},
		})
	}

	return nil
}

func (v *VideoSrv) Publish(ctx context.Context, req *videoPb.PublishRequest, res *videoPb.PublishResponse) error {
	token := req.Token
	data := req.Data
	title := req.Title
	uid, _ := util.GetUserIdFromToken(token)
	VideoUrl, err := util.UploadVideo(data)
	if err != nil {
		PublishResponseData(res, 1, "发布失败")
		return err
	}
	video := model.Video{
		AuthorID: uid,
		PlayUrl:  VideoUrl,
		CoverUrl: "http://rz2n87yck.hn-bkt.clouddn.com/cover_21c95d84-9960-4dd1-a59e-54b7f6ea804d.jpg",
		Title:    title,
	}
	if err := dao.NewVideoDao(ctx).CreateVideo(&video); err != nil {
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
		Author, _ := dao.NewVideoDao(ctx).FindUser(video)
		vid := int(video.ID)
		FavoriteCount, _ := dao.NewVideoDao(ctx).GetFavoriteCount(vid)
		CommentCount, _ := dao.NewVideoDao(ctx).GetCommentCount(vid)
		IsFavorite, _ := dao.NewVideoDao(ctx).GetIsFavorite(vid, token)
		FollowCount, _ := dao.NewVideoDao(ctx).GetFollowCount(uid)
		FollowerCount, _ := dao.NewVideoDao(ctx).GetFollowerCount(uid)
		WorkCount, _ := dao.NewVideoDao(ctx).GetWorkCount(uid)
		UserFavoriteCount, _ := dao.NewVideoDao(ctx).GetUserFavoriteCount(uid)
		TotalFavorited, _ := dao.NewVideoDao(ctx).GetTotalFavorited(uid)
		IsFollow, _ := dao.NewVideoDao(ctx).GetIsFollowed(uid, token)
		res.VideoList = append(res.VideoList, &videoPb.Video{
			Id:            int64(video.ID),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			Title:         video.Title,
			FavoriteCount: FavoriteCount,
			CommentCount:  CommentCount,
			IsFavorite:    IsFavorite,
			Author: &videoPb.User{
				Id:              int64(uid),
				Name:            Author.Username,
				Avatar:          Author.Avatar,
				BackgroundImage: Author.BackgroundImage,
				Signature:       Author.Signature,
				FollowCount:     FollowCount,
				FollowerCount:   FollowerCount,
				WorkCount:       WorkCount,
				FavoriteCount:   UserFavoriteCount,
				TotalFavorited:  TotalFavorited,
				IsFollow:        IsFollow,
			},
		})
	}
	return nil
}

func PublishResponseData(res *videoPb.PublishResponse, StatusCode int32, StatusMsg string) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
}
