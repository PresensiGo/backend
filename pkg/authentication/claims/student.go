package claims

import "github.com/golang-jwt/jwt/v5"

type Student struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	NIS      string `json:"nis" validate:"required"`
	SchoolId uint   `json:"school_id" validate:"required"`

	jwt.RegisteredClaims
}
