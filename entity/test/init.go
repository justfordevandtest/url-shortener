package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shorturl/entity"
)

type PackageTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
}

func (suite *PackageTestSuite) TearDownTest() {
	givenUser.Password = givenPwd
}

var (
	givenPwd = "rabbitfinance"
	givenUser = &entity.User{
		Username: "me",
		Password: givenPwd,
	}
)