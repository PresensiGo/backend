package dto

type User struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
	SchoolId uint   `json:"school_id" validate:"required"`
}
