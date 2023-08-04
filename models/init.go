package models

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	fmt.Println("Gorm init")
	InitDB()
}

func InitDB() {
	mysqluser, _ := beego.AppConfig.String("mysqluser")
	mysqlpass, _ := beego.AppConfig.String("mysqlpass")
	mysqlurls, _ := beego.AppConfig.String("mysqlurls")
	mysqldb, _ := beego.AppConfig.String("mysqldb")

	dbStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqluser, mysqlpass, mysqlurls, mysqldb)

	var err error
	DB, err = gorm.Open(mysql.Open(dbStr),
		&gorm.Config{
			//Logger: logger,
		},
	)
	if err != nil {
		fmt.Println(err)
		//这里先忽略生产环境的错误，方便测试
		if beego.BConfig.RunMode != "prod" {
			panic("failed to connect database")
		}
		return
	}

	// 自动迁移（创建）数据库表
	if err := DB.AutoMigrate(&User{}, &Video{}, &Comment{}, &Relation{}); err != nil {
		panic("failed to create tables")
	}

}
