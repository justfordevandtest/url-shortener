// +build integration

package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shorturl/repository/user"
)

type PackageTestSuite struct {
	suite.Suite
	ctx  context.Context
	repo *user.StaticRepo
}

func (suite *PackageTestSuite) SetupTest() {
	var err error
	suite.repo, err = user.New()
	suite.NoError(err)
}

var (
	givenPwd = "pwd"
)
