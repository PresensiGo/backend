package handlers

import (
	"net/http"
	"strconv"

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
