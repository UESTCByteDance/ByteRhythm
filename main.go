package main

import (
	_ "ByteRhythm/models"
	_ "ByteRhythm/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	//web.InsertFilter("*", web.BeforeRouter, utils.FilterToken)
	web.Run()
}