package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
	baseHandler
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (handler *IndexHandler) Handle(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
