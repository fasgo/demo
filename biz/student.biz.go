package biz

import (
	"context"
	"fmt"
	"github.com/fasgo/base"
	"github.com/fasgo/demo/api"
)

type StudentServiceService struct {
	*api.UnimplementedStudentServiceServer
}

var _ api.StudentServiceServer = (*StudentServiceService)(nil)

func (s *StudentServiceService) Add(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	fmt.Println(base.Json(req))
	rsp = req
	return
}
func (s *StudentServiceService) Del(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	fmt.Println(base.Json(req))
	rsp = req
	return
}
func (s *StudentServiceService) Upd(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	fmt.Println(base.Json(req))
	rsp = req
	return
}
func (s *StudentServiceService) Get(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	fmt.Println(base.Json(req))
	rsp = req
	return
}
func (s *StudentServiceService) All(ctx context.Context, req *api.AllReq) (rsp *api.AllRsp, err error) {
	fmt.Println(base.Json(req))
	rsp = new(api.AllRsp)
	return
}
