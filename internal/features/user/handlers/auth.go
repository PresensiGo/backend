package handlers

import (
	"net/http"

	"api/internal/features/user/dto/requests"
	"api/internal/features/user/services"
	"api/internal/shared/dto/responses"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	service *services.Auth
}

func NewAuth(service *services.Auth) *Auth {
	return &Auth{service}
}

// @id			login
// @tags		account
// @param		body body requests.Login true "body"
// @success		200 {object} responses.Login
// @router		/api/v1/auth/login [post]
func (h *Auth) Login(c *gin.Context) {
	var request requests.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if response, err := h.service.Login(
		request.Email,
		request.Password,
	); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id			login2
// @tags		account
// @param		body body requests.Login2 true "body"
// @success		200 {object} responses.Login2
// @router		/api/v1/auth/login-2 [post]
func (h *Auth) Login2(c *gin.Context) {
	var req requests.Login2
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if response, err := h.service.Login2(req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id			logout
// @tags		account
// @param		body body requests.Logout true "Logout Request"
// @success		200	{object} responses.Logout
// @router		/api/v1/auth/logout [post]
func (h *Auth) Logout(c *gin.Context) {
	var req requests.Logout
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if response, err := h.service.Logout(req.RefreshToken); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id			logout2
// @tags		account
// @param		body body requests.Logout2 true "Logout Request"
// @success		200	{object} responses.Logout2
// @router		/api/v1/auth/logout-2 [post]
func (h *Auth) Logout2(c *gin.Context) {
	var req requests.Logout2
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if response, err := h.service.Logout2(req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// @id			refreshToken
// @tags		account
// @param		body body requests.RefreshToken true "body"
// @success		200 {object} responses.RefreshToken
// @router		/api/v1/auth/refresh-token [post]
func (h *Auth) RefreshToken(c *gin.Context) {
	var request requests.RefreshToken
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(request.RefreshToken) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if response, err := h.service.RefreshToken(request.RefreshToken); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
