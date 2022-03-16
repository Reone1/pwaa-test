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

type TwitterGetAccessRequestQuery struct {
	CallbackURL string `json:"callback_url"`
}

// ShowAccount godoc
// @Summary      트위터 request Token
// @Description  트위터 request Token
// @Tags         twitter
// @Accept       json 
// @Param        query query TwitterGetAccessRequestQuery false "callback_url"
// @Success      200  {string}  requestToken
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /twitter/requset-token [get]
func (control *UserController) TwitterGetRequest(c *gin.Context){
	var query = c.Request.URL.Query()
	requestToken, requestSecret, err := userService.GetTwitterAuthToken(query.Get("callback_url"))
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}
	c.JSON(200, gin.H{
		"oauth_token": requestToken,
		"oauth_token_secret": requestSecret,
	})
}

type TwitterGetTokenRequestBody struct {
	OAuthToken string `json:"oauth_token"`
	OAuthTokenSecret string `json:"oauth_token_secret"` 
	OAuthVerifier string `json:"oauth_verifier"` 
	CallbackURL  string `json:"callbackURL"`
}

// ShowAccount godoc
// @Summary      트위터 access Token
// @Description  트위터 access Token
// @Tags         twitter
// @Accept       json 
// @Param        query query TwitterGetTokenRequestBody false "토큰이 필요합니다."
// @Success      200  {string}  token
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /twitter/access-token [get]
func (control *UserController) TwitterGetAccess(c *gin.Context){
	var query = c.Request.URL.Query()

	token, err := userService.GetTwitterAccessToken(query.Get("oauth_token"), query.Get("oauth_token_secret"), query.Get("oauth_verifier"), query.Get("callbackURL"))

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return;
	}
	c.JSON(200, token)
}

type KakaoTokenRequestBody struct {
	GrantType string `json:"grant_type"`
	ClientId string `json:"client_id"`
	RedirectUri string `json:"redirect_uri"`
	Code string `json:"code"`
}
// ShowAccount godoc
// @Summary      kakao access Token
// @Description  kakao access Token
// @Tags         kakao
// @Accept       json 
// @Param        body body KakaoTokenRequestBody false "토."
// @Success      200  {string}  token
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /kakao/access-token [post]
func (control *UserController) KakaoGetAccessToken(c *gin.Context){
	var body KakaoTokenRequestBody

	if err := c.ShouldBindJSON(&body) ;err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	token, err := userService.GetKakaoOauthToken(body.GrantType, body.ClientId, body.RedirectUri, body.Code)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}
	id, err := userService.GetKakaoUser(token)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}
	user, err := userService.FindKakaoUser(id)
	
	if err !=  nil{
		httputil.NewError(c, http.StatusNotFound, err)
		return 
	}
	c.JSON(200, token)
}