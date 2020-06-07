package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/nhannhan159/weather-app-go/app"
)

// @title Weather app API
// @version 1.0
// @description This is nhannhan159 go application.

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	mainApp := app.App{}
	mainApp.Initialize()
	mainApp.Run()
}
