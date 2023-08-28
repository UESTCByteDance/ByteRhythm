package config

import (
	"fmt"
	"sync"

	"gopkg.in/ini.v1"
)

var (
	wg                     sync.WaitGroup
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassWord             string
	DBName                 string
	Charset                string
	Avatar                 string
	Background             string
	Signature              string
	EtcdHost               string
	EtcdPort               string
	RedisHost              string
	RedisPort              string
	JaegerHost             string
	JaegerPort             string
	HttpHost               string
	HttpPort               string
	UserServiceAddress     string
	VideoServiceAddress    string
	MessageServiceAddress  string
	CommentServiceAddress  string
	RelationServiceAddress string
	FavoriteServiceAddress string
	Bucket                 string
	AccessKey              string
	SecretKey              string
	Domain                 string
	RabbitMQ               string
	RabbitMQHost           string
	RabbitMQPort           string
	RabbitMQUser           string
	RabbitMQPassWord       string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("load config failed,err:%v\n", err)
	}
	wg.Add(9)
	go LoadMySQL(file)
	go LoadUser(file)
	go LoadEtcd(file)
	go LoadJaeger(file)
	go LoadGin(file)
	go LoadService(file)
	go LoadQiNiuYun(file)
	go LoadRabbitMQ(file)
	go LoadRedis(file)
	wg.Wait()
}

func LoadMySQL(file *ini.File) {
	DBHost = file.Section("MySQL").Key("DBHost").String()
	DBPort = file.Section("MySQL").Key("DBPort").String()
	DBUser = file.Section("MySQL").Key("DBUser").String()
	DBPassWord = file.Section("MySQL").Key("DBPassWord").String()
	DBName = file.Section("MySQL").Key("DBName").String()
	Charset = file.Section("MySQL").Key("Charset").String()
	wg.Done()
}

func LoadUser(file *ini.File) {
	Avatar = file.Section("User").Key("Avatar").String()
	Background = file.Section("User").Key("Background").String()
	Signature = file.Section("User").Key("Signature").String()
	wg.Done()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("Etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("Etcd").Key("EtcdPort").String()
	wg.Done()
}

func LoadJaeger(file *ini.File) {
	JaegerHost = file.Section("Jaeger").Key("JaegerHost").String()
	JaegerPort = file.Section("Jaeger").Key("JaegerPort").String()
	wg.Done()
}

func LoadGin(file *ini.File) {
	HttpHost = file.Section("Gin").Key("HttpHost").String()
	HttpPort = file.Section("Gin").Key("HttpPort").String()
	wg.Done()
}

func LoadService(file *ini.File) {
	UserServiceAddress = file.Section("Service").Key("UserServiceAddress").String()
	VideoServiceAddress = file.Section("Service").Key("VideoServiceAddress").String()
	MessageServiceAddress = file.Section("Service").Key("MessageServiceAddress").String()
	CommentServiceAddress = file.Section("Service").Key("CommentServiceAddress").String()
	RelationServiceAddress = file.Section("Service").Key("RelationServiceAddress").String()
	FavoriteServiceAddress = file.Section("Service").Key("FavoriteServiceAddress").String()
	wg.Done()
}

func LoadQiNiuYun(file *ini.File) {
	Bucket = file.Section("QiNiuYun").Key("Bucket").String()
	AccessKey = file.Section("QiNiuYun").Key("AccessKey").String()
	SecretKey = file.Section("QiNiuYun").Key("SecretKey").String()
	Domain = file.Section("QiNiuYun").Key("Domain").String()
	wg.Done()
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("RabbitMQ").Key("RabbitMQ").String()
	RabbitMQHost = file.Section("RabbitMQ").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("RabbitMQ").Key("RabbitMQPort").String()
	RabbitMQUser = file.Section("RabbitMQ").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("RabbitMQ").Key("RabbitMQPassWord").String()
	wg.Done()
}

func LoadRedis(file *ini.File) {
	RedisHost = file.Section("Redis").Key("RedisHost").String()
	RedisPort = file.Section("Redis").Key("RedisPort").String()
	wg.Done()
}
