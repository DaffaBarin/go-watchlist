package watchlists

import (
	"go-watchlist/app/middlewares"
	"go-watchlist/business/watchlists"
	controller "go-watchlist/controllers"
	"go-watchlist/controllers/watchlists/request"
	"go-watchlist/controllers/watchlists/response"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type WatchlistController struct {
	WatchlistService watchlists.Service
}

func NewWatchlistController(service watchlists.Service) *WatchlistController {
	return &WatchlistController{
		WatchlistService: service,
	}
}

func (ctrl *WatchlistController) Create(c echo.Context) error {
	req := request.CreateWatchlist{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	user := middlewares.GetUser(c)
	data, err := ctrl.WatchlistService.Create(user.ID, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomainCreateWatchlistResponse(data))
}

func (ctrl *WatchlistController) GetAllByUserID(c echo.Context) error {
	user := middlewares.GetUser(c)
	data, err := ctrl.WatchlistService.GetAllByUserID(user.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *WatchlistController) GetByID(c echo.Context) error {
	user := middlewares.GetUser(c)
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := ctrl.WatchlistService.GetByID(user.ID, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *WatchlistController) InsertMedia(c echo.Context) error {
	req := request.InsertMedia{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	data, err := ctrl.WatchlistService.InsertMedia(id, req.MediaID)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainInsertWatchlistResponse(data))
}

func (ctrl *WatchlistController) UpdateMedia(c echo.Context) error {

	watchlist, _ := strconv.Atoi(c.Param("watchlist"))
	media, _ := strconv.Atoi(c.Param("media"))
	user := middlewares.GetUser(c)
	data, err := ctrl.WatchlistService.UpdateMedia(user.ID, watchlist, media)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainUpdateMediaResponse(data))
}
