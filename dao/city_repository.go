package dao

import (
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type ICityRepository interface {
	domain.IRepository
}

type cityRepository struct {
	BaseRepository
}

func NewCityRepository(daoManager domain.IDaoManager) ICityRepository {
	this := &cityRepository{
		BaseRepository: BaseRepository{
			db: daoManager.GetDb().Table(domain.TableCity),
		},
	}
	return this
}

func (repository *cityRepository) AutoMigrate() {
	repository.db.AutoMigrate(&model.City{})
}

func (repository *cityRepository) Create(entity interface{}) error {
	return nil
}

func (repository *cityRepository) Update(entity interface{}) error {
	return nil
}

func (repository cityRepository) FindAll() (interface{}, error) {
	var entities []model.City
	err := repository.db.Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (repository cityRepository) FindByID(id int) (interface{}, error) {
	entity := &model.City{}
	err := repository.db.First(entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
