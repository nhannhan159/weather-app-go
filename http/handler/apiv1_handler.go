package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
)

type IApiV1Handler interface {
	domain.IHandler
}

type apiV1Handler struct {
	baseHandler
	userHandler    IUserHandler
	cityHandler    ICityHandler
	weatherHandler IWeatherHandler
}

func NewApiV1Handler(
	userHandler IUserHandler,
	cityHandler ICityHandler,
	weatherHandler IWeatherHandler,
) IApiV1Handler {
	return &apiV1Handler{
		userHandler:    userHandler,
		cityHandler:    cityHandler,
		weatherHandler: weatherHandler,
	}
}

func (handler *apiV1Handler) Handle(router *gin.Engine) {
	apiV1Group := router.Group("/api/v1")
	{
		handler.userHandler.HandleGroup(apiV1Group)
		handler.cityHandler.HandleGroup(apiV1Group)
		handler.weatherHandler.HandleGroup(apiV1Group)
	}
}
