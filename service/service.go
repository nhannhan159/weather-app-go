package service

import (
	"github.com/nhannhan159/weather-app-go/domain"
)

type baseService struct {
	config     domain.GlobalConfig
	repository domain.IRepository
}
