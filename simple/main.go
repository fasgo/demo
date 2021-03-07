package main

import (
	"fmt"
	"github.com/fasgo/demo/simple/hdl"
	"github.com/fasgo/protoapi"
	"time"
)

func main() {

	svr := protoapi.NewServerWith(&protoapi.Config{
		HttpAddr: ":80",
	})

	// access审计, 如果没有Grpc的情况下, svr.HttpAccessOption与svr.Use()效果等同, 后者是root group.
	svr.HttpServerOption(func(ctx *protoapi.Context) {
		// 设置头部
		ctx.Request.Header.Set("x-user-id", "1111")
		ctx.Request.Header.Set("x-team-id", "2222")
		ctx.Request.Header.Set("x-enterprise-id", "2222")

		// 统计时间, 并判断是否带有abort的查询参数, 有的话中止, 否则继续
		start := time.Now().UnixNano()
		if ctx.QueryValue("abort") != "" {
			ctx.Abort()
		} else {
			ctx.Next()
		}
		end := time.Now().UnixNano()
		fmt.Printf("access: path=%v, status=%v, used(ns)=%v\n", ctx.Request.URL.String(), ctx.StatusCode(), end-start)
	})

	// 访问students注册的handleFun都会触发StudentsInterceptor
	students := svr.Group("/simple/students", hdl.StudentsInterceptor)
	students.POST("", hdl.Create)
	students.PUT("/:name", hdl.Update)
	students.DELETE("/:name", hdl.Delete)
	students.GET("/:name", hdl.Retrieve)

	// 访问 Get /simple/students/:name/list 不会触发StudentsInterceptor, 因为其不在students的group router
	svr.GET("/simple/students/:name/list", hdl.Retrieve)
	// 暴露静态资源文件
	svr.Static("/myfiles", "/")

	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}
