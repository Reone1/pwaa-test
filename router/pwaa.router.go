package router

import controllers "pwaa-test.com/controller"

func init() {
	controller := new(controllers.PwaaController)
	logRouter := router.Group("/pwaa")
	{
		logRouter.GET("/", controller.GetPwaa)
	}
}