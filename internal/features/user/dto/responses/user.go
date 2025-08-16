package responses

import "api/internal/features/user/domains"

type ImportAccounts struct {
	Message string `json:"message" validate:"required"`
}

type GetAllUsers struct {
	Users []domains.User `json:"users" validate:"required"`
}

type UpdateAccountPassword struct {
	User domains.User `json:"user" validate:"required"`
}

type DeleteAccount struct {
	Message string `json:"message" validate:"required"`
}
