package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(entity interface{}) error
	Update(entity interface{}) error
	FindAll() ([]interface{}, error)
	FindByID(id int) (interface{}, error)
}

type BaseRepository struct {
	db        *gorm.DB
	newEntity func() interface{}
}

func (repository BaseRepository) cloneEntities() []interface{} {
	slice := make([]interface{}, 1)
	slice[0] = repository.newEntity()

	return slice
}

func (repository *BaseRepository) Create(entity interface{}) error {
	return nil
}

func (repository *BaseRepository) Update(entity interface{}) error {
	return nil
}

func (repository BaseRepository) FindAll() ([]interface{}, error) {
	entities := repository.cloneEntities()
	err := repository.db.Find(&entities).Error

	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (repository BaseRepository) FindByID(id int) (interface{}, error) {
	entity := repository.newEntity()
	err := repository.db.First(&entity).Error

	if err != nil {
		return nil, err
	}

	return entity, nil
}
