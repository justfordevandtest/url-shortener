package test

import "shorturl/config"

func (suite *PackageTestSuite) TestGet() {
	conf := config.Get()

	suite.Equal("http://localhost:8080/api/v1", conf.BaseURL)
	suite.Equal("rabbit finance test", conf.JWTRealm)
	suite.Equal("secret", conf.JWTSecret)
}
