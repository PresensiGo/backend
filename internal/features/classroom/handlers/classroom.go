package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/classroom/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Classroom struct {
	service *services.Classroom
}

func NewClassroom(service *services.Classroom) *Classroom {
	return &Classroom{service}
}

// @id 			getAllClassrooms
// @tags 		classroom
// @success 	200 {object} responses.GetAll
// @router 		/api/v1/classrooms [get]
func (h *Classroom) GetAll(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.GetAll(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Id			getAllClassroomWithMajors
// @Tags		classroom
// @Param 		batch_id path int true "Batch Id"
// @Success		200	{object}	responses.GetAllClassroomWithMajors
// @Router		/api/v1/classrooms/batches/{batch_id} [get]
func (h *Classroom) GetAllWithMajors(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllWithMajor(uint(batchId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
