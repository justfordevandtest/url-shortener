package test

import (
	"context"
	"shorturl/entity"
	"shorturl/repository/urlcache"
	"time"

	"github.com/stretchr/testify/suite"
	"shorturl/config"
)

type PackageTestSuite struct {
	suite.Suite
	ctx   context.Context
	cache *url.RedisCache
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	conf := config.Get()

	var err error
	suite.cache, err = url.New(conf.RedisCacheAddr)
	suite.NoError(err)
}

func (suite *PackageTestSuite) SetupTest() {
	conf := config.Get()

	var err error
	suite.cache, err = url.New(conf.RedisCacheAddr)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TearDownTest() {
	_ = suite.cache.Client.FlushAll(suite.ctx)
}

func (suite *PackageTestSuite) TearDownSuite() {
	_ = suite.cache.Client.FlushAll(suite.ctx)
}

var (
	givenExpired = time.Now().Unix() + 5
	givenURL     = &entity.URL{
		ID:      "Lb",
		URL:     "https://rabbit.co.th",
		Expired: &givenExpired,
	}
	givenURLOther = &entity.URL{
		ID:      "Lc",
		URL:     "https://rabbit.co.th/a",
		Expired: &givenExpired,
	}
)
