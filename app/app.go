package app

import (
	"github.com/nhannhan159/weather-app-go/common"
	"github.com/nhannhan159/weather-app-go/domain"
)

type App struct {
	GlobalConfig *common.GlobalConfig
	Resources    *domain.Resources
	Repositories *Repositories
	Services     *Services
	HTTPHandlers *HTTPHandlers
	GRPCHandlers *GRPCHandlers
}

func (app *App) Initialize() {
	app.initializeConfig()
	app.initializeResources()
	app.initializeDependencies()
}

func (app App) Run() {
	app.runResources()
}
