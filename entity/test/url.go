package test

import (
	"shorturl/entity"
	"time"
)

func (suite *PackageTestSuite) TestMakeURL() {
	givenNumb := uint64(99)
	givenURL := "https://example.com"
	givenExpired := time.Now().Unix() + 5

	url := entity.MakeURL(givenNumb, givenURL, &givenExpired)

	suite.IsType(&entity.URL{}, url)
	suite.Equal("Lb", url.ID)
	suite.Equal(givenExpired, *url.Expired)
}
