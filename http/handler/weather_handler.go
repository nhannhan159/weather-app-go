package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type IWeatherService interface {
	FindAll(*domain.Resources) ([]model.Weather, error)
	FindByID(*domain.Resources, int) (*model.Weather, error)
}

type WeatherHandler struct {
	baseRestHandler
	weatherService IWeatherService
}

func NewWeatherHandler(weatherService IWeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

func (handler *WeatherHandler) HandleFindAll(ctx *gin.Context) {
	res, err := handler.weatherService.FindAll(handler.getResourcesFromContext(ctx))
	if err != nil {
		handler.handleBabRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (handler *WeatherHandler) HandleFindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handler.handleBabRequest(ctx, errors.New("id must be an integer"))
		return
	}

	res, err2 := handler.weatherService.FindByID(handler.getResourcesFromContext(ctx), id)
	if err2 != nil {
		handler.handleBabRequest(ctx, err2)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
