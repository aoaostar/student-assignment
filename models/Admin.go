package models

import (
	"StudentAssignment/models/types"
	"StudentAssignment/pkg/Db"
	"gorm.io/gorm"
)

type Admins struct {
	gorm.Model
	Username string  `gorm:"column:username;type:varchar(255);uniqueIndex;NOT NULL" json:"username"`
	Password string  `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
}
type ApiAdmins struct {
	ID        uint            `json:"id"`
	CreatedAt types.LocalTime `json:"created_at"`
	UpdatedAt types.LocalTime `json:"updated_at"`
	Username  string          `json:"username"`
}

func (u *Admins) Create() error {

	return Db.Db.Create(u).Error
}
