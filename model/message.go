package model

import "time"

type Message struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement"  json:"id"`
	FromUserID int       `gorm:"column:from_user_id;index:idx_message"  json:"from_user_id"`
	FromUser   User      `gorm:"foreignKey:FromUserID;AssociationForeignKey:ID"  json:"from_user"`
	ToUserID   int       `gorm:"column:to_user_id;index:idx_message"  json:"to_user_id"`
	ToUser     User      `gorm:"foreignKey:ToUserID;AssociationForeignKey:ID"  json:"to_user"`
	Content    string    `gorm:"column:content;size:1024"  json:"content"`
	CreatedAt  time.Time `  json:"created_at"`
}
