package middleware

import "github.com/kataras/iris/middleware/basicauth"

var BaicAuth  = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin" : "password",
	},
})
