package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/service"
)

func cityHandler(group *gin.RouterGroup, serviceContext *service.Context) {
	cityService := serviceContext.CityService
	cityGroup := group.Group("/city")
	{
		cityGroup.GET("/", cityService.HandleFindAll)
		cityGroup.GET("/:id", cityService.HandleFindByID)
	}
}
