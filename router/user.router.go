package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controllers "pwaa-test.com/controller"
)

func init() {
	fmt.Print("User Router called")
	controller := new(controllers.UserController)
	// AuthGaurd, err := authGaurd.SetAuthGaurd()
	// if err != nil {
	// 	log.Panic("authgaurd Error")
	// 	return
	// }
	logRouter := router.Group("/user")
	// logRouter.Use(AuthGaurd.MiddlewareFunc())
	{
		logRouter.POST("/", controller.CreateOne)
		logRouter.GET("/", controller.GetUser)	
	}
	
	twitterRoute := logRouter.Group("/twitter")
	{
		twitterRoute.GET("/access-token", func (c *gin.Context) {
			c.JSON(200, gin.H{
				"message":"t",
			})
		})
		twitterRoute.GET("/", func (c *gin.Context) {
			c.JSON(200, gin.H{
				"message":"token",
			})
		})
	}

	kakaoRoute := logRouter.Group("/kakao")
	{
		kakaoRoute.GET("/", func (c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "kakao Oauth",
			})
		})
	}
}