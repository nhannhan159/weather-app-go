package domain

import "github.com/gin-gonic/gin"

type IRootHandler interface {
	Handle(router *gin.Engine)
}

type IHandler interface {
	IRootHandler
	HandleGroup(routerGroup *gin.RouterGroup)
}

type IHttpManager interface {
	RegisterHandler(handlers IRootHandler)
	RegisterResources(resources *Resources)
	Run()
}
