package middlewares

import (
	"github.com/gin-gonic/gin"
)

func StudentMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
