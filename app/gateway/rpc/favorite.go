package rpc

import (
	"ByteRhythm/idl/favorite/favoritePb"
	"context"
)

func FavoriteAction(ctx context.Context, req *favoritePb.FavoriteActionRequest) (res *favoritePb.FavoriteActionResponse, err error) {
	res, err = FavoriteService.FavoriteAction(ctx, req)
	if err != nil {
		return
	}
	return
}

func FavoriteList(ctx context.Context, req *favoritePb.FavoriteListRequest) (res *favoritePb.FavoriteListResponse, err error) {
	res, err = FavoriteService.FavoriteList(ctx, req)
	if err != nil {
		return
	}
	return
}
