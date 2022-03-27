package controllers

// twitter Controller
// 1. get Code endpoint
// 2. get accessToken endpoint
// 3. get userData endpoint

// apple controller
// 1. get User id endpoint
// 2. get

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
	"pwaa-test.com/module/utils/jwt"
	service "pwaa-test.com/services"
)
var OAuthService = new(service.OAuthService)
type AuthController struct {}
type OAuthNotFoundUser struct{
	Key string `json:"key"`
	UserType string `json:"type"`
	Message string `json:"message"`
}

// ShowAccount godoc
// @Summary      kakao access Token
// @Description  kakao에서 코드를 이용해 로그인을 할 수 있도록 합니다.
// @Tags         oauth
// @Accept       json 
// @Success      200  {object}  loginResponseBody
// @Success      202  {object}  OAuthNotFoundUser
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /oauth/kakao [get]
func (controller *AuthController) GetKakaoCode(c *gin.Context) {
	// kakao controller
	// 1. get Code endpoint
	query := c.Request.URL.Query()

	token, err := OAuthService.GetKakaoOauthToken(query.Get("code"))
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
		
	userId, err := OAuthService.GetKakaoUser(token)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	
	kakaoUser, err := userService.FindOauthUser("kakao", userId)
	if err != nil {
		c.JSON(202, OAuthNotFoundUser{
			UserType: "kakao",
			Key: fmt.Sprint(userId),
			Message: "not found user",
		})
		return
	}
	accessToken, err := userService.GetToken(kakaoUser.ID.Hex())

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}

	c.JSON(200, loginResponseBody{
		Token: accessToken,
	})
}

type AppleLoginRequestBody struct {
	UserId string `json:"user"`
}
// ShowAccount godoc
// @Summary      apple Login
// @Description  apple Login
// @Tags         oauth
// @Accept       json 
// @Param        body body AppleLoginRequestBody false "Apple 로그인"
// @Success      200  {object}  loginResponseBody
// @Success      202  {object}  OAuthNotFoundUser
// @Failure      404  {object}  httputil.HTTPError
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /oauth/apple/login [post]
func (controller *AuthController) AppleLogin(c *gin.Context) {
	var body AppleLoginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}
	user, err := userService.FindOauthUser("apple", body.UserId)
	if err !=  nil{
		c.JSON(202, OAuthNotFoundUser{
			UserType: "apple",
			Key: body.UserId,
			Message: "not found user",
		})
		return 
	}
	token, err := new(jwt.Module).CreateToken(user.ID.Hex())
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
// @Tags         oauth
// @Accept       json 
// @Param        query query TwitterGetAccessRequestQuery false "callback_url"
// @Success      200  {string}  requestToken
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /twitter/requset-token [get]
func (control *UserController) TwitterGetRequest(c *gin.Context){
	var query = c.Request.URL.Query()
	requestToken, requestSecret, err := OAuthService.GetTwitterAuthToken(query.Get("callback_url"))
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
// @Tags         oauth
// @Accept       json
// @Param        query query TwitterGetTokenRequestBody false "twitter oauth 토큰이 필요합니다."
// @Success      200  {string}  token
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /twitter/access-token [get]
func (control *UserController) TwitterGetAccess(c *gin.Context){
	var query = c.Request.URL.Query()

	token, err := OAuthService.GetTwitterAccessToken(query.Get("oauth_token"), query.Get("oauth_token_secret"), query.Get("oauth_verifier"), query.Get("callbackURL"))

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return;
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}