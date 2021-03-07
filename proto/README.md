# package demo

## 安装protogen

### 依赖:

1. go 1.15+
2. git 2.x+

### 安装

1. 下载二进制
   - windows: bin/windows_amd64/protogen.exe
   - linux:   bin/linux_amd64/protogen
   - darwin:  bin/darwin_amd64/protogen

   其他OS请从源码编译!

2. 添加至本地path

   注意: 第一次执行会在protogen位置创建".protogen"目录缓存protoc, protoc-gen-go, protoc-gen-go-grpc, protoc-gen-go-protoapi等插件, 所以用户需要拥有protogen所在目录的可写权限.

### 从源码编译:

执行go get github.com/fasgo/protogen@latest

windows示例

```
set GO111MODULE = on
go get github.com/fasgo/protogen@latest
%GOPATH%\bin\protogen --help
```

linux/macos示例

```
export GO111MODULE=on
go get github.com/fasgo/protogen@latest
$GOPATH/bin/protogen --help
```

## 使用protogen:

使用protogen编译指定proto目录或proto文件.

需要注意: 目录会自动递归子目录!

```
Usage of protogen [options] {files...} :
  -bson
        对message的field添加bson:"xxx"标签
  -clean
        清除本地插件缓存.protogen
  -debug
        打印调试信息
  -goprivate string
        GO私有仓库,默认$GOPRIVATE
  -goproxy string
        GO代理服务,默认$GOPROXY或https://goproxy.cn (default "https: //goproxy.cn")
  -grpcv1
        grpc接口兼容v1(require_unimplemented_servers=false)
  -h    打印帮助信息
  -help
        打印帮助信息
  -proto_path string
        PB查找目录,多值用逗号分隔
  -unomitempty
        对message的field剔除json:"xxx"标签的omitempty选项
  -version
        打印版本信息
```

## 开发project

参考: [https://github.com/fasgo/demo](https://github.com/fasgo/demo)

假设项目结构:

```
$WORKSPACE
  |__api/: 访问访问层(使用protogen维护)
  |__biz/: 业务逻辑层
  |__dao/: 数据持久层
  |__main.go: 主程序
  |__conf.toml: 配置文件
```

1. 定义proto

在api定义student.proto: 数据结构, 服务接口(grpc, restful, websocket)

```
syntax = "proto3";

package api;

import "github.com/fasgo/protoapi/http.proto";

// 定义班级信息
message Clazz {
  int64 cno = 1; // 班级编号
  string name = 2; // 班级名称
  string desc = 3; // 班级描述
}

// 定义学生信息
// +TABLE -name=student -engine=InnoDB -comment=学生测试表
message Student {
  // +COLUMN -name=s_no -comment="学生编号"
  string sno = 1 ; // 学生编号
  string name = 2; // 学生名称
  uint32 gender = 3; // 性别
  int32  grade = 4; // 排名
  repeated float  score = 5; // 多门评分
  Clazz  clazz = 6; // 班级信息
}

// 定义all请求
message AllReq {
  int32 from = 1; // 分页开始下标(从0开始)
  int32 size = 2; // 分页大小
  string search = 3; // 模糊查询的输入内容
  string field = 4; // 查询结果排序字段
  bool desc = 5; // 查询结果排序是否DESC
}

// 定义all响应
message AllRsp{
  int32  total = 1; //查询结果总数
  repeated Student data = 2; // 查询结果数据
}

// 定义Tag服务(请使用Go命名规范), 可以通过@Path指定URI, 默认/<package>/<service>/<method>
service TagService {

  rpc All(AllReq) returns (AllRsp){
    option (protoapi.http) = {method: GET, path:"/demo/students"}; // 生成restful接口"get /demo/students"
    option (protoapi.http) = {method: WEBSOCKET, path:"/demo/students/ws"}; // 生成websocket接口"ws /demo/students/ws"
  }
}

```

上述除grpc接口外, 使用"http.get", "http.websocket"分别定义了restful与websocket接口.

2. 编译proto

在项目目录下执行"protogen api"编译生成相应的grpc及http源码.

```bigquery
$WORKSPACE/
  |__api/
     |__demo.pb.go
     |__demo.proto
     |__demo_grpc.pb.go
     |__demo_protoapi.pb.go
```

其中在xxx_protoapi.pb.go最后的注释会有service接口实现参考, 可以直接复制节省代码时间

```bigquery
/*--------------------------SERVICES IMPLEMENT BEGIN--------------------------

import (
	"context"
)
type TagServiceService struct {
	*api.UnimplementedTagServiceServer
}
var _ api.TagServiceServer = (*TagServiceService)(nil)
func (s *TagServiceService) All(ctx context.Context, req *api.AllReq) (rsp *api.AllRsp, err error) {
	return
}
----------------------------SERVICES IMPLEMENT END----------------------------*/
```

初次执行需要下载protoc, protoc-gen-go-grpc, protoc-gen-go-http等依赖工具.

```
protogen
2021-02-04 14:56:31 [I] - fetch https://maven.aliyun.com/repository/central/com/google/protobuf/protoc/3.14.0/protoc-3.14.0-windows-x86_64.exe
2021-02-04 14:56:37 [I] - goget google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
2021-02-04 14:56:54 [I] - goget google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.1
2021-02-04 14:57:07 [I] - goget github.com/fasgo/http/cmd/protoc-gen-go-http@v0.0.4
2021-02-04 14:57:27 [I] - goget github.com/fasgo/include@v0.0.1
```

3. 实现service

实现service接口

```
$WORKSPACE/
  |__biz/
     |__student.biz.go
```

由第2步知protogen会在"xxx_http.pb.go"最后注释部分生成service的参考实现, 业务开发通过"复制粘贴"节省代码时间.

```
type TagServiceService struct {
	*api.UnimplementedTagServiceServer // grpcv2后的新特性(不用太在意!)
}

var _ api.TagServiceServer = (*TagServiceService)(nil)

func (s *TagServiceService) All(ctx context.Context, req *api.AllReq) (rsp *api.AllRsp, err error) {
	fmt.Printf("from %v\n", req.From)
	rsp = new(api.AllRsp)
	rsp.Total = 10
	for i := 0; i < 10; i++ {
		rsp.Data = append(rsp.Data, &api.Student{
			Sno:    kits.ToString(i),
			Name:   "学生" + kits.ToString(i),
			Gender: 1,
			Grade:  int32(i),
			Score: []float32{
				88.8,
				99.9,
				77.7,
			},
			Clazz: &api.Clazz{
				Cno:  int64(i) % 3,
				Name: "班级" + kits.ToString(int64(i)%3),
				Desc: "这是一个测试班级",
			},
		})
	}
	return
}


```

4. 注册service

注册service实现

```
$WORKSPACE/
  |__main.go
```

protogen为每个proto的service生成"XXXRegistry", 业务开发通过RegisterService()将其与具体的service实现绑定即可

```
func init() {
	registry.Init(&registry.Config{
		Endpoint: "10.13.144.164:2379",
		Username: "root",
		Password: "123456",
		Debug:    true,
	})
}

func main() {

	httpAddr := ":8080"
	grpcAddr := ":9080"
	if len(os.Args) > 1 {
		grpcAddr = os.Args[1]
	}
	if len(os.Args) > 2 {
		httpAddr = os.Args[2]
	}

	// 1. 创建server. 方式1: 通过API设置参数. 方式2: 通过conf.toml设置参数
	svr := protoapi.NewServerWith(&protoapi.Config{
		Name:            "demo",        // 注册服务名称
		HttpAddr:        httpAddr,      // 开启http访问
		GrpcAddr:        grpcAddr,      // 开启grpc访问
		WbskCheckOrigin: protoapi.DOWN, // websocket不启用origin检测
	})
	/*
		svr := protoapi.NewServer() // 使用conf.toml的[protoapi]配置
	*/

	// 2. 绑定实现
	svr.RegisterService(api.TagServiceRegistry, new(biz.TagServiceService))

	// 3. 启动监听
	if err := svr.ListenAndServe(); err != nil {
		base.DefaultLogger.Errorf("server error: %+v", err)
	}
}

```

5. 测试service

启动main函数,分别用grpc, restful, websocket客户端访问即可