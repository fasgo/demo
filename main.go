package main

import (
	"github.com/fasgo/demo/api"
	"github.com/fasgo/demo/biz"
	"github.com/fasgo/protoapi"
)

func main() {
	// 1. 创建server. 可以从conf配置， 也可以编程配置
	//s := serverWithConfig()
	s := serverWithConfToml()

	// 2. 注册服务实现. 每个service会有对应的Registry/Implement注册到server
	s.RegisterService(api.StudentServiceRegistry, new(biz.StudentServiceService))

	// 3. 监听与服务
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

// 从conf.toml加载配置
func serverWithConfToml() *protoapi.Server {
	// load config from conf.toml
	s := protoapi.NewServer()
	return s
}

// 自己定义配置
func serverWithConfig() *protoapi.Server {
	c := &protoapi.Config{
		HttpAddr: ":80",
		GrpcAddr: ":90",
	}
	s := protoapi.NewServerWith(c)
	return s
}
