package api

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"github.com/kataras/iris/v12"
)

type Class struct {

}

func (i Class) All(c iris.Context)  {

	class := new([]models.Classes)
	apiClass := new([]models.ApiClasses)
	Db.Db.Model(class).Find(apiClass)

	util.JSON(c,"ok","success", apiClass)

}

func (i Class) Get(c iris.Context)  {

	class := new(models.Classes)
	util.JSON(c,"ok","success", class)

}
func (i Class) Update(c iris.Context)  {

}
func (i Class) Create(c iris.Context)  {

}
func (i Class) Delete(c iris.Context)  {

}