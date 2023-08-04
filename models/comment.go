package models

import "gorm.io/gorm"

// Comment 属于 Video, VideoID是外键(belongs to)
// Comment 也属于 User,UserID是外键(belongs to)
type Comment struct {
	gorm.Model
	Video   Video  `gorm:"foreignkey:VideoID"`
	VideoID int    `gorm:"index:idx_videoid;not null"`
	User    User   `gorm:"foreignkey:UserID"`
	UserID  int    `gorm:"index:idx_userid;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}

func (Comment) TableName() string {
	return "comment"
}
