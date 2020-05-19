package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/model"
)

type Repository interface {
	GetTableName() string
	AutoMigrate()
	FindAll() []model.Weather
	FindById(id int) model.Weather
}

type BaseRepository struct {
	db *gorm.DB
}
