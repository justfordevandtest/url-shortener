package test

import (
	"errors"
	"shorturl/entity"
)

func (suite *PackageTestSuite) TestReadByCredentialSuccess() {
	suite.validator.On("Validate", givenCredentialInput).Once().Return(nil)
	suite.userRepo.On("Read", givenCredentialInput.Username).Once().Return(givenUser, nil)
	output, err := suite.comp.ReadByCredential(givenCredentialInput)

	suite.NoError(err)
	suite.Equal(givenCredentialInput.Username, output.Username)
}

func (suite *PackageTestSuite) TestReadByCredentialFailValidation() {
	suite.validator.On("Validate", givenCredentialInput).Once().Return(givenValidatorListErr)
	output, err := suite.comp.ReadByCredential(givenCredentialInput)

	suite.Nil(output)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ValidatorListErr))
}

func (suite *PackageTestSuite) TestReadByCredentialFailRepo() {
	suite.validator.On("Validate", givenCredentialInput).Once().Return(nil)
	suite.userRepo.On("Read", givenCredentialInput.Username).Once().Return(nil, errors.New(""))
	output, err := suite.comp.ReadByCredential(givenCredentialInput)

	suite.Nil(output)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ReadRecordErr))
}

func (suite *PackageTestSuite) TestReadByCredentialFailPasswordCheck() {
	suite.validator.On("Validate", givenIncorrectPasswordInput).Once().Return(nil)
	suite.userRepo.On("Read", givenIncorrectPasswordInput.Username).Once().Return(givenUser, nil)
	output, err := suite.comp.ReadByCredential(givenIncorrectPasswordInput)

	suite.Nil(output)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.UnauthorizedErr))
}