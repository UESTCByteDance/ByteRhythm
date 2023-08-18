package model

import "time"

type Follow struct {
	Id             int       `gorm:"column(id);primaryKey;autoIncrement"  json:"id"`
	UserId         int       `gorm:"column(user_id) "  json:"user_id"`
	User           User      `gorm:"foreignKey:UserId;AssociationForeignKey:Id"  json:"user"`
	FollowedUserId int       `gorm:"column(followed_user_id)"  json:"followed_user_id"`
	FollowedUser   User      `gorm:"foreignKey:FollowedUserId;AssociationForeignKey:Id"  json:"followed_user"`
	CreatedAt      time.Time `json:"created_at"`
}
