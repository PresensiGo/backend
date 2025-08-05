package handlers

import (
	"net/http"
	"strconv"

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

// @id 			updateBatch
// @tags 		batch
// @param 		batch_id path int true "batch id"
// @param 		body body requests.Update true "body"
// @success 	200 {object} domains.Batch
// @router 		/api/v1/batches/{batch_id} [put]
func (h *Batch) Update(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.Update
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Update(uint(batchId), req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			deleteBatch
// @tags 		batch
// @param 		batch_id path int true "batch id"
// @success 	200 {string} string
// @router 		/api/v1/batches/{batch_id} [delete]
func (h *Batch) Delete(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(batchId)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
