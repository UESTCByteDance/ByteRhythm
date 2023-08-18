package model

import "time"

type Comment struct {
	Id        int       `gorm:"column(id);primaryKey;autoIncrement"  json:"id"`
	UserId    int       `gorm:"column(user_id)"  json:"user_id"`
	User      User      `gorm:"foreignKey:UserId;AssociationForeignKey:Id"  json:"user"`
	VideoId   int       `gorm:"column(video_id)"  json:"video_id"`
	Video     Video     `gorm:"foreignKey:VideoId;AssociationForeignKey:Id"  json:"video"`
	Content   string    `gorm:"column(content);size:1024"  json:"content"`
	CreatedAt time.Time ` json:"created_at"`
}
