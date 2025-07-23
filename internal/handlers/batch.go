package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Batch struct {
	service *services.Batch
}

func NewBatch(service *services.Batch) *Batch {
	return &Batch{service}
}

// @ID			getAllBatches
// @Tags		batch
// @Success	200	{object}	responses.GetAllBatchesResponse
// @Router		/api/v1/batch [get]
func (h *Batch) GetAll(c *gin.Context) {
	response, err := h.service.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
