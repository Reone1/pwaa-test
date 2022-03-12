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
func (control *UserController) TestUserLogin(c *gin.Context) {
	token, err := userService.TestLogin()
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200, loginResponseBody{
		Token: token,
	})
}

type TwitterGetAccessRequestBody struct {
	CallbackURL string `json:"callback_url"`
}

// ShowAccount godoc
// @Summary      트위터 request Token
// @Description  트위터 request Token
// @Tags         test
// @Accept       json 
// @Param        body body TwitterGetAccessRequestBody false "callback_url"
// @Success      200  {string}  requestToken
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /twitter/requset-token [post]
func (control *UserController) TwitterGetAccess(c *gin.Context){
	var body TwitterGetAccessRequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}

	requestToken, _, err := userService.GetTwitterAuthToken(body.CallbackURL)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}
	c.JSON(200, requestToken)
}
type TwitterGetTokenRequestBody struct {
	OAuthToken string `json:"oauth_token"`
	OAuthTokenSecret string `json:"oauth_token_secret"` 
	OAuthVerifier string `json:"oauth_verifier"` 
	CallbackURL  string `json:"callbackURL"`
}
func (control *UserController) TwitterGetToken(c *gin.Context){
	var body TwitterGetTokenRequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}

	token, err := userService.GetTwitterAccessToken(body.OAuthToken, body.OAuthTokenSecret, body.OAuthVerifier, body.CallbackURL)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}