package utils

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	// 导入orm包
	"github.com/beego/beego/v2/client/orm"
	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库连接
func init() {
	//0.获取数据库配置
	mysqlName, _ := beego.AppConfig.String("mysqlName")
	mysqlUser, _ := beego.AppConfig.String("mysqlUser")
	mysqlPass, _ := beego.AppConfig.String("mysqlPass")
	mysqlAdds, _ := beego.AppConfig.String("mysqlAdds")
	mysqlPort, _ := beego.AppConfig.String("mysqlPort")
	mysqlChar, _ := beego.AppConfig.String("mysqlChar")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", mysqlUser, mysqlPass, mysqlAdds, mysqlPort, mysqlName, mysqlChar)

	fmt.Println(dataSource)

	//1.注册数据库驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println("注册数据库驱动出错" + err.Error())
		return
	}
	//2.连接数据库
	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		fmt.Println("数据库连接出错" + err.Error())
		return
	}
	//3.设置最大数据库连接
	orm.SetMaxOpenConns("default", 10)
	//4.设置最大数据库空闲连接
	orm.SetMaxIdleConns("default", 10)

	fmt.Println("数据库连接成功")

	orm.Debug = true

	//err = orm.RunSyncdb("default", false, true)
	//if err != nil {
	//	return
	//}
}
