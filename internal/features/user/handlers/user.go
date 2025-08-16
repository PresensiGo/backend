package handlers

import (
	"net/http"
	"path/filepath"

	"api/internal/features/user/services"
	"api/internal/shared/dto/responses"
	"api/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type User struct {
	service *services.User
}

func NewUser(service *services.User) *User {
	return &User{
		service: service,
	}
}

// @tags 		account
// @success 	200 {object} responses.GetAllUsers
// @router 		/api/v1/accounts [get]
func (h *User) GetAll(c *gin.Context) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	result, err := h.service.GetAll(user.SchoolId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @accept 		multipart/form-data
// @param 		file formData file true "file"
// @success 	200 {object} responses.ImportAccounts
// @router 		/api/v1/accounts/import [post]
func (h *User) ImportAccounts(c *gin.Context) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".xlsx" && ext != ".xls" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	src, err := file.Open()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer src.Close()

	if response, err := h.service.ImportAccounts(user.SchoolId, src); err != nil {
		c.AbortWithStatusJSON(
			err.Code, responses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
