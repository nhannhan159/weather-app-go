package middleware

import (
	"github.com/gin-contrib/opengintracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func TracerMiddleware() gin.HandlersChain {
	tracer := opentracing.GlobalTracer()
	return []gin.HandlerFunc{
		opengintracing.NewSpan(tracer, "start http request"),
		opengintracing.InjectToHeaders(tracer, true),
	}
}
