package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Auth struct {
}

func (m *Auth) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Login")
}

func (m Auth) Login(c iris.Context) {
	c.ViewData("Title", "管理员登录")
	c.View("admin/login.html")
}
