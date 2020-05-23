package handler

import (
	"net/http"

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

func (handler *baseHandler) handleBabRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
}
