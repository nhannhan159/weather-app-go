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

func (handler *CityHandler) HandleFindAll(ctx *gin.Context) {
	res, err := handler.cityService.FindAll(handler.getResourcesFromContext(ctx))
	if err != nil {
		handler.handleBabRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

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
