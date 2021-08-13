package test

import (
	"errors"
	"net/http"
)

func (suite *PackageTestSuite) TestListSuccess() {
	suite.comp.On("List", givenListInput).Once().Return(givenListOutput, nil)
	resp := performRequest(suite.router, "GET", listPath, nil)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestListServiceErr() {
	suite.comp.On("List", givenListInput).Once().Return(nil, errors.New("test"))
	resp := performRequest(suite.router, "GET", listPath, nil)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}