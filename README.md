# fasgo

## 初衷

打造一套[googleapis](https://github.com/googleapis/googleapis)的技术骨架, 使用者借此构建高速的api服务!

## 目标

fasgo使用protocol buffers version3(proto3)作为接口描述语言(IDL)定义服务API接口与消息结构. 同一个服务接口自动创建多种不同的访问方式:

1. JSON over HTTP: 你可以使用restful方式访问服务api, 例如网页端.
2. JSON over Websocket: 你可以使用websocket方式访问服务api, 例如app端.
3. Protocol Buffers over gRPC: 你可以使用grpc方式访问服务api, 例如微服务间.

## 局限

proto3的数据类型不够丰富:

- 基本类型
  + 内置类型:
    * 布尔: bool
    * 整型: int32/64, sint32/64, uint32/64, sfixed32/64, fixed32/64
    * 浮点: float, double
    * 字串: string
    * 字节: bytes
  
  + 用户类型
    * message
    * enum [protojson]
    * onceof [protojson]

- 复合类型
  + 数组类型: repeated
  + 键值类型: map<key,val>
  
一般业务接口或数据存储而言, 上述数据类型(布尔,数值,字串, 及其数组,键值)已经满足! 但是如果你的业务强烈需要datetime等类型, 请仔细考虑! 


## 组件

- base

  基本工具, 配置管理等

- log

  日志组件

- http

  http服务器组件: grpc, rest, websocket.

- sqlx

  mysql, mariadb等sql db组件, 结合protogen自动生成orm相关代码, 节省书写sql的时间!

- mgox

  mongodb组件

- redis

  redis/pika组件

- kafka

  kafka组件

- etcd

  etcd组件

- cmd

  命令行组件

- protogen

  内部封装protoc, protoc-gen-go, protoc-gen-go-grpc等, 一键帮你生成所谓的源码!

- install

  所用mongodb, redis, etcd等安装脚本与方式!


你可以基于fasgo上述组件快速搭建自己的性能api服务器!

## 示例



