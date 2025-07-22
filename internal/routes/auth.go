package routes

import (
	"api/internal/injectors"
	"api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	group := r.Group("/auth")
	handler := injectors.InitAuthHandler()

	group.POST("/login", handler.Login)
	group.POST("/register", handler.Register)
	group.POST("/refresh-token", handler.RefreshToken)

	group.Use(middleware.AuthMiddleware()).
		GET("/logout", handler.Logout)
}
