package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/major/dto/requests"
	"api/internal/features/major/services"
	"api/internal/shared/dto/responses"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Major struct {
	service *services.Major
}

func NewMajor(service *services.Major) *Major {
	return &Major{service}
}

// @id 			createMajor
// @tags 		major
// @param		batch_id path int true "batch id"
// @param		body body requests.CreateMajor true "body"
// @success 	200 {object} domains.Major
// @router 		/api/v1/batches/{batch_id}/majors [post]
func (h *Major) Create(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.CreateMajor
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(uint(batchId), req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			getAllMajors
// @tags 		major
// @success 	200 {array} domains.Major
// @router 		/api/v1/majors [get]
func (h *Major) GetAllMajors(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.GetAllMajors(authUser.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id			getAllMajorsByBatchId
// @tags		major
// @param		batch_id path int true "batch id"
// @success		200 {object} responses.GetAllMajorsByBatchId
// @router		/api/v1/batches/{batch_id}/majors [get]
func (h *Major) GetAllMajorsByBatchId(c *gin.Context) {
	batchId, err := strconv.Atoi(c.Param("batch_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.GetAllMajorsByBatchId(uint(batchId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id 			updateMajor
// @tags 		major
// @param		batch_id path int true "batch id"
// @param		major_id path int true "major id"
// @param		body body requests.UpdateMajor true "body"
// @success 	200 {object} domains.Major
// @router 		/api/v1/batches/{batch_id}/majors/{major_id} [put]
func (h *Major) Update(c *gin.Context) {
	majorId, err := strconv.Atoi(c.Param("major_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.UpdateMajor
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Update(uint(majorId), req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @id 			deleteMajor
// @tags 		major
// @param		batch_id path int true "batch id"
// @param		major_id path int true "major id"
// @success 	200 {string} string
// @router 		/api/v1/batches/{batch_id}/majors/{major_id} [delete]
func (h *Major) Delete(c *gin.Context) {
	majorId, err := strconv.Atoi(c.Param("major_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(majorId)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
