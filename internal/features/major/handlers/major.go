package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/major/services"
	"github.com/gin-gonic/gin"
)

type Major struct {
	service *services.Major
}

func NewMajor(service *services.Major) *Major {
	return &Major{service}
}

// @Id			getAllMajors
// @Tags		major
// @Param		batch_id	path		int	true	"Batch Id"
// @Success	200			{object}	responses.GetAllMajors
// @Router		/api/v1/majors/batch/{batch_id} [get]
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
