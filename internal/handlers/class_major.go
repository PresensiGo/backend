package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ClassMajor struct {
	service *services.ClassMajor
}

func NewClassMajor(service *services.ClassMajor) *ClassMajor {
	return &ClassMajor{service}
}

// @ID			getAllClassMajors
// @Tags		classMajor
// @Param		batch_id	path		int	true	"Batch ID"
// @Success		200			{object}	responses.GetAllClassMajors
// @Router		/api/v1/class_majors/batch/{batch_id} [get]
func (h *ClassMajor) GetAll(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAll(uint(batchId))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
