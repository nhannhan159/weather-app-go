package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/nhannhan159/weather-app-go/app"
	"github.com/nhannhan159/weather-app-go/config"
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/handler"
)

func main() {
	globalConfig := config.InitializeConfig()
	daoContext := dao.InitializeDao(configValue.Database)
	handlerContext := handler.InitializeHandler(configValue, daoContext)
	app.InitializeServer(handlerContext)
}
