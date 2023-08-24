package service

import (
	"ByteRhythm/app/user/dao"
	"ByteRhythm/config"
	"ByteRhythm/idl/user/userPb"
	"ByteRhythm/model"

	"ByteRhythm/util"
	"context"
	"sync"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) Login(ctx context.Context, req *userPb.UserRequest, res *userPb.UserResponse) (err error) {
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.Username)
	if user.ID == 0 || util.Md5(req.Password) != user.Password {
		UserResponseData(res, 1, "用户名或密码错误")
		return nil
	}
	token := util.GenerateToken(user, 0)
	uid := int64(user.ID)
	UserResponseData(res, 0, "登录成功", uid, token)
	return nil
}

func (u *UserSrv) Register(ctx context.Context, req *userPb.UserRequest, res *userPb.UserResponse) (err error) {
	username := req.Username
	password := req.Password
	if len(username) > 32 || len(password) > 32 {
		UserResponseData(res, 1, "用户名或密码不能超过32位")
		return nil
	}
	if user, _ := dao.NewUserDao(ctx).FindUserByUserName(username); user.ID != 0 {
		UserResponseData(res, 1, "用户名已存在")
		return nil
	}

	user := BuildUserModel(username, password)
	if id, err := dao.NewUserDao(ctx).CreateUser(&user); err != nil {
		UserResponseData(res, 1, "注册失败")
		return err
	} else {
		token := util.GenerateToken(&user, 0)
		UserResponseData(res, 0, "注册成功", id, token)
		return nil
	}
}

func (u *UserSrv) UserInfo(ctx context.Context, req *userPb.UserInfoRequest, res *userPb.UserInfoResponse) error {
	token := req.Token
	uid := int(req.UserId)

	user, _ := dao.NewUserDao(ctx).FindUserById(uid)
	if user.ID == 0 {
		UserInfoResponseData(res, 1, "用户不存在")
		return nil
	}

	User := BuildUserPbModel(ctx, user, token)
	UserInfoResponseData(res, 0, "获取用户信息成功", User)
	return nil
}

func UserResponseData(res *userPb.UserResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.UserId = params[0].(int64)
		res.Token = params[1].(string)
	}
}

func UserInfoResponseData(res *userPb.UserInfoResponse, StatusCode int32, StatusMsg string, params ...interface{}) {
	res.StatusCode = StatusCode
	res.StatusMsg = StatusMsg
	if len(params) != 0 {
		res.User = params[0].(*userPb.User)
	}
}

func BuildUserModel(username string, password string) model.User {
	config.Init()
	avatar := config.Avatar
	background := config.Background
	signature := config.Signature
	return model.User{
		Username:        username,
		Password:        util.Md5(password),
		Avatar:          avatar,
		BackgroundImage: background,
		Signature:       signature,
	}
}
func BuildUserPbModel(ctx context.Context, user *model.User, token string) *userPb.User {
	uid := int(user.ID)
	FollowCount, _ := dao.NewUserDao(ctx).GetFollowCount(uid)
	FollowerCount, _ := dao.NewUserDao(ctx).GetFollowerCount(uid)
	WorkCount, _ := dao.NewUserDao(ctx).GetWorkCount(uid)
	FavoriteCount, _ := dao.NewUserDao(ctx).GetFavoriteCount(uid)
	TotalFavorited, _ := dao.NewUserDao(ctx).GetTotalFavorited(uid)
	IsFollow, _ := dao.NewUserDao(ctx).GetIsFollowed(uid, token)
	return &userPb.User{
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
