package main

import (
	"ByteRhythm/app/gateway/router"
	"ByteRhythm/app/gateway/rpc"
	"ByteRhythm/config"
	"fmt"
	"time"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

func main() {
	config.Init()
	rpc.InitRPC()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// 得到一个微服务实例
	webService := web.NewService(
		web.Name("HttpService"), // 微服务名字
		web.Address(fmt.Sprintf("%s:%s", config.HttpHost, config.HttpPort)),
		web.Registry(etcdReg),           // etcd注册件
		web.Handler(router.NewRouter()), // 路由
		web.RegisterTTL(time.Second*30), // 服务注册时间
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	_ = webService.Init()
	_ = webService.Run()
}
