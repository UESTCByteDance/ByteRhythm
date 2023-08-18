package model

import "time"

type Message struct {
	Id         int       `gorm:"column(id);primaryKey;autoIncrement"  json:"id"`
	FromUserId int       `gorm:"column(from_user_id)"  json:"from_user_id"`
	FromUser   User      `gorm:"foreignKey:FromUserId;AssociationForeignKey:Id"  json:"from_user"`
	ToUserId   int       `gorm:"column(to_user_id)"  json:"to_user_id"`
	ToUser     User      `gorm:"foreignKey:ToUserId;AssociationForeignKey:Id"  json:"to_user"`
	Content    string    `gorm:"column(content);size(1024)"  json:"content"`
	CreateTime time.Time `  json:"created_at"`
}
