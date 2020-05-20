package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/service"
)

func InitializeRoutes(router *gin.Engine, serviceContext *service.Context) {
	router.GET("/", indexHandler)
	apiV1Group := router.Group("/api/v1")
	{
		userHandler(apiV1Group)
		cityHandler(apiV1Group, serviceContext)
		weatherHandler(apiV1Group, serviceContext)
	}
}
