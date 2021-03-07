package main

import (
	"fmt"
	"github.com/fasgo/demo/proto/api"
	"github.com/fasgo/base/kits"
	"github.com/fasgo/registry"
	"github.com/fasgo/registry/http"
)

func init() {
	registry.Init(&registry.Config{
		Endpoint: "10.13.144.164:2379",
		Username: "root",
		Password: "123456",
		Debug:    true,
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
