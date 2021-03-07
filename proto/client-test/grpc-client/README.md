# 参考资料
- grpc name resolver原理及实践:
  - https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487040&idx=1&sn=35e54214535da2f2203de2b7f09010d1&source=41#wechat_redirect

- grpc客户端负载均衡/重试/健康检查:
  - 185.199.111.153 github.io
  - http://yangxikun.github.io/golang/2019/10/19/golang-grpc-client-side-lb.html

- 使用dns做resolver以及MAX_CONNECTION_AGE处理dns ttl的问题:
  - https://rafaeleyng.github.io/grpc-load-balancing-with-grpc-go
  -

- 介绍grpc.WithDefaultServiceConfig()的参数ServiceConfig的message形式
  - https://github.com/grpc/grpc/blob/master/doc/service_config.md
  - https://github.com/grpc/proposal/blob/master/A2-service-configs-in-dns.md

- 介绍client与server连接机制
  - http://yangxikun.github.io/golang/2019/10/19/golang-grpc-client-side-lb.html

- 解析grpc.ClientConn源码
  - https://grpc.io/docs/
  - https://zhuanlan.zhihu.com/p/104060740

- grpc name resolution
  - https://github.com/grpc/grpc/blob/master/doc/naming.md

- 基于WithDefaultServiceConfig的一个示例
  - https://github.com/mbobakov/grpc-consul-resolver


# 关键概念
- Resolver
  - passthrough
  - dns
  - manual

- Balancer
  - pickerfirst
  - roundrobin
  - grpclb

- Picker
  - pickerfirst
  - roundrobin
  - grpclb


# 实现

## 目的

定制resolver实现:
1. etcd服务发现/注册(TBD)
2. addr多连接支持(N个),替代连接池

## 思路

支持2种scheme:
1. etcd:///endpoint#N, 其中N表示创建N个连接(默认1个)
2. pass:///ip1:port1[#N1],ip2:port2[#N2]..., 其中N1,N2表示创建连接数量

对于1的前缀必然是etcd
对于2的前缀可选是extd, pass, addr, 暂定pass, 相对于passthrough而言

问题是如何解析target...
1. scheme
2. authority
3. endpoint
   针对endpoint再做解析最后生成相应的结果Address

## 问题
1. waitForResolvedAddrs阻塞

   解决: 参考passthrough的源码并进行修改

2. 测试server端的连接是否有2条? 并且client是否真正roundrobin?

   解决: 在server添加creds连接拦截器, 打印每个连接的handshake信息

## 源码

- server
```
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	...
)

func main() {

	grpcAddr := ":9080"
	if len(os.Args) > 1 {
		grpcAddr = os.Args[1]
	}

	// 1. 创建server
	//通过code设置
	svr := http.NewServerWith(&http.Config{
		//HttpAddr:        ":8080",   // 开启http访问
		GrpcAddr:        grpcAddr,  // 开启grpc访问
		WbskCheckOrigin: http.DOWN, // websocket不启用origin检测
	})

	svr.GrpcServerOption(grpc.Creds(new(TransportCredentialsTest)))
	//通过conf设置
	/*svr := http.NewServer()
	 */

	// 2. 注册service. 绑定实现
	svr.RegisterService(api.TagServiceRegistry, new(biz.TagServiceService))

	// 3. 启动server. 提供服务
	if err := svr.ListenAndServe(); err != nil {
		base.DefaultLogger.Errorf("server error: %+v", err)
	}
}

type TransportCredentialsTest struct {
}

func (tc *TransportCredentialsTest) ClientHandshake(ctx context.Context, name string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	fmt.Println("ClientHandshake#########################")
	return nil, nil, nil
}
func (tc *TransportCredentialsTest) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	fmt.Println("ServerHandshake#########################")
	fmt.Printf("Remote Addr %v, %v\n", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
	ai := AuthInfoTest("test")
	return conn, &ai, nil
}
func (tc *TransportCredentialsTest) Info() credentials.ProtocolInfo {
	fmt.Println("Info#########################")
	return credentials.ProtocolInfo{}
}

func (tc *TransportCredentialsTest) Clone() credentials.TransportCredentials {
	return tc
}

func (tc *TransportCredentialsTest) OverrideServerName(string) error {
	fmt.Println("OverrideServerName#########################")
	return nil
}

type AuthInfoTest string

func (ai *AuthInfoTest) AuthType() string {
	return string(*ai)
}

```

- client
```
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
	...
)

func main() {
	// 默认是pickerfirst
	cc, err := grpc.Dial("pass:///:9080#2,:9090#1", grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig":[{ "round_robin":{}}]}`))

	if err != nil {
		panic(err)
	}
	defer cc.Close()

	cl := api.NewTagServiceClient(cc)
	for i := 0; ; i++ {
		rsp, err := cl.All(context.Background(), &api.AllReq{
			Search: "all",
			From:   int32(i),
			Size:   10,
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v: %v\n", i, kits.ToJson(rsp.Data))
		time.Sleep(500 * time.Millisecond)
	}

}

```
- resolver
```
package main

import (
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
	"strconv"
	"strings"
)

/*
支持2种scheme:
1. etcd:///endpoint#N
2. pass:///endpoint1#N1,endpoint2#N2....
*/

func init() {
	resolver.Register(new(passAnchorBuilder))
}

type AnchorAddress struct {
	Addr string // 服务地址
	Anch int    // 锚记数量
}

/*
格式: addr1#anch1,addr2#anch2...
*/
func ParseAnchorAddress(endpoint string) (rt []*AnchorAddress) {
	var (
		addr string
		anch int
	)
	for _, val := range strings.Split(endpoint, ",") {
		idx := strings.IndexByte(val, '#')
		if idx > 0 {
			addr = val[:idx]
			anch, _ = strconv.Atoi(val[idx+1:])
		} else {
			addr = val
		}
		if anch < 1 {
			anch = 1
		}
		rt = append(rt, &AnchorAddress{
			Addr: addr,
			Anch: anch,
		})
	}
	return
}

type passAnchorResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *passAnchorResolver) ResolveNow(resolver.ResolveNowOptions) {

}

func (r *passAnchorResolver) Close() {

}

func (r *passAnchorResolver) start() {
	var state resolver.State
	for _, item := range ParseAnchorAddress(r.target.Endpoint) {
		for i := 0; i < item.Anch; i++ {
			state.Addresses = append(state.Addresses, resolver.Address{
				Addr:       item.Addr,
				Attributes: attributes.New("idx", i),
			})
		}
	}
	r.cc.UpdateState(state)

	/*下述代码会在ClientConn.conns生成多个连接对象，但无法配合roundrobin做相关负载均衡*/
	//for _, item := range ParseAnchorAddress(r.target.Endpoint) {
	//	for i := 0; i < item.Anch; i++ {
	//		r.cc.UpdateState(resolver.State{Addresses: []resolver.Address{
	//			{
	//				Addr:       item.Addr,
	//				Attributes: attributes.New("idx", i),
	//			},
	//		}})
	//	}
	//}
}

type passAnchorBuilder struct {
}

func (b *passAnchorBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	r := &passAnchorResolver{
		target: target,
		cc:     cc,
	}
	r.start()

	return r, nil
}

func (b *passAnchorBuilder) Scheme() string {
	return "pass"
}

```

## 总结
    1. balancer默认是pickerfirst,不是roundrobin
    2. resolver.start()逻辑不能放在ResolveNow(),具体参考passthrough
    3. ClientConn.UpdateState()多次调用会在ClientConn.conns生成多个连接对象,但无法与roundrobin共用
    4. ClientConn.UpdateState()的State的Address必须指定不同的attribute对象,否则会覆盖去重!
    5. client-server端效果达到预期,自动容错,负载均衡(根据#比例)