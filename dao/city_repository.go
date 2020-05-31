package dao

import (
	"github.com/nhannhan159/weather-app-go/common"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/model"
)

type CityRepository struct {
	baseRepository
}

func NewCityRepository(daoManager domain.IDaoManager) *CityRepository {
	this := &CityRepository{
		baseRepository: baseRepository{
			db: daoManager.GetDB().Table(common.TableCity),
		},
	}
	return this
}

func (repository *CityRepository) AutoMigrate() {
	repository.db.AutoMigrate(&model.City{})
}

func (repository *CityRepository) Create(entity *model.City) (*model.City, error) {
	panic("implement me")
}

func (repository *CityRepository) Update(entity *model.City) (*model.City, error) {
	panic("implement me")
}

func (repository *CityRepository) Delete(id int) error {
	panic("implement me")
}

func (repository CityRepository) FindAll() ([]model.City, error) {
	var entities []model.City
	err := repository.db.Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (repository CityRepository) FindByID(id int) (*model.City, error) {
	entity := &model.City{}
	err := repository.db.First(entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
