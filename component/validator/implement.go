package validator

import (
	"github.com/go-playground/validator/v10"
	"shorturl/component/shortener"
)

type GoPlayGroundValidator struct {
	validate      *validator.Validate
	blacklistRepo BlacklistRepository
}

func New(blacklistRepo BlacklistRepository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate:      validator.New(),
		blacklistRepo: blacklistRepo,
	}

	v.validate.RegisterStructValidation(v.ShortenInputValidation, &shortener.ShortenInput{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
