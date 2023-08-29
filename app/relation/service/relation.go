package service

import (
	"ByteRhythm/app/relation/dao"
	"ByteRhythm/idl/relation/relationPb"
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"
	"sync"
	"time"
)

type RelationSrv struct {
}

var RelationSrvIns *RelationSrv
var RelationSrvOnce sync.Once

// GetRelationSrv 懒汉式的单例模式 lazy-loading --> 懒汉式
func GetRelationSrv() *RelationSrv {
	RelationSrvOnce.Do(func() {
		RelationSrvIns = &RelationSrv{}
	})
	return RelationSrvIns
}

func (r RelationSrv) ActionRelation(ctx context.Context, req *relationPb.RelationActionRequest, res *relationPb.RelationActionResponse) error {
	token := req.Token
	toUserId := req.ToUserId
	actionType := req.ActionType

	fromUserId, err := util.GetUserIdFromToken(token)
	if err != nil {
		return nil
	}

	if int64(fromUserId) == toUserId {
		RelationActionResponseData(res, 1, "不能对自己进行该操作！")
		return nil
	}
	if actionType == 1 {
		// 关注
		user, _ := dao.NewRelationDao(ctx).FindUserById(int(toUserId))
		if user.ID == 0 {
			RelationActionResponseData(res, 1, "用户不存在!")
			return nil
		}

		Relation := BuildRelationModel(int(toUserId), fromUserId)
		if RowsAffected, err := dao.NewRelationDao(ctx).AddFollow(&Relation); err == nil {
			if RowsAffected > 0 {
				RelationActionResponseData(res, 0, "关注成功!")
				return nil
			} else {
				RelationActionResponseData(res, 1, "已经关注过该用户了！")
				return nil
			}
		} else {
			RelationActionResponseData(res, 1, "关注失败!")
			return err
		}

	} else if actionType == 2 {
		// 取消关注
		Relation := BuildRelationModel(int(toUserId), fromUserId)
		if _, err := dao.NewRelationDao(ctx).CancelFollow(&Relation); err == nil {
			RelationActionResponseData(res, 0, "取消关注成功！")
			return nil
		}
		RelationActionResponseData(res, 1, "取消关注失败！")
		return err

	}
	RelationActionResponseData(res, 1, "检查参数信息！")
	return nil

}

func (r RelationSrv) ListFollowRelation(ctx context.Context, req *relationPb.RelationFollowRequest, res *relationPb.RelationFollowResponse) error {
	userId := req.UserId
	token := req.Token
	follows, err := dao.NewRelationDao(ctx).FindAllFollow(int(userId))
	if err != nil {
		RelationFollowResponseData(res, 1, "获取关注列表失败！")
		return err
	}

	for _, follow := range follows {
		user, _ := dao.NewRelationDao(ctx).FindUserById(follow.UserID)
		if user.ID == 0 {
			RelationFollowResponseData(res, 1, "用户不存在!")
			return nil
		}
		res.UserList = append(res.UserList, BuildUserPbModel(ctx, user, token))
	}
	RelationFollowResponseData(res, 0, "获取关注列表成功！")
	return nil
}

func (r RelationSrv) ListFollowerRelation(ctx context.Context, req *relationPb.RelationFollowerRequest, res *relationPb.RelationFollowerResponse) error {
	userId := req.UserId
	token := req.Token
	follows, err := dao.NewRelationDao(ctx).FindAllFollower(int(userId))
	if err != nil {
		RelationFollowerResponseData(res, 1, "获取粉丝列表失败！")
		return err
	}

	for _, follow := range follows {
		user, _ := dao.NewRelationDao(ctx).FindUserById(follow.FollowedUserID)
		if user.ID == 0 {
			RelationFollowerResponseData(res, 1, "用户不存在!")
			return nil
		}
		res.UserList = append(res.UserList, BuildUserPbModel(ctx, user, token))
	}
	RelationFollowerResponseData(res, 0, "获取粉丝列表成功！")
	return nil
}

func (r RelationSrv) ListFriendRelation(ctx context.Context, req *relationPb.RelationFriendRequest, res *relationPb.RelationFriendResponse) error {
	userId := req.UserId
	token := req.Token
	follows, err := dao.NewRelationDao(ctx).FindAllFollower(int(userId))
	if err != nil {
		RelationFriendResponseData(res, 1, "获取粉丝列表失败！")
		return err
	}

	for _, follow := range follows {
		count, _ := dao.NewRelationDao(ctx).GetFriendCount(follow.FollowedUserID, int(userId))
		if count == 1 {
			user, _ := dao.NewRelationDao(ctx).FindUserById(follow.FollowedUserID)
			if user.ID == 0 {
				RelationFriendResponseData(res, 1, "用户不存在!")
				return nil
			}
			res.UserList = append(res.UserList, BuildUserPbModel(ctx, user, token))
		} else {
			continue
		}
	}
	RelationFriendResponseData(res, 0, "获取好友列表成功")
	return nil
}

func RelationActionResponseData(res *relationPb.RelationActionResponse, StatusCode int32, StatusMsg string) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
}

func RelationFollowResponseData(res *relationPb.RelationFollowResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.UserList = params[0].([]*relationPb.User)
	}
}

func RelationFollowerResponseData(res *relationPb.RelationFollowerResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.UserList = params[0].([]*relationPb.User)
	}
}

func RelationFriendResponseData(res *relationPb.RelationFriendResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.UserList = params[0].([]*relationPb.User)
	}
}

func BuildRelationModel(userID int, followedUserID int) model.Follow {
	return model.Follow{
		UserID:         userID,
		FollowedUserID: followedUserID,
		CreatedAt:      time.Now(),
	}
}

func BuildUserPbModel(ctx context.Context, user *model.User, token string) *relationPb.User {
	uid := int(user.ID)
	FollowCount, _ := dao.NewRelationDao(ctx).GetFollowCount(uid)
	FollowerCount, _ := dao.NewRelationDao(ctx).GetFollowerCount(uid)
	WorkCount, _ := dao.NewRelationDao(ctx).GetWorkCount(uid)
	FavoriteCount, _ := dao.NewRelationDao(ctx).GetFavoriteCount(uid)
	TotalFavorited, _ := dao.NewRelationDao(ctx).GetTotalFavorited(uid)
	IsFollow, _ := dao.NewRelationDao(ctx).GetIsFollowed(uid, token)
	return &relationPb.User{
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
