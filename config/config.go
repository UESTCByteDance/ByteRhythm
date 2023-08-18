package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassWord         string
	DBName             string
	Charset            string
	Avatar             string
	Background         string
	Signature          string
	EtcdHost           string
	EtcdPort           string
	HttpHost           string
	HttpPort           string
	UserServiceAddress string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("load config failed,err:%v\n", err)
	}
	LoadMySQL(file)
	LoadUser(file)
	LoadEtcd(file)
	LoadGin(file)
	LoadService(file)
}

func LoadMySQL(file *ini.File) {
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPassWord = file.Section("mysql").Key("DBPassWord").String()
	DBName = file.Section("mysql").Key("DBName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadUser(file *ini.File) {
	Avatar = file.Section("user").Key("Avatar").String()
	Background = file.Section("user").Key("Background").String()
	Signature = file.Section("user").Key("Signature").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadGin(file *ini.File) {
	HttpHost = file.Section("gin").Key("HttpHost").String()
	HttpPort = file.Section("gin").Key("HttpPort").String()
}
func LoadService(file *ini.File) {
	UserServiceAddress = file.Section("service").Key("UserServiceAddress").String()
}
