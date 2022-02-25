package authGaurd

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"pwaa-test.com/models/entity"
	service "pwaa-test.com/services"
)

var identityKey = "id"
var AuthMiddleware *jwt.GinJWTMiddleware
func SetAuthGaurd() (*jwt.GinJWTMiddleware, error) {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := claims[identityKey].(string)
			userService := new(service.UserService)
			userService.Find(id)
			return &entity.User{
				
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals struct{
				accessToken string
			}
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			token := loginVals.accessToken
	
			if token == ""{ 
				return &entity.User{
					Type:  "userID",
					Key:  "stirng",
				}, nil
			}
	
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*entity.User); ok && v.Type == "kakao" {
				return true
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
	return AuthMiddleware, err
}
