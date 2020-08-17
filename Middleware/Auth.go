package middleware

import (
	"time"

	"go-api/models"
	"go-api/repositories"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "login"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Auth() (*jwt.GinJWTMiddleware, error) {
	rep := repositories.NewUserRepository()
	return jwt.New(&jwt.GinJWTMiddleware{

		Realm:       "go-api",
		Key:         []byte("secretkey"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Login,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Login: claims[identityKey].(string),
			}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var json login
			if err := c.ShouldBind(&json); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			var user models.User
			if err := rep.LoginUser(&user, json.Username); err == nil {
				if checkPasswordHash(json.Password, user.Senha) {
					return &user, nil
				}
			}

			return nil, jwt.ErrFailedAuthentication
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			token, ok := data.(*models.User)
			if ok {
				var user models.User
				if err := rep.LoginUser(&user, token.Login); err == nil {
					return true
				}
			}

			return false
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
