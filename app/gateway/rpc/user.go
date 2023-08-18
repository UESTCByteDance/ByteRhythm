package rpc

import (
	"ByteRhythm/idl/pb"
	"context"
)

func Login(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	res, err = UserService.Login(ctx, req)
	if err != nil {
		res.StatusCode = 1
		res.StatusMsg = err.Error()
		return
	}
	return
}

func Register(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	res, err = UserService.Register(ctx, req)
	if err != nil {
		res.StatusCode = 1
		res.StatusMsg = err.Error()
		return
	}
	return
}

func UserInfo(ctx context.Context, req *pb.UserInfoRequest) (res *pb.UserInfoResponse, err error) {
	res, err = UserService.UserInfo(ctx, req)
	if err != nil {
		res.StatusCode = 1
		res.StatusMsg = err.Error()
		return
	}
	return
}
