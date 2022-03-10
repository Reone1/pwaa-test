package router

import (
	controllers "pwaa-test.com/controller"
	authGaurd "pwaa-test.com/module/utils/auth-gaurd"
)

func init() {
	controller := new(controllers.BottleController)
	bottleRouter := router.Group("/bottle")
	bottleRouter.Use(authGaurd.AuthMiddleware)
	{
		bottleRouter.POST("/", controller.Create)
		bottleRouter.GET("/", controller.GetOne)
		bottleRouter.GET("/list", controller.GetMany)
		bottleRouter.GET("/img", controller.GetBottleImg)
		bottleRouter.POST("/status", controller.UpdateBottleStatus)
	}
}