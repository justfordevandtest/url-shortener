package test

import (
	"errors"
	"shorturl/entity"
)

func (suite *PackageTestSuite) TestAccessURLSuccessCacheMiss() {
	suite.urlCache.On("Read", givenAccessIn.ID).Once().Return(nil)
	suite.urlRepo.On("Read", givenAccessIn.ID).Once().Return(givenURL, nil)
	suite.urlRepo.On("IncrHit", givenAccessIn.ID).Once().Return(nil)
	URL, err := suite.comp.AccessURL(givenAccessIn)

	suite.NoError(err)
	suite.Equal(givenURL.URL, URL.URL)
}

func (suite *PackageTestSuite) TestAccessURLSuccessCacheHit() {
	suite.urlCache.On("Read", givenAccessIn.ID).Once().Return(givenPopularURL)
	suite.urlRepo.On("IncrHit", givenAccessIn.ID).Once().Return(nil)
	suite.urlCache.On("Write", givenPopularURL).Once().Return(nil)
	URL, err := suite.comp.AccessURL(givenAccessIn)

	suite.NoError(err)
	suite.Equal(givenURL.URL, URL.URL)
}

func (suite *PackageTestSuite) TestAccessURLSuccessNoExpired() {
	suite.urlCache.On("Read", givenAccessIn.ID).Once().Return(nil)
	suite.urlRepo.On("Read", givenAccessIn.ID).Once().Return(givenURL, nil)
	suite.urlRepo.On("IncrHit", givenAccessIn.ID).Once().Return(nil)
	URL, err := suite.comp.AccessURL(givenAccessIn)

	suite.NoError(err)
	suite.Equal(givenURL.URL, URL.URL)
}

func (suite *PackageTestSuite) TestAccessURLFailCacheWrite() {
	suite.urlCache.On("Read", givenAccessIn.ID).Once().Return(givenPopularURL)
	suite.urlRepo.On("IncrHit", givenAccessIn.ID).Once().Return(nil)
	suite.urlCache.On("Write", givenPopularURL).Once().Return(errors.New(""))
	URL, err := suite.comp.AccessURL(givenAccessIn)

	suite.Nil(URL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.CacheWriteErr))
}

func (suite *PackageTestSuite) TestAccessURLFailExpired() {
	suite.urlCache.On("Read", givenAccessIn.ID).Once().Return(nil)
	suite.urlRepo.On("Read", givenAccessIn.ID).Once().Return(givenExpiredURL, nil)
	URL, err := suite.comp.AccessURL(givenAccessIn)

	suite.Nil(URL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.ExpiredURLErr))
}

func (suite *PackageTestSuite) TestAccessURLFailNotFound() {
	suite.urlCache.On("Read", givenAccessIn.ID).Once().Return(nil)
	suite.urlRepo.On("Read", givenAccessIn.ID).Once().Return(nil, errors.New(""))
	URL, err := suite.comp.AccessURL(givenAccessIn)

	suite.Nil(URL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.NotFoundRecordErr))
}

func (suite *PackageTestSuite) TestAccessURLFailIncrHit() {
	suite.urlCache.On("Read", givenAccessIn.ID).Once().Return(nil)
	suite.urlRepo.On("Read", givenAccessIn.ID).Once().Return(givenURL, nil)
	suite.urlRepo.On("IncrHit", givenAccessIn.ID).Once().Return(errors.New(""))
	URL, err := suite.comp.AccessURL(givenAccessIn)

	suite.Nil(URL)
	suite.Error(err)
	suite.IsType(&entity.Error{}, err)
	suite.True(err.(*entity.Error).IsType(entity.UpdateRecordErr))
}

