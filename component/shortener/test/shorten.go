package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"shorturl/component/shortener"
	"shorturl/entity"
)

func (suite *PackageTestSuite) TestShortenURLSuccess() {
	suite.validator.On("Validate", givenShortenIn).Once().Return(nil)
	suite.urlRepo.On("CountID", mock.Anything).Once().Return(0, nil)
	suite.urlRepo.On("Create", mock.Anything).Once().Return(nil)
	shortURL, err := suite.comp.ShortenURL(givenShortenIn)

	suite.NoError(err)
	suite.IsType(&shortener.ShortenOutput{}, shortURL)
}

func (suite *PackageTestSuite) TestShortenURLSuccessNoExpired() {
	suite.validator.On("Validate", givenShortenIn).Once().Return(nil)
	suite.urlRepo.On("CountID", mock.Anything).Once().Return(0, nil)
	suite.urlRepo.On("Create", mock.Anything).Once().Return(nil)
	shortURL, err := suite.comp.ShortenURL(givenShortenIn)

	suite.NoError(err)
	suite.IsType(&shortener.ShortenOutput{}, shortURL)
}

func (suite *PackageTestSuite) TestShortenURLFailInvalidURL() {
	suite.validator.On("Validate", givenShortenIn).Once().Return(givenValidatorShortenErr)
	shortURL, err := suite.comp.ShortenURL(givenShortenIn)

	suite.Nil(shortURL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ValidatorShortenErr))
}

func (suite *PackageTestSuite) TestShortenURLFailBlacklistURL() {
	suite.validator.On("Validate", givenShortenIn).Once().Return(givenValidatorShortenErr)
	shortURL, err := suite.comp.ShortenURL(givenShortenIn)

	suite.Nil(shortURL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ValidatorShortenErr))
}

func (suite *PackageTestSuite) TestShortenURLFailCountRecord() {
	suite.validator.On("Validate", givenShortenIn).Once().Return(nil)
	suite.urlRepo.On("CountID", mock.Anything).Once().Return(0, errors.New(""))
	shortURL, err := suite.comp.ShortenURL(givenShortenIn)

	suite.Nil(shortURL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.CountRecordErr))
}

func (suite *PackageTestSuite) TestShortenURLFailCreateRecord() {
	suite.validator.On("Validate", givenShortenIn).Once().Return(nil)
	suite.urlRepo.On("CountID", mock.Anything).Once().Return(0, nil)
	suite.urlRepo.On("Create", mock.Anything).Once().Return(errors.New(""))
	shortURL, err := suite.comp.ShortenURL(givenShortenIn)

	suite.Nil(shortURL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.CreateRecordErr))
}
