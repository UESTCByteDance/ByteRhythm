package models

import (
	"time"
)

type Message struct {
	Id         int       `orm:"column(id);pk;auto" description:"消息id" json:"id"`
	FromUserId *User     `orm:"column(from_user_id);rel(fk)" description:"消息发送者id" json:"from_user_id"`
	ToUserId   *User     `orm:"column(to_user_id);rel(fk)" description:"消息接收者id" json:"to_user_id"`
	Content    string    `orm:"column(content);size(1024)" description:"消息内容" json:"content"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间" json:"create_time"`
}
