package handlers

import (
	"net/http"
	"strconv"

	"api/internal/dto"
	"api/internal/dto/responses"
	"api/internal/services"
	"github.com/gin-gonic/gin"
)

type Student struct {
	service *services.Student
}

func NewStudent(service *services.Student) *Student {
	return &Student{service}
}

// @Id			getAllStudentsByClassroomId
// @Tags		student
// @Param		classroom_id path int true "Classroom Id"
// @Success		200	{object} responses.GetAllStudentsByClassroomId
// @Router		/api/v1/students/classrooms/{classroom_id} [get]
func (h *Student) GetAllByClassroomId(c *gin.Context) {
	classroomId, err := strconv.Atoi(c.Param("classroom_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllStudentsByClassroomId(uint(classroomId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @id getAllStudents
// @tags student
// @param keyword query string true "Keyword"
// @success 200 {object} responses.GetAllStudents
// @router /api/v1/students [get]
func (h *Student) GetAll(c *gin.Context) {
	keyword := c.Query("keyword")
	if len(keyword) == 0 {
		c.JSON(
			http.StatusOK, responses.GetAllStudents{
				Students: make([]dto.Student, 0),
			},
		)
		return
	}

	response, err := h.service.GetAll(keyword)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
