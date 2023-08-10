package object

type VideoInfo struct {
	Author        *UserInfo `json:"author"`
	CommentCount  int       `json:"comment_count"`
	CoverURL      string    `json:"cover_url"`
	FavoriteCount int       `json:"favorite_count"`
	ID            int       `json:"id"`
	IsFavorite    bool      `json:"is_favorite"`
	PlayURL       string    `json:"play_url"`
	Title         string    `json:"title"`
}
