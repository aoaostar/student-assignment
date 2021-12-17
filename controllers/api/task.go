package api

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type Task struct {
}

func (i Task) All(c iris.Context) {
	task := new(models.Tasks)
	sess := sessions.Get(c)
	user := sess.Get("user").(*models.Users)
	tasks := []models.ApiTasks{}

	Db.Db.Model(task).Where("class_id=?", user.ClassId).Find(&tasks)
	type m struct {
		models.ApiTasks
		Job models.ApiJobs `json:"job"`
	}
	var data []m
	for _, apiTasks := range tasks {
		job := new(models.Jobs)
		job.GetByTaskId(int64(apiTasks.ID))
		apiJob := new(models.ApiJobs)
		copier.Copy(apiJob, job)
		m2 := m{
			ApiTasks: apiTasks,
			Job:  *apiJob,
		}
		data = append(data, m2)
	}
	util.JSON(c, "ok", "success", data)

}
func (i Task) Get(c iris.Context) {

}
func (i Task) Update(c iris.Context) {

}
func (i Task) Create(c iris.Context) {

}
func (i Task) Delete(c iris.Context) {

}
