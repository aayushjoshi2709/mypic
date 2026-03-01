package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "RequestID"

func generateRequestID() string {
	return fmt.Sprintf("mypic-%s", uuid.New().String())
}

func RequestIdMiddleware(ctx *gin.Context) {
	requestId := ctx.GetHeader("X-Request-ID")
	if requestId == "" {
		requestId = generateRequestID()
	}
	ctx.Set(RequestIDKey, requestId)
	ctx.Writer.Header().Set("X-Request-ID", requestId)
	ctx.Next()
}
