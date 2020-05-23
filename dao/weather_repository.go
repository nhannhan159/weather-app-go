package dao

import (
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type IWeatherRepository interface {
	domain.IRepository
}

type weatherRepository struct {
	BaseRepository
}

func NewWeatherRepository(daoManager domain.IDaoManager) IWeatherRepository {
	this := &weatherRepository{
		BaseRepository: BaseRepository{
			db: daoManager.GetDB().Table(domain.TableWeather),
		},
	}
	return this
}

func (repository *weatherRepository) AutoMigrate() {
	repository.db.AutoMigrate(&model.Weather{})
}

func (repository *weatherRepository) Create(entity interface{}) error {
	return nil
}

func (repository *weatherRepository) Update(entity interface{}) error {
	return nil
}

func (repository *weatherRepository) Delete(id int) error {
	panic("implement me")
}

func (repository weatherRepository) FindAll() (interface{}, error) {
	var entities []model.Weather
	err := repository.db.Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (repository weatherRepository) FindByID(id int) (interface{}, error) {
	entity := &model.Weather{}
	err := repository.db.First(entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
