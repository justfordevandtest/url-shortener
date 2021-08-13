package test

import "strings"

func (suite *PackageTestSuite) TestList() {
	items, err := suite.repo.List()

	suite.NoError(err)
	suite.Equal(len(strings.Split(suite.conf.Blacklist, ",")), len(items))
}
