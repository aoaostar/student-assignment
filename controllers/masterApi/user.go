package masterApi

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
)

type User struct {
}
type (
	deleteUserParams struct {
		Id int64 `json:"id"`
	}
)

func (i User) All(c iris.Context) {
	users := new([]models.Users)
	apiUsers := new([]models.ApiUsers)
	Db.Db.Preload("Class").Find(&users)
	copier.Copy(apiUsers, users)
	util.JSON(c, "ok", "success", apiUsers)

}

func (i User) Get(c iris.Context) {
	user := new(models.Users)
	apiUsers := new(models.ApiUsers)
	Db.Db.Model(user).Where("id=?", c.URLParam("id")).First(&apiUsers)
	util.JSON(c, "ok", "success", apiUsers)

}
func (i User) Update(c iris.Context) {

	util.JSON(c, "ok", "保存成功")
}
func (i User) Delete(c iris.Context) {
	params := new(deleteUserParams)

	err := c.ReadForm(params)
	if err != nil && !iris.IsErrPath(err) {
		util.JSON(c, "error", err.Error())
		return
	}
	user := new(models.Users)
	err = user.Get(params.Id)
	if user.ID == 0 || err != nil {
		util.JSON(c, "ok", "删除成功")
		return
	}
	err = user.Delete()
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	util.JSON(c, "ok", "删除成功")
}
