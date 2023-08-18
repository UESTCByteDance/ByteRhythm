package dao

import (
	"context"
	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

func NewVideoDao(ctx context.Context) *VideoDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &VideoDao{NewDBClient(ctx)}
}
