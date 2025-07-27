package handlers

import (
	"api/internal/dto/requests"
	"api/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	service *services.Auth
}

func NewAuth(service *services.Auth) *Auth {
	return &Auth{service}
}

// @ID			login
// @Tags		auth
// @Param		body	body		requests.Login	true	"Login request"
// @Success		200		{object}	responses.Login
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

// @ID			register
// @Tags		auth
// @Param		body body		requests.Register	true	"Login request"
// @Success		200	{object}	responses.Register
// @Router		/api/v1/auth/register [post]
func (h *Auth) Register(c *gin.Context) {
	var request requests.Register
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Register(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @ID			logout
// @Tags		auth
// @Param		body body requests.Logout true "Logout Request"
// @Success		200	{object} responses.Logout
// @Router		/api/v1/auth/logout [post]
func (h *Auth) Logout(c *gin.Context) {
	var req requests.Logout
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Logout(req.RefreshToken)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @ID			refreshToken
// @Tags		auth
// @Param		body	body		requests.RefreshToken	true	"Refresh token req"
// @Success		200		{object}	responses.RefreshToken
// @Router		/api/v1/auth/refresh-token [post]
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

// @ID			refreshTokenTTL
// @Tags		auth
// @Param		body	body		requests.RefreshTokenTTL	true	"Refresh token req"
// @Success		200		{string}	string
// @Router		/api/v1/auth/refresh-token-ttl [post]
func (h *Auth) RefreshTokenTTL(c *gin.Context) {
	var req requests.RefreshTokenTTL
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.service.RefreshTokenTTL(req.RefreshToken); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
