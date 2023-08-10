package models

import (
	"time"
)

type Favorite struct {
	Id         int       `orm:"column(id);pk;auto" description:"喜欢id" json:"id"`
	UserId     *User     `orm:"column(user_id);rel(fk)" description:"点赞用户id" json:"user_id"`
	VideoId    *Video    `orm:"column(video_id);rel(fk)" description:"点赞视频id" json:"video_id"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间" json:"create_time"`
}
