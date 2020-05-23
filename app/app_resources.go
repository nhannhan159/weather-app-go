package app

import (
	"github.com/nhannhan159/weather-app-go/common"
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/http"
)

func (app *App) initializeResources() {
	app.Resources = &domain.Resources{
		DaoManager:  dao.NewDaoManager(app.GlobalConfig.Database),
		HTTPManager: http.NewHTTPManager(app.GlobalConfig.Server),
		Logger:      common.NewLogger(),
	}
}

func (app *App) runResources() {
	app.Resources.DaoManager.AutoMigrate()

	httpManager := app.Resources.HTTPManager
	httpManager.RegisterResources(app.Resources)

	for _, handler := range app.HTTPHandlers.registeredHandlers {
		httpManager.RegisterHandler(handler)
	}

	httpManager.Run()
}
