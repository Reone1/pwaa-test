package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}
func NewLoginError(c *gin.Context, status int, err *HTTPLoginError) {
	c.JSON(status, err)
}


type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}


type HTTPLoginError struct {
	Key    string    `json:"Key" example:"user Primary Id"`
	UserType string `json:"type" example:"User type kakao, twitter, apple"`
	Message string `json:"message" example:"error message"`
}
