package router

import (
	"go-web-gin-example/router/handler"
	. "go-web-gin-example/api/userApi"
)

func initUserRouter() {
	// public api
	router.GET("/login", LoginApi)

	user := router.Group("/user", handler.UserHandler())
	{
		user.GET("/refresh", RefreshUserInfoApi)
	}
}
