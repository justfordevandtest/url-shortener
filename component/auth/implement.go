package auth

import (
	"shorturl/entity"
)

type impl struct {
	userRepo   UserRepo
	validator Validator
}

func New(userRepo UserRepo, v Validator) (comp Comp) {
	return &impl{
		userRepo:   userRepo,
		validator: v,
	}
}

func (i *impl) ReadByCredential(input *CredentialInput) (output *UserOutput, err error) {
	err = i.validator.Validate(input)
	if err != nil {
		return nil, entity.ValidatorListErr(err)
	}

	user, err := i.userRepo.Read(input.Username)
	if err != nil {
		return nil, entity.ReadRecordErr(err)
	}

	if !user.CheckPasswordHash(input.Password) {
		return nil, entity.UnauthorizedErr(nil)
	}

	return &UserOutput{Username: user.Username}, nil
}
