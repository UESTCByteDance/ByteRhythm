package dao

import (
	"ByteRhythm/model"
	"context"

	"gorm.io/gorm"
)

type MessageDao struct {
	*gorm.DB
}

func NewMessageDao(ctx context.Context) *MessageDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &MessageDao{NewDBClient(ctx)}
}

func (m *MessageDao) CreateMessage(message *model.Message) (id int64, err error) {
	err = m.Model(&model.Message{}).Create(&message).Error
	if err != nil {
		return
	}
	return int64(message.ID), nil
}

func (m *MessageDao) FindAllMessages(fromUserID int64, toUserID int64) (messages []*model.Message, err error) {
	err = m.Model(&model.Message{}).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", fromUserID, toUserID, toUserID, fromUserID).Order("id ASC").Find(&messages).Error
	if err != nil {
		return
	}
	return messages, nil
}
