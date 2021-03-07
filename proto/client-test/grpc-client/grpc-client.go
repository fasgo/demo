package main

import (
	"context"
	"fmt"
	"github.com/fasgo/demo/proto/api"
	"github.com/fasgo/base/kits"
	"github.com/fasgo/registry"
	"github.com/fasgo/registry/grpc"
	"time"
)

func init() {
	registry.Init(&registry.Config{
		Endpoint: "10.13.144.164:2379",
		Username: "root",
		Password: "123456",
		Debug:    false,
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
