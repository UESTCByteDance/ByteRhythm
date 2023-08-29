package model

import "time"

type Video struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AuthorID  int       `gorm:"column:author_id;index"  json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorID;AssociationForeignKey:ID"  json:"author"`
	PlayUrl   string    `gorm:"column:play_url"  json:"play_url"`
	CoverUrl  string    `gorm:"column:cover_url"  json:"cover_url"`
	Title     string    `gorm:"column:title"  json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
