package handler

import (
	"api/features/auth"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}

// Login godoc
//
//	@Accept		json
//	@Produce	json
//	@Tags		auth
//	@Param		body	body		auth.LoginRequest	true	"Login request"
//	@Success	200		{object}	auth.LoginResponse
//	@Router		/api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
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
//	@Tags		auth
//	@Param		body	body		auth.RegisterRequest	true	"Login request"
//	@Success	200		{object}	auth.RegisterResponse
//	@Router		/api/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var request auth.RegisterRequest
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
	var request auth.RefreshTokenRequest
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
