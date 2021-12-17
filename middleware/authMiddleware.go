package middleware

import (
	"StudentAssignment/pkg/util"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"net/http"
)

func AuthAdminMiddleware(c iris.Context) {
	sess := sessions.Get(c)
	user := sess.Get("admin")
	if user == nil {
		if c.IsAjax() || c.Method() != http.MethodGet {
			c.StatusCode(http.StatusUnauthorized)
			util.JSON(c, "error", "请登录")
			c.Skip()
		}
		c.Redirect("/admin/auth")
	}
	c.Next()

}

func AuthMiddleware(c iris.Context) {
	sess := sessions.Get(c)
	user := sess.Get("user")
	if user == nil {
		if c.IsAjax() || c.Method() != http.MethodGet {
			c.StatusCode(http.StatusUnauthorized)
			util.JSON(c, "error", "请登录")
			c.Skip()
		}
		c.Redirect("/auth/login")
	}
	c.Next()

}