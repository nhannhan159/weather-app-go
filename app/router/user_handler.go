package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userHandler(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user",
			})
		})
	}

}
