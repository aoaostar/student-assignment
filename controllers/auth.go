package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type Auth struct {

}
func (m *Auth) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login", "Login")
	b.Handle("GET", "/logout", "Logout")
	b.Handle("GET", "/register", "Register")
}

func (m Auth) Login(c iris.Context)  {
	c.ViewData("Title", "登录")
	c.View("auth/login.html")
}

func (m Auth) Logout(c iris.Context)  {
	session := sessions.Get(c)
	session.Clear()
	c.Redirect("/")
}

func (m Auth) Register(c iris.Context)  {
	c.ViewData("Title", "注册")
	c.View("auth/register.html")
}