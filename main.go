package main

import (
	"github.com/gohade/hade/app/console"
	"github.com/gohade/hade/app/http"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/provider/app"
	"github.com/gohade/hade/framework/provider/distributed"
	"github.com/gohade/hade/framework/provider/env"
	"github.com/gohade/hade/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&env.HadeEnvProvider{})
	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
