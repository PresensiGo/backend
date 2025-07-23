package handlers

import (
	"api/internal/dto/requests"
	"api/internal/services"
	"api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	service *services.Auth
}

func NewAuth(service *services.Auth) *Auth {
	return &Auth{service}
}

// @ID			Login
// @Accept		json
// @Produce	json
// @Tags		auth
// @Param		body	body		requests.Login	true	"Login request"
// @Success	200		{object}	responses.Login
// @Router		/api/v1/auth/login [post]
func (h *Auth) Login(c *gin.Context) {
	var request requests.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Login(
		request.Email,
		request.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @ID			Register
// @Tags		auth
// @Param		body	body		requests.Register	true	"Login request"
// @Success	200		{object}	responses.Register
// @Router		/api/v1/auth/register [post]
func (h *Auth) Register(c *gin.Context) {
	var request requests.Register
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Register(
		request.Name,
		request.Email,
		request.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @ID			Logout
// @Tags		auth
// @Success	200	{object}	responses.Logout
// @Router		/api/v1/auth/logout [get]
func (h *Auth) Logout(c *gin.Context) {
	authUser := utils.GetAuthenticatedUser(c)

	response, err := h.service.Logout(authUser.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Auth) RefreshToken(c *gin.Context) {
	var request requests.RefreshToken
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.RefreshToken(request.RefreshToken)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
