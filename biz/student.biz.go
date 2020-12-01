package biz
import (
	"context"
	"github.com/fasgo/demo/api"
)
type StudentServiceService struct {
	*api.UnimplementedStudentServiceServer
}
var _ api.StudentServiceServer = (*StudentServiceService)(nil)
func (s *StudentServiceService) Add(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return
}
func (s *StudentServiceService) Del(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return
}
func (s *StudentServiceService) Upd(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return
}
func (s *StudentServiceService) Get(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return
}
func (s *StudentServiceService) All(ctx context.Context, req *api.AllReq) (rsp *api.AllRsp, err error) {
	return
}
