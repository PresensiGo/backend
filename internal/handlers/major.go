package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Major struct {
	service *services.Major
}

func NewMajor(service *services.Major) *Major {
	return &Major{service}
}

// @ID			getAllMajors
// @Tags		major
// @Success	200	{object}	responses.GetAllMajors
// @Router		/api/v1/major [get]
func (h *Major) GetAll(c *gin.Context) {
	batchId, err := strconv.ParseUint(c.Param("batch_id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllMajors(batchId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
