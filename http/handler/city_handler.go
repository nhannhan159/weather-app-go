package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
	"github.com/nhannhan159/weather-app-go/service"
)

type ICityHandler interface {
	domain.IHandler
}

type cityHandler struct {
	baseHandler
	cityService service.ICityService
}

func NewCityHandler(cityService service.ICityService) ICityHandler {
	return &cityHandler{
		cityService: cityService,
	}
}

func (handler *cityHandler) HandleGroup(group *gin.RouterGroup) {
	cityGroup := group.Group("/city")
	{
		cityGroup.GET("/", handler.handleFindAll)
		cityGroup.GET("/:id", handler.handleFindByID)
	}
}

func (handler *cityHandler) handleFindAll(ctx *gin.Context) {
	resources := handler.getResourcesFromContext(ctx)
	resources.Logger.Info("test logger")
	res, err := handler.cityService.FindAll(handler.getResourcesFromContext(ctx))
	if err != nil {
		handler.handleBabRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (handler *cityHandler) handleFindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handler.handleBabRequest(ctx, errors.New("id must be an integer"))
		return
	}

	res, err2 := handler.cityService.FindByID(handler.getResourcesFromContext(ctx), id)
	if err2 != nil {
		handler.handleBabRequest(ctx, err2)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
