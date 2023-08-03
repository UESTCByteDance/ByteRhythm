package models

type Favorite struct {
	ID      int `orm:"column(id);auto;pk" json:"id"`
	UserID  int `orm:"column(user_id)" json:"user_id"`   // 用户ID
	VideoID int `orm:"column(video_id)" json:"video_id"` // 视频ID
}
