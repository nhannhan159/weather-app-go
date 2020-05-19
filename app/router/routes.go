package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/nhannhan159/weather-app-go/app/router/api_v1"
	"github.com/nhannhan159/weather-app-go/handler"
)

func InitializeRoutes(router *gin.Engine, handlerContext *handler.HandlerContext) {
	router.GET("/", index)
	apiV1Group := router.Group("/api/v1")
	{
		apiV1.User(apiV1Group)
		apiV1.Weather(apiV1Group, handlerContext)
	}
}
