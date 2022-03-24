package controllers

import (
	"encoding/json"
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
	
	type KakaoOauthTokenResponse struct {
		Token string `json:"access_token"`
	}
	res := KakaoOauthTokenResponse{}
	if err := json.Unmarshal([]byte(token), &res); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	userId, err := userService.GetKakaoUser(res.Token)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, gin.H{"userId": userId})
}





// twitter Controller
// 1. get Code endpoint
// 2. get accessToekn endpoint
// 3. get userData endpoint

// apple controller
// 1. get User id endpoint
// 2. get