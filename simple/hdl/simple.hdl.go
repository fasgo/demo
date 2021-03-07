package hdl

import (
	"fmt"
	"github.com/fasgo/demo/simple/biz"
	"github.com/fasgo/demo/simple/mdl"
	"github.com/fasgo/protoapi"
)

// Post /simple/students
func Create(c *protoapi.Context) {
	var s *mdl.Student
	err := c.Scheme(&s, "json")
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	rsp, err := biz.Create(c, s)
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	c.WriteApplyResult(rsp)
}

// Put /simple/students/:name
func Update(c *protoapi.Context) {
	var s *mdl.Student
	err := c.Scheme(&s, "json")
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	rsp, err := biz.Update(c, s)
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	c.WriteApplyResult(rsp)
}

// Delete /simple/students/:name
func Delete(c *protoapi.Context) {
	var s = new(mdl.Student)
	err := c.Scheme(&s, "json")
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	rsp, err := biz.Delete(c, s)
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	c.WriteApplyResult(rsp)
}

// Get /simple/students/:name
func Retrieve(c *protoapi.Context) {
	var s = new(mdl.Student)
	err := c.Scheme(&s, "json")
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	rsp, err := biz.Retrieve(c, s)
	if err != nil {
		c.WriteErrorResult(err)
		return
	}

	c.WriteApplyResult(rsp)
}

func StudentsInterceptor(c *protoapi.Context) {
	fmt.Println("students interceptor")
}
