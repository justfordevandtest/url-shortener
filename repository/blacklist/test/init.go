package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shorturl/config"
	"shorturl/repository/blacklist"
	"strings"
)

type PackageTestSuite struct {
	suite.Suite
	ctx  context.Context
	conf *config.Config
	repo *blacklist.StaticRepo
}

func (suite *PackageTestSuite) SetupTest() {
	suite.conf = config.Get()

	var err error
	suite.repo, err = blacklist.New(strings.Split(suite.conf.Blacklist, ","))
	suite.NoError(err)
}
