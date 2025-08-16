package requests

type UpdateAccountPassword struct {
	Password string `json:"password" validate:"required"`
}
