package models

import (
	"StudentAssignment/models/types"
	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	//ClassId string  `gorm:"column:class_id;type:varchar(255);uniqueIndex;NOT NULL" json:"class_id"`
	Users []Users `gorm:"foreignKey:class_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ApiClasses struct {
	ID        uint            `json:"id"`
	CreatedAt types.LocalTime `json:"created_at"`
	UpdatedAt types.LocalTime `json:"updated_at"`
	Name      string          `json:"name"`
}
