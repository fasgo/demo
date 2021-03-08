package main

import (
	"github.com/fasgo/base"
	"github.com/fasgo/demo/proto/api"
	"github.com/fasgo/demo/proto/biz"
	"github.com/fasgo/demo/proto/interceptor"
	"github.com/fasgo/protoapi"
	"github.com/fasgo/registry"
	"google.golang.org/grpc"
	"os"
)

func init() {
	// 初始化服务注册/发现
	registry.Init(&registry.Config{
		Grpc: map[string]map[string]string{
			"demo": {
				"127.0.0.1:9080": "3", // 值为连接数, 覆盖grpc.Dial()中指定的默认值
			},
		},
		Http: map[string]map[string]string{
			"demo": {
				"127.0.0.1:8080": "",
			},
		},
		Debug: true,
	})
	// 初始化资源(I18N)文件
	protoapi.InitResource("zh", `E:\pbxworkspace\src\ksogit.kingsoft.net\fasgo\demo\proto\resources`)
}

func main() {

	httpAddr := "127.0.0.1:8080"
	grpcAddr := "127.0.0.1:9080"
	if len(os.Args) > 1 {
		grpcAddr = os.Args[1]
	}
	if len(os.Args) > 2 {
		httpAddr = os.Args[2]
	}

	// 1. 创建server. 方式1: 通过API设置参数. 方式2: 通过conf.toml设置参数
	svr := protoapi.NewServerWith(&protoapi.Config{
		Name:              "demo",   // 注册服务名称
		HttpAddr:          httpAddr, // 开启http访问
		GrpcAddr:          grpcAddr, // 开启grpc访问
		WbskOriginDisable: true,     // websocket不启用origin检测
	})
	/*
		svr := protoapi.NewServer() // 使用conf.toml的[protoapi]配置
	*/

	// 2. 绑定实现
	svr.RegisterService(api.TagServiceRegistry, new(biz.TagServiceService))

	// 连接拦截. 建立http2连接时
	svr.GrpcServerOption(grpc.Creds(new(interceptor.TransportCredentialsTest)))
	// 请求拦截. 每次http2请求时
	svr.GrpcServerOption(grpc.UnaryInterceptor(interceptor.RequestInteceptor))

	// 3. 启动监听
	if err := svr.ListenAndServe(); err != nil {
		base.DefaultLogger.Errorf("server error: %+v", err)
	}
}
