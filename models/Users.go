package models

import (
	"StudentAssignment/models/types"
	"StudentAssignment/pkg/Db"
	"fmt"
	"gorm.io/gorm"
	"os"
)

type Users struct {
	gorm.Model
	ClassId int64 `gorm:"column:class_id;type:bigint(20) unsigned;NOT NULL" json:"class_id"`
	//UserId   string  `gorm:"column:user_id;type:varchar(255);uniqueIndex;NOT NULL" json:"user_id"`
	Username string  `gorm:"column:username;type:varchar(255);NOT NULL" json:"username"`
	Password string  `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	Jobs     []Jobs  `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class    Classes `gorm:"foreignKey:class_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type ApiUsers struct {
	ID        uint            `json:"id"`
	CreatedAt types.LocalTime `json:"created_at"`
	UpdatedAt types.LocalTime `json:"updated_at"`
	ClassId   int64           `json:"class_id"`
	Username  string          `json:"username"`
	Class     ApiClasses      `json:"class"`
}

func (u *Users) Create() error {

	return Db.Db.Create(u).Error
}

func (u *Users) Get(id int64) error {

	return Db.Db.Where("id=?", id).First(&u).Error
}

func (u *Users) Delete() error {
	tx := Db.Db.Begin()
	jobs := new([]Jobs)
	err := tx.Where("user_id=?", u.ID).Delete(jobs).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Delete(u).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	_ = os.RemoveAll(fmt.Sprintf("./storage/uploads/%v", u.ID))

	tx.Commit()
	return nil

}

func (u *Users) Save() error {
	return Db.Db.Save(u).Error

}
