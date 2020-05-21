package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
)

type IUserHandler interface {
	domain.IHandler
}

type userHandler struct {
	baseHandler
}

func NewUserHandler() IUserHandler {
	return &userHandler{}
}

func (handler *userHandler) HandleGroup(routerGroup *gin.RouterGroup) {
	userGroup := routerGroup.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user",
			})
		})
	}

}
