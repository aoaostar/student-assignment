package masterApi

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/util"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type Admin struct {

}


func (i Admin) Get(c iris.Context)  {

	sess := sessions.Get(c)
	admin := sess.Get("admin").(*models.Admins)
	apiAdmin := new(models.ApiAdmins)
	copier.Copy(&apiAdmin,&admin)
	util.JSON(c,"ok","success",apiAdmin)

}
func (i Admin) Update(c iris.Context)  {

}
func (i Admin) Create(c iris.Context)  {

}
func (i Admin) Delete(c iris.Context)  {

}