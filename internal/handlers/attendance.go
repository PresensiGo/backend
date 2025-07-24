package handlers

import (
	"api/internal/dto/requests"
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Attendance struct {
	service *services.Attendance
}

func NewAttendance(service *services.Attendance) *Attendance {
	return &Attendance{service}
}

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
