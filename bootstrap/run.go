package bootstrap
import (
	"StudentAssignment/config"
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
)

func Run()  {
	app := iris.Default()
	pugEngine := iris.HTML("./views", ".html")
	pugEngine.Delims("{[","]}")
	pugEngine.Layout("layouts/layout.html")
	pugEngine.Reload(true)
	app.RegisterView(pugEngine)

	app.HandleDir("/","public")

	InitConfig()
	InitLog()
	InitValidator()
	InitSession(app)
	InitRouter(app)
	InitDatabase()

	err := app.Run(iris.Addr(config.Conf.Listen))
	if err != nil {
		log.Fatal(err)
	}

}
