package router

import (
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
		userRouter.POST("/login", controller.Login)
	}

	
	twitterRoute := router.Group("/twitter")
	{
		twitterRoute.GET("/request-token",controller.TwitterGetRequest)
		twitterRoute.GET("/access-token", controller.TwitterGetAccess)
	}

	kakaoRoute := router.Group("/kakao")
	{
		kakaoRoute.GET("/login", controller.KakaoGetAccessToken)
	}
	apppleRouter := router.Group("/apple")
	{
		apppleRouter.GET("/login", controller.KakaoGetAccessToken)
	}
}