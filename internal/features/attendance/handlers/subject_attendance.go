package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/attendance/dto/requests"
	"api/internal/features/attendance/services"
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

// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @param 		body body requests.CreateSubjectAttendance true "body"
// @success 	200 {object} responses.CreateSubjectAttendance
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances [post]
func (h *SubjectAttendance) Create(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateSubjectAttendance
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(uint(classroomId), req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @tags 		attendance
// @param 		batch_id path int true "batch id"
// @param 		major_id path int true "major id"
// @param 		classroom_id path int true "classroom id"
// @success 	200 {object} responses.GetAllSubjectAttendances
// @router 		/api/v1/batches/{batch_id}/majors/{major_id}/classrooms/{classroom_id}/subject-attendances [get]
func (h *SubjectAttendance) GetAll(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.GetAll(uint(classroomId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
