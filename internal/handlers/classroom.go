package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Classroom struct {
	service *services.Classroom
}

func NewClassroom(service *services.Classroom) *Classroom {
	return &Classroom{service}
}

// @ID			getAllClassrooms
// @Tags		class
// @Success	200	{object}	responses.GetAllClassrooms
// @Router		/api/v1/class [get]
func (h *Classroom) GetAll(c *gin.Context) {
	majorId, err := strconv.ParseUint(c.Param("major_id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllClassrooms(majorId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
