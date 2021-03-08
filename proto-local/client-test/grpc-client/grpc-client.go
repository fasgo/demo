package main

import (
	"context"
	"fmt"
	"github.com/fasgo/base/kits"
	"github.com/fasgo/demo/proto/api"
	"github.com/fasgo/registry"
	"github.com/fasgo/registry/grpc"
	"time"
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
		Debug: false,
	})
}

func main() {

	// 指定服务名称与每个addr的http2连接数量(轮询策略)
	cc, err := grpc.Dial("demo", 2)

	if err != nil {
		panic(err)
	}
	defer cc.Close()

	cl := api.NewTagServiceClient(cc)

	//TestGracefulRestart(cl)
	TestInteceptor(cl)

}

func TestInteceptor(cl api.TagServiceClient) {
	for i := 0; ; i++ {
		rsp, err := cl.All(context.Background(), &api.AllReq{
			From:   0,
			Size:   100,
			Search: "data",
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v: %v\n", i, kits.ToJson(rsp))
		time.Sleep(time.Second)
	}
}

func TestGracefulRestart(cl api.TagServiceClient) {
	for i := 0; ; i++ {
		rsp, err := cl.Get(context.Background(), &api.Student{
			Sno: "1234",
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v: %v\n", i, kits.ToJson(rsp))
	}
}
