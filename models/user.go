package models

import "gorm.io/gorm"

// User Gorm Data Structures
// User和Video是多对多关系，两个模型之间会有一个连接表 user_favorite_videos
type User struct {
	gorm.Model
	UserName       string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string  `gorm:"type:varchar(256);not null" json:"password"`
	FavoriteVideos []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos"`
	FollowingCount int     `gorm:"default:0" json:"following_count"`
	FollowerCount  int     `gorm:"default:0" json:"follower_count"`
}

func (User) TableName() string {
	return "user"
}
