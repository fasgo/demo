package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fasgo/demo/api"
	"io"
	"net/http"
	"os"
)

func main() {
	tag := &api.Student{
		Name: "demo",
	}
	bs, _ := json.Marshal(tag)
	rsp, err := http.Post("http://127.0.0.1/demo/students", "application/json", bytes.NewReader(bs))
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	fmt.Println(rsp.Proto, rsp.Status)
	io.Copy(os.Stdout, rsp.Body)
}

/*
输出:
Load conf success: E:\temp\conf.toml
HTTP/1.1 200 OK
{"code":0,"data":{"name":"demo"},"tag":"StudentService.Add"}
*/