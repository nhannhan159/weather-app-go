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
	this.newEntity = this.cloneEntity
	this.newEntities = this.cloneEntities
	this.db.AutoMigrate(&model.Weather{})

	return this
}

func (repository WeatherRepository) cloneEntity() interface{} {
	return &model.Weather{}
}

func (repository WeatherRepository) cloneEntities() interface{} {
	return []model.Weather{}
}
