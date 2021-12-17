package models

import (
	"StudentAssignment/models/types"
	"StudentAssignment/pkg/Db"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type Jobs struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL" json:"user_id"`
	TaskId   int64  `gorm:"column:task_id;type:bigint(20) unsigned;NOT NULL" json:"task_id"`
	Content  string `gorm:"column:content;type:varchar(255);" json:"content"`
	Filepath string `gorm:"column:filepath;type:varchar(255);" json:"filepath"`
	Filename string `gorm:"column:filename;type:varchar(255);" json:"filename"`
	User     Users  `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Task     Tasks  `gorm:"foreignKey:task_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type ApiJobs struct {
	ID        uint            `json:"id"`
	CreatedAt types.LocalTime `json:"created_at"`
	UpdatedAt types.LocalTime `json:"updated_at"`
	UserId    int64           `json:"user_id"`
	TaskId    int64           ` json:"task_id"`
	Content   string          `json:"content"`
	Filepath  string          ` json:"filepath"`
	Filename  string          ` json:"filename"`
	Task      ApiTasks        ` json:"task"`
}
type MasterJobs struct {
	ID        uint            `json:"id"`
	CreatedAt types.LocalTime `json:"created_at"`
	UpdatedAt types.LocalTime `json:"updated_at"`
	UserId    int64           `json:"user_id"`
	TaskId    int64           ` json:"task_id"`
	Content   string          `json:"content"`
	Filepath  string          ` json:"filepath"`
	Filename  string          ` json:"filename"`
	Task      ApiTasks        ` json:"task"`
	User      ApiUsers        ` json:"user"`
}

func (j *Jobs) Update(filename string, content string) error {

	params := iris.Map{}

	if filename != "" {
		params["filename"] = filename
	}
	if content != "" {
		params["content"] = content
	}
	params["content"] = content
	return Db.Db.Model(j).Where("id=?", j.ID).Updates(params).Error
}

func (j *Jobs) Get(id int64) error {
	return Db.Db.Where("id=?", id).First(j).Error
}

func (j *Jobs) GetByTaskId(taskId int64) error {
	return Db.Db.Where("task_id=?", taskId).First(j).Error
}
