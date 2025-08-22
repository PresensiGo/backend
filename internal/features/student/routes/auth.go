package routes

import (
	"api/internal/features/student/handlers"
	"api/pkg/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterStudentAuth(g *gin.RouterGroup, handler *handlers.StudentAuth) {
	relativePath := "/auth/students"

	{
		group := g.Group(relativePath)

		group.POST("/login", handler.Login)
		group.POST("/refresh-token", handler.RefreshTokenStudent)
	}

	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())

		group.POST("/accounts/:student_token_id/eject", handler.Eject)
	}
}
