package routes

import (
	"api/features/auth"
	"api/handler"
	"api/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	group := r.Group("/auth")

	authService := auth.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)

	group.POST("/login", authHandler.Login)
	group.POST("/register", authHandler.Register)
	group.POST("/refresh-token", authHandler.RefreshToken)

	group.Use(middleware.AuthMiddleware()).
		GET("/logout", authHandler.Logout)
}
