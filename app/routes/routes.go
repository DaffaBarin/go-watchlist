package routes

import (
	"errors"
	middlewareApp "go-watchlist/app/middlewares"
	controller "go-watchlist/controllers"
	"go-watchlist/controllers/admins"
	"go-watchlist/controllers/medias"
	"go-watchlist/controllers/users"
	"go-watchlist/controllers/watchlists"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig           middleware.JWTConfig
	AdminController     admins.AdminController
	UserController      users.UserController
	MediaController     medias.MediaController
	WatchlistController watchlists.WatchlistController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	admins := e.Group("admins")
	admins.POST("/register", cl.AdminController.Register)
	admins.POST("/login", cl.AdminController.Login)

	adminsPost := admins
	adminsPost.POST("/media", cl.MediaController.Create, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationAdmin())

	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)
	users.POST("/watchlist", cl.WatchlistController.Create, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationUser())
	users.GET("/watchlist", cl.WatchlistController.GetAllByUserID, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationUser())
	users.GET("/watchlist/:id", cl.WatchlistController.GetByID, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationUser())
	users.POST("/watchlist/:id", cl.WatchlistController.InsertMedia, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationUser())
	users.PUT("/watchlist/:watchlist/:media", cl.WatchlistController.UpdateMedia, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationUser())

	medias := e.Group("medias")
	medias.GET("/", cl.MediaController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationUser())
	medias.GET("/:id", cl.MediaController.GetByID, middleware.JWTWithConfig(cl.JwtConfig), RoleValidationUser())
}

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("unathorized"))
			}
		}
	}
}

func RoleValidationUser() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" || claims.Role == "user" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("unathorized"))
			}
		}
	}
}
