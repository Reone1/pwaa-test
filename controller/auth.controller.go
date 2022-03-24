package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
)

type AuthController struct {}
type KakaoCodeRequest struct {
	Code string `json:"code"`
}
func (controller *AuthController) GetKakaoCode(c *gin.Context) {
	// kakao controller
	// 1. get Code endpoint
	var body KakaoCodeRequest
	if err := c.ShouldBindJSON(&body); err != nil{
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	token, err := userService.GetKakaoOauthToken(body.Code)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200, token)

}
func (controller *AuthController)GetKakaoToken(c *gin.Context) {
	// 2. get Access Token endpoint
}
func (controller *AuthController)GetKakaoUser(c *gin.Context) {
  // 3. get UserData Endpoint

}




// twitter Controller
// 1. get Code endpoint
// 2. get accessToekn endpoint
// 3. get userData endpoint

// apple controller
// 1. get User id endpoint
// 2. get