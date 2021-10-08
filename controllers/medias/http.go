package medias

import (
	"errors"
	"go-watchlist/business/medias"
	controller "go-watchlist/controllers"
	"go-watchlist/controllers/medias/request"
	"go-watchlist/controllers/medias/response"
	"go-watchlist/drivers/thirdparties/tmdb"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type MediaController struct {
	MediaService medias.Service
}

func NewMediaController(service medias.Service) *MediaController {
	return &MediaController{
		MediaService: service,
	}
}

func (ctrl *MediaController) Create(c echo.Context) error {
	req := request.MediasCreate{}
	res := request.Medias{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	// panggil third party, ubah req ke data tmdb
	if req.Media_Type == "Movie" {
		res, _ = tmdb.TransformMovie(req)
	} else if req.Media_Type == "Tv" {
		res, _ = tmdb.TransformTV(req)
	}
	data, err := ctrl.MediaService.Create(res.ToDomain())
	if data.ID == 0 {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, errors.New("id invalid"))
	}
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainInsert(data))
}

func (ctrl *MediaController) GetAll(c echo.Context) error {

	data, err := ctrl.MediaService.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *MediaController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := ctrl.MediaService.GetByID(id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)
}
