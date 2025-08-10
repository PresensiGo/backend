package routes

import (
	"api/internal/features/student/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterStudentAuth(g *gin.RouterGroup, handler *handlers.StudentAuth) {
	group := g.Group("/auth/students")

	group.POST("/login", handler.Login)
	group.POST("/refresh-token", handler.RefreshToken)
}
