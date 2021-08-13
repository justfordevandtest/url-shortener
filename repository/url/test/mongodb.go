// +build integration

package test

func (suite *PackageTestSuite) TestCreate() {
	err := suite.repo.Create(givenURL)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TestList() {
	var err error
	err = suite.repo.Create(givenURL)
	suite.NoError(err)
	err = suite.repo.Create(givenURLOther)
	suite.NoError(err)

	total, items, err := suite.repo.List(1, 10, nil)

	suite.NoError(err)
	suite.Equal(2, total)
	suite.Equal(2, len(items))
}

func (suite *PackageTestSuite) TestSearchList() {
	var err error
	err = suite.repo.Create(givenURL)
	suite.NoError(err)
	err = suite.repo.Create(givenURLOther)
	suite.NoError(err)

	total, items, err := suite.repo.List(1, 10, givenFilters)

	suite.NoError(err)
	suite.Equal(1, total)
	suite.Equal(1, len(items))
}

func (suite *PackageTestSuite) TestSearchListNotFound() {
	var err error
	err = suite.repo.Create(givenURL)
	suite.NoError(err)
	err = suite.repo.Create(givenURLOther)
	suite.NoError(err)

	total, items, err := suite.repo.List(1, 10, givenFiltersNotFound)

	suite.NoError(err)
	suite.Equal(0, total)
	suite.Equal(0, len(items))
}

func (suite *PackageTestSuite) TestRead() {
	_ = suite.repo.Create(givenURL)
	url, err := suite.repo.Read(givenURL.ID)

	suite.NoError(err)
	suite.Equal(givenURL.ID, url.ID)
	suite.Equal(givenURL.URL, url.URL)
}

func (suite *PackageTestSuite) TestReadNotFound() {
	_ = suite.repo.Create(givenURL)
	url, err := suite.repo.Read(givenURLOther.ID)

	suite.Nil(url)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestIncrHit() {
	err := suite.repo.Create(givenURL)
	suite.NoError(err)

	err = suite.repo.IncrHit(givenURL.ID)
	suite.NoError(err)

	url, err := suite.repo.Read(givenURL.ID)

	suite.NoError(err)
	suite.Equal(givenURL.HitCount + 1, url.HitCount)
}

func (suite *PackageTestSuite) TestIncrHitNotFound() {
	err := suite.repo.Create(givenURL)
	suite.NoError(err)

	err = suite.repo.IncrHit(givenURLOther.ID)
	suite.Error(err)

	url, err := suite.repo.Read(givenURL.ID)

	suite.NoError(err)
	suite.Equal(givenURL.HitCount, url.HitCount)
}

func (suite *PackageTestSuite) TestDelete() {
	err := suite.repo.Create(givenURL)
	suite.NoError(err)

	err = suite.repo.Delete(givenURL.ID)
	suite.NoError(err)

	url, err := suite.repo.Read(givenURL.ID)

	suite.NoError(err)
	suite.Equal(int64(0), *url.Expired)
}

func (suite *PackageTestSuite) TestDeleteNotFound() {
	err := suite.repo.Create(givenURL)
	suite.NoError(err)

	err = suite.repo.Delete(givenURLOther.ID)
	suite.Error(err)

	url, err := suite.repo.Read(givenURL.ID)

	suite.NoError(err)
	suite.Equal(*givenURL.Expired, *url.Expired)
}
