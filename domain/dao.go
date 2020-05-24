package domain

import (
	"github.com/jinzhu/gorm"
)

type IRepository interface {
	AutoMigrate()
}

type IDaoManager interface {
	GetDB() *gorm.DB
	AddRepository(repository IRepository)
	AutoMigrate()
	TearDown()
}
