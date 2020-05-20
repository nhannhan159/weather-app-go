package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/dao"
)

type WeatherService struct {
	BaseService
}

func NewWeatherService(daoContext *dao.Context) *WeatherService {
	return &WeatherService{
		BaseService: BaseService{
			repository: daoContext.WeatherRepository,
		},
	}
}

func (service WeatherService) HandleFindAll(context *gin.Context) {
	res, err := service.repository.FindAll()

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, res)
}

func (service WeatherService) HandleFindByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		fmt.Println("id must be an integer")
		context.JSON(http.StatusBadRequest, "error")
	}

	res, err2 := service.repository.FindByID(id)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, err2)
		return
	}

	context.JSON(http.StatusOK, res)
}
