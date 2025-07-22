package handlers

import (
	"api/internal/dto/requests"
	"api/internal/services"
	"api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

// Login godoc
//
//	@Id			Login
//	@Accept		json
//	@Produce	json
//	@Tags		auth
//	@Param		body	body		requests.LoginRequest	true	"Login request"
//	@Success	200		{object}	responses.LoginResponse
//	@Router		/api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var request requests.LoginRequest
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

// Register godoc
//
//	@Id			Register
//	@Tags		auth
//	@Param		body	body		requests.RegisterRequest	true	"Login request"
//	@Success	200		{object}	responses.RegisterResponse
//	@Router		/api/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var request requests.RegisterRequest
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

func (h *AuthHandler) Logout(c *gin.Context) {
	userData := utils.GetAuthenticatedUser(c)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully", "user": userData})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var request requests.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := h.service.RefreshToken(request.AccessToken)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}
