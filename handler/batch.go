package handler

import (
	"api/features/batch"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BatchHandler struct {
	service *batch.Service
}

func NewBatchHandler(service *batch.Service) *BatchHandler {
	return &BatchHandler{service}
}

func (h *BatchHandler) CreateBatch(c *gin.Context) {
	var request batch.CreateBatchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.Create(request.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
