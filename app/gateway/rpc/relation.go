package rpc

import (
	"ByteRhythm/idl/relation/relationPb"
	"context"
)

func ActionRelation(ctx context.Context, req *relationPb.RelationActionRequest) (res *relationPb.RelationActionResponse, err error) {
	res, err = RelationService.ActionRelation(ctx, req)
	if err != nil {
		return
	}
	return
}

func ListFollowRelation(ctx context.Context, req *relationPb.RelationFollowRequest) (res *relationPb.RelationFollowResponse, err error) {
	res, err = RelationService.ListFollowRelation(ctx, req)
	if err != nil {
		return
	}
	return
}

func ListFollowerRelation(ctx context.Context, req *relationPb.RelationFollowerRequest) (res *relationPb.RelationFollowerResponse, err error) {
	res, err = RelationService.ListFollowerRelation(ctx, req)
	if err != nil {
		return
	}
	return
}

func ListFriendRelation(ctx context.Context, req *relationPb.RelationFriendRequest) (res *relationPb.RelationFriendResponse, err error) {
	res, err = RelationService.ListFriendRelation(ctx, req)
	if err != nil {
		return
	}
	return
}
