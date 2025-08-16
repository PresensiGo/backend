package requests

type UpdateAccountPassword struct {
	Password string `json:"password" validate:"required"`
}

type UpdateAccountRole struct {
	Role string `json:"role" validate:"required"`
}
