package rpc

import (
	"ByteRhythm/idl/user/userPb"
	"ByteRhythm/idl/video/videoPb"

	"go-micro.dev/v4"
)

var (
	UserService  userPb.UserService
	VideoService videoPb.VideoService
)

func InitRPC() {
	UserMicroService := micro.NewService(micro.Name("UserService.client"))
	userService := userPb.NewUserService("UserService", UserMicroService.Client())
	UserService = userService

	VideoMicroService := micro.NewService(micro.Name("VideoService.client"))
	videoService := videoPb.NewVideoService("VideoService", VideoMicroService.Client())
	VideoService = videoService
}
