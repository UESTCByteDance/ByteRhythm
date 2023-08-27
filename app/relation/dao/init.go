package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"ByteRhythm/config"
)

var db *gorm.DB
var RedisClient *redis.Client

func InitMySQL() {
	host := config.DBHost
	port := config.DBPort
	database := config.DBName
	username := config.DBUser
	password := config.DBPassWord
	charset := config.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}

func InitRedis() {
	// 初始化 Redis 客户端
	host := config.RedisHost
	port := config.RedisPort
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port, // Redis 服务器地址
		Password: "",                // Redis 访问密码（如果有的话）
		DB:       1,                 // Redis 数据库索引
	})
}

func Database(connString string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	db = DB
	migration()
	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	DB := db
	return DB.WithContext(ctx)
}
