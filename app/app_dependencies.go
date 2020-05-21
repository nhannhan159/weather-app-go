package app

import (
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/http/handler"
	"github.com/nhannhan159/weather-app-go/service"
)

type Repositories struct {
	CityRepository    dao.ICityRepository
	WeatherRepository dao.IWeatherRepository
}

type Services struct {
	CityService    service.ICityService
	WeatherService service.IWeatherService
}

type HTTPHandlers struct {
	indexHandler       handler.IIndexHandler
	userHandler        handler.IUserHandler
	cityHandler        handler.ICityHandler
	weatherHandler     handler.IWeatherHandler
	apiV1Handler       handler.IApiV1Handler
	registeredHandlers []domain.IHandler
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
	apiV1Handler := handler.NewApiV1Handler(userHandler, cityHandler, weatherHandler)

	app.HTTPHandlers = &HTTPHandlers{
		indexHandler:   indexHandler,
		userHandler:    userHandler,
		cityHandler:    cityHandler,
		weatherHandler: weatherHandler,
		apiV1Handler:   apiV1Handler,
		registeredHandlers: []domain.IHandler{
			indexHandler,
			apiV1Handler,
		},
	}
}
