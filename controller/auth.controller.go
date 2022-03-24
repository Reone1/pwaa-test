package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
)

type AuthController struct {}

func (controller *AuthController) GetKakaoCode(c *gin.Context) {
	// kakao controller
	// 1. get Code endpoint
	query := c.Request.URL.Query()

	token, err := userService.GetKakaoOauthToken(query.Get("code"))
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
		
	userId, err := userService.GetKakaoUser(token)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	
	kakaoUser, err := userService.FindOauthUser("kakao", userId)
	if err != nil {
		httputil.NewLoginError(c, http.StatusNotFound, &httputil.HTTPLoginError{
			UserType: "kakao",
			Key: fmt.Sprint(userId),
			Message: "not found user",
		})
		return
	}
	accessToken, err := userService.GetToken("kakao",kakaoUser.Key)

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return 
	}

	c.JSON(200, loginResponseBody{
		Token: accessToken,
	})
}





// twitter Controller
// 1. get Code endpoint
// 2. get accessToekn endpoint
// 3. get userData endpoint

// apple controller
// 1. get User id endpoint
// 2. get