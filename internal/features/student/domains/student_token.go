package domains

import (
	"time"

	"api/internal/features/student/models"
)

type StudentToken struct {
	Id           uint      `json:"id" validate:"required"`
	StudentId    uint      `json:"student_id" validate:"required"`
	DeviceId     string    `json:"device_id" validate:"required"`
	RefreshToken string    `json:"refresh_token" validate:"required"`
	TTL          time.Time `json:"ttl" validate:"required"`
}

func FromStudentTokenModel(m *models.StudentToken) *StudentToken {
	return &StudentToken{
		Id:           m.ID,
		StudentId:    m.StudentId,
		DeviceId:     m.DeviceId,
		RefreshToken: m.RefreshToken,
		TTL:          m.TTL,
	}
}

func (s *StudentToken) ToModel() *models.StudentToken {
	return &models.StudentToken{
		StudentId:    s.StudentId,
		DeviceId:     s.DeviceId,
		RefreshToken: s.RefreshToken,
		TTL:          s.TTL,
	}
}
