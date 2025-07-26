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

// @ID 			createAttendance
// @Tags 		attendance
// @Param 		body body requests.CreateAttendance true "Body"
// @Success 	200 {string} string
// @Router		/api/v1/attendances/ [post]
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

// @ID 			getAllAttendances
// @Tags 		attendance
// @Param 		classroom_id path int true "Classroom ID"
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
