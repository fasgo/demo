// Code generated by protoc-gen-go-http. DO NOT EDIT!
// Versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// Modified at 2020-12-21 16:14:19

package api

import (
	json "encoding/json"
	http "github.com/fasgo/http"
	grpc "google.golang.org/grpc"
	io "io"
)

func StudentServiceRegistry(impl interface{}) (*grpc.ServiceDesc, []*http.HttpDesc) {
	var svc = impl.(StudentServiceServer)
	var hds []*http.HttpDesc
	var hd *http.HttpDesc
	{
		hd = new(http.HttpDesc)
		hd.Package = `api`
		hd.Service = `StudentService`
		hd.Method = `Add`
		hd.PostPath = `/demo/students`
		hd.PostFunc = func(ctx *http.Context, in io.Reader) (interface{}, error) {
			req := new(Student)
			if ctx.Request.URL.RawQuery != `` {
				http.QueryUint64(ctx, `sno`, &req.Sno)
				http.QueryString(ctx, `name`, &req.Name)
				http.QueryUint32(ctx, `age`, &req.Age)
				http.QueryBool(ctx, `male`, &req.Male)
				http.QueryString(ctx, `desc`, &req.Desc)
			}
			err := json.NewDecoder(in).Decode(req)
			if err != nil && err != io.EOF {
				return nil, err
			}
			return svc.Add(ctx, req)
		}
		hds = append(hds, hd)
	}
	{
		hd = new(http.HttpDesc)
		hd.Package = `api`
		hd.Service = `StudentService`
		hd.Method = `Del`
		hd.DeletePath = `/demo/students/:sno`
		hd.DeleteFunc = func(ctx *http.Context, in io.Reader) (interface{}, error) {
			req := new(Student)
			http.ParamUint64(ctx, `sno`, &req.Sno)
			if ctx.Request.URL.RawQuery != `` {
				http.QueryUint64(ctx, `sno`, &req.Sno)
				http.QueryString(ctx, `name`, &req.Name)
				http.QueryUint32(ctx, `age`, &req.Age)
				http.QueryBool(ctx, `male`, &req.Male)
				http.QueryString(ctx, `desc`, &req.Desc)
			}
			err := json.NewDecoder(in).Decode(req)
			if err != nil && err != io.EOF {
				return nil, err
			}
			return svc.Del(ctx, req)
		}
		hds = append(hds, hd)
	}
	{
		hd = new(http.HttpDesc)
		hd.Package = `api`
		hd.Service = `StudentService`
		hd.Method = `Upd`
		hd.PutPath = `/demo/students/:sno`
		hd.PutFunc = func(ctx *http.Context, in io.Reader) (interface{}, error) {
			req := new(Student)
			http.ParamUint64(ctx, `sno`, &req.Sno)
			if ctx.Request.URL.RawQuery != `` {
				http.QueryUint64(ctx, `sno`, &req.Sno)
				http.QueryString(ctx, `name`, &req.Name)
				http.QueryUint32(ctx, `age`, &req.Age)
				http.QueryBool(ctx, `male`, &req.Male)
				http.QueryString(ctx, `desc`, &req.Desc)
			}
			err := json.NewDecoder(in).Decode(req)
			if err != nil && err != io.EOF {
				return nil, err
			}
			return svc.Upd(ctx, req)
		}
		hds = append(hds, hd)
	}
	{
		hd = new(http.HttpDesc)
		hd.Package = `api`
		hd.Service = `StudentService`
		hd.Method = `Get`
		hd.GetPath = `/demo/students/:sno`
		hd.GetFunc = func(ctx *http.Context, in io.Reader) (interface{}, error) {
			req := new(Student)
			http.ParamUint64(ctx, `sno`, &req.Sno)
			if ctx.Request.URL.RawQuery != `` {
				http.QueryUint64(ctx, `sno`, &req.Sno)
				http.QueryString(ctx, `name`, &req.Name)
				http.QueryUint32(ctx, `age`, &req.Age)
				http.QueryBool(ctx, `male`, &req.Male)
				http.QueryString(ctx, `desc`, &req.Desc)
			}
			err := json.NewDecoder(in).Decode(req)
			if err != nil && err != io.EOF {
				return nil, err
			}
			return svc.Get(ctx, req)
		}
		hd.WebsocketPath = `/demo/student/ws`
		hd.WebsocketFunc = func(ctx *http.Context, in io.Reader) (interface{}, error) {
			req := new(Student)
			if ctx.Request.URL.RawQuery != `` {
				http.QueryUint64(ctx, `sno`, &req.Sno)
				http.QueryString(ctx, `name`, &req.Name)
				http.QueryUint32(ctx, `age`, &req.Age)
				http.QueryBool(ctx, `male`, &req.Male)
				http.QueryString(ctx, `desc`, &req.Desc)
			}
			err := json.NewDecoder(in).Decode(req)
			if err != nil && err != io.EOF {
				return nil, err
			}
			return svc.Get(ctx, req)
		}
		hds = append(hds, hd)
	}
	{
		hd = new(http.HttpDesc)
		hd.Package = `api`
		hd.Service = `StudentService`
		hd.Method = `All`
		hd.GetPath = `/demo/students`
		hd.GetFunc = func(ctx *http.Context, in io.Reader) (interface{}, error) {
			req := new(AllReq)
			if ctx.Request.URL.RawQuery != `` {
				http.QueryInt32(ctx, `from`, &req.From)
				http.QueryInt32(ctx, `size`, &req.Size)
				http.QueryString(ctx, `search`, &req.Search)
				http.QueryString(ctx, `field`, &req.Field)
				http.QueryBool(ctx, `desc`, &req.Desc)
			}
			err := json.NewDecoder(in).Decode(req)
			if err != nil && err != io.EOF {
				return nil, err
			}
			return svc.All(ctx, req)
		}
		hds = append(hds, hd)
	}
	return &_StudentService_serviceDesc, hds
}

/*--------------------------SERVICES IMPLEMENT BEGIN--------------------------

import (
	"context"
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
----------------------------SERVICES IMPLEMENT END----------------------------*/
