package responses

import "api/internal/features/user/domains"

type GetAllUsers struct {
	Users []domains.User `json:"users" validate:"required"`
}
