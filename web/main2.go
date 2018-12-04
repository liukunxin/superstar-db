package main

import (
	"github.com/liukunxin/superstar-db/bootstrap"
	"github.com/liukunxin/superstar-db/web/middleware/identity"
	"github.com/liukunxin/superstar-db/web/routes"
)

func main() {
	app := bootstrap.New("Superstar database", "Quincy")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8081")
}
