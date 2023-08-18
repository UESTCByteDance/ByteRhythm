package service

import (
	"ByteRhythm/idl/video/videoPb"
	"context"
	"sync"
)

type VideoSrv struct {
}

var VideoSrvIns *VideoSrv
var VideoSrvOnce sync.Once

func GetVideoSrv() *VideoSrv {
	VideoSrvOnce.Do(func() {
		VideoSrvIns = &VideoSrv{}
	})
	return VideoSrvIns
}

func (v *VideoSrv) Feed(ctx context.Context, req *videoPb.FeedRequest, res *videoPb.FeedResponse) error {
	//TODO implement me
	panic("implement me")
}

func (v *VideoSrv) Publish(ctx context.Context, req *videoPb.PublishRequest, res *videoPb.PublishResponse) error {
	//TODO implement me
	panic("implement me")
}

func (v *VideoSrv) PublishList(ctx context.Context, req *videoPb.PublishListRequest, res *videoPb.PublishListResponse) error {
	//TODO implement me
	panic("implement me")
}
