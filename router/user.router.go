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
	
	twitterRoute := router.Group("/twitter")
	{
		twitterRoute.GET("/request-token",controller.TwitterGetAccess)
		twitterRoute.GET("/access-token", controller.TwitterGetToken)
	}

	kakaoRoute := router.Group("/kakao")
	{
		kakaoRoute.GET("/", func (c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "kakao Oauth",
			})
		})
	}
}