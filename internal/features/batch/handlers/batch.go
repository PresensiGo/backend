package handlers

import (
	"net/http"

	"api/internal/features/batch/dto/requests"
	"api/internal/features/batch/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Batch struct {
	service *services.Batch
}

func NewBatch(service *services.Batch) *Batch {
	return &Batch{service}
}

// @id 			createBatch
// @tags 		batch
// @param 		body body requests.Create true "body"
// @success 	200 {object} domains.Batch
// @router 		/api/v1/batches [post]
func (h *Batch) Create(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.Create
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.Create(authUser.SchoolId, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Id			getAllBatches
// @Tags		batch
// @Success		200	{object} responses.GetAllBatches
// @Router		/api/v1/batches [get]
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
