package auth

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/component/auth"
	"time"
)

var (
	identityKey = "Username"
)

func New(comp auth.Comp, realm string, secret string) (middleware *jwt.GinJWTMiddleware, err error) {
	if len(realm) < 1 {
		return nil, errors.New("missing realm")
	}
	if len(secret) < 1 {
		return nil, errors.New("missing secret")
	}

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       realm,
		Key:         []byte(secret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "Username",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*auth.UserOutput); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &auth.UserOutput{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: Login(comp),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:    "header: Authorization, query: token, cookie:token",
		TokenHeadName:  "Bearer",
		TimeFunc:       time.Now,
		SendCookie:     true,
		SecureCookie:   false,
		CookieHTTPOnly: true,
		CookieDomain:   "localhost:8080",
		CookieName:     "token",
		CookieSameSite: http.SameSiteDefaultMode,
	})
}

func Login(comp auth.Comp) (handler func(c *gin.Context) (interface{}, error)) {
	return func(c *gin.Context) (interface{}, error) {
		var input auth.CredentialInput
		if err := c.ShouldBind(&input); err != nil {
			return nil, jwt.ErrMissingLoginValues
		}

		user, err := comp.ReadByCredential(&input)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		return user, nil
	}
}
