package test

func (suite *PackageTestSuite) TestRead() {
	user, err := suite.repo.Read("me")

	suite.NoError(err)
	suite.Equal("me", user.Username)
	suite.True(user.CheckPasswordHash(givenPwd))
}
