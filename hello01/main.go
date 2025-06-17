package main

import (
	"flag"
	"fmt"

	"hello01/internal/config"
	"hello01/internal/handler"
	"hello01/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/hello01-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) //加载配置文件
	fmt.Println(c.Address)
	server := rest.MustNewServer(c.RestConf) //开启http服务
	defer server.Stop()

	ctx := svc.NewServiceContext(c)       //服务上下文，依赖注入，将需要的配置信息进行注入
	handler.RegisterHandlers(server, ctx) //注册路由

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
