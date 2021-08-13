package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shorturl/component/shortener"
	"shorturl/component/shortener/mocks"
	"shorturl/config"
	"shorturl/entity"
	"time"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	config    *config.Config
	comp      shortener.Comp
	urlRepo   *mocks.URLRepo
	urlCache  *mocks.URLCache
	validator *mocks.Validator
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.config = config.Get()
	suite.urlRepo = &mocks.URLRepo{}
	suite.urlCache = &mocks.URLCache{}
	suite.validator = &mocks.Validator{}
	suite.comp = shortener.
		New(suite.config.BaseURL, suite.config.CacheThreshold, suite.urlRepo, suite.urlCache, suite.validator)
}

var (
	givenExpired        = time.Now().Unix() + 5
	givenExpiredAlready = time.Now().Unix() - 1
	givenAccessIn       = &shortener.AccessInput{
		ID: "Lb",
	}
	givenShortenIn = &shortener.ShortenInput{
		URL:     "https://rabbit.co.th",
		Expired: &givenExpired,
	}
	givenURL = &entity.URL{
		ID:      "Lb",
		URL:     "https://rabbit.co.th",
		Expired: &givenExpired,
	}
	givenExpiredURL = &entity.URL{
		ID:      "Lb",
		URL:     "https://rabbit.co.th",
		Expired: &givenExpiredAlready,
	}
	givenPopularURL = &entity.URL{
		ID:      "Lb",
		URL:     "https://rabbit.co.th",
		Expired: &givenExpired,
		HitCount: 10,
	}
	givenValidatorShortenErr = entity.ValidatorShortenErr(nil)
)
