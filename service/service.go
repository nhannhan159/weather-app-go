package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/config"
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/dao/repository"
)

type RestService interface {
	HandleFindAll(context *gin.Context)
	HandleFindByID(context *gin.Context)
}

type BaseService struct {
	repository repository.Repository
}

type Context struct {
	GlobalConfig   config.GlobalConfig
	CityService    RestService
	WeatherService RestService
}

func InitializeService(globalConfig config.GlobalConfig, daoContext *dao.Context) *Context {
	return &Context{
		GlobalConfig:   globalConfig,
		CityService:    NewCityService(daoContext),
		WeatherService: NewWeatherService(daoContext),
	}
}
