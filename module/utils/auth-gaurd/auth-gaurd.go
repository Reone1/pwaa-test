package authGaurd

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
	"pwaa-test.com/module/utils/jwt"
	service "pwaa-test.com/services"
)

func AuthMiddleware(c *gin.Context) {
	if len(c.Request.Header["Auth-Token"]) == 0 {
		httputil.NewError	(c, http.StatusUnauthorized, errors.New("not 'Authorization' header"))
		c.Abort()
	}
	jwtModule := new(jwt.Module)
	fmt.Print(c.Request.Header)
	auth := c.Request.Header["Auth-Token"][0]
	if len(auth) == 0 {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("not 'Authorization' header"))
		c.Abort()
	}
	token := strings.Split(auth, " ")[1]

	id := jwtModule.DecodeToken(token)
	userService := new(service.UserService)
	user, err := userService.FindById(id)
	if err != nil {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("can found USER"))
		c.Abort()
	}
	c.Set("userId", user.ID.Hex())
}