package http

import (
	_ "github.com/nhannhan159/weather-app-go/docs"
	"github.com/nhannhan159/weather-app-go/http/handler"
)

func (server *HTTPManager) initializeRouter() {
	handlers := server.handlerCollection
	router := server.engine

	router.GET("/swagger/*any", handler.SwaggerHandler(server.config))

	router.GET("/", handlers.IndexHandler.Handle)
	router.GET("/user", handlers.UserHandler.Handle)

	apiV1Group := router.Group("/api/v1")
	{
		apiV1Group.GET("/city", handlers.CityHandler.HandleFindAll)
		apiV1Group.GET("/city/:id", handlers.CityHandler.HandleFindByID)

		apiV1Group.GET("/weather", handlers.WeatherHandler.HandleFindAll)
		apiV1Group.GET("/weather/:id", handlers.WeatherHandler.HandleFindByID)
	}
}
