package rpc

import (
	"ByteRhythm/idl/pb"

	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
)

func InitRPC() {
	UserMicroService := micro.NewService(micro.Name("UserService.client"))
	userService := pb.NewUserService("UserService", UserMicroService.Client())
	UserService = userService
}
