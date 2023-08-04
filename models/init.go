package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	username, _ := web.AppConfig.String("username")
	password, _ := web.AppConfig.String("password")
	host, _ := web.AppConfig.String("host")
	port, _ := web.AppConfig.String("port")
	database, _ := web.AppConfig.String("database")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local", username, password, host, port, database)
	err := orm.RegisterDataBase("default", "mysql", datasource)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	orm.RegisterModel(new(User), new(Comment), new(Favorite), new(Follow), new(Video))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}
