package test

import (
	"errors"
	"net/http"
)

func (suite *PackageTestSuite) TestDeleteSuccess() {
	suite.comp.On("Delete", givenDelInput).Once().Return(nil)
	resp := performRequest(suite.router, "DELETE", delPath, nil)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestDeleteServiceErr() {
	suite.comp.On("Delete", givenDelInput).Once().Return(errors.New("test"))
	resp := performRequest(suite.router, "DELETE", delPath, nil)

	suite.Equal(http.StatusInternalServerError, resp.Code)
}