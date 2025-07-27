package handlers

import (
	"api/internal/services"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Reset struct {
	service *services.Reset
}

func NewReset(service *services.Reset) *Reset {
	return &Reset{service}
}

// @Id			reset
// @Tags		reset
// @Success		200	{string} string
// @Router		/api/v1/reset [get]
func (h *Reset) Reset(c *gin.Context) {
	authUser := authentication.GetAuthenticatedUser(c)
	if authUser.SchoolId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.ResetBySchoolId(uint(authUser.SchoolId)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
