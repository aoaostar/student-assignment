package masterApi

import (
	"StudentAssignment/models"
	"StudentAssignment/models/types"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"github.com/kataras/iris/v12"
	"time"
)

type Task struct {
}
type (
	createTaskParams struct {
		Name     string          `json:"name" validate:"required"`
		ClassId  int64           `json:"class_id" validate:"required,numeric"`
		Content  string          `json:"content" validate:"required"`
		Deadline types.LocalTime `json:"deadline" validate:"required,datetime=2006-01-02 15:04:05"`
	}
	updateTaskParams struct {
		Id int64 `json:"id"`
		createTaskParams
	}
	deleteTaskParams struct {
		Id int64 `json:"id"`
	}
)

func (i Task) All(c iris.Context) {
	task := new(models.Tasks)
	tasks := new([]models.ApiTasks)
	Db.Db.Model(task).Find(&tasks)

	util.JSON(c, "ok", "success", tasks)

}
func (i Task) Get(c iris.Context) {
	task := new(models.Tasks)
	apiTasks := new(models.ApiTasks)
	Db.Db.Model(task).Where("id=?", c.URLParam("id")).First(&apiTasks)

	util.JSON(c, "ok", "success", apiTasks)
}
func (i Task) Update(c iris.Context) {

	params := new(updateTaskParams)

	err := c.ReadJSON(params)
	if err != nil && !iris.IsErrPath(err) {
		util.JSON(c, "error", err.Error())
		return
	}
	task := new(models.Tasks)
	err = task.Get(params.Id)
	if task.ID == 0 || err != nil {
		util.JSON(c, "error", "该任务不存在")
		return
	}
	task.ClassId = params.ClassId
	task.Name = params.Name
	task.Content = params.Content
	deadline, _ := params.Deadline.Value()
	task.Deadline = deadline.(time.Time)
	err = task.Save()
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	util.JSON(c, "ok", "保存成功")
}
func (i Task) Create(c iris.Context) {

	params := new(createTaskParams)

	err := c.ReadJSON(params)
	if err != nil && !iris.IsErrPath(err) {
		util.JSON(c, "error", err.Error())
		return
	}
	task := new(models.Tasks)
	task.ClassId = params.ClassId
	task.Name = params.Name
	task.Content = params.Content
	deadline, _ := params.Deadline.Value()

	task.Deadline = deadline.(time.Time)
	err = task.Create()
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}

	util.JSON(c, "ok", "发布成功")
}
func (i Task) Delete(c iris.Context) {
	params := new(deleteTaskParams)

	err := c.ReadForm(params)
	if err != nil && !iris.IsErrPath(err) {
		util.JSON(c, "error", err.Error())
		return
	}
	task := new(models.Tasks)
	err = task.Get(params.Id)
	if task.ID == 0 || err != nil {
		util.JSON(c, "ok", "删除成功")
		return
	}
	err = task.Delete()
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	util.JSON(c, "ok", "删除成功")
}
