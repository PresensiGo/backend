package routes

import (
	"api/internal/features/user/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAuth(r *gin.RouterGroup, handler *handlers.Auth) {
	group := r.Group("/auth")

	group.POST("/login", handler.Login)
	group.POST("/register", handler.Register)
	group.POST("/refresh-token", handler.RefreshToken)

	authorized := group.Group("/")
	authorized.Use(middlewares.AuthMiddleware())

	authorized.POST("/logout", handler.Logout)
	authorized.POST("/refresh-token-ttl", handler.RefreshTokenTTL)
}
