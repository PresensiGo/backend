package responses

import "api/internal/features/user/domains"

type ImportAccounts struct {
	Message string `json:"message" validate:"required"`
}

type GetAllUsers struct {
	Users []domains.User `json:"users" validate:"required"`
}

type GetAccount struct {
	User domains.User `json:"user" validate:"required"`
} // @name GetAccountRes

type UpdateAccountPassword struct {
	User domains.User `json:"user" validate:"required"`
}

type UpdateAccountRole struct {
	User domains.User `json:"user" validate:"required"`
}

type DeleteAccount struct {
	Message string `json:"message" validate:"required"`
}
