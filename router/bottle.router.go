package router

import controllers "pwaa-test.com/controller"

func init() {
	controller := new(controllers.BottleController)
	logRouter := router.Group("/bottle")
	{
		logRouter.POST("/", controller.Create)
		logRouter.GET("/", controller.GetOne)
	}
}