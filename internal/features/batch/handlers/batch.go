package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/batch/dto/requests"
	"api/internal/features/batch/services"
	"api/internal/shared/dto/responses"
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
// @param 		body body requests.CreateBatch true "body"
// @success 	200 {object} domains.Batch
// @router 		/api/v1/batches [post]
func (h *Batch) Create(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateBatch
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

// @id			getAllBatches
// @tags		batch
// @success		200	{object} responses.GetAllBatches
// @router		/api/v1/batches [get]
func (h *Batch) GetAllBatches(c *gin.Context) {
	if response, err := h.service.GetAllBatches(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			getBatch
// @tags		batch
// @param		batch_id path int true "batch id"
// @success		200	{object} responses.GetBatch
// @router		/api/v1/batches/{batch_id} [get]
func (h *Batch) GetBatch(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetBatch(uint(batchId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			updateBatch
// @tags 		batch
// @param 		batch_id path int true "batch id"
// @param 		body body requests.UpdateBatch true "body"
// @success 	200 {object} domains.Batch
// @router 		/api/v1/batches/{batch_id} [put]
func (h *Batch) Update(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateBatch
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
