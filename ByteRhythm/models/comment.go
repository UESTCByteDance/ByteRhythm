package models

import "time"

type Comment struct {
	Id         int       `orm:"column(id);pk;auto" description:"评论id" json:"id"`
	UserId     *User     `orm:"column(user_id);rel(fk)" description:"评论用户id" json:"user_id"`
	VideoId    *Video    `orm:"column(video_id);rel(fk)" description:"评论视频id" json:"video_id"`
	Content    string    `orm:"column(content);size(512);null" description:"评论内容" json:"content"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间" json:"create_time"`
}
