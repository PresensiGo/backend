package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	service *services.Student
}

func NewStudent(service *services.Student) *Student {
	return &Student{service}
}

func (h *Student) GetAllStudents(c *gin.Context) {
	classId, err := strconv.ParseUint(c.Param("class_id"), 10, 8)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllStudents(classId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
