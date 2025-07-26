package routes

import (
	"api/internal/injectors"
	"api/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuth(r *gin.RouterGroup) {
	group := r.Group("/auth")
	handler := injectors.InitAuthHandler()

	group.POST("/login", handler.Login)
	group.POST("/register", handler.Register)
	group.POST("/refresh-token", handler.RefreshToken)

	authorized := group.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	authorized.POST("/logout", handler.Logout)
	authorized.POST("/refresh-token-ttl", handler.RefreshTokenTTL)
}
