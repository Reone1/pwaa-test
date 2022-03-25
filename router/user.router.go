package router

import (
	controllers "pwaa-test.com/controller"
	authGaurd "pwaa-test.com/module/utils/auth-gaurd"
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
		userRouter.POST("/signin", userController.CreateOne)
		userRouter.GET("/", userController.GetUser).Use(authGaurd.AuthMiddleware)
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