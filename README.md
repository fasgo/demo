# Go使用protocolbuffer快速开发api服务DEMO

- 本教程就是以"protogen + protoapi"为基础介绍如何快速构建api服务!
- step-by-step教程:[Go使用protocolbuffer快速构建api服务教程](https://www.cnblogs.com/zolo/p/14066444.html)

# 搭建环境

## 第1步: 下载安装Go 1.15

[GO下载页面: https://golang.google.cn/dl/](https://golang.google.cn/dl/)

详情不再赘述! Go都不会安装, 哪还能说啥?

## 第2步: 下载安装protogen

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

## 第3步: 下载项目源码

```
git clone https://github.com/fasgo/demo.git
```

step-by-step教程:[Go使用protocolbuffer快速构建api服务教程](https://www.cnblogs.com/zolo/p/14066444.html)