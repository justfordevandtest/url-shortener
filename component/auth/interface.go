package auth

import "shorturl/entity"

//go:generate mockery --name=Comp
type Comp interface {
	ReadByCredential(input *CredentialInput) (output *UserOutput, err error)
}

//go:generate mockery --name=UserRepo
type UserRepo interface {
	Read(username string) (user *entity.User, err error)
}

//go:generate mockery --name=Validator
type Validator interface {
	Validate(item interface{}) (err error)
}
