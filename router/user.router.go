package router

import (
	controllers "pwaa-test.com/controller"
)

func init() {
	userController := new(controllers.UserController)
	authController := new(controllers.AuthController)

	testRouter := router.Group("/test")
	{
		testRouter.POST("/login", userController.TestUserLogin)
	}
	userRouter := router.Group("/user")
	{
		userRouter.POST("/", userController.CreateOne)
		userRouter.GET("/", userController.GetUser)
		userRouter.POST("/login", userController.Login)
	}

	
	twitterRoute := router.Group("/twitter")
	{
		twitterRoute.GET("/request-token",userController.TwitterGetRequest)
		twitterRoute.GET("/access-token", userController.TwitterGetAccess)
	}

	authRoute := router.Group("/oauth")
	{
		authRoute.GET("/kakao", authController.GetKakaoCode)
	}

	apppleRouter := router.Group("/apple")
	{
		apppleRouter.POST("/login", userController.AppleLogin)
	}
}