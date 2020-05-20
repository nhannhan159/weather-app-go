package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/model"
)

type WeatherRepository struct {
	BaseRepository
}

func NewWeatherRepository(db *gorm.DB) *WeatherRepository {
	this := &WeatherRepository{
		BaseRepository: BaseRepository{
			db: db.Table("weather"),
		},
	}
	this.db.AutoMigrate(&model.Weather{})
	return this
}

func (this WeatherRepository) cloneEntity() model.DaoModel {
	return &model.Weather{}
}

func (this WeatherRepository) cloneEntities() []model.DaoModel {
	slice := make([]model.DaoModel, 1)
	slice[0] = this.cloneEntity()
	return slice
}
