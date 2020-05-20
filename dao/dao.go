package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/config"
	"github.com/nhannhan159/weather-app-go/dao/repository"
)

type Context struct {
	DB                *gorm.DB
	CityRepository    repository.Repository
	WeatherRepository repository.Repository
}

func InitializeDao(dbConfig config.DatabaseConfig) *Context {
	db, err := gorm.Open(dbConfig.Dialect, dbConfig.DSN)
	if err != nil {
		panic(err)
	}

	return &Context{
		DB:                db,
		CityRepository:    repository.NewCityRepository(db),
		WeatherRepository: repository.NewWeatherRepository(db),
	}
}
