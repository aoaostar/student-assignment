package masterApi

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"path"
)

type Job struct {
}

type updateJobParams struct {
	Id      int64   `json:"id"`
	Content *string `json:"content"`
}


func (i Job) All(c iris.Context) {

	job := new(models.Jobs)
	jobs := new([]models.Jobs)
	masterJobs := new([]models.MasterJobs)

	Db.Db.Model(job).Preload("Task").Preload("User").Find(jobs)

	copier.Copy(masterJobs, jobs)
	util.JSON(c, "ok", "success", masterJobs)

}

func (i Job) Get(c iris.Context) {

	job := new(models.Jobs)
	masterJob := new(models.MasterJobs)

	Db.Db.Model(job).Where("id=?", c.URLParam("id")).Preload("Task").Preload("User").First(job)
	copier.Copy(masterJob, job)
	util.JSON(c, "ok", "success", masterJob)

}
func (i Job) Update(c iris.Context) {

	params := new(updateJobParams)

	err := c.ReadForm(params)
	if err != nil && !iris.IsErrPath(err) {
		util.JSON(c, "error", err.Error())
		return
	}
	job := new(models.Jobs)
	err = job.Get(params.Id)
	if job.ID == 0 || err != nil {
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
