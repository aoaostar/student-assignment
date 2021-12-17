package models

import (
	"StudentAssignment/models/types"
	"StudentAssignment/pkg/Db"
	"gorm.io/gorm"
	"time"
)

type Tasks struct {
	gorm.Model
	Name     string    `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	ClassId  int64     `gorm:"column:class_id;type:bigint(20);NOT NULL" json:"class_id"`
	Content  string    `gorm:"column:content;type:varchar(255);NOT NULL" json:"content"`
	Deadline time.Time `gorm:"column:deadline;type:datetime(3);NOT NULL" json:"deadline"`
}

type ApiTasks struct {
	ID        uint            `json:"id"`
	CreatedAt types.LocalTime `json:"created_at"`
	UpdatedAt types.LocalTime `json:"updated_at"`
	Name      string          `json:"name"`
	ClassId   int64           `json:"class_id"`
	Content   string          `json:"content"`
	Deadline  types.LocalTime `json:"deadline"`
}

func (t *Tasks) Get(id int64) error {
	return Db.Db.Where("id=?", id).First(&t).Error
}

func (t *Tasks) Create() error {

	return Db.Db.Create(t).Error
}

func (t *Tasks) Delete() error {
	tx := Db.Db.Begin()
	jobs := new([]Jobs)
	err := tx.Where("task_id=?", t.ID).Delete(jobs).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Delete(t).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (t *Tasks) Save() error {
	return Db.Db.Save(t).Error

}
