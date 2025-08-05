package domains

type School struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}
