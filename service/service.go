package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/config"
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/dao/repository"
)

type RestService interface {
	HandleFindAll(context *gin.Context)
	HandleFindById(context *gin.Context)
}

type BaseService struct {
	repository repository.Repository
}

type ServiceContext struct {
	GlobalConfig   config.GlobalConfig
	CityService    RestService
	WeatherService RestService
}

func InitializeService(globalConfig config.GlobalConfig, daoContext *dao.DaoContext) *ServiceContext {
	return &ServiceContext{
		GlobalConfig:   globalConfig,
		CityService:    NewCityService(daoContext),
		WeatherService: NewWeatherService(daoContext),
	}
}
