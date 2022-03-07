package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pwaa-test.com/models/entity"
	httputil "pwaa-test.com/module/utils"
	service "pwaa-test.com/services"
)

var userService = new(service.UserService)
type UserController struct {}

func (control *UserController) GetUser(c *gin.Context) {
	fmt.Print("getUser controller called")
	c.JSON(200, gin.H{
		"message": c.Request.URL,
	})
}

func (control *UserController) CreateOne(c *gin.Context) {
	if err := userService.Create(&entity.User{}); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200,gin.H{
		"message": "success Create User",
	})
}

type loginResponseBody struct {
	Token string `json:"token" example:"token string (JWT)"`
}
// ShowAccount godoc
// @Summary      테스트 유저 로그인
// @Description  테스트 유저 로그인
// @Tags         test
// @Accept       json 
// @Success      200  {object}  loginResponseBody
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /test/login [post]
func (contorl *UserController) TestUserLogin(c *gin.Context) {
	token, err := userService.TestLogin()
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200, loginResponseBody{
		Token: token,
	})
}