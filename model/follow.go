package model

import "time"

type Follow struct {
	ID             uint      `gorm:"column:id;primaryKey;autoIncrement"  json:"id"`
	UserID         int       `gorm:"column:user_id"  json:"user_id"`
	User           User      `gorm:"foreignKey:UserID;AssociationForeignKey:ID"  json:"user"`
	FollowedUserID int       `gorm:"column:followed_user_id"  json:"followed_user_id"`
	FollowedUser   User      `gorm:"foreignKey:FollowedUserID;AssociationForeignKey:ID"  json:"followed_user"`
	CreatedAt      time.Time `json:"created_at"`
}
