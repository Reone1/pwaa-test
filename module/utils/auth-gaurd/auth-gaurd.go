package authGaurd

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
	"pwaa-test.com/module/utils/jwt"
	service "pwaa-test.com/services"
)

func AuthMiddleware(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("not 'Authorization' header"))
		c.Abort()
	}
	jwtModule := new(jwt.Module)
	auth := c.Request.Header["Authorization"][0]
	token := strings.Split(auth, " ")[1]

	id := jwtModule.DecodeToken(token)
	userService := new(service.UserService)
	user, err := userService.Find(id)
	if err != nil {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("can found USER"))
		c.Abort()
	}
	c.Set("userId", user.ID.Hex())
}