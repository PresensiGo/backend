package handlers

import (
	"api/internal/dto/requests"
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Attendance struct {
	service *services.Attendance
}

func NewAttendance(service *services.Attendance) *Attendance {
	return &Attendance{service}
}

// @Id 			createAttendance
// @Tags 		attendance
// @Param 		body body requests.CreateAttendance true "Body"
// @Success 	200 {string} string
// @Router		/api/v1/attendances [post]
func (h *Attendance) Create(c *gin.Context) {
	var request requests.CreateAttendance
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.service.Create(request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @Id 			getAllAttendances
// @Tags 		attendance
// @Param 		classroom_id path int true "Classroom Id"
// @Success 	200 {object} responses.GetAllAttendances
// @Router		/api/v1/attendances/classrooms/{classroom_id} [get]
func (h *Attendance) GetAll(c *gin.Context) {
	classId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAll(uint(classId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Id 			getAttendance
// @Tags 		attendance
// @Param 		attendance_id path int true "Attendance Id"
// @Success 	200 {object} responses.GetAttendance
// @Router		/api/v1/attendances/{attendance_id} [get]
func (h *Attendance) GetById(c *gin.Context) {
	attendanceId, err := strconv.Atoi(c.Param("attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetById(uint(attendanceId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Id 			deleteAttendance
// @Tags 		attendance
// @Param 		attendance_id path int true "Attendance Id"
// @Success 	200 {string} string
// @Router		/api/v1/attendances/{attendance_id} [delete]
func (h *Attendance) Delete(c *gin.Context) {
	attendanceID, err := strconv.Atoi(c.Param("attendance_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(attendanceID)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
