package api

import (
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"StudentAssignment/pkg/validator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type Auth struct {
}
type (
	adminLoginParams struct {
		Username string `json:"username" validate:"required,min=4,max=25"`
		Password string `json:"password" validate:"required,min=5,max=25"`
	}
	loginParams struct {
		UserId   int64  `json:"user_id" validate:"required,numeric"`
		Password string `json:"password" validate:"required,min=5,max=25"`
	}
	registerParams struct {
		ClassId         int64  `json:"class_id" validate:"required,numeric"`
		UserId          int64  `json:"user_id" validate:"required,numeric"`
		Username        string `json:"username" validate:"required,min=2,max=6"`
		Password        string `json:"password" validate:"required,min=5,max=25"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}
)

func (i Auth) AdminLogin(c iris.Context) {

	params := new(adminLoginParams)
	err := c.ReadJSON(params)
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	err = validator.Validate.Struct(params)
	if err != nil {

		util.JSON(c, "error", err.Error())
		return
	}
	admin := new(models.Admins)

	Db.Db.Model(admin).Where("username=?", params.Username).First(admin)
	if admin.ID == 0 {
		util.JSON(c, "error", "用户名错误，请重试")
		return
	}
	if !util.ValidatePasswords(admin.Password, params.Password) {
		util.JSON(c, "error", "密码错误，请重试")
		return
	}
	sess := sessions.Get(c)
	sess.Set("admin", admin)

	util.JSON(c, "ok", "登录成功", params)

}
func (i Auth) Login(c iris.Context) {

	params := new(loginParams)
	err := c.ReadJSON(params)
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	err = validator.Validate.Struct(params)
	if err != nil {

		util.JSON(c, "error", err.Error())
		return
	}
	user := new(models.Users)

	Db.Db.Model(user).Where("id=?", params.UserId).Preload("Class").First(user)
	if user.ID == 0 {
		util.JSON(c, "error", "用户名错误，请重试")
		return
	}
	if !util.ValidatePasswords(user.Password, params.Password) {
		util.JSON(c, "error", "密码错误，请重试")
		return
	}
	sess := sessions.Get(c)
	sess.Set("user", user)

	util.JSON(c, "ok", "登录成功", params)

}

func (i Auth) Register(c iris.Context) {
	params := new(registerParams)
	err := c.ReadJSON(params)
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	err = validator.Validate.Struct(params)
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}

	user := new(models.Users)
	Db.Db.Model(user).Where("id=?", params.UserId).First(user)
	if user.ID != 0 {
		util.JSON(c, "error", "该用户已存在，如有疑问请联系管理员")
		return
	}
	user.ID = uint(params.UserId)
	user.ClassId = params.ClassId
	user.Username = params.Username
	user.Password = util.HashAndSalt(params.Password)
	err = user.Create()
	if err != nil {
		util.JSON(c, "error", err.Error())
		return
	}
	util.JSON(c, "ok", "注册成功", params)
}
