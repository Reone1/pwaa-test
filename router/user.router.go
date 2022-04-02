package router

import (
	controllers "pwaa-test.com/controller"
	authGaurd "pwaa-test.com/module/utils/auth-gaurd"
)


func UserRouter() {
	userController := new(controllers.UserController)
	authController := new(controllers.AuthController)

	testRouter := router.Group("/test")
	{
		testRouter.PUT("/oauth/signout", authGaurd.AuthMiddleware, userController.OAuthSignOut)
		testRouter.POST("/login", userController.TestUserLogin)
	}
	userRouter := router.Group("/user")
	{
		userRouter.POST("/signin", userController.CreateOne)
		userRouter.GET("/", authGaurd.AuthMiddleware, userController.GetUser)
		userRouter.PUT("/privacy", authGaurd.AuthMiddleware, userController.UpdateUserPrivacy)
		userRouter.POST("/login", userController.Login)
	}

	
	twitterRoute := router.Group("/twitter")
	{
		twitterRoute.GET("/request-token",userController.TwitterGetRequest)
		twitterRoute.GET("/access-token", userController.TwitterGetAccess)
	}

	oauthRoute := router.Group("/oauth")
	{
		kakaoRoute := oauthRoute.Group("/kakao")
		{
			kakaoRoute.GET("/", authController.GetKakaoCode)
		}
		apppleRouter := oauthRoute.Group("/apple")
		{
			apppleRouter.POST("/login", authController.AppleLogin)
		}
	}
}