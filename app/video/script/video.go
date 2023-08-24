package script

import (
	"ByteRhythm/app/video/mq"
	"ByteRhythm/app/video/service"
	"ByteRhythm/idl/video/videoPb"
	"context"
	"encoding/json"
)

func VideoCreateSync(ctx context.Context) {
	Sync := new(SyncVideo)
	err := Sync.SyncVideoCreate(ctx)
	if err != nil {
		return
	}
}

type SyncVideo struct {
}

func (s *SyncVideo) SyncVideoCreate(ctx context.Context) error {
	RabbitMQName := "video-create-queue"
	msg, err := mq.ConsumeMessage(ctx, RabbitMQName)
	if err != nil {
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msg {
			// 落库
			var req *videoPb.PublishRequest
			err = json.Unmarshal(d.Body, &req)
			if err != nil {
				return
			}
			err = service.VideoMQ2MySQL(ctx, req)
			if err != nil {
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}
