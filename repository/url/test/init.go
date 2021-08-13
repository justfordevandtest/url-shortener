package test

import (
	"context"
	"shorturl/entity"
	"shorturl/repository/url"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"shorturl/config"
)

type PackageTestSuite struct {
	suite.Suite
	ctx  context.Context
	repo *url.MongoDBRepo
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	conf := config.Get()

	var err error
	suite.repo, err = url.New(suite.ctx, conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBURLCollName)
	suite.NoError(err)
}

func (suite *PackageTestSuite) SetupTest() {
	conf := config.Get()

	var err error
	suite.repo, err = url.New(suite.ctx, conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBURLCollName)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TearDownTest() {
	_, _ = suite.repo.Coll.DeleteMany(suite.ctx, bson.M{})
}

func (suite *PackageTestSuite) TearDownSuite() {
	_ = suite.repo.DB.Drop(suite.ctx)
}

var (
	givenExpired        = time.Now().Unix() + 5
	givenURL = &entity.URL{
		ID:      "Lb",
		URL:     "https://rabbit.co.th",
		Expired: &givenExpired,
	}
	givenURLOther = &entity.URL{
		ID:      "Lc",
		URL:     "https://rabbit.co.th/a",
		Expired: &givenExpired,
	}
	givenFilters = map[string]interface{}{
		"ID": "Lb",
	}
	givenFiltersNotFound = map[string]interface{}{
		"ID": "notfound",
	}
)