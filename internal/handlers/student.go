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

// @ID			getAllStudents
// @Tags		student
// @Param		classroom_id path int true "Classroom ID"
// @Success		200	{object}	responses.GetAllStudents
// @Router		/api/v1/students/classrooms/{classroom_id} [get]
func (h *Student) GetAll(c *gin.Context) {
	classroomID, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllStudents(uint(classroomID))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
