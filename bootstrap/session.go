package bootstrap

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

const cookieNameForSessionID = "aoaostar_session_id"

func InitSession(app *iris.Application) {

	sess := sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
	app.Use(sess.Handler())
}
