package service

import (
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type IWeatherRepository interface {
	FindAll() ([]model.Weather, error)
	FindByID(int) (*model.Weather, error)
	Create(*model.Weather) (*model.Weather, error)
	Update(*model.Weather) (*model.Weather, error)
	Delete(int) error
}

type WeatherService struct {
	repository IWeatherRepository
}

func NewWeatherService(weatherRepository IWeatherRepository) *WeatherService {
	return &WeatherService{
		repository: weatherRepository,
	}
}

func (service WeatherService) FindAll(resources *domain.Resources) ([]model.Weather, error) {
	resources.BizLogger.Infof("find all weather")
	return service.repository.FindAll()
}

func (service WeatherService) FindByID(resources *domain.Resources, id int) (*model.Weather, error) {
	resources.BizLogger.Infof("find weather by id: %v", id)
	return service.repository.FindByID(id)
}

func (service WeatherService) Create(resources *domain.Resources, entity *model.Weather) (*model.Weather, error) {
	resources.BizLogger.Infof("create weather: %+v", entity)
	return service.repository.Create(entity)
}

func (service WeatherService) Update(resources *domain.Resources, entity *model.Weather) (*model.Weather, error) {
	resources.BizLogger.Infof("update weather: %+v", entity)
	return service.repository.Update(entity)
}

func (service WeatherService) Delete(resources *domain.Resources, id int) error {
	resources.BizLogger.Infof("delete weather by id: %v", id)
	return service.repository.Delete(id)
}
