package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"pwaa-test.com/models/entity"
	httputil "pwaa-test.com/module/utils"
	"pwaa-test.com/module/utils/jwt"
	service "pwaa-test.com/services"
)

var userService = new(service.UserService)
type UserController struct {}
// ShowAccount godoc
// @Summary      유저 정보 조회
// @Description  유저 정보 조회
// @Tags         user
// @Accept       json 
// @Success      200  {object}  entity.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /user [get]
// @Security ApiKeyAuth
func (control *UserController) GetUser(c *gin.Context) {
	data, ok := c.Get("userId")
	if !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not authorized"))
		return	
	}
	userId := fmt.Sprintf("%v", data)
	user, err := userService.FindById(userId)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return	
	}
	c.JSON(200, entity.User{
		NickName: user.NickName,
		Type: user.Type,
		DefaultModel: mgm.DefaultModel{
			DateFields: user.DateFields,
		},
	})
}

type SignInRequestBody struct {
	NickName string `json:"nickName"`
	Key string `json:"key"`
	UserType string `json:"type"`
}

type SigninResponse struct {
	Token string `json:"token"`
	NickName string `json:"nickName"`
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
	user, err := userService.FindBykey(body.UserType, body.Key)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, errors.New("create user Error"))
		return
	}
	token, err := new(jwt.Module).CreateToken(user.ID.Hex())
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, errors.New("token Error"))
		return
	}
	c.JSON(200,SigninResponse{
		Token: token,
		NickName: user.NickName,
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
	user, err := userService.FindBykey("normal", body.Key)
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
	user, err := userService.FindByNickName("test")
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
	c.JSON(200, loginResponseBody{
		Token: token,
	})
}