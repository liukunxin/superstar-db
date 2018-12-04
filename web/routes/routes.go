package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/liukunxin/superstar-db/bootstrap"
	"github.com/liukunxin/superstar-db/services"
	"github.com/liukunxin/superstar-db/web/controllers"
	"github.com/liukunxin/superstar-db/web/middleware"
)

func Configure(b *bootstrap.Bootstrapper)  {
	superstarService := services.NewSuperstarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BaicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))
}
