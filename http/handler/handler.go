package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
)

type baseHandler struct {
}

func (handler *baseHandler) Handle(router *gin.Engine) {
}

func (handler *baseHandler) HandleGroup(routerGroup *gin.RouterGroup) {
}

func (handler *baseHandler) getResourcesFromContext(ctx *gin.Context) *domain.Resources {
	res, exists := ctx.Get(domain.AppResources)
	if !exists {
		panic("cannot get resources from context")
	}
	return res.(*domain.Resources)
}
