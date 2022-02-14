package main

import (
	"time"

	services "code.huawink.com/beiwanglu/api-gateway/services"
	weblib "code.huawink.com/beiwanglu/api-gateway/weblib"
	wrappers "code.huawink.com/beiwanglu/api-gateway/wrappers"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	//user
	//注册etcd服务
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//micro注册微服务
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	//用户服务调取实例
	userService := services.NewUserService("rpcUserService", userMicroService.Client())
	// task
	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	//用户服务调取实例
	taskService := services.NewTaskService("rpcTaskService", taskMicroService.Client())
	server := web.NewService(
		web.Name("httpService"),
		web.Address(":4000"),
		web.Handler(weblib.NewRouter(userService, taskService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	_ = server.Init()
	_ = server.Run()
}
