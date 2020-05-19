package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/model"
)

type Repository interface {
	GetTableName() string
	AutoMigrate()
	FindAll() []model.DaoModel
	FindById(id int) model.DaoModel
}

type BaseRepository struct {
	db *gorm.DB
}
