package object

type MessageDto struct {
	Id         int    ` json:"id"`
	FromUserId int    ` json:"from_user_id"`
	ToUserId   int    ` json:"to_user_id"`
	Content    string ` json:"content"`
	CreateTime int    ` json:"create_time"`
}
