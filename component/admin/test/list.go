package test

import (
	"errors"
	"shorturl/entity"
)

func (suite *PackageTestSuite) TestListSuccess() {
	suite.validator.On("Validate", givenListInput).Once().Return(nil)
	suite.urlRepo.On("List", givenListInput.Page, givenListInput.PerPage, givenEmptyFilters).Once().Return(givenTotal, givenList, nil)
	output, err := suite.comp.List(givenListInput)

	suite.NoError(err)
	suite.Equal(givenTotal, output.Total)
	suite.Equal(givenListInput.PerPage, len(output.Items))
}

func (suite *PackageTestSuite) TestListSuccessSearchID() {
	suite.validator.On("Validate", givenListFilterIDInput).Once().Return(nil)
	suite.urlRepo.On("List", givenListInput.Page, givenListInput.PerPage, givenIDFilters).Once().Return(givenTotal, givenList, nil)
	output, err := suite.comp.List(givenListFilterIDInput)

	suite.NoError(err)
	suite.Equal(givenTotal, output.Total)
	suite.Equal(givenListInput.PerPage, len(output.Items))
}

func (suite *PackageTestSuite) TestListSuccessSearchKeyword() {
	suite.validator.On("Validate", givenListFilterKeywordInput).Once().Return(nil)
	suite.urlRepo.On("List", givenListInput.Page, givenListInput.PerPage, givenKeywordFilters).Once().Return(givenTotal, givenList, nil)
	output, err := suite.comp.List(givenListFilterKeywordInput)

	suite.NoError(err)
	suite.Equal(givenTotal, output.Total)
	suite.Equal(givenListInput.PerPage, len(output.Items))
}

func (suite *PackageTestSuite) TestListFailValidation() {
	suite.validator.On("Validate", givenListInput).Once().Return(givenValidatorListErr)
	output, err := suite.comp.List(givenListInput)

	suite.Nil(output)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ValidatorListErr))
}

func (suite *PackageTestSuite) TestListFailRepo() {
	suite.validator.On("Validate", givenListInput).Once().Return(nil)
	suite.urlRepo.On("List", givenListInput.Page, givenListInput.PerPage, givenEmptyFilters).Once().Return(0, nil, errors.New(""))
	output, err := suite.comp.List(givenListInput)

	suite.Nil(output)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ListRecordsErr))
}