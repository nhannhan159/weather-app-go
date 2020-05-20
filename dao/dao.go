package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/config"
	"github.com/nhannhan159/weather-app-go/dao/repository"
)

type DaoContext struct {
	DB                *gorm.DB
	CityRepository    repository.Repository
	WeatherRepository repository.Repository
}

func InitializeDao(dbConfig config.DatabaseConfig) *DaoContext {
	db, err := gorm.Open(dbConfig.Dialect, dbConfig.DSN)
	if err != nil {
		panic(err)
	}

	return &DaoContext{
		DB:                db,
		CityRepository:    repository.NewCityRepository(db),
		WeatherRepository: repository.NewWeatherRepository(db),
	}
}
