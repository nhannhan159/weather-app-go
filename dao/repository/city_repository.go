package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/model"
)

type CityRepository struct {
	BaseRepository
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	this := &CityRepository{
		BaseRepository: BaseRepository{
			db: db.Table("city"),
		},
	}
	this.newEntity = this.cloneEntity
	this.newEntities = this.cloneEntities
	this.db.AutoMigrate(&model.City{})

	return this
}

func (repository CityRepository) cloneEntity() interface{} {
	return &model.City{}
}

func (repository CityRepository) cloneEntities() interface{} {
	return []model.City{}
}
