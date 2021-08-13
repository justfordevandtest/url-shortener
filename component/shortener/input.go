package shortener

type ShortenInput struct {
	URL     string `validate:"required,url"`
	Expired *int64
}

type AccessInput struct {
	ID string
}
