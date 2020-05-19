package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/dao"
)

type WeatherHandler struct {
	BaseHandler
}

func NewWeatherHandler(daoContext *dao.DaoContext) *WeatherHandler {
	return &WeatherHandler{
		BaseHandler: BaseHandler{
			repository: daoContext.WeatherRepository,
		},
	}
}

func (handler WeatherHandler) HandleFindAll(context *gin.Context) {
	context.JSON(200, handler.repository.FindAll())
}

func (handler WeatherHandler) HandleFindById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		fmt.Println("id must be an integer")
		context.JSON(400, "error")
	}
	context.JSON(200, handler.repository.FindById(id))
}
