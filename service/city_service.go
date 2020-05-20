package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/dao"
)

type CityService struct {
	BaseService
}

func NewCityService(daoContext *dao.Context) *CityService {
	return &CityService{
		BaseService: BaseService{
			repository: daoContext.CityRepository,
		},
	}
}

func (service CityService) HandleFindAll(context *gin.Context) {
	res, err := service.repository.FindAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, res)
}

func (service CityService) HandleFindByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		fmt.Println("id must be an integer")
		context.JSON(http.StatusBadRequest, "error")
	}

	res, err2 := service.repository.FindByID(id)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, err2)
		return
	}

	context.JSON(http.StatusOK, res)
}
