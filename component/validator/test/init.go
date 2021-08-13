package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shorturl/component/shortener"
	"shorturl/component/validator"
	"shorturl/component/validator/mocks"
)

type PackageTestSuite struct {
	suite.Suite
	ctx           context.Context
	validator     *validator.GoPlayGroundValidator
	blacklistRepo *mocks.BlacklistRepository
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.blacklistRepo = &mocks.BlacklistRepository{}
	suite.validator = validator.New(suite.blacklistRepo)
}

type SimpleTestStruct struct {
	Title string `validate:"required"`
	Body  string `validate:"max=5"`
}

func (suite *PackageTestSuite) TestValidatorValidateValid() {
	given := &SimpleTestStruct{
		Title: "test",
		Body:  "AAA",
	}

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestValidatorValidateInvalid() {
	given := &SimpleTestStruct{
		Title: "",
		Body:  "AAAAAAA",
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}

var (
	givenBlacklist = []string{
		"bad(.+)",
	}
	givenInput = shortener.ShortenInput{
		URL:     "https://sub.badURL.com",
		Expired: nil,
	}
	givenInvalidBlacklist = []string{"a(b"}
)