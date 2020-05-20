package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/dao"
)

type WeatherService struct {
	BaseService
}

func NewWeatherService(daoContext *dao.DaoContext) *WeatherService {
	return &WeatherService{
		BaseService: BaseService{
			repository: daoContext.WeatherRepository,
		},
	}
}

func (this WeatherService) HandleFindAll(context *gin.Context) {
	res, err := this.repository.FindAll()
	if err != nil {
		context.JSON(400, err)
		return
	}
	context.JSON(200, res)
}

func (this WeatherService) HandleFindById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		fmt.Println("id must be an integer")
		context.JSON(400, "error")
	}
	res, err2 := this.repository.FindById(id)
	if err2 != nil {
		context.JSON(400, err2)
		return
	}
	context.JSON(200, res)
}
