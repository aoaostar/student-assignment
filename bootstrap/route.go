package bootstrap

import (
	"StudentAssignment/config"
	"StudentAssignment/controllers"
	"StudentAssignment/controllers/admin"
	"StudentAssignment/controllers/api"
	"StudentAssignment/controllers/masterApi"

	"StudentAssignment/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {

	indexParty := app.Party("/")
	authParty := app.Party("/auth")
	jobParty := app.Party("/job")
	adminAuthParty := app.Party("/admin/auth")
	adminParty := app.Party("/admin")

	indexParty.Use(middleware.AuthMiddleware)
	jobParty.Use(middleware.AuthMiddleware)
	adminParty.Use(middleware.AuthAdminMiddleware)

	mvc.New(indexParty).Handle(new(controllers.Index))
	mvc.New(authParty).Handle(new(controllers.Auth))
	mvc.New(jobParty).Handle(new(controllers.Job))
	mvc.New(adminAuthParty).Handle(new(admin.Auth))
	mvc.New(adminParty).Handle(new(admin.Index))

	authApiParty := app.Party("/api/auth")
	{
		authController := new(api.Auth)
		authApiParty.Post("/login", authController.Login)
		authApiParty.Post("/admin_login", authController.AdminLogin)
		authApiParty.Post("/register", authController.Register)
	}
	classApiParty := app.Party("/api")
	{
		classController := new(api.Class)
		classApiParty.Get("/classes", classController.All)
		classApiParty.Get("/class", classController.Get)
	}
	apiParty := app.Party("/api")
	{
		apiParty.Use(middleware.AuthMiddleware)
		userController := new(api.User)
		{
			apiParty.Get("/user", userController.Get)
			apiParty.Put("/user", userController.Update)
		}
		jobController := new(api.Job)
		{
			apiParty.Get("/job", jobController.Get)
			apiParty.Put("/job", iris.LimitRequestBodySize(config.Conf.UploadSize*1<<20), jobController.Update)
			apiParty.Post("/job", jobController.Create)
			apiParty.Delete("/job", jobController.Delete)
		}
		taskController := new(api.Task)
		{
			apiParty.Get("/tasks", taskController.All)
			apiParty.Get("/task", taskController.Get)
			apiParty.Put("/task", taskController.Update)
			apiParty.Post("/task", taskController.Create)
			apiParty.Delete("/task", taskController.Delete)
		}
	}

	masterApiParty := app.Party("/api/master")
	{

		masterApiParty.Use(middleware.AuthAdminMiddleware)
		adminController := new(masterApi.Admin)
		masterApiParty.Get("/admin", adminController.Get)

		userController := new(masterApi.User)
		{
			masterApiParty.Get("/users", userController.All)
			masterApiParty.Get("/user", userController.Get)
			masterApiParty.Put("/user", userController.Update)
			masterApiParty.Delete("/user", userController.Delete)
		}
		jobController := new(masterApi.Job)
		{
			masterApiParty.Get("/jobs", jobController.All)
			masterApiParty.Get("/job", jobController.Get)
			masterApiParty.Put("/job", iris.LimitRequestBodySize(config.Conf.UploadSize*1<<20), jobController.Update)
			masterApiParty.Post("/job", jobController.Create)
			masterApiParty.Delete("/job", jobController.Delete)
		}
		taskController := new(masterApi.Task)
		{
			masterApiParty.Get("/tasks", taskController.All)
			masterApiParty.Get("/task", taskController.Get)
			masterApiParty.Put("/task", taskController.Update)
			masterApiParty.Post("/task", taskController.Create)
			masterApiParty.Delete("/task", taskController.Delete)
		}
	}
}
