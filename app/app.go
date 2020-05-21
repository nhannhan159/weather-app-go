package app

import (
	"github.com/nhannhan159/weather-app-go/domain"
)

type App struct {
	GlobalConfig *domain.GlobalConfig
	Resources    *domain.Resources
	Repositories *Repositories
	Services     *Services
	HTTPHandlers *HTTPHandlers
}

func (app *App) Initialize() {
	app.initializeConfig()
	app.initializeResources()
	app.initializeDependencies()
}

func (app App) Run() {
	app.runResources()
}
