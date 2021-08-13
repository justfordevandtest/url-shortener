package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shorturl/component/admin"
	"shorturl/component/admin/mocks"
	"shorturl/config"
	"shorturl/entity"
)

type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	config    *config.Config
	comp      admin.Comp
	urlRepo   *mocks.URLRepo
	validator *mocks.Validator
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.config = config.Get()
	suite.urlRepo = &mocks.URLRepo{}
	suite.validator = &mocks.Validator{}
	suite.comp = admin.New(suite.urlRepo, suite.validator)
}

var (
	givenListInput = &admin.ListInput{
		Page:    1,
		PerPage: 3,
		Filters: nil,
	}
	givenTotal = 10
	givenList = []entity.URL{
		{
			ID:      "aa",
			URL:     "https://rabbit.co.th/a",
			Expired: nil,
		},
		{
			ID:      "ab",
			URL:     "https://rabbit.co.th/b",
			Expired: nil,
		},
		{
			ID:      "ac ",
			URL:     "https://rabbit.co.th/c",
			Expired: nil,
		},
	}
	givenValidatorListErr = entity.ValidatorListErr(nil)
	givenDelInput = &admin.DelInput{
		ID: "Lb",
	}
	givenURL = &entity.URL{
		ID:      "Lb",
		URL:     "https://rabbit.co.th",
		Expired: nil,
	}
)