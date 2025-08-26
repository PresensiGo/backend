package routes

import (
	"api/internal/features/student/handlers"
	"api/pkg/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterStudent(g *gin.RouterGroup, handler *handlers.Student) {
	{
		relativePath := "/batches/:batch_id/majors/:major_id/classrooms/:classroom_id"
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())

		group.GET("/students", handler.GetAllStudentsByClassroomId)
		group.GET("/student-accounts", handler.GetAllAccountsByClassroomId)
		group.PUT("/students/:student_id", handler.UpdateStudent)
		group.DELETE("/students/:student_id", handler.DeleteStudent)
	}

	{
		relativePath := "/students"
		group := g.Group(relativePath).Use(middlewares.StudentMiddleware())

		group.GET("/profile", handler.GetProfileStudent)
	}
}
