package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/model"
)

type WeatherRepository struct {
	BaseRepository
}

func NewWeatherRepository(db *gorm.DB) *WeatherRepository {
	repository := &WeatherRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
	repository.AutoMigrate()
	return repository
}

func (repository WeatherRepository) GetTableName() string {
	return "weather"
}

func (repository WeatherRepository) AutoMigrate() {
	repository.db.Table(repository.GetTableName()).AutoMigrate(&model.Weather{})
}

func (repository WeatherRepository) FindAll() []model.Weather {
	weathers := []model.Weather{}
	err := repository.db.Table(repository.GetTableName()).Find(&weathers).Error
	if err != nil {
		panic(err)
	}
	return weathers
}

func (repository WeatherRepository) FindById(id int) model.Weather {
	weather := model.Weather{}
	err := repository.db.Table(repository.GetTableName()).First(&weather).Error
	if err != nil {
		panic(err)
	}
	return weather
}
