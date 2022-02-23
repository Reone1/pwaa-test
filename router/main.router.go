package router

import "github.com/gin-gonic/gin"

var router *gin.Engine = gin.Default()

func SetRouter() *gin.Engine{
	if router == nil {
		router = gin.New()
	}
	return router
}