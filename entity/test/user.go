package test

func (suite *PackageTestSuite) TestHashPassword() {
	givenUser.HashPassword()
	suite.NotEqual(givenPwd, givenUser.Password)
}

func (suite *PackageTestSuite) TestCheckPassword() {
	givenUser.HashPassword()
	suite.True(givenUser.CheckPasswordHash(givenPwd))
}
