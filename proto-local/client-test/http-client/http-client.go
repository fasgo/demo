package main

import (
	"fmt"
	"github.com/fasgo/base/kits"
	"github.com/fasgo/demo/proto/api"
	"github.com/fasgo/registry"
	"github.com/fasgo/registry/http"
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
}

func main() {
	var rsp = new(api.AllRsp)
	_, _, _, err := http.JsonResult("demo", http.GET, "/demo/students", nil, &api.AllReq{
		Search: "all",
		From:   0,
		Size:   10,
	}, &rsp)
	if err != nil {
		panic(err)
	}
	fmt.Println("total: ", rsp.Total)
	for i, v := range rsp.Data {
		fmt.Printf("%v: %v\n", i, kits.ToJson(v))
	}
}
