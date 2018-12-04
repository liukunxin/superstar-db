package main

import (
	"github.com/liukunxin/superstar-db/bootstrap"
	"github.com/liukunxin/superstar-db/web/middleware/identity"
	"github.com/liukunxin/superstar-db/web/routes"
)

func newApp() *bootstrap.Bootstrapper  {
	app := bootstrap.New("Superstar database", "Quincy")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main()  {
	app := newApp()
	app.Listen(":8080")
}
