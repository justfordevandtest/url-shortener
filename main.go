package main

import (
	"github.com/gin-gonic/gin"
	ginLogRus "github.com/toorop/gin-logrus"
	"shorturl/config"
)

// @title URL Shortener API
// @version 1.0
// @description Create a URL-shortener service to shorten URLs.\n\nAPI clients will be able to create short URLs from a full length URL.\n\nIt will also support redirecting the short urls to the correct url.

// @contact.name API Support
// @contact.url https://github.com/justfordevandtest/rabbit-finance-test
// @contact.email ekkasith.w@gmail.com

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl http://localhost:8080/api/v1/admin/login

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	appConfig := config.Get()

	logger := setupLog()

	router := gin.New()
	router.Use(ginLogRus.Logger(logger), gin.Recovery())

	newApp(appConfig).RegisterRoute(router)
	_ = router.Run()
}