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
	this.db.AutoMigrate(&model.City{})
	return this
}

func (this CityRepository) cloneEntity() model.DaoModel {
	return &model.City{}
}
