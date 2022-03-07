package router

import (
	"github.com/gin-gonic/gin"
	controllers "pwaa-test.com/controller"
)

func init() {
	controller := new(controllers.UserController)
	testRouter := router.Group("/test")
	{
		testRouter.POST("/login", controller.TestUserLogin)
	}
	userRouter := router.Group("/user")
	{
		userRouter.POST("/", controller.CreateOne)
		userRouter.GET("/", controller.GetUser)
	}
	
	twitterRoute := userRouter.Group("/twitter")
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

	kakaoRoute := userRouter.Group("/kakao")
	{
		kakaoRoute.GET("/", func (c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "kakao Oauth",
			})
		})
	}
}