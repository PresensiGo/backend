package requests

type CreateSubject struct {
	Name string `json:"name" validate:"required"`
}
