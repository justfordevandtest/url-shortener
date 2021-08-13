package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shorturl/component/auth/mocks"
	"shorturl/config"
)

type PackageTestSuite struct {
	suite.Suite
	ctx  context.Context
	conf *config.Config
	comp *mocks.Comp
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.conf = config.Get()
	suite.comp = &mocks.Comp{}
}
