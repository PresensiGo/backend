package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/major/dto/requests"
	"api/internal/features/major/services"
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
// @param		body body requests.Create true "body"
// @success 	200 {object} domains.Major
// @router 		/api/v1/majors [post]
func (h *Major) Create(c *gin.Context) {
	var req requests.Create
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(req)
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

// @Id			getAllMajorsByBatchId
// @Tags		major
// @Param		batch_id	path		int	true	"Batch Id"
// @Success		200			{object}	responses.GetAllMajorsByBatchId
// @Router		/api/v1/majors/batch/{batch_id} [get]
func (h *Major) GetAllByBatchId(c *gin.Context) {
	batchId, err := strconv.ParseUint(c.Param("batch_id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllMajorsByBatchId(batchId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @id 			updateMajor
// @tags 		major
// @param		major_id path int true "major id"
// @param		body body requests.Update true "body"
// @success 	200 {object} domains.Major
// @router 		/api/v1/majors/{major_id} [put]
func (h *Major) Update(c *gin.Context) {
	majorId, err := strconv.Atoi(c.Param("major_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var req requests.Update
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
