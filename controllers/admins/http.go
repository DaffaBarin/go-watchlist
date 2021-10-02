package admins

import (
	"go-watchlist/business/admins"
	controller "go-watchlist/controllers"
	"go-watchlist/controllers/admins/request"
	"go-watchlist/controllers/admins/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminService admins.Service
}

func NewAdminController(service admins.Service) *AdminController {
	return &AdminController{
		AdminService: service,
	}
}

func (ctrl *AdminController) Register(c echo.Context) error {
	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.AdminService.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainRegister(data))
}

func (ctrl *AdminController) Login(c echo.Context) error {
	req := request.AdminsLogin{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.AdminService.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainLogin(data))
}
