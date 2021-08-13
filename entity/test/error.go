package test

import (
	"errors"
	"shorturl/entity"
)

func (suite *PackageTestSuite) TestTypeOfErr() {
	givenCause := errors.New("test")
	givenErr := entity.ExpiredURLErr(givenCause)

	actualErr := entity.TypeOfErr(givenErr)

	suite.Equal(givenErr, actualErr)
}

func (suite *PackageTestSuite) TestTypeOfErrSimple() {
	givenCause := errors.New("test")

	actualErr := entity.TypeOfErr(givenCause)

	suite.Equal(actualErr.Cause, givenCause)
}

func (suite *PackageTestSuite) TestError() {
	givenCause := errors.New("test")
	givenErr := entity.ExpiredURLErr(givenCause)

	actualErr := givenErr.Error()

	expectErr := "test (EXPIRED:URL)"
	suite.Equal(expectErr, actualErr)
}

func (suite *PackageTestSuite) TestIsType() {
	givenCause := errors.New("test")
	givenErr := entity.ExpiredURLErr(givenCause)

	actual := givenErr.IsType(entity.ExpiredURLErr)

	suite.Equal(true, actual)
}
