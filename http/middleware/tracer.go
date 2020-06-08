package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func TracerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx, ctx.HandlerName())
		defer func() {
			go span.Finish()
		}()

		ctx.Next()
	}
}
