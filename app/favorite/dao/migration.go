package dao

import (
	"ByteRhythm/model"
	"log"
)

func migration() {
	err := db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.Follow{}, &model.Video{}, &model.Favorite{}, &model.Comment{})
	if err != nil {
		log.Fatal(err)
	}
}
