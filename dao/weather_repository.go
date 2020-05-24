package dao

import (
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type WeatherRepository struct {
	baseRepository
}

func NewWeatherRepository(daoManager domain.IDaoManager) *WeatherRepository {
	this := &WeatherRepository{
		baseRepository: baseRepository{
			db: daoManager.GetDB().Table(domain.TableWeather),
		},
	}
	return this
}

func (repository *WeatherRepository) AutoMigrate() {
	repository.db.AutoMigrate(&model.Weather{})
}

func (repository *WeatherRepository) Create(entity *model.Weather) (*model.Weather, error) {
	panic("implement me")
}

func (repository *WeatherRepository) Update(entity *model.Weather) (*model.Weather, error) {
	panic("implement me")
}

func (repository *WeatherRepository) Delete(id int) error {
	panic("implement me")
}

func (repository WeatherRepository) FindAll() ([]model.Weather, error) {
	var entities []model.Weather
	err := repository.db.Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (repository WeatherRepository) FindByID(id int) (*model.Weather, error) {
	entity := &model.Weather{}
	err := repository.db.First(entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
