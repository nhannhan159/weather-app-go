package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/handler"
)

func Weather(group *gin.RouterGroup, handlerContext *handler.HandlerContext) {
	weatherHandler := handlerContext.WeatherHandler
	weatherGroup := group.Group("/weather")
	{
		weatherGroup.GET("/", weatherHandler.HandleFindAll)
		weatherGroup.GET("/:ID", weatherHandler.HandleFindById)
	}

}
