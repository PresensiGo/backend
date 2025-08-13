package requests

type CreateMajor struct {
	Name string `json:"name" validate:"required"`
}

type UpdateMajor struct {
	Name string `json:"name" validate:"required"`
}
