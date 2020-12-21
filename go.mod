module github.com/fasgo/demo

go 1.15

require (
	github.com/fasgo/base v0.2.0-a4
	github.com/fasgo/http v0.0.10
	github.com/golang/protobuf v1.4.1
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/fasgo/http v0.0.9 => ../http
)