package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/dao"
)

type WeatherHandler struct {
	BaseHandler
}

func NewWeatherHandler(daoContext *dao.DaoContext) *WeatherHandler {
	return &WeatherHandler{
		repository: daoContext.WeatherRepository,
	}
}

func (handler WeatherHandler) HandleFindAll(context *gin.Context) {
	context.JSON(200, handler.repository.FindAll())
}

func (handler WeatherHandler) HandleFindById(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, handler.repository.FindById(id))
}
