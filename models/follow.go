package models

type Follow struct {
	ID          int `orm:"column(id);auto;pk" json:"id"`
	FollowerID  int `orm:"column(follower_id)" json:"follower_id"`   // 粉丝的用户ID
	FollowingID int `orm:"column(following_id)" json:"following_id"` // 被关注的用户ID
}
