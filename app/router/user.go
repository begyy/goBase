package router

import (
	"github.com/labstack/echo/v4"
	"goBase/app/api"
	"goBase/app/middleware"
)

func ConfigureUserRoutes(e *echo.Echo) {

	e.POST("/api/v1/user/sign-up/", api.SignUp)
	e.POST("/api/v1/user/sign-in/", api.SignIn)

	userGroup := e.Group("/api/v1/user")
	userGroup.Use(middleware.AuthenticationMiddleware())
	userGroup.GET("/me/", api.UserMe)
}
