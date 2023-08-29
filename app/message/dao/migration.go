package dao

import (
	"ByteRhythm/model"
	"log"
)

func migration() {
	err := db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.Message{}, &model.User{})
	if err != nil {
		log.Fatal(err)
	}
}
