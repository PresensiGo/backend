package handlers

import (
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BatchHandler struct {
	service *services.BatchService
}

func NewBatchHandler(service *services.BatchService) *BatchHandler {
	return &BatchHandler{service}
}

// @Id			getAllBatches
// @Tags		batch
// @Success	200	{object}	responses.GetAllBatchesResponse
// @Router		/api/v1/batch [get]
func (h *BatchHandler) GetAll(c *gin.Context) {
	response, err := h.service.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
