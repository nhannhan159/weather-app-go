package domain

import "github.com/gin-gonic/gin"

type IHttpManager interface {
	RegisterHandler(*HandlerCollection)
	RegisterResources(*Resources)
	Run()
}

type IRestHandler interface {
	HandleFindAll(ctx *gin.Context)
	HandleFindByID(ctx *gin.Context)
	HandleCreate(ctx *gin.Context)
	HandleUpdate(ctx *gin.Context)
	HandleDelete(ctx *gin.Context)
}

type ICustomHandler interface {
	Handle(ctx *gin.Context)
}

type HandlerCollection struct {
	IndexHandler   ICustomHandler
	UserHandler    ICustomHandler
	CityHandler    IRestHandler
	WeatherHandler IRestHandler
}
