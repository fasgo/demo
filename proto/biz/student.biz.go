package biz

import (
	"context"
	"fmt"
	"github.com/fasgo/demo/proto/api"
	"github.com/fasgo/base"
	"github.com/fasgo/base/kits"
)

type TagServiceService struct {
	*api.UnimplementedTagServiceServer // grpcv2后的新特性(不用太在意!)
}

var _ api.TagServiceServer = (*TagServiceService)(nil)

func (s *TagServiceService) All(ctx context.Context, req *api.AllReq) (rsp *api.AllRsp, err error) {
	fmt.Printf("from %v\n", req.From)

	if req.Search == "message" {
		return nil, base.Error(1001, "test message %v", 1234)
	}

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

func (s *TagServiceService) Get(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	fmt.Println("request: ", kits.ToJson(req))
	rsp = req
	rsp.Name = "teset -> " + rsp.Name
	return
}
