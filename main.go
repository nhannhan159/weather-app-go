package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/nhannhan159/weather-app-go/app"
	"github.com/nhannhan159/weather-app-go/config"
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/service"
)

func main() {
	globalConfig := config.InitializeConfig()
	daoContext := dao.InitializeDao(globalConfig.Database)
	serviceContext := service.InitializeService(globalConfig, daoContext)
	app.InitializeServer(globalConfig.Server, serviceContext)
}
