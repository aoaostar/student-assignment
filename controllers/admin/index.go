package admin



import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Index struct {
}

func (m *Index) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Index")
	b.Handle("GET", "/tasks", "Tasks")
	b.Handle("GET", "/task/edit/{taskId}", "TaskEdit")
	b.Handle("GET", "/task/create", "TaskCreate")
	b.Handle("GET", "/jobs", "Jobs")
	b.Handle("GET", "/job/edit/{jobId}", "JobEdit")
	b.Handle("GET", "/users", "Users")
}

func (m Index) Index(c iris.Context) {
	c.ViewData("Title", "控制台")
	c.View("admin/index.html")
}

func (m Index) Tasks(c iris.Context) {
	c.ViewData("Title", "所有任务")
	c.View("admin/task/index.html")
}
func (m Index) TaskEdit(c iris.Context) {
	c.ViewData("Title", "修改任务")
	c.ViewData("taskId", c.Params().Get("taskId"))
	c.View("admin/task/edit.html")
}
func (m Index) TaskCreate(c iris.Context) {
	c.ViewData("Title", "发布任务")
	c.View("admin/task/create.html")
}

func (m Index) Jobs(c iris.Context) {
	c.ViewData("Title", "所有作业")
	c.View("admin/job/index.html")
}
func (m Index) JobEdit(c iris.Context) {
	c.ViewData("Title", "修改作业")
	c.ViewData("jobId", c.Params().Get("jobId"))
	c.View("admin/job/edit.html")
}

func (m Index) Users(c iris.Context) {
	c.ViewData("Title", "学生")
	c.View("admin/user.html")
}
