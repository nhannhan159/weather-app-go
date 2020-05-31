package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/nhannhan159/weather-app-go/common"
	"github.com/nhannhan159/weather-app-go/domain"
)

type baseRepository struct {
	db *gorm.DB
}

type DaoManager struct {
	db           *gorm.DB
	repositories []domain.IRepository
}

func NewDaoManager(dbConfig common.DatabaseConfig) *DaoManager {
	db, err := initializeDb(dbConfig)
	if err != nil {
		panic(err)
	}

	return &DaoManager{
		db:           db,
		repositories: []domain.IRepository{},
	}
}

func initializeDb(dbConfig common.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(dbConfig.Dialect, dbConfig.DSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dao *DaoManager) GetDB() *gorm.DB {
	return dao.db
}

func (dao *DaoManager) AddRepository(repository domain.IRepository) {
	dao.repositories = append(dao.repositories, repository)
}

func (dao *DaoManager) AutoMigrate() {
	for _, repository := range dao.repositories {
		repository.AutoMigrate()
	}
}

func (dao *DaoManager) TearDown() {
	err := dao.db.Close()
	if err != nil {
		panic(err)
	}
}
