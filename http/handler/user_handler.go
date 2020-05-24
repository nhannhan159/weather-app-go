package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	baseHandler
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (handler *UserHandler) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user",
	})
}
