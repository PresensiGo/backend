package attendance

import (
	"api/internal/features/attendance/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func RegisterAttendance(g *gin.RouterGroup) {
	handlers := injector.InitAttendanceHandlers()

	routes.RegisterAttendance(g, handlers.Attendance)
	routes.RegisterLateness(g, handlers.Lateness)

	routes.RegisterGeneralAttendance(g, handlers.GeneralAttendance)
}
