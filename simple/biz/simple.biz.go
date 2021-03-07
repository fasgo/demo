package biz

import (
	"fmt"
	"context"
	"github.com/fasgo/demo/simple/mdl"
	"github.com/fasgo/base/kits"
)

func Create(c context.Context, s *mdl.Student) (r *mdl.Student, err error) {
	fmt.Println("create a student:", kits.ToJson(s))
	r = s
	return
}

func Update(c context.Context, s *mdl.Student) (r *mdl.Student, err error) {
	fmt.Println("Update a student:", kits.ToJson(s))
	r = s
	return
}

func Delete(c context.Context, s *mdl.Student) (r *mdl.Student, err error) {
	fmt.Println("Delete a student:", kits.ToJson(s))
	r = s
	return
}

func Retrieve(c context.Context, s *mdl.Student) (r *mdl.Student, err error) {
	fmt.Println("Retrieve a student:", kits.ToJson(s))
	r = s
	return
}
