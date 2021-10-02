package routes

import (
	// middlewareApp "go-watchlist/app/middlewares"
	"go-watchlist/controllers/admins"
	"go-watchlist/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	AdminController admins.AdminController
	UserController  users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	admins := e.Group("admins")
	admins.POST("/register", cl.AdminController.Register)
	admins.POST("/login", cl.AdminController.Login)

	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)
}
