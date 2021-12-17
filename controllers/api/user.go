package api

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/util"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type User struct {

}


func (i User) Get(c iris.Context)  {

	sess := sessions.Get(c)
	user := sess.Get("user").(*models.Users)
	apiUser := new(models.ApiUsers)
	copier.Copy(&apiUser,&user)
	util.JSON(c,"ok","success",apiUser)

}
func (i User) Update(c iris.Context)  {

}
func (i User) Create(c iris.Context)  {

}
func (i User) Delete(c iris.Context)  {

}