package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/service"
)

func weatherHandler(group *gin.RouterGroup, serviceContext *service.ServiceContext) {
	weatherService := serviceContext.WeatherService
	weatherGroup := group.Group("/weather")
	{
		weatherGroup.GET("/", weatherService.HandleFindAll)
		weatherGroup.GET("/:ID", weatherService.HandleFindById)
	}

}
