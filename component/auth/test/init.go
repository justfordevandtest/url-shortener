package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"shorturl/component/auth"
	"shorturl/component/auth/mocks"
	"shorturl/config"
	"shorturl/entity"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	config    *config.Config
	comp      auth.Comp
	userRepo  *mocks.UserRepo
	validator *mocks.Validator
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.config = config.Get()
	suite.userRepo = &mocks.UserRepo{}
	suite.validator = &mocks.Validator{}
	suite.comp = auth.New(suite.userRepo, suite.validator)
}

var (
	givenCredentialInput = &auth.CredentialInput{
		Username: "foo",
		Password: "bar",
	}
	givenIncorrectPasswordInput = &auth.CredentialInput{
		Username: givenCredentialInput.Username,
		Password: "baz",
	}
	givenUser = &entity.User{
		Username: givenCredentialInput.Username,
		Password: hashPassword(givenCredentialInput.Password),
	}
	givenValidatorListErr = entity.ValidatorListErr(nil)
)

func hashPassword(password string) (hash string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}