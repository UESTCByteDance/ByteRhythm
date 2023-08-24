package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/metadata"
)

func Jaeger(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var md = make(metadata.Metadata, 1)
		opName := ctx.Request.URL.Path + "-" + ctx.Request.Method
		parentSpan := tracer.StartSpan(opName)
		defer parentSpan.Finish()
		injectErr := tracer.Inject(parentSpan.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
		if injectErr != nil {
			log.Fatalf("%s: Couldn't inject metadata", injectErr)
		}
		newCtx := metadata.NewContext(ctx.Request.Context(), md)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
