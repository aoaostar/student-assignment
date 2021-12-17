package controllers

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"os"
)

type Index struct {
}

func (m *Index) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Index")
	b.Handle("GET", "/me", "Me")
	b.Handle("GET", "/download/{userId}/{filepath}", "Download")
}

func (m Index) Index(c iris.Context) {
	c.ViewData("Title", "首页")
	c.View("index.html")
}
func (m Index) Me(c iris.Context) {
	c.ViewData("Title", "我的资料")
	c.View("me.html")
}

func (m Index) Download(c iris.Context) {
	params := c.Params()
	userId := params.Get("userId")
	filepath := params.Get("filepath")
	job := new(models.Jobs)

	Db.Db.Model(job).Where("user_id=? and filepath=?", userId, filepath).First(&job)

	if job.ID == 0 {
		util.JSON(c, "error", "该文件不存在")
		return
	}

	filename := fmt.Sprintf("./storage/uploads/%v/%v", userId, filepath)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		util.JSON(c, "error", "该文件不存在")
		return
	}

	c.SendFile(filename, job.Filename)
}
