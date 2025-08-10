package attendance

import (
	"api/internal/features/attendance/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterAttendance(g *gin.RouterGroup) {
	handlers := injector.InitAttendanceHandlers()

	routes.RegisterGeneralAttendance(g, handlers.GeneralAttendance)
	routes.RegisterSubjectAttendance(g, handlers.SubjectAttendance)
}
