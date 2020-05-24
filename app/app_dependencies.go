package app

import (
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/http/handler"
	"github.com/nhannhan159/weather-app-go/service"
)

type Repositories struct {
	CityRepository    *dao.CityRepository
	WeatherRepository *dao.WeatherRepository
}

type Services struct {
	CityService    *service.CityService
	WeatherService *service.WeatherService
}

type HTTPHandlers struct {
	indexHandler   *handler.IndexHandler
	userHandler    *handler.UserHandler
	cityHandler    *handler.CityHandler
	weatherHandler *handler.WeatherHandler
}

func (app *App) initializeDependencies() {
	app.initialRepositories()
	app.initialServices()
	app.initialHandlers()
}

func (app *App) initialRepositories() {
	cityRepository := dao.NewCityRepository(app.Resources.DaoManager)
	weatherRepository := dao.NewWeatherRepository(app.Resources.DaoManager)

	app.Resources.DaoManager.AddRepository(cityRepository)
	app.Resources.DaoManager.AddRepository(weatherRepository)

	app.Repositories = &Repositories{
		CityRepository:    cityRepository,
		WeatherRepository: weatherRepository,
	}
}

func (app *App) initialServices() {
	app.Services = &Services{
		CityService:    service.NewCityService(app.Repositories.CityRepository),
		WeatherService: service.NewWeatherService(app.Repositories.WeatherRepository),
	}
}

func (app *App) initialHandlers() {
	indexHandler := handler.NewIndexHandler()
	userHandler := handler.NewUserHandler()
	cityHandler := handler.NewCityHandler(app.Services.CityService)
	weatherHandler := handler.NewWeatherHandler(app.Services.WeatherService)

	app.HTTPHandlers = &HTTPHandlers{
		indexHandler:   indexHandler,
		userHandler:    userHandler,
		cityHandler:    cityHandler,
		weatherHandler: weatherHandler,
	}
}
