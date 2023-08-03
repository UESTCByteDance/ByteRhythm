package models

import "time"

type Video struct {
	ID            int       `orm:"column(id);auto;pk" json:"id"`
	User          int       `orm:"column(user_id)" json:"user_id"`
	Title         string    `orm:"column(title);size(255)" json:"title"`
	Introduction  string    `orm:"column(introduction);size(1000)" json:"introduction"`
	PlayURL       string    `orm:"column(play_url);size(255)" json:"play_url"`
	CoverURL      string    `orm:"column(cover_url);size(255)" json:"cover_url"`
	FavoriteCount int       `orm:"column(favorite_count);default(0)" json:"favorite_count"`
	CommentCount  int       `orm:"column(comment_count);default(0)" json:"comment_count"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime)" json:"-"`
}
