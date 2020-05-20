package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
