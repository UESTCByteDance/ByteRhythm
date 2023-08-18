package service

import (
	"ByteRhythm/app/user/dao"
	"ByteRhythm/config"
	"ByteRhythm/idl/user/userPb"
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"
	"fmt"
	"sync"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

// GetUserSrv 懒汉式的单例模式 lazy-loading --> 懒汉式
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) Login(ctx context.Context, req *userPb.UserRequest, res *userPb.UserResponse) (err error) {
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.Username)
	if user.Id == 0 || util.Md5(req.Password) != user.Password {
		res.StatusCode = 1
		res.StatusMsg = "用户名或密码错误"
		return nil
	}
	token := util.GenerateToken(user, 0)
	uid := int64(user.Id)
	res.StatusCode = 0
	res.StatusMsg = "登录成功"
	res.UserId = uid
	res.Token = token
	return nil
}

func (u *UserSrv) Register(ctx context.Context, req *userPb.UserRequest, res *userPb.UserResponse) (err error) {
	username := req.Username
	password := req.Password
	if len(username) > 32 || len(password) > 32 {
		res.StatusCode = 1
		res.StatusMsg = "用户名或密码不能超过32位"
		return nil
	}
	if user, _ := dao.NewUserDao(ctx).FindUserByUserName(username); user.Id != 0 {
		res.StatusCode = 1
		res.StatusMsg = "用户名已存在"
		return nil
	}

	config.Init()
	avatar := config.Avatar
	background := config.Background
	signature := config.Signature
	user := model.User{
		Username:        username,
		Password:        util.Md5(password),
		Avatar:          avatar,
		BackgroundImage: background,
		Signature:       signature,
	}
	if id, err := dao.NewUserDao(ctx).CreateUser(&user); err != nil {
		res.StatusCode = 1
		res.StatusMsg = "注册失败"
		return err
	} else {
		token := util.GenerateToken(&user, 0)
		res.StatusCode = 0
		res.StatusMsg = "注册成功"
		res.UserId = id
		res.Token = token
		return nil
	}
}

func (u *UserSrv) UserInfo(ctx context.Context, req *userPb.UserInfoRequest, res *userPb.UserInfoResponse) error {
	token := req.Token
	uid := req.UserId
	if token != "" {
		if err := util.ValidateToken(token); err != nil {
			fmt.Println(err)
			res.StatusCode = 1
			res.StatusMsg = "token验证失败"
			return nil
		}
	}

	user, _ := dao.NewUserDao(ctx).FindUserById(uid)
	if user.Id == 0 {
		res.StatusCode = 1
		res.StatusMsg = "获取用户信息失败"
		return nil
	}

	//返回User
	res.StatusCode = 0
	res.StatusMsg = "获取用户信息成功"

	FollowCount, _ := dao.NewUserDao(ctx).GetFollowCount(uid)
	FollowerCount, _ := dao.NewUserDao(ctx).GetFollowerCount(uid)
	WorkCount, _ := dao.NewUserDao(ctx).GetWorkCount(uid)
	FavoriteCount, _ := dao.NewUserDao(ctx).GetFavoriteCount(uid)
	TotalFavorited, _ := dao.NewUserDao(ctx).GetTotalFavorited(uid)
	IsFollow, _ := dao.NewUserDao(ctx).GetIsFollowed(uid, token)
	res.User = &userPb.User{
		Id:              int64(user.Id),
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
	return nil

}
