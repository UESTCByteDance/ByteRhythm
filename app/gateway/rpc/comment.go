package rpc

import (
	"ByteRhythm/idl/comment/commentPb"
	"context"
)

func CommentAction(ctx context.Context, req *commentPb.CommentActionRequest) (res *commentPb.CommentActionResponse, err error) {
	res, err = CommentService.CommentAction(ctx, req)
	if err != nil {
		return
	}
	return
}

func CommentList(ctx context.Context, req *commentPb.CommentListRequest) (res *commentPb.CommentListResponse, err error) {
	res, err = CommentService.CommentList(ctx, req)
	if err != nil {
		return
	}
	return
}
