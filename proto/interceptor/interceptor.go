package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"time"
)

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

func RequestInteceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("RequestInteceptor#################################")
	start := time.Now().UnixNano()
	rsp, err := handler(ctx, req)
	end := time.Now().UnixNano()
	fmt.Println("access method: %v, used(ns): %v", info.FullMethod, end-start)
	return rsp, err
}
