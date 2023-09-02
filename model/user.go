package model

import "time"

type User struct {
	ID              uint      `gorm:"column:id;primaryKey;autoIncrement;index" json:"id"`
	Username        string    `gorm:"column:username;unique;index"  json:"username"`
	Password        string    `gorm:"column:password"  json:"password"`
	Avatar          string    `gorm:"column:avatar" json:"avatar"`
	BackgroundImage string    `gorm:"column:background_image"  json:"background_image"`
	Signature       string    `gorm:"column:signature"  json:"signature"`
	CreatedAt       time.Time `json:"created_at"`
}
