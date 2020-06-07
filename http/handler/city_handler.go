package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type ICityService interface {
	FindAll(*domain.Resources) ([]model.City, error)
	FindByID(*domain.Resources, int) (*model.City, error)
}

type CityHandler struct {
	baseRestHandler
	cityService ICityService
}

func NewCityHandler(cityService ICityService) *CityHandler {
	return &CityHandler{
		cityService: cityService,
	}
}

// CityHandler.HandleFindAll godoc
// @Summary Get city list
// @Description get city list
// @Accept  json
// @Produce  json
// @Success 200 {array} model.City
// @Failure 400 {object} handler.HTTPError
// @Failure 404 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /city [get]
func (handler *CityHandler) HandleFindAll(ctx *gin.Context) {
	res, err := handler.cityService.FindAll(handler.getResourcesFromContext(ctx))
	if err != nil {
		handler.handleBabRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// CityHandler.HandleFindByID godoc
// @Summary Get city by id
// @Description get city by id
// @ID get-city-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "City ID"
// @Success 200 {object} model.City
// @Failure 400 {object} handler.HTTPError
// @Failure 404 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /city/{id} [get]
func (handler *CityHandler) HandleFindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handler.handleBabRequest(ctx, errors.New("id must be an integer"))
		return
	}

	res, err2 := handler.cityService.FindByID(handler.getResourcesFromContext(ctx), id)
	if err2 != nil {
		handler.handleBabRequest(ctx, err2)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
