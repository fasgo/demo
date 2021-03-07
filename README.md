# demo

这是protoapi的演示项目

## 介绍文档 [protoapi介绍(fasgo)](https://www.kdocs.cn/l/cmJ338WjJQ8e). 

***如果发现任何问题或BUG, 敬请不吝告知 QQ(1255422783), 非常感谢!***

# demo/simple

演示HTTP服务器的相关功能:

## HTTP服务器
- Martini-Link API
- 路径参数 URL-Params
- 分组路由 Group Router
- 请求拦截器 Interceptor
- 上下文辅助工具 Context utilities to http request, http.Request
- 资源及本地化 Resource and Localize(I18N/I10N), 
- 静态文件服务器 Expose static resources


# demo/proto

演示GRPC服务器, HTTP-GRPC网关, 微服务集群相关功能:

## GRPC服务器
- 脚手架 protogen
- GRPC连接拦截/请求拦截
- 资源及本地化 Resource and Localize(I18N/I10N)

## HTTP-GRPC网关
- 使用restful方式请求GRPC服务
- 使用websocket方式请求GRPC服务

## 微服务集群
- 服务自动注册/发现(需要etcd服务器)
- 客户端负载均衡/透明容错