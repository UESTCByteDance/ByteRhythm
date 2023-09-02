package model

import "time"

type Comment struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement"  json:"id"`
	UserID    int       `gorm:"column:user_id"  json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;AssociationForeignKey:ID"  json:"user"`
	VideoID   int       `gorm:"column:video_id;index"  json:"video_id"`
	Video     Video     `gorm:"foreignKey:VideoID;AssociationForeignKey:ID"  json:"video"`
	Content   string    `gorm:"column:content;size:1024"  json:"content"`
	CreatedAt time.Time ` json:"created_at"`
}
