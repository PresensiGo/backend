package services

import (
	"net/http"

	"api/internal/features/school/dto/responses"
	schoolRepo "api/internal/features/school/repositories"
	"api/pkg/authentication"
	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
)

type School struct {
	schoolRepo *schoolRepo.School
}

func NewSchool(schoolRepo *schoolRepo.School) *School {
	return &School{
		schoolRepo: schoolRepo,
	}
}

func (s *School) GetSchool(c *gin.Context) (*responses.GetSchool, *failure.App) {
	user := authentication.GetAuthenticatedUser(c)
	if user.SchoolId == 0 {
		return nil, failure.NewApp(http.StatusForbidden, "Anda tidak memiliki akses!", nil)
	}

	if school, err := s.schoolRepo.Get(user.SchoolId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetSchool{
			School: *school,
		}, nil
	}
}
