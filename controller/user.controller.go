package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pwaa-test.com/models/entity"
	httputil "pwaa-test.com/module/utils"
	"pwaa-test.com/module/utils/jwt"
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

type SignInRequestBody struct {
	NickName string `json:"nickName"`
	Key string `json:"key"`
	UserType string `json:"type"`
}

type SigninResponse struct {
	Message string `json:"message"`
}
// ShowAccount godoc
// @Summary      유저 닉네임 생성
// @Description  유저 닉네임 생성
// @Tags         user
// @Accept       json 
// @Param        body body SignInRequestBody false "Signin request Body"
// @Success      200  {object}  SigninResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /signin [post]
func (control *UserController) CreateOne(c *gin.Context) {
	var body SignInRequestBody
	if err := c.ShouldBindJSON(&body); err != nil{
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	
	if err := userService.Create(&entity.User{
		Type: body.UserType,
		Key: body.Key,
		NickName: body.NickName,
		}); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200,SigninResponse{
		Message: "ok",
	})
}

type LoginRequestBody struct {
	Key string `json:"key" exmaple:"user key from server"`
}
type loginResponseBody struct {
	Token string `json:"token" example:"token string (JWT)"`
}
// ShowAccount godoc
// @Summary      로그인
// @Description  로그인
// @Tags         user
// @Accept       json 
// @Param        body body LoginRequestBody false "유저 로그인 토큰 발급"
// @Success      200  {object}  loginResponseBody
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /user/login [post]
func (control *UserController) Login(c *gin.Context) {
	var body LoginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil{
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	user, err := userService.FindBykey(body.Key)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	jwtModule := new(jwt.Module)
	token, err := jwtModule.CreateToken(user.ID.Hex())
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200,gin.H{
		"token": token,
	})
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

	token, err := userService.GetTwitterAccessToken(query.Get("oauth_token"), query.Get("oauth_token_secret"), query.Get("oauth_verifier"), query.Get("callbackURL"))

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return;
	}
	c.JSON(200, token)
}

type KakaoTokenRequestBody struct {
	ClientId string `json:"client_id"`
	RedirectUri string `json:"redirect_uri"`
	Code string `json:"code"`
}

// ShowAccount godoc
// @Summary      kakao access Token
// @Description  kakao access Token
// @Tags         oauth
// @Accept       json 
// @Param        body body KakaoTokenRequestBody false "kakao 로그인"
// @Success      200  {object}  loginResponseBody
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPLoginError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /kakao/login [post]
func (control *UserController) KakaoGetAccessToken(c *gin.Context){
	var body KakaoTokenRequestBody

	if err := c.ShouldBindJSON(&body) ;err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	kakaoToken, err := userService.GetKakaoOauthToken(body.Code)

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}

	id, err := userService.GetKakaoUser(kakaoToken)

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}

	kakaoUser, err := userService.FindOauthUser("kakao", id)

	if err !=  nil{
		httputil.NewLoginError(c, http.StatusNotFound, &httputil.HTTPLoginError{
			UserType: "kakao",
			Key: id,
			Message: "not found user",
		})
		return 
	}

	accessToken, err := userService.GetToken("kakao", kakaoUser.ID.Hex())

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
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPLoginError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /apple/login [post]
func (controller *UserController) AppleLogin(c *gin.Context) {
	var body AppleLoginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}
	user, err := userService.FindOauthUser("apple", body.UserId)
	if err !=  nil{
		httputil.NewLoginError(c, http.StatusNotFound, &httputil.HTTPLoginError{
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