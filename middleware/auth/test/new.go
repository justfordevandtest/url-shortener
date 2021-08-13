package test

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"shorturl/middleware/auth"
)

func (suite *PackageTestSuite) TestNewSuccess() {
	m, err := auth.New(suite.comp, suite.conf.JWTRealm, suite.conf.JWTSecret)

	suite.NoError(err)
	suite.IsType(&jwt.GinJWTMiddleware{}, m)
}

func (suite *PackageTestSuite) TestNewFailMissingRealm() {
	m, err := auth.New(suite.comp, "", suite.conf.JWTSecret)

	suite.Nil(m)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestNewFailMissingSecret() {
	m, err := auth.New(suite.comp, suite.conf.JWTRealm, "")

	suite.Nil(m)
	suite.Error(err)
}