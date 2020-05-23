package service

import (
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/domain"
)

type IWeatherService interface {
	domain.IService
}

type weatherService struct {
	baseService
}

func NewWeatherService(weatherRepository dao.IWeatherRepository) IWeatherService {
	return &weatherService{
		baseService: baseService{
			repository: weatherRepository,
		},
	}
}

func (service weatherService) FindAll(resources *domain.Resources) (interface{}, error) {
	return service.repository.FindAll()
}

func (service weatherService) FindByID(resources *domain.Resources, id int) (interface{}, error) {
	return service.repository.FindByID(id)
}

func (service weatherService) Create(resources *domain.Resources, entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (service weatherService) Update(resources *domain.Resources, entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (service weatherService) Delete(resources *domain.Resources, id int) (interface{}, error) {
	panic("implement me")
}
