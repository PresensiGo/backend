package handlers

import (
	"net/http"

	"api/internal/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Batch struct {
	service *services.Batch
}

func NewBatch(service *services.Batch) *Batch {
	return &Batch{service}
}

// @Id			getAllBatches
// @Tags		batch
// @Success		200	{object} responses.GetAllBatches
// @Router		/api/v1/batch [get]
func (h *Batch) GetAll(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllBySchoolId(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
