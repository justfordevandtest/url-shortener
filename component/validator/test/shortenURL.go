package test

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func (suite *PackageTestSuite) TestShortenInputURLFailBlacklist() {
	suite.blacklistRepo.On("List").Once().Return(givenBlacklist, nil)
	err := suite.validator.Validate(givenInput)
	vErrs := err.(validator.ValidationErrors)

	suite.Error(err)
	suite.Equal("blacklist", vErrs[0].Tag())
}

func (suite *PackageTestSuite) TestShortenInputURLFailFetchBlacklist() {
	suite.blacklistRepo.On("List").Once().Return(nil, errors.New(""))
	err := suite.validator.Validate(givenInput)
	vErrs := err.(validator.ValidationErrors)

	suite.Error(err)
	suite.Equal("fetch-blacklist", vErrs[0].Tag())
}

func (suite *PackageTestSuite) TestShortenInputURLFailInvalidRegex() {
	suite.blacklistRepo.On("List").Once().Return(givenInvalidBlacklist, nil)
	err := suite.validator.Validate(givenInput)
	vErrs := err.(validator.ValidationErrors)

	suite.Error(err)
	suite.Equal("invalid-blacklist", vErrs[0].Tag())
}