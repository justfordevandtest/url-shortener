package admin

type ListInput struct {
	Page    int    `validate:"min=1"`
	PerPage int    `validate:"min=1"`
	ID      string
	Keyword string
}

type DelInput struct {
	ID string `validate:"required"`
}
