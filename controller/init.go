package app

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"shorturl/controller/admin"
	"shorturl/controller/shortener"
	"shorturl/docs"
)

type App struct {
	shortenerCtrl *shortener.Ctrl
	adminCtrl     *admin.Ctrl
	authMiddle    *jwt.GinJWTMiddleware
}

func New(shortenerCtrl *shortener.Ctrl, adminCtrl *admin.Ctrl, authMiddle *jwt.GinJWTMiddleware) *App {
	return &App{
		shortenerCtrl: shortenerCtrl,
		adminCtrl:     adminCtrl,
		authMiddle:    authMiddle,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	publicRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		publicRoutes.POST("/shorten", app.shortenerCtrl.Shorten)
		publicRoutes.GET("/:id", app.shortenerCtrl.Access)

		adminRoutes := publicRoutes.Group("/admin")
		adminRoutes.POST("/login", app.authMiddle.LoginHandler)
		adminRoutes.GET("/refresh", app.authMiddle.RefreshHandler)
		adminRoutes.Use(app.authMiddle.MiddlewareFunc())
		{
			adminRoutes.GET("", app.adminCtrl.List)
			adminRoutes.DELETE("/:id", app.adminCtrl.Delete)
		}
	}

	return app
}
