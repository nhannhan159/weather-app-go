package domain

import (
	"github.com/jinzhu/gorm"
)

type IRepository interface {
	AutoMigrate()
	Create(entity interface{}) error
	Update(entity interface{}) error
	Delete(id int) error
	FindAll() (interface{}, error)
	FindByID(id int) (interface{}, error)
}

type IDaoManager interface {
	GetDB() *gorm.DB
	AddRepository(repository IRepository)
	AutoMigrate()
	TearDown()
}
