package test

import (
	"errors"
	"shorturl/entity"
)

func (suite *PackageTestSuite) TestDeleteSuccess() {
	suite.validator.On("Validate", givenDelInput).Once().Return(nil)
	suite.urlRepo.On("Read", givenDelInput.ID).Once().Return(givenURL, nil)
	suite.urlRepo.On("Delete", givenURL.ID).Once().Return(nil)
	err := suite.comp.Delete(givenDelInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteFailValidation() {
	suite.validator.On("Validate", givenDelInput).Once().Return(givenValidatorListErr)
	err := suite.comp.Delete(givenDelInput)

	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ValidatorListErr))
}

func (suite *PackageTestSuite) TestDeleteFailRepoRead() {
	suite.validator.On("Validate", givenDelInput).Once().Return(nil)
	suite.urlRepo.On("Read", givenDelInput.ID).Once().Return(nil, errors.New(""))
	err := suite.comp.Delete(givenDelInput)

	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ReadRecordErr))
}

func (suite *PackageTestSuite) TestDeleteFailRepoDeletion() {
	suite.validator.On("Validate", givenDelInput).Once().Return(nil)
	suite.urlRepo.On("Read", givenDelInput.ID).Once().Return(givenURL, nil)
	suite.urlRepo.On("Delete", givenURL.ID).Once().Return(errors.New(""))
	err := suite.comp.Delete(givenDelInput)

	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.DeleteRecordErr))
}