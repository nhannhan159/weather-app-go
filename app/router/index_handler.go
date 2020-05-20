package router

import "github.com/gin-gonic/gin"

func indexHandler(conext *gin.Context) {
	conext.JSON(200, gin.H{
		"message": "pong",
	})
}
