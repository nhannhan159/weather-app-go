package router

import "github.com/gin-gonic/gin"

func userHandler(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "user",
			})
		})
	}

}
