package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	ID            int       `orm:"column(id);auto;pk" json:"id"`
	Username      string    `orm:"column(username);size(50);unique" json:"username"`
	Password      string    `orm:"column(password);size(500)" json:"-"`
	Email         string    `orm:"column(email);size(100)" json:"email"`
	Phone         string    `orm:"column(phone);size(20)" json:"phone"`
	Avatar        string    `orm:"column(avatar);size(255);default(/static/avatar/default.jpg)" json:"avatar"`
	Background    string    `orm:"column(background);size(255);default(/static/background/default.jpg)" json:"background"`
	Signature     string    `orm:"column(signature);size(255)" json:"signature"`
	Gender        int       `orm:"column(gender)" json:"gender"`
	Birthday      time.Time `orm:"column(birthday);type(date)" json:"birthday"`
	Location      string    `orm:"column(location);size(255)" json:"location"`
	FriendCount   int       `orm:"column(friend_count);default(0)" json:"friend_count"`
	FollowerCount int       `orm:"column(follower_count);default(0)" json:"follower_count"`
	FollowCount   int       `orm:"column(follow_count);default(0)" json:"follow_count"`
	LikeCount     int       `orm:"column(like_count);default(0)" json:"like_count"`
	WorkCount     int       `orm:"column(work_count);default(0)" json:"work_count"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Comment), new(Favorite), new(Follow), new(Video))
}
