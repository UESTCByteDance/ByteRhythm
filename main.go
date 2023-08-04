package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"my-tiktok/models"
	_ "my-tiktok/routers"
)

func main() {
	models.Init()
	beego.Run()
}
