package service

import (
	"github.com/nhannhan159/weather-app-go/dao"
	"github.com/nhannhan159/weather-app-go/domain"
)

type ICityService interface {
	domain.IService
}

type cityService struct {
	baseService
}

func NewCityService(cityRepository dao.ICityRepository) ICityService {
	return &cityService{
		baseService: baseService{
			repository: cityRepository,
		},
	}
}

func (service cityService) FindAll(resources *domain.Resources) (interface{}, error) {
	return service.repository.FindAll()
}

func (service cityService) FindByID(resources *domain.Resources, id int) (interface{}, error) {
	return service.repository.FindByID(id)
}
