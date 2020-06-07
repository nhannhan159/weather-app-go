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

// WeatherHandler.HandleFindAll godoc
// @Summary Get weather list
// @Description get weather list
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Weather
// @Failure 400 {object} handler.HTTPError
// @Failure 404 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /weather [get]
func (handler *WeatherHandler) HandleFindAll(ctx *gin.Context) {
	res, err := handler.weatherService.FindAll(handler.getResourcesFromContext(ctx))
	if err != nil {
		handler.handleBabRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// WeatherHandler.HandleFindByID godoc
// @Summary Get weather by id
// @Description get weather by id
// @ID get-weather-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Weather ID"
// @Success 200 {object} model.Weather
// @Failure 400 {object} handler.HTTPError
// @Failure 404 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /weather/{id} [get]
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
