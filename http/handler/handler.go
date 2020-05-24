package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
)

type baseHandler struct {
}

type baseRestHandler struct {
	baseHandler
}

func (handler *baseRestHandler) HandleFindAll(ctx *gin.Context) {
	panic("not yet implemented")
}

func (handler *baseRestHandler) HandleFindByID(ctx *gin.Context) {
	panic("not yet implemented")
}

func (handler *baseRestHandler) HandleCreate(ctx *gin.Context) {
	panic("not yet implemented")
}

func (handler *baseRestHandler) HandleUpdate(ctx *gin.Context) {
	panic("not yet implemented")
}

func (handler *baseRestHandler) HandleDelete(ctx *gin.Context) {
	panic("not yet implemented")
}

func (handler *baseHandler) getResourcesFromContext(ctx *gin.Context) *domain.Resources {
	res, exists := ctx.Get(domain.AppResources)
	if !exists {
		panic("cannot get resources from context")
	}

	return res.(*domain.Resources)
}

func (handler *baseHandler) handleBabRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
}
