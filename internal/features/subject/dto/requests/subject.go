package requests

type CreateSubject struct {
	Name string `json:"name" validate:"required"`
}

type UpdateSubject struct {
	Name string `json:"name" validate:"required"`
}
