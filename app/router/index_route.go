package router

import "github.com/gin-gonic/gin"

func index(conext *gin.Context) {
	conext.JSON(200, gin.H{
		"message": "pong",
	})
}
