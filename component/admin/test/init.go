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
		ID:      "",
		Keyword: "",
	}
	givenListFilterIDInput = &admin.ListInput{
		Page:    1,
		PerPage: 3,
		ID:      "Lb",
		Keyword: "",
	}
	givenListFilterKeywordInput = &admin.ListInput{
		Page:    1,
		PerPage: 3,
		ID:      "",
		Keyword: "Lb",
	}
	givenEmptyFilters = make(map[string]interface{})
	givenIDFilters = map[string]interface{}{
		"ID": "Lb",
	}
	givenKeywordFilters = map[string]interface{}{
		"URL": "Lb",
	}
	givenTotal = 10
	givenList  = []entity.URL{
		{
			ID:      "aa",
			URL:     "https://example.com/a",
			Expired: nil,
		},
		{
			ID:      "ab",
			URL:     "https://example.com/b",
			Expired: nil,
		},
		{
			ID:      "ac ",
			URL:     "https://example.com/c",
			Expired: nil,
		},
	}
	givenValidatorListErr = entity.ValidatorListErr(nil)
	givenDelInput         = &admin.DelInput{
		ID: "Lb",
	}
	givenURL = &entity.URL{
		ID:      "Lb",
		URL:     "https://example.com",
		Expired: nil,
	}
)
