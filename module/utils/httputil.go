package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
		Body: err,
	}
	c.JSON(status, er)
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
	Body interface {} `json:"body" example:"Server Error Body"`
}

