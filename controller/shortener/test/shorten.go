package test

import (
	"bytes"
	"errors"
	"net/http"
)

func (suite *PackageTestSuite) TestShortenSuccess() {
	suite.comp.On("ShortenURL", givenShortenInput).Once().Return(givenOutput, nil)
	resp := performRequest(suite.router, "POST", "/shorten", buildJsonRequestBody(givenShortenInput))

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) TestShortenBadJson() {
	resp := performRequest(suite.router, "POST", "/shorten", bytes.NewBuffer([]byte("test")))
	suite.Equal(http.StatusBadRequest, resp.Code)
}

func (suite *PackageTestSuite) TestShortenServiceErr() {
	suite.comp.On("ShortenURL", givenShortenInput).Once().Return(nil, errors.New("test"))
	resp := performRequest(suite.router, "POST", "/shorten", buildJsonRequestBody(givenShortenInput))

	suite.Equal(http.StatusInternalServerError, resp.Code)
}

func (suite *PackageTestSuite) TestShortenValidationErr() {
	suite.comp.On("ShortenURL", givenShortenInput).Once().Return(nil, givenValidateErr)
	resp := performRequest(suite.router, "POST", "/shorten", buildJsonRequestBody(givenShortenInput))

	suite.Equal(http.StatusUnprocessableEntity, resp.Code)
}
