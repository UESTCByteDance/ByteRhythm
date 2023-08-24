package dao

import (
	"ByteRhythm/model"
	"log"
)

func migration() {
	err := db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Follow{})
	if err != nil {
		log.Fatal(err)
	}
}
