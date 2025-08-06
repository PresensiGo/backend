package handlers

import (
	"net/http"

	"api/internal/features/attendance/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type SubjectAttendance struct {
	service *services.SubjectAttendance
}

func NewSubjectAttendance(service *services.SubjectAttendance) *SubjectAttendance {
	return &SubjectAttendance{
		service: service,
	}
}

// @id 			getAllSubjectAttendances
// @tags 		attendance
// @success 	200 {object} responses.GetAllSubjectAttendances
// @router 		/api/v1/subject-attendances [get]
func (h *SubjectAttendance) GetAll(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	result, err := h.service.GetAll(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
