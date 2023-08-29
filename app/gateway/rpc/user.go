package rpc

import (
	"ByteRhythm/idl/user/userPb"
	"context"
)

func Login(ctx context.Context, req *userPb.UserRequest) (res *userPb.UserResponse, err error) {
	res, err = UserService.Login(ctx, req)
	if err != nil {
		return
	}
	return
}

func Register(ctx context.Context, req *userPb.UserRequest) (res *userPb.UserResponse, err error) {
	res, err = UserService.Register(ctx, req)
	if err != nil {
		return
	}
	return
}

func UserInfo(ctx context.Context, req *userPb.UserInfoRequest) (res *userPb.UserInfoResponse, err error) {
	res, err = UserService.UserInfo(ctx, req)
	if err != nil {
		return
	}
	return
}
