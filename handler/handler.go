package handler

import (
	"github.com/nhannhan159/weather-app-go/config"
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/dao/repository"
)

type RestHandler interface {
	HandleFindAll(context *gin.Context)
	HandleFindById(context *gin.Context)
}

type BaseHandler struct {
	repository repository.Repository
}

type HandlerContext struct {
	GlobalConfig   config.GlobalConfig
	WeatherHandler RestHandler
}

func InitializeHandler(globalConfig config.GlobalConfig, daoContext dao.DaoContext) *HandlerContext {
	return &HandlerContext{
		GlobalConfig: globalConfig,
		WeatherHandler: 
	}
}
