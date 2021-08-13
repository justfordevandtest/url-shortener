package test

import (
	"net/http"
	"shorturl/presenter"
)

type ItemStruct struct {
	Title string
	Body  string
}

func (suite *PackageTestSuite) TestMakeSuccessResp() {
	st := struct {
		Title string `validate:"required"`
		Body  string `validate:"required"`
	}{
		Title: "",
		Body:  "",
	}
	presenter.MakeSuccessResp(http.StatusOK, st)
	suite.Equal(http.StatusOK, suite.ctx.Writer.Status())
}