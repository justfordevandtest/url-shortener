package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"shorturl/component/shortener"
	"shorturl/component/shortener/mocks"
	"shorturl/config"
	shortenerCtrl "shorturl/controller/shortener"
	"shorturl/entity"
	"time"
)

type PackageTestSuite struct {
	suite.Suite
	router *gin.Engine
	ctx    *gin.Context
	conf   *config.Config
	ctrl   *shortenerCtrl.Ctrl
	comp   *mocks.Comp
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.comp = &mocks.Comp{}
	suite.ctrl = shortenerCtrl.New(suite.comp)

	suite.router = gin.New()
	suite.router.Handle(http.MethodPost, "/shorten", suite.ctrl.Shorten)
	suite.router.Handle(http.MethodGet, "/:id", suite.ctrl.Access)
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func buildJsonRequestBody(input interface{}) (body io.Reader) {
	jsonBytes, _ := json.Marshal(input)
	return bytes.NewBuffer(jsonBytes)
}

var (
	givenExpired      = time.Now().Unix() + 5
	givenShortenInput = &shortener.ShortenInput{
		URL:     "https://example.com",
		Expired: &givenExpired,
	}
	givenAccessInput = &shortener.AccessInput{
		ID: "Lb",
	}
	givenOutput      = &shortener.ShortenOutput{URL: "https://example.com"}
	givenExpiredErr  = entity.ExpiredURLErr(nil)
	givenNotFoundErr = entity.NotFoundRecordErr(nil)
	givenValidateErr = entity.ValidatorShortenErr(validator.ValidationErrors{})
)