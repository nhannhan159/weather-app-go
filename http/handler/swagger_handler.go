package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nhannhan159/weather-app-go/common"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerHandler(config common.ServerConfig) gin.HandlerFunc {
	url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", config.Host, config.Port))
	return ginSwagger.WrapHandler(swaggerFiles.Handler, url)
}
