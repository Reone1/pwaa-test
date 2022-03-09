package router

import "github.com/gin-gonic/gin"

var router *gin.Engine = gin.Default()

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		// c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
		}

		c.Next()
	}
}
func SetRouter() *gin.Engine{
	if router == nil {
		router = gin.New()
	}
	router.Use(CORSMiddleware())
	return router
}