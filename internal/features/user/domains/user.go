package domains

import "api/internal/features/user/models"

type UserRole string

var (
	AdminUserRole   UserRole = "admin"
	TeacherUserRole UserRole = "teacher"
)

type User struct {
	Id       uint     `json:"id" validate:"required"`
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required"`
	Password string   `json:"password" validate:"required"`
	Role     UserRole `json:"role" validate:"required"`
	SchoolId uint     `json:"school_id" validate:"required"`
} // @name User

func FromUserModel(m *models.User) *User {
	return &User{
		Id:       m.ID,
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
		Role:     UserRole(m.Role),
		SchoolId: m.SchoolId,
	}
}

func (u *User) ToModel() *models.User {
	return &models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     string(u.Role),
		SchoolId: u.SchoolId,
	}
}
