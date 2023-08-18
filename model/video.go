package model

import "time"

type Video struct {
	Id        int       `gorm:"column(id);primaryKey;autoIncrement" json:"id"`
	AuthorId  int       `gorm:"column(author_id);"  json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorId;AssociationForeignKey:Id"  json:"author"`
	PlayUrl   string    `gorm:"column(play_url)"  json:"play_url"`
	CoverUrl  string    `gorm:"column(cover_url)"  json:"cover_url"`
	Title     string    `gorm:"column(title)"  json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
