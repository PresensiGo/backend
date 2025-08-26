package requests

type UpdateStudent struct {
	NIS    string `json:"nis" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
}
