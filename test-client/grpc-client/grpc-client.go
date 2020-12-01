package main

import (
	"context"
	"fmt"
	"github.com/fasgo/demo/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:90", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewStudentServiceClient(conn)

	req := &api.Student{
		Name: "小王",
	}

	rsp, err := client.Add(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rsp)
}

/*
输出:
Load conf success: E:\temp\conf.toml
name:"小王"
*/
