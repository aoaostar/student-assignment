package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Job struct {

}
func (m *Job) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Index")
	b.Handle("GET", "/edit/{taskId}", "Edit")
}

func (m Job) Index(c iris.Context)  {
	c.ViewData("Title", "作业")
	c.View("job/index.html")
}

func (m Job) Edit(c iris.Context)  {

	c.ViewData("Title", "修改作业")
	c.ViewData("taskId",c.Params().Get("taskId"))

	c.View("job/edit.html")
}
