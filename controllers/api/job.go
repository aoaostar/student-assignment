package api

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"fmt"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"io"
	"os"
	"path"
	"strconv"
)

type Job struct {
}
type updateParams struct {
	Id      int64   `json:"id"`
	Content *string `json:"content"`
}

func (i Job) Get(c iris.Context) {

	sess := sessions.Get(c)
	user := sess.Get("user").(*models.Users)
	taskId, _ := strconv.ParseInt(c.URLParam("task_id"), 10, 64)
	job := new(models.Jobs)
	job.TaskId = taskId
	job.UserId = int64(user.ID)
	v1, _ := uuid.NewV1()
	job.Filepath = v1.String()
	apiJobs := new(models.ApiJobs)

	Db.Db.Model(job).Where("task_id=?",taskId).Preload("Task").FirstOrCreate(job)

	copier.Copy(&apiJobs, &job)
	util.JSON(c, "ok", "success", apiJobs)

}
func (i Job) Update(c iris.Context) {
	sess := sessions.Get(c)
	user := sess.Get("user").(*models.Users)

	params := new(updateParams)

	err := c.ReadForm(params)
	if err != nil && !iris.IsErrPath(err) {
		util.JSON(c, "error", err.Error())
		return
	}
	job := new(models.Jobs)
	err = job.Get(params.Id)
	if job.UserId != int64(user.ID) || err != nil {
		util.JSON(c, "error", "该job不存在")
		return
	}
	file, header, err := c.FormFile("file")

	var filename string
	if err == nil {
		filename = header.Filename
		filepath := fmt.Sprintf("./storage/uploads/%v/%s", job.UserId, job.Filepath)
		dir := path.Dir(filepath)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
		openFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			util.JSON(c, "error", err.Error())
			return
		}
		_, err = io.Copy(openFile, file)
		if err != nil {
			util.JSON(c, "error", err.Error())
			return
		}

	}
	err = job.Update(filename, *params.Content)
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	util.JSON(c, "ok", "保存成功")
	return

}
func (i Job) Create(c iris.Context) {

}
func (i Job) Delete(c iris.Context) {

}
