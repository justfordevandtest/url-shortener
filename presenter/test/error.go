package test

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
	"shorturl/entity"
	"shorturl/presenter"
)

type ValidatorTestStruct struct {
	Title string `validate:"required" json:"title"`
	Body  string `validate:"required"`
}

func (suite *PackageTestSuite) TestMakeErrResp() {
	err := entity.ExpiredURLErr(errors.New("test"))
	resp := presenter.MakeErrResp(err)
	suite.Equal(http.StatusGone, resp.Code)
}

func (suite *PackageTestSuite) TestMakeValidationShortenErrResp() {
	s := ValidatorTestStruct{
		Title: "",
		Body:  "",
	}
	cause := validator.New().Struct(s)

	resp := presenter.MakeErrResp(entity.ValidatorShortenErr(cause))
	suite.Equal(http.StatusUnprocessableEntity, resp.Code)
}

func (suite *PackageTestSuite) TestMakeReadRecordErrResp() {
	resp := presenter.MakeErrResp(entity.ReadRecordErr(nil))
	suite.Equal(http.StatusNotFound, resp.Code)
}

func (suite *PackageTestSuite) TestMakeNotFoundRecordErrResp() {
	resp := presenter.MakeErrResp(entity.NotFoundRecordErr(nil))
	suite.Equal(http.StatusNotFound, resp.Code)
}

func (suite *PackageTestSuite) TestMakeInternalErrResp() {
	resp := presenter.MakeErrResp(errors.New("test"))
	suite.Equal(http.StatusInternalServerError, resp.Code)
}
