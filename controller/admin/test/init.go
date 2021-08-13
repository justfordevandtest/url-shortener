package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"shorturl/component/admin"
	"shorturl/component/admin/mocks"
	"shorturl/config"
	adminCtrl "shorturl/controller/admin"
	"shorturl/entity"
)

type PackageTestSuite struct {
	suite.Suite
	router *gin.Engine
	ctx    *gin.Context
	conf   *config.Config
	ctrl   *adminCtrl.Ctrl
	comp   *mocks.Comp
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.comp = &mocks.Comp{}
	suite.ctrl = adminCtrl.New(suite.comp)

	suite.router = gin.New()
	suite.router.Handle(http.MethodGet, "/admin", suite.ctrl.List)
	suite.router.Handle(http.MethodDelete, "/admin/:id", suite.ctrl.Delete)
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var (
	listPath       = "/admin?Page=1&PerPage=3"
	delPath        = "/admin/Lb"
	givenListInput = &admin.ListInput{
		Page:    1,
		PerPage: 3,
		ID:      "",
		Keyword: "",
	}
	givenList = []entity.URL{
		{
			ID:      "aa",
			URL:     "https://example.com/a",
			Expired: nil,
		},
		{
			ID:      "ab",
			URL:     "https://example.com/b",
			Expired: nil,
		},
		{
			ID:      "ac ",
			URL:     "https://example.com/c",
			Expired: nil,
		},
	}
	givenTotal      = 10
	givenListOutput = &admin.ListOutput{
		Total: givenTotal,
		Items: givenList,
	}
	givenDelInput = &admin.DelInput{
		ID: "Lb",
	}
)
