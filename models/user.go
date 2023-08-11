package models

import (
	"time"
)

type User struct {
	Id              int       `orm:"column(id);pk;auto" description:"用户id" json:"id"`
	Username        string    `orm:"column(username);unique;size(255)" description:"用户名" json:"username"`
	Password        string    `orm:"column(password);size(255)" description:"密码" json:"password"`
	Avatar          string    `orm:"column(avatar);size(255);default(http://rz2n87yck.hn-bkt.clouddn.com/avatar.jpg)" description:"用户头像" json:"avatar"`
	BackgroundImage string    `orm:"column(background_image);size(255);default(http://rz2n87yck.hn-bkt.clouddn.com/background.jpg)" description:"用户个人页顶部大图" json:"background_image"`
	Signature       string    `orm:"column(signature);size(255);null" description:"个人简介" json:"signature"`
	CreateTime      time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间" json:"create_time"`
}
