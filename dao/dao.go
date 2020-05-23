package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/domain"
)

type daoManager struct {
	db           *gorm.DB
	repositories []domain.IRepository
}

func NewDaoManager(dbConfig domain.DatabaseConfig) domain.IDaoManager {
	db, err := initializeDb(dbConfig)
	if err != nil {
		panic(err)
	}

	return &daoManager{
		db:           db,
		repositories: []domain.IRepository{},
	}
}

func initializeDb(dbConfig domain.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(dbConfig.Dialect, dbConfig.DSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dao daoManager) GetDB() *gorm.DB {
	return dao.db
}

func (dao daoManager) AddRepository(repository domain.IRepository) {
	dao.repositories = append(dao.repositories, repository)
}

func (dao *daoManager) AutoMigrate() {
	for _, repository := range dao.repositories {
		repository.AutoMigrate()
	}
}

func (dao *daoManager) TearDown() {
	err := dao.db.Close()
	if err != nil {
		panic(err)
	}
}
