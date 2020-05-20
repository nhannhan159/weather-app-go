package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/model"
)

type Repository interface {
	Create(entity model.DaoModel) error
	Update(entity model.DaoModel) error
	FindAll() ([]model.DaoModel, error)
	FindById(id int) (model.DaoModel, error)
}

type BaseRepository struct {
	db *gorm.DB
}

func (this BaseRepository) cloneEntity() model.DaoModel {
	return nil
}

func (this BaseRepository) cloneEntities() []model.DaoModel {
	return nil
}

func (this *BaseRepository) Create(entity model.DaoModel) error {
	return nil
}

func (this *BaseRepository) Update(entity model.DaoModel) error {
	return nil
}

func (this BaseRepository) FindAll() ([]model.DaoModel, error) {
	entities := this.cloneEntities()
	err := this.db.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (this BaseRepository) FindById(id int) (model.DaoModel, error) {
	entity := this.cloneEntity()
	err := this.db.First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
