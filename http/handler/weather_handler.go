package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/service"
)

type IWeatherHandler interface {
	domain.IHandler
}

type weatherHandler struct {
	baseHandler
	weatherService service.IWeatherService
}

func NewWeatherHandler(weatherService service.IWeatherService) IWeatherHandler {
	return &weatherHandler{
		weatherService: weatherService,
	}
}

func (handler *weatherHandler) HandleGroup(group *gin.RouterGroup) {
	cityGroup := group.Group("/weather")
	{
		cityGroup.GET("/", handler.handleFindAll)
		cityGroup.GET("/:id", handler.handleFindByID)
	}
}

func (handler *weatherHandler) handleFindAll(ctx *gin.Context) {
	res, err := handler.weatherService.FindAll(handler.getResourcesFromContext(ctx))
	if err != nil {
		handler.handleBabRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (handler *weatherHandler) handleFindByID(ctx *gin.Context) {
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
