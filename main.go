package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/nhannhan159/weather-app-go/app"
)

func main() {
	mainApp := app.App{}
	mainApp.Initialize()
	mainApp.Run()
}
