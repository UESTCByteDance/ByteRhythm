package main

import (
	"ByteRhythm/app/comment/dao"
	"ByteRhythm/app/comment/service"
	"ByteRhythm/config"
	"ByteRhythm/idl/comment/commentPb"
	"github.com/go-micro/plugins/v4/registry/etcd"

	"fmt"

	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitMySQL()
	dao.InitRedis()

	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("CommentService"), // 微服务名字
		micro.Address(config.CommentServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
	)

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = commentPb.RegisterCommentServiceHandler(microService.Server(), service.GetCommentSrv())
	// 启动微服务
	_ = microService.Run()
}
