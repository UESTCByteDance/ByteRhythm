package models

import "time"

type Video struct {
	Id         int       `orm:"column(id);pk;auto" json:"id"`
	AuthorId   *User     `orm:"column(author_id);rel(fk)" description:"作者id" json:"author_id"`
	PlayUrl    string    `orm:"column(play_url);size(255)" description:"视频播放地址" json:"play_url"`
	CoverUrl   string    `orm:"column(cover_url);size(255)" description:"视频封面地址" json:"cover_url"`
	Title      string    `orm:"column(title);size(255)" description:"视频标题" json:"title"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间" json:"create_time"`
}
