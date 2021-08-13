package shortener

import "shorturl/entity"

//go:generate mockery --name=Comp
type Comp interface {
	ShortenURL(input *ShortenInput) (output *ShortenOutput, err error)
	AccessURL(input *AccessInput) (output *ShortenOutput, err error)
}

//go:generate mockery --name=URLRepo
type URLRepo interface {
	Create(ent *entity.URL) (err error)
	Read(ID string) (ent *entity.URL, err error)
	CountID(ID string) (count int, err error)
	IncrHit(ID string) (err error)
}

//go:generate mockery --name=URLCache
type URLCache interface {
	Read(ID string) (ent *entity.URL)
	Write(ent *entity.URL) (err error)
}

//go:generate mockery --name=Validator
type Validator interface {
	Validate(item interface{}) (err error)
}
