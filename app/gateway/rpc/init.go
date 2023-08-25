package rpc

import (
	"ByteRhythm/idl/comment/commentPb"
	"ByteRhythm/idl/favorite/favoritePb"
	"ByteRhythm/idl/message/messagePb"
	"ByteRhythm/idl/relation/relationPb"
	"ByteRhythm/idl/user/userPb"
	"ByteRhythm/idl/video/videoPb"

	"go-micro.dev/v4"
)

var (
	UserService     userPb.UserService
	VideoService    videoPb.VideoService
	MessageService  messagePb.MessageService
	RelationService relationPb.RelationService
	CommentService  commentPb.CommentService
	FavoriteService favoritePb.FavoriteService
)

func InitRPC() {
	UserMicroService := micro.NewService(micro.Name("UserService.client"))
	userService := userPb.NewUserService("UserService", UserMicroService.Client())
	UserService = userService

	VideoMicroService := micro.NewService(micro.Name("VideoService.client"))
	videoService := videoPb.NewVideoService("VideoService", VideoMicroService.Client())
	VideoService = videoService

	MessageMicroService := micro.NewService(micro.Name("MessageService.client"))
	messageService := messagePb.NewMessageService("MessageService", MessageMicroService.Client())
	MessageService = messageService

	RelationMicroService := micro.NewService(micro.Name("RelationService.client"))
	relationService := relationPb.NewRelationService("RelationService", RelationMicroService.Client())
	RelationService = relationService

	CommentMicroService := micro.NewService(micro.Name("CommentService.client"))
	commentService := commentPb.NewCommentService("CommentService", CommentMicroService.Client())
	CommentService = commentService

	FavoriteMicroService := micro.NewService(micro.Name("FavoriteService.client"))
	favoriteService := favoritePb.NewFavoriteService("FavoriteService", FavoriteMicroService.Client())
	FavoriteService = favoriteService
}
