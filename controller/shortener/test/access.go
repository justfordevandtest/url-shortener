package test

import (
	"fmt"
	"net/http"
)

func (suite *PackageTestSuite) TestAccessSuccess() {
	suite.comp.On("AccessURL", givenAccessInput).Once().Return(givenOutput, nil)
	resp := performRequest(suite.router, "GET", fmt.Sprintf("/%s", givenAccessInput.ID), nil)

	suite.Equal(http.StatusFound, resp.Code)
}

func (suite *PackageTestSuite) TestAccessFailExpired() {
	suite.comp.On("AccessURL", givenAccessInput).Once().Return(nil, givenExpiredErr)
	resp := performRequest(suite.router, "GET", fmt.Sprintf("/%s", givenAccessInput.ID), nil)

	suite.Equal(http.StatusGone, resp.Code)
}

func (suite *PackageTestSuite) TestAccessFailNotFound() {
	suite.comp.On("AccessURL", givenAccessInput).Once().Return(nil, givenNotFoundErr)
	resp := performRequest(suite.router, "GET", fmt.Sprintf("/%s", givenAccessInput.ID), nil)

	suite.Equal(http.StatusNotFound, resp.Code)
}
