package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	authComp "shorturl/component/auth"
	"shorturl/component/validator"
	"shorturl/config"
	"shorturl/controller/admin"
	"shorturl/controller/shortener"
	"shorturl/middleware/auth"
	"strings"

	adminComp "shorturl/component/admin"
	shortenerComp "shorturl/component/shortener"
	app "shorturl/controller"
	blacklistRepo "shorturl/repository/blacklist"
	urlRepo "shorturl/repository/url"
	urlCache "shorturl/repository/urlcache"
	userRepo "shorturl/repository/user"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	uRepo, err := urlRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBURLCollName)
	panicIfErr(err)
	bRepo, err := blacklistRepo.New(strings.Split(appConfig.Blacklist, ","))
	panicIfErr(err)
	usrRepo, err := userRepo.New()
	panicIfErr(err)

	uCache, err := urlCache.New(appConfig.RedisCacheAddr)
	panicIfErr(err)

	v := validator.New(bRepo)

	shortenerCmp := shortenerComp.New(appConfig.BaseURL, appConfig.CacheThreshold, uRepo, uCache, v)
	adminCmp := adminComp.New(uRepo, v)
	authCmp := authComp.New(usrRepo, v)

	sCtrl := shortener.New(shortenerCmp)
	aCtrl := admin.New(adminCmp)

	authMiddle, err := auth.New(authCmp, appConfig.JWTRealm, appConfig.JWTSecret)
	panicIfErr(err)

	return app.New(sCtrl, aCtrl, authMiddle)
}

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
