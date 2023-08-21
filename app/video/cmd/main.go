package main

import (
	"ByteRhythm/app/video/dao"
	"ByteRhythm/app/video/service"
	"ByteRhythm/config"
	"ByteRhythm/idl/video/videoPb"
	"fmt"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitMySQL()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("VideoService"), // 微服务名字
		micro.Address(config.VideoServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = videoPb.RegisterVideoServiceHandler(microService.Server(), service.GetVideoSrv())
	// 启动微服务
	_ = microService.Run()
}
