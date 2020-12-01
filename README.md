如果无法查看图片, 请跳转博客正文: [https://www.cnblogs.com/zolo/p/14066444.html](https://www.cnblogs.com/zolo/p/14066444.html)

# 前述

本教程只说"实践",不谈"理论". 如果哪位觉得有趣, 可以加本人QQ(1255422783)详细交流!

# 问题

对于后端开发, 经常"众口难调". 一套业务逻辑却要三套不同实现API!

- 网页端要"http(json) api"(如restful api)
- 移动端要"websocket api"
- 服务端要"grpc api"

# 正题

本教程主要介绍如何使用"protogen + protoapi"的开发步骤. 
源码仓库:[https://github.com/fasgo/demo](https://github.com/fasgo/demo)

## 开发步骤

### 第1步: Goland启用"GoModule"支持
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_201201042153enable-go-module.jpg)


### 第2步: 根据需求定义proto
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_2012010422445-define-proto.jpg)

### 第3步: 进入项目执行protogen
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_2012010422152-execute-protogen.jpg)

****注意****: 如果是首次执行需要下载PB所需的插件, 过程为稍慢, 请耐心等待!
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_2012010422314-first-protogen.jpg)

### 第4步: 复制service模板
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_2012010422506-copy-service-template.jpg)

****注意****: 在xxxx_http.pb.go(例如stduent_http.pb.go)最后会生成service的实现模板, 直接将其复制到自己的biz代码即可!

### 第5步: 实现service逻辑
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_2012010422567-implement-service.jpg)

### 第6步: 注册service实例, 并启动服务
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_2012010423038-register-service.jpg)

### 第7步: 客户端连接测试
- grpc client
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_2012010528099-grpc-client.jpg)

- http client
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_20120105281510-http-client.jpg)

- websocket client
![image](https://images.cnblogs.com/cnblogs_com/zolo/907331/o_20120105282111-websoket-client.jpg)

## 搭建环境

### 第1步: 下载安装Go 1.15

[GO下载页面: https://golang.google.cn/dl/](https://golang.google.cn/dl/)

详情不再赘述! Go都不会安装, 哪还能说啥?

### 第2步: 下载安装protogen

- protogen介绍
    
    protogen是作者开发的一个快速构建工具,封装了protoc等一系列pb常用工具的下载,安装与使用! 一键完成proto的所有烦琐配置!   
    
- protogen安装: 
    
    "强烈建议添加protogen到系统环境PATH!"
    ----------------------------
    第3步教程假定你已经将protogen添加到PATH.
    
    protogen采用go get安装方式, 整体流程分4步:
    1. 指定GoPath 
    2. 启用GoModule 
    3. 设置GoProxy 
    4. GoGet下载安装
    
    具体脚本操作: 
    
    - windows安装: (以临时目录为例)
    ```
    set GOPATH=c:\Windows\Temp
    set GO111MODULE=on
    set GOPROXY=https://goproxy.cn
    go get github.com/fasgo/protogen@latest
   ```
  
   安装位置: c:\Windows\Temp\bin\protogen
   ```
    c:\Windows\Temp\bin\protogen.exe --help
    Usage of protogen [options] {files...} :
      -clean
            清除缓存.protogen
      -compatible
            gRPC接口兼容v1.x
      -debug
            开启调试
      -goproxy string
            GO代理服务,默认$GOPROXY(https://goproxy.cn) (default "https://goproxy.cn")
      -h    打印帮助
      -help
            打印帮助
      -proto_path string
            PB查找路径,多值逗号分隔
      -version
            打印版本
    ```
  
    - linux(darwin)安装: (以临时目录为例)
    ```
    export GOPATH=/tmp
    export GO111MODULE=on
    export GOPROXY=https://goproxy.cn
    go get github.com/fasgo/protogen@latest
    ```
    安装位置: /tmp/bin/protogen
    ```
    /tmp/bin/protogen --help
    Usage of protogen [options] {files...} :
      -clean
            清除缓存.protogen
      -compatible
            gRPC接口兼容v1.x
      -debug
            开启调试
      -goproxy string
            GO代理服务,默认$GOPROXY(https://goproxy.cn) (default "https://goproxy.cn")
      -h    打印帮助
      -help
            打印帮助
      -proto_path string
            PB查找路径,多值逗号分隔
      -version
            打印版本
    ```

### 第3步: 下载项目源码

```
git clone https://github.com/fasgo/demo.git e:\temp\
```

step-by-step教程:[Go使用protocolbuffer快速构建api服务教程](https://www.cnblogs.com/zolo/p/14066444.html)
