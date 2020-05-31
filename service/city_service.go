package service

import (
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type ICityRepository interface {
	FindAll() ([]model.City, error)
	FindByID(int) (*model.City, error)
	Create(*model.City) (*model.City, error)
	Update(*model.City) (*model.City, error)
	Delete(int) error
}

type CityService struct {
	repository ICityRepository
}

func NewCityService(cityRepository ICityRepository) *CityService {
	return &CityService{
		repository: cityRepository,
	}
}

func (service CityService) FindAll(resources *domain.Resources) ([]model.City, error) {
	resources.BizLogger.Infof("find all city")
	return service.repository.FindAll()
}

func (service CityService) FindByID(resources *domain.Resources, id int) (*model.City, error) {
	resources.BizLogger.Infof("find city by id: %v", id)
	return service.repository.FindByID(id)
}

func (service CityService) Create(resources *domain.Resources, entity *model.City) (*model.City, error) {
	resources.BizLogger.Infof("create city: %+v", entity)
	return service.repository.Create(entity)
}

func (service CityService) Update(resources *domain.Resources, entity *model.City) (*model.City, error) {
	resources.BizLogger.Infof("update city: %+v", entity)
	return service.repository.Update(entity)
}

func (service CityService) Delete(resources *domain.Resources, id int) error {
	resources.BizLogger.Infof("delete city by id: %v", id)
	return service.repository.Delete(id)
}
