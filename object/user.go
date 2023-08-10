package object

type UserInfo struct {
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	FavoriteCount   int    `json:"favorite_count"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	ID              int    `json:"id"`
	IsFollow        bool   `json:"is_follow"`
	Name            string `json:"name"`
	Signature       string `json:"signature"`
	TotalFavorited  int    `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
}
