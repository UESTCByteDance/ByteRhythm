package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
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
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("load config failed,err:%v\n", err)
	}
	LoadMySQL(file)
	LoadUser(file)
	LoadEtcd(file)
	LoadJaeger(file)
	LoadGin(file)
	LoadService(file)
	LoadQiNiuYun(file)
}

func LoadMySQL(file *ini.File) {
	DBHost = file.Section("MySQL").Key("DBHost").String()
	DBPort = file.Section("MySQL").Key("DBPort").String()
	DBUser = file.Section("MySQL").Key("DBUser").String()
	DBPassWord = file.Section("MySQL").Key("DBPassWord").String()
	DBName = file.Section("MySQL").Key("DBName").String()
	Charset = file.Section("MySQL").Key("Charset").String()
}

func LoadUser(file *ini.File) {
	Avatar = file.Section("User").Key("Avatar").String()
	Background = file.Section("User").Key("Background").String()
	Signature = file.Section("User").Key("Signature").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("Etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("Etcd").Key("EtcdPort").String()
}

func LoadJaeger(file *ini.File) {
	JaegerHost = file.Section("Jaeger").Key("JaegerHost").String()
	JaegerPort = file.Section("Jaeger").Key("JaegerPort").String()
}

func LoadGin(file *ini.File) {
	HttpHost = file.Section("Gin").Key("HttpHost").String()
	HttpPort = file.Section("Gin").Key("HttpPort").String()
}

func LoadService(file *ini.File) {
	UserServiceAddress = file.Section("Service").Key("UserServiceAddress").String()
	VideoServiceAddress = file.Section("Service").Key("VideoServiceAddress").String()
	MessageServiceAddress = file.Section("Service").Key("MessageServiceAddress").String()
	CommentServiceAddress = file.Section("Service").Key("CommentServiceAddress").String()
	RelationServiceAddress = file.Section("Service").Key("RelationServiceAddress").String()
	FavoriteServiceAddress = file.Section("Service").Key("FavoriteServiceAddress").String()
}

func LoadQiNiuYun(file *ini.File) {
	Bucket = file.Section("QiNiuYun").Key("Bucket").String()
	AccessKey = file.Section("QiNiuYun").Key("AccessKey").String()
	SecretKey = file.Section("QiNiuYun").Key("SecretKey").String()
	Domain = file.Section("QiNiuYun").Key("Domain").String()
}
