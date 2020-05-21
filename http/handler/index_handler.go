package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
)

type IIndexHandler interface {
	domain.IHandler
}

type indexHandler struct {
	baseHandler
}

func NewIndexHandler() IIndexHandler {
	return &indexHandler{}
}

func (handler *indexHandler) Handler(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
