package test

func (suite *PackageTestSuite) TestRedisReadWrite() {
	err := suite.cache.Write(givenURL)
	suite.NoError(err)

	url := suite.cache.Read(givenURL.ID)
	suite.NotNil(url)
}

func (suite *PackageTestSuite) TestRedisReadFail() {
	err := suite.cache.Write(givenURL)
	suite.NoError(err)

	url := suite.cache.Read(givenURLOther.ID)
	suite.Nil(url)
}

func (suite *PackageTestSuite) TestRedisReadFailUnmarshal() {
	err := suite.cache.Client.Set(suite.ctx, givenURL.ID, "", 0).Err()
	suite.NoError(err)

	url := suite.cache.Read(givenURL.ID)
	suite.Nil(url)
}
