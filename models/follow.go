package models

import "time"

type Follow struct {
	Id             int       `orm:"column(id);pk;auto" description:"关注id" json:"id"`
	UserId         *User     `orm:"column(user_id);rel(fk)" description:"关注用户id" json:"user_id"`
	FollowedUserId *User     `orm:"column(followed_user_id);rel(fk)" description:"粉丝用户id" json:"followed_user_id"`
	CreateTime     time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间" json:"create_time"`
}
