package script

import (
	"ByteRhythm/app/favorite/service"
	"ByteRhythm/idl/favorite/favoritePb"
	"ByteRhythm/mq"
	"context"
	"encoding/json"
)

func FavoriteCreateSync(ctx context.Context) {
	Sync := new(SyncFavorite)
	err := Sync.SyncFavoriteCreate(ctx)
	if err != nil {
		return
	}
}

func FavoriteDeleteSync(ctx context.Context) {
	Sync := new(SyncFavorite)
	err := Sync.SyncFavoriteDelete(ctx)
	if err != nil {
		return
	}
}

type SyncFavorite struct {
}

func (s *SyncFavorite) SyncFavoriteCreate(ctx context.Context) error {
	RabbitMQName := "favorite-create-queue"
	msg, err := mq.ConsumeMessage(ctx, RabbitMQName)
	if err != nil {
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msg {
			// 落库
			var req *favoritePb.FavoriteActionRequest
			err = json.Unmarshal(d.Body, &req)
			if err != nil {
				return
			}
			err = service.FavoriteMQ2DB(ctx, req)
			if err != nil {
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}

func (s *SyncFavorite) SyncFavoriteDelete(ctx context.Context) error {
	RabbitMQName := "favorite-delete-queue"
	msg, err := mq.ConsumeMessage(ctx, RabbitMQName)
	if err != nil {
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msg {
			// 落库
			var req *favoritePb.FavoriteActionRequest
			err = json.Unmarshal(d.Body, &req)
			if err != nil {
				return
			}
			err = service.FavoriteMQ2DB(ctx, req)
			if err != nil {
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}
