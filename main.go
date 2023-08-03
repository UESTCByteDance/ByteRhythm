package main

import (
	_ "ByteRhythm/models"
	_ "ByteRhythm/routers"
	"ByteRhythm/utils"
	"github.com/beego/beego/v2/client/orm"

	"github.com/beego/beego/v2/server/web"
)

func main() {

	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	utils.InitMySQL()
	orm.RunCommand()
	web.Run()
}
