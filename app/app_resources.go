package app

import (
	"github.com/nhannhan159/weather-app-go/common"
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/grpc"
	"github.com/nhannhan159/weather-app-go/http"
)

func (app *App) initializeResources() {
	app.Resources = &domain.Resources{
		DaoManager:  dao.NewDaoManager(app.GlobalConfig.Database),
		HTTPManager: http.NewHTTPManager(app.GlobalConfig),
		GRPCManager: grpc.NewGRPCServerManager(app.GlobalConfig.GRPC),
		GinLogger:   common.NewGinLogger(app.GlobalConfig).Sugar(),
		BizLogger:   common.NewBizLogger(app.GlobalConfig).Sugar(),
		GRPCLogger:  common.NewGRPCLogger(app.GlobalConfig),
	}
	common.InitializeTracer(app.GlobalConfig)
}

func (app *App) runResources() {
	app.Resources.DaoManager.AutoMigrate()

	grpcManager := app.Resources.GRPCManager
	grpcManager.RegisterResources(app.Resources)
	grpcManager.RegisterHandler(app.GRPCHandlers.cityHandler)
	grpcManager.RegisterHandler(app.GRPCHandlers.weatherHandler)
	go grpcManager.Run()

	httpManager := app.Resources.HTTPManager
	httpManager.RegisterResources(app.Resources)
	httpManager.RegisterHandler(&domain.HandlerCollection{
		IndexHandler:   app.HTTPHandlers.indexHandler,
		UserHandler:    app.HTTPHandlers.userHandler,
		CityHandler:    app.HTTPHandlers.cityHandler,
		WeatherHandler: app.HTTPHandlers.weatherHandler,
	})
	httpManager.Run()
}
