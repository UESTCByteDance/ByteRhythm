package models

import "time"

type Comment struct {
	ID        int       `orm:"column(id);auto;pk" json:"id"`
	Content   string    `orm:"column(content);size(1000)" json:"content"`
	UserID    int       `orm:"column(user_id)" json:"user_id"`
	VideoID   int       `orm:"column(video_id)" json:"video_id"`
	ParentID  int       `orm:"column(parent_id);default(0)" json:"parent_id"` // 父级评论ID，为0表示一级评论
	LikeCount int       `orm:"column(like_count);default(0)" json:"like_count"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
}
