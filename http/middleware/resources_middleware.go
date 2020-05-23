package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/domain"
)

func ResourcesMiddleware(resources *domain.Resources) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(domain.AppResources, resources)
		ctx.Next()
	}
}
