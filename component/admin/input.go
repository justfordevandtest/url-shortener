package admin

type ListInput struct {
	Page    int `validate:"min=1"`
	PerPage int `validate:"min=1"`
	Filters map[string]interface{}
}

type DelInput struct {
	ID string `validate:"required"`
}
